package libtc

import (
	"fmt"
	"strconv"
	"strings"
)

type FilterU32Match struct {
	// raw values
	Size   int
	Mask   uint32
	Value  uint32
	Offset int
	// known fields
	Name  string
	Param string
}

func (fm FilterU32Match) Empty() bool {
	return fm.Mask == 0 && len(fm.Param) == 0
}

func (fm FilterU32Match) String() []string {
	if fm.Empty() {
		return split("u8 0x00 0x00 at 0")
	}
	if len(fm.Param) != 0 {
		return split(fmt.Sprintf("%s %s", fm.Name, fm.Param))
	} else {
		var size = fm.Size
		if size == 0 {
			size = 32
		}
		return split(fmt.Sprintf("u%d %s %s at %d",
			size,
			hex(fm.Value, fm.Size),
			hex(fm.Mask, fm.Size),
			fm.Offset))
	}
}

type FilterU32 struct {
	Table     int
	Bucket    int
	Index     int
	Div       int
	LinkTable int
	Hash      FilterU32Match
	Sample    FilterU32Match
	Match     []FilterU32Match
	Options   []Option
}

func (f *FilterU32) Handle() string {
	if f.Div != 0 {
		return fmt.Sprintf("%x:", f.Table)
	} else {
		return fmt.Sprintf("%x:%x:%x", f.Table, f.Bucket, f.Index)
	}
}

func (f *FilterU32) Id() int {
	if f.Div != 0 {
		return f.Table
	} else {
		return f.Index
	}
}

func (f *FilterU32) MakeAddParams() []string {
	if f.Div != 0 {
		return []string{
			HANDLE_PARAM, fmt.Sprintf("%x:", f.Table),
			FILTER_U32,
			U32_DIV_PARAM, strconv.Itoa(f.Div),
		}
	}
	var params = []string{
		HANDLE_PARAM, fmt.Sprintf("::%x", f.Index),
		FILTER_U32,
		U32_TABLE_PARAM, fmt.Sprintf("%x:%x", f.Table, f.Bucket),
	}
	if f.LinkTable >= 0 {
		params = append(params, U32_LINK_PARAM)
		params = append(params, fmt.Sprintf("%x:", f.LinkTable))
	}
	if len(f.Match) == 0 {
		params = append(params, U32_MATCH_PARAM)
		params = append(params, FilterU32Match{}.String()...)
	} else {
		for _, match := range f.Match {
			params = append(params, U32_MATCH_PARAM)
			params = append(params, match.String()...)
		}
	}
	if f.Hash.Mask != 0 {
		params = append(params, U32_HASH_PARAM)
		params = append(params, split(fmt.Sprintf("mask %s at %d",
			hex(f.Hash.Mask, 0), f.Hash.Offset))...)
	}
	if !f.Sample.Empty() {
		params = append(params, U32_SAMPLE_PARAM)
		params = append(params, f.Sample.String()...)
	}

	return append(params, opt2params(f.Options)...)
}

func (f *FilterU32) MakeDelParams() []string {
	if f.Div != 0 {
		return []string{
			HANDLE_PARAM, fmt.Sprintf("%x:", f.Table),
			FILTER_U32,
		}
	} else {
		return []string{
			HANDLE_PARAM, fmt.Sprintf("%x:%x:%x", f.Table, f.Bucket, f.Index),
			FILTER_U32,
		}
	}
}

func (f *FilterU32) ParseParams(lines []string, flt *Filter) error {

	for _, line := range lines {
		var params = split(line)
		for n := 0; n < len(params); n++ {
			switch params[n] {
			case U32_FLOW_PARAM:
				flt.ToClass = findString(params[n+1], 1, 16)
			case U32_HANDLE_PARAM:
				var s = params[n+1]
				f.Index = findString(s, 2, 16)
				f.Bucket = findString(s, 1, 16)
				f.Table = findString(s, 0, 16)
			case U32_DIV_PARAM:
				f.Div = s2i(params[n+1])
				f.Index = -1
				f.Bucket = -1
			case U32_MATCH_PARAM:
				var valmask = strings.Split(params[n+1], "/")
				if len(valmask) != 2 {
					break
				}
				f.Match = append(f.Match, FilterU32Match{
					Mask:   hex2uint(valmask[1]),
					Value:  hex2uint(valmask[0]),
					Offset: s2i(params[n+3]),
				})
			}
		}
	}
	return nil
}

type FilterFW struct {
	Mark int
}

func (f *FilterFW) Handle() string {
	return fmt.Sprintf("%d", f.Mark)
}

func (f *FilterFW) Id() int {
	return f.Mark
}

func (f *FilterFW) MakeAddParams() []string {
	return []string{
		HANDLE_PARAM, f.Handle(),
		FILTER_FW,
	}
}

func (f *FilterFW) MakeDelParams() []string {
	return f.MakeAddParams()
}

func (f *FilterFW) ParseParams(lines []string, flt *Filter) error {

	for _, line := range lines {
		var params = split(line)
		for i := 0; i < len(params); i++ {
			switch params[i] {
			case FW_FLOW_PARAM:
				flt.ToClass = findString(params[i+1], 1, 16)
			case FW_HANDLE_PARAM:
				f.Mark = s2i(params[i+1])
			}
		}
	}
	return nil
}
