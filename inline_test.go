package libwimark

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Embedded1 struct {
	A1   int    `json:"a1"`
	A2   string `json:"a2"`
	Type string `json:"type"`
}

type Embedded2 struct {
	A1   string `json:"a1"`
	A2   int    `json:"a2"`
	Type string `json:"type"`
}

type EmbeddedUni struct {
	I    int    `json:"i"`
	S    string `json:"s"`
	Type string `json:"type"`
}

type EmbeddedFull struct {
	B    bool   `json:"b"`
	Type string `json:"type"`
}

type GeneralStruct struct {
	I  int                     `json:"i"`
	S  string                  `json:"s"`
	E1 map[string]Embedded1    `json:"-" inline:"yes,type:a"`
	E2 map[string]Embedded2    `json:"-" inline:"yes,type:b"`
	E  map[string]EmbeddedFull `json:"-" inline:"yes"`
	U  EmbeddedUni             `json:"-" inline:"yes,unique,u,type:u"`
}

var SampleStruct = GeneralStruct{
	I: 1,
	S: "string",
	E1: map[string]Embedded1{
		"A": Embedded1{
			A1:   42,
			A2:   "24",
			Type: "a",
		},
	},
	E2: map[string]Embedded2{
		"B": Embedded2{
			A1:   "24",
			A2:   42,
			Type: "b",
		},
	},
	E: map[string]EmbeddedFull{
		"E": EmbeddedFull{
			B:    true,
			Type: "type",
		},
	},
	U: EmbeddedUni{
		I:    10,
		S:    "10",
		Type: "u",
	},
}

var SampleString = `
{
	"A": {
		"a1": 42,
		"a2": "24",
		"type": "a"
	},
	"B": {
		"a1": "24",
		"a2": 42,
		"type": "b"
	},
	"E": {
		"b": true,
		"type": "type"
	},
	"i": 1,
	"s": "string",
	"u": {
		"i": 10,
		"s": "10",
		"type": "u"
	}
}
`

func TestMarshal(t *testing.T) {
	b, e := MarshalInline(&SampleStruct)
	assert.True(t, e == nil, fmt.Sprint(e))

	var exp = SampleString
	exp = strings.Replace(exp, "\n", "", -1)
	exp = strings.Replace(exp, "\t", "", -1)
	exp = strings.Replace(exp, " ", "", -1)
	assert.Equal(t, exp, string(b))
}

func TestUnmarshal(t *testing.T) {
	var f GeneralStruct = SampleStruct
	var fld = map[string]interface{}{
		"type:a": &map[string]Embedded1{},
		"type:b": &map[string]Embedded2{},
		"type:u": &EmbeddedUni{},
		"":       &map[string]EmbeddedFull{},
	}
	e := UnmarshalInline([]byte(SampleString), &f, fld)
	assert.True(t, e == nil, fmt.Sprint(e))
	assert.Equal(t, SampleStruct, f)
}
