package libwimark

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSecuritySuite(t *testing.T) {
	var fixture = []byte(`["tkip"]`)
	var expectation = TKIP
	var result []SecuritySuite
	var err = json.Unmarshal(fixture, &result)

	if err != nil {
		panic(err)
	}
	assert.Equal(t, expectation, result[0])
}

func TestEnumSecurity(t *testing.T) {
	var fixture = []byte(`{"type": "wpa2personal", "data": {"psk": "qwerty", "suite": "tkip"}}`)
	var d = &WPA2PersonalData{}
	d.Suite = TKIP
	d.PSK = "qwerty"
	var expectation = EnumSecurity{T: WPA2Personal, D: d}
	var result EnumSecurity
	var err = json.Unmarshal(fixture, &result)

	if err != nil {
		panic(err)
	}
	assert.True(t, reflect.DeepEqual(expectation, result))
}

func TestWLAN(t *testing.T) {
	var fixture = []byte(`
            {
                "name": "myhotspot",
                "ssid": "qwertyasdfgh",
                "description": "This is my pretty little honeypot",
                "security": {
                    "type": "wpa2personal",
                    "data": {
                        "psk": "qwerty",
                        "suite": "aes"
                    }
                },
                "vlan": 5,
                "radius_accounting": []
            }
        `)
	var d = &WPA2PersonalData{}
	d.Suite = AES
	d.PSK = "qwerty"
	var expectation = WLAN{
		Name:        "myhotspot",
		SSID:        "qwertyasdfgh",
		Description: "This is my pretty little honeypot",
		Security: &EnumSecurity{
			T: WPA2Personal,
			D: d,
		},
		VLAN:             5,
		RadiusAccounting: []UUID{},
	}

	var result WLAN
	var err = json.Unmarshal(fixture, &result)

	if err != nil {
		panic(err)
	}
	assert.True(t, reflect.DeepEqual(expectation, result))
}
