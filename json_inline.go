// json_inline.go
package libwimark

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ms "github.com/mitchellh/mapstructure"
	deepcopy "github.com/mohae/deepcopy"
)

type inlineTag struct {
	Enabled bool
	Unique  bool
	Attrs   map[string]string
	Key     string
	Name    string
}

func parseInlineTag(tag string) inlineTag {
	var res = inlineTag{Attrs: map[string]string{}}
	var keys = []string{}

	var tags = strings.Split(tag, ",")
	for _, t := range tags {
		switch strings.ToLower(t) {
		case "yes":
			res.Enabled = true
		case "no":
			res.Enabled = false
		case "unique":
			res.Unique = true
		default:
			index := strings.Index(t, ":")
			if index != -1 {
				keys = append(keys, t)
				res.Attrs[t[:index]] = t[index+1:]
			} else {
				res.Name = t
			}
		}
	}
	res.Key = strings.Join(keys, ",")
	return res
}

func mapDecode(from interface{}) (map[string]interface{}, error) {
	// wont use mapstructure,
	// just because it ignores 'omitempty' and such json tags
	var m map[string]interface{}
	b, err := json.Marshal(from)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &m)
	return m, err
}

func strDecode(from interface{}, to interface{}) error {
	var dc = ms.DecoderConfig{TagName: "json", Result: to}
	var dec *ms.Decoder
	dec, _ = ms.NewDecoder(&dc)
	return dec.Decode(from)
}

func extractJsonTag(field reflect.StructField) string {
	jstag := field.Tag.Get("json")
	ind := strings.Index(jstag, ",")
	if ind != -1 {
		jstag = jstag[:ind]
	}
	return jstag
}

func MarshalInline(val interface{}) (b []byte, e error) {

	m, err := mapDecode(val)
	if err != nil {
		return nil, err
	}
	var v = reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return json.Marshal(m)
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := parseInlineTag(field.Tag.Get("inline"))
		if tag.Enabled {
			f, e := mapDecode(v.Field(i).Interface())
			if e != nil {
				return nil, e
			}
			if tag.Unique {
				if len(tag.Name) == 0 {
					m[field.Name] = f
				} else {
					m[tag.Name] = f
				}
			} else {
				for n, v := range f {
					m[n] = v
				}
			}
		}
	}

	return json.Marshal(m)
}

func UnmarshalInline(b []byte, val interface{}, tmpl map[string]interface{}) error {

	var doc map[string]interface{}
	var err error
	if err = json.Unmarshal(b, &doc); err != nil {
		return err
	}
	if doc == nil {
		return nil
	}
	if err = json.Unmarshal(b, val); err != nil {
		return err
	}
	// if 'val' is not a ptr, we'll be kicked out by json.Unmarshal
	var v = reflect.ValueOf(val).Elem()
	if v.Kind() != reflect.Struct {
		return nil
	}

	// remove non-inline
	var key_inlines = []int{}
	var def_inlines = []int{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := parseInlineTag(field.Tag.Get("inline"))
		if !tag.Enabled {
			delete(doc, extractJsonTag(field))
		} else if len(tag.Attrs) != 0 {
			key_inlines = append(key_inlines, i)
		} else {
			def_inlines = append(def_inlines, i)
		}
	}
	var none_attrs = deepcopy.Copy(doc).(map[string]interface{})
	for _, i := range key_inlines {
		field := v.Type().Field(i)
		tag := parseInlineTag(field.Tag.Get("inline"))
		var attrs = map[string]interface{}{}
		for key, val := range doc {
			vmap, ok := val.(map[string]interface{})
			if !ok {
				continue
			}
			ok = true
			for check_key, check_val := range tag.Attrs {
				value, has_value := vmap[check_key]
				if !has_value || value != check_val {
					ok = false
				}
			}
			if !ok {
				continue
			}
			delete(none_attrs, key)
			if tag.Unique {
				attrs = vmap
				break
			} else {
				attrs[key] = val
			}
		}
		var f = v.Field(i)
		tmp, ok := tmpl[tag.Key]
		if !ok {
			return errors.New(fmt.Sprintf("No template for %s", tag.Key))
		}
		err = strDecode(attrs, tmp)
		if err != nil {
			return err
		}
		f.Set(reflect.ValueOf(tmp).Elem())
	}
	for _, i := range def_inlines {
		var f = v.Field(i)
		tmp, ok := tmpl[""]
		if !ok {
			return errors.New("No template for default key")
		}
		err = strDecode(none_attrs, tmp)
		if err != nil {
			return err
		}
		f.Set(reflect.ValueOf(tmp).Elem())
	}

	return nil
}
