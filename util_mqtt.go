package libwimark

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	ms "github.com/mitchellh/mapstructure"
	deepcopy "github.com/mohae/deepcopy"
	cache "github.com/patrickmn/go-cache"
)

func MQTTConnectSyncOpts(opts *mqtt.ClientOptions) (mqtt.Client, error) {
	var client = mqtt.NewClient(opts)
	var token = client.Connect()
	token.Wait()

	return client, token.Error()
}

func MQTTMustConnectSyncOpts(opts *mqtt.ClientOptions) mqtt.Client {
	var client, err = MQTTConnectSyncOpts(opts)
	if err != nil {
		panic(err)
	}

	return client
}

func MQTTConnectSync(addr string) (mqtt.Client, error) {
	var opts = mqtt.NewClientOptions()
	opts.AddBroker(addr)
	return MQTTConnectSyncOpts(opts)
}

func MQTTMustConnectSync(addr string) mqtt.Client {
	var client, err = MQTTConnectSync(addr)
	if err != nil {
		panic(err)
	}

	return client
}

func onDisconnect(client mqtt.Client, err error) {
	time.Sleep(time.Second * 1)
	panic("MQTT broker connection lost")
}

func MQTTServiceStart(addr string, s Module, v Version, meta interface{}) (mqtt.Client, error) {
	return MQTTServiceStartWithId(addr, s, v, "", meta)
}

func MQTTServiceStartWithId(addr string, s Module, v Version, id string, meta interface{}) (mqtt.Client, error) {
	ts := time.Now().Unix()

	// prepare disconnect event for will message
	var opts = mqtt.NewClientOptions().AddBroker(addr)
	eventDisconnetTopic := EventTopic{
		SenderModule: s,
		SenderID:     "",
		Type:         SystemEventTypeServiceDisconnected,
	}

	level := SystemEventLevelINFO

	eventObject := SystemEventObject{
		Type: SystemEventTypeServiceDisconnected,
		Data: nil,
	}
	eventDisconnetPayload := SystemEvent{
		Subject_id: s.String(),
		Timestamp:  ts,
		Level:      level,
	}
	eventDisconnetPayload.SystemEventObject = eventObject

	b, err := json.Marshal(eventDisconnetPayload)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error: (%s) during service start", err.Error()))
	}
	opts.SetWill(eventDisconnetTopic.TopicPath(), string(b), 2, false)

	opts.SetClientID(s.String())

	opts.SetConnectionLostHandler(onDisconnect)

	client, err := MQTTConnectSyncOpts(opts)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error: (%s) while connecting to broker", err.Error()))
	}

	//sending retain status message
	statusTopic := StatusTopic{
		SenderModule: s,
		SenderID:     id,
	}

	statusPayload := ModuleStatus{
		Version: v.Version,
		Commit:  v.Commit,
		Build:   v.Build,
		Service: s,
		Id:      id,
		State:   ServiceStateConnected,
		Meta:    meta,
	}

	b, err = json.Marshal(statusPayload)
	if err != nil {
		return client, errors.New(fmt.Sprintf("error: (%s) while marshalling to status message ", err.Error()))
	}
	err = MQTTPublishRetainedSync(client, statusTopic, b)
	if err != nil {
		return client, errors.New(fmt.Sprintf("error: (%s) while publishing retained in service start", err.Error()))
	}

	// sending connect event
	eventConnectTopic := EventTopic{
		SenderModule: s,
		SenderID:     id,
		Type:         SystemEventTypeServiceConnected,
	}

	level = SystemEventLevelINFO

	eventData := ServiceConnectedData{
		Version: v.Version,
		Commit:  v.Commit,
		Build:   v.Build,
	}
	eventObject = SystemEventObject{
		Type: SystemEventTypeServiceConnected,
		Data: eventData,
	}
	eventConnectedPayload := SystemEvent{
		Subject_id: s.String(),
		Timestamp:  ts,
		Level:      level,
	}
	eventConnectedPayload.SystemEventObject = eventObject

	b, err = json.Marshal(eventConnectedPayload)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error: (%s) marshalling connected event in service start", err.Error()))
	}

	err = MQTTPublishSync(client, eventConnectTopic, b)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error: (%s) publishing connected event in service start", err.Error()))
	}
	return client, nil
}

type MQTTMessage interface {
	Topic() Topic
	Payload() ([]byte, error)
	Retained() bool
}

type MQTTDocumentMessage struct {
	T Topic
	D interface{}
	R bool
}

func (self MQTTDocumentMessage) Topic() Topic {
	return self.T
}

func (self MQTTDocumentMessage) Payload() ([]byte, error) {
	if self.D != nil {
		return json.Marshal(self.D)
	} else {
		return nil, nil
	}
}

func (self MQTTDocumentMessage) Retained() bool {
	return self.R
}

type MQTTRawMessage struct {
	T Topic
	D []byte
	R bool
}

func (self MQTTRawMessage) Topic() Topic {
	return self.T
}

func (self MQTTRawMessage) Payload() ([]byte, error) {
	return self.D, nil
}

func (self MQTTRawMessage) Retained() bool {
	return self.R
}

func MQTTPublishMsg(client mqtt.Client, log_cb func(string), msg MQTTMessage) {
	var topic_str = msg.Topic().TopicPath()
	var payload, pErr = msg.Payload()

	if pErr != nil {
		log_cb(fmt.Sprintf("Error marshalling outgoind payload. Topic: %s, Error: %s", topic_str, pErr))
	} else {
		log_cb(fmt.Sprintf("Sending message - Topic: %s, Payload: %s\n", topic_str, payload))
		client.Publish(topic_str, 2, msg.Retained(), payload)
	}
}

func MQTTMakePublishChan(client mqtt.Client, log_cb func(string)) chan<- MQTTMessage {
	var publishChan = make(chan MQTTMessage)
	go func() {
		for msg := range publishChan {
			MQTTPublishMsg(client, log_cb, msg)
		}
	}()

	return publishChan
}

type MsgCb func(mqtt.Message)

func MQTTSubscribeSync(client mqtt.Client, topics []Topic, cb MsgCb) error {
	if topics != nil {
		var topiclist = map[string]byte{}
		for _, t := range topics {
			if t != nil {
				topiclist[t.TopicPath()] = 2
			}
		}

		var token = client.SubscribeMultiple(topiclist, func(_ mqtt.Client, msg mqtt.Message) {
			cb(msg)
		})

		token.Wait()
		var err = token.Error()
		if err != nil {
			return err
		}
	}

	return nil
}

func MQTTMustSubscribeSync(client mqtt.Client, topics []Topic, cb MsgCb) {
	var err = MQTTSubscribeSync(client, topics, cb)
	if err != nil {
		panic(err)
	}
}

func MQTTUnsubscribeSync(client mqtt.Client, topics []Topic) error {
	var err error
	if topics != nil {
		var topiclist = []string{}
		for _, t := range topics {
			if t != nil {
				topiclist = append(topiclist, t.TopicPath())
			}
		}
		var token = client.Unsubscribe(topiclist...)
		token.Wait()
		err = token.Error()
		if err != nil {
			return err
		}
	}

	return err
}

func MQTTMustUnsubscribeSync(client mqtt.Client, topics []Topic) {
	var err = MQTTUnsubscribeSync(client, topics)
	if err != nil {
		panic(err)
	}
}

func MQTTPublishSync(client mqtt.Client, topic Topic, payload interface{}) error {
	var token = client.Publish(topic.TopicPath(), 2, false, payload)
	token.Wait()
	var err = token.Error()
	return err
}

func MQTTPublishRetainedSync(client mqtt.Client, topic Topic, payload interface{}) error {
	var token = client.Publish(topic.TopicPath(), 2, true, payload)
	token.Wait()
	var err = token.Error()
	return err
}

func CacheSetRequest(c *cache.Cache, id string, d time.Duration, v MQTTMessage) {
	c.Set(id, v, d)
}

func CacheGetRequest(c *cache.Cache, id string) MQTTMessage {
	var v, ok = c.Get(id)
	if !ok {
		return nil
	}
	return v.(MQTTMessage)
}

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
