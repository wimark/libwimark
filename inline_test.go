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
	Num  string `json:"num"`
}

type Embedded3 struct {
	S    string `json:"s"`
	Type string `json:"type"`
	Num  string `json:"num"`
}

type EmbeddedUni struct {
	I    int    `json:"i,omitempty"`
	S    string `json:"s,omitempty"`
	Type string `json:"type,omitempty"`
}

type EmbeddedFull struct {
	B    bool   `json:"b"`
	Type string `json:"type"`
}

type GeneralStruct struct {
	I  int                     `json:"i"`
	S  string                  `json:"s"`
	E1 map[string]Embedded1    `json:"-" inline:"yes,type:a"`
	E2 map[string]Embedded2    `json:"-" inline:"yes,type:b,num:2"`
	E3 map[string]Embedded3    `json:"-" inline:"yes,type:b,num:3"`
	E  map[string]EmbeddedFull `json:"-" inline:"yes"`
	U1 EmbeddedUni             `json:"-" inline:"yes,unique,u1,type:u1"`
	U2 EmbeddedUni             `json:"-" inline:"yes,unique,omitempty,u2,type:u2"`
	U3 EmbeddedUni             `json:"-" inline:"yes,unique,omitempty,u3,type:u3"`
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
		"B2": Embedded2{
			A1:   "24",
			A2:   42,
			Type: "b",
			Num:  "2",
		},
	},
	E3: map[string]Embedded3{
		"B3": Embedded3{
			S:    "sss",
			Type: "b",
			Num:  "3",
		},
	},
	E: map[string]EmbeddedFull{
		"E": EmbeddedFull{
			B:    true,
			Type: "type",
		},
	},
	U1: EmbeddedUni{
		I:    10,
		S:    "10",
		Type: "u1",
	},
	U2: EmbeddedUni{
		I:    20,
		S:    "20",
		Type: "u2",
	},
	U3: EmbeddedUni{},
}

var SampleString = `
{
	"A": {
		"a1": 42,
		"a2": "24",
		"type": "a"
	},
	"B2": {
		"a1": "24",
		"a2": 42,
		"num": "2",
		"type": "b"
	},
	"B3": {
		"num": "3",
		"s": "sss",
		"type": "b"
	},
	"E": {
		"b": true,
		"type": "type"
	},
	"i": 1,
	"s": "string",
	"u1": {
		"i": 10,
		"s": "10",
		"type": "u1"
	},
	"u2": {
		"i": 20,
		"s": "20",
		"type": "u2"
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
		"type:a":       &map[string]Embedded1{},
		"type:b,num:2": &map[string]Embedded2{},
		"type:b,num:3": &map[string]Embedded3{},
		"type:u1":      &EmbeddedUni{},
		"type:u2":      &EmbeddedUni{},
		"type:u3":      &EmbeddedUni{},
		"":             &map[string]EmbeddedFull{},
	}
	e := UnmarshalInline([]byte(SampleString), &f, fld)
	assert.True(t, e == nil, fmt.Sprint(e))
	assert.Equal(t, SampleStruct, f)
}
