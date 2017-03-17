package libwimark

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCompare(t *testing.T) {
	var fixture1 = SecuritySuite{TKIP{}}
	var fixture2 = SecuritySuite{TKIP{}}
	var fixture3 = SecuritySuite{AES{}}

	assert.Equal(t, fixture1, fixture2)
	assert.NotEqual(t, fixture1, fixture3)

	assert.True(t, fixture1 == fixture2)
	assert.True(t, fixture1 != fixture3)
}

func TestSecuritySuite(t *testing.T) {
	var fixture = []byte(`["tkip"]`)
	var expectation = SecuritySuite{TKIP{}}
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
	d.Suite = SecuritySuite{TKIP{}}
	d.PSK = "qwerty"
	var expectation = EnumSecurity{T: SecurityType{WPA2Personal{}}, D: d}
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
	d.Suite = SecuritySuite{AES{}}
	d.PSK = "qwerty"
	var expectation = WLAN{
		Name:        "myhotspot",
		SSID:        "qwertyasdfgh",
		Description: "This is my pretty little honeypot",
		Security: &EnumSecurity{
			T: SecurityType{WPA2Personal{}},
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

func TestCPE(t *testing.T) {
	var fixture = []byte(`
        {
            "name": "mycpe",
            "description": "this is my CPE",
            "model": "MODELID001",
            "interfaces": {
                "wlan0": {
                    "addr": "macaddr0",
                    "type": "wifi",
                    "data": {
                        "name": "WLAN Interface 0",
                        "mac": "macaddr1",
                        "frequency": 2.4,
                        "band_mode": "n",
                        "bandwidth": "20",
                        "channel": "5",
                        "tx_power": 100,
                        "wlans": ["WLAN001", "WLAN002"]
                    }
                }
            },
            "config_status": "ok"
        }
    `)
	var d = &WiFiData{}
	d.Name = "WLAN Interface 0"
	d.Mac = "macaddr1"
	d.Frequency = 2.4
	d.BandMode = "n"
	d.Bandwidth = "20"
	d.Channel = "5"
	d.TxPower = 100
	d.WLANs = []UUID{"WLAN001", "WLAN002"}

	var i CPEInterface
	i.Addr = "macaddr0"
	i.T = CPEInterfaceType{InterfaceWiFi{}}
	i.D = d

	var expectation CPE
	expectation.Name = "mycpe"
	expectation.Description = "this is my CPE"
	expectation.Model = UUID("MODELID001")
	expectation.Interfaces = map[string]CPEInterface{}
	expectation.Interfaces["wlan0"] = i
	expectation.ConfigStatus = ConfigurationStatus{StatusOK{}}

	var result CPE
	var err = json.Unmarshal(fixture, &result)

	if err != nil {
		panic(err)
	}

	assert.True(t, reflect.DeepEqual(expectation, result))
}
