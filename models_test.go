package libwimark

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
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
	var fixture = []byte(`{"type": "wpa2personal", "data": {"psk": "qwerty", "suites": ["tkip"]}}`)
	var d = &WPA2PersonalData{}
	d.Suites = []SecuritySuite{SecuritySuite{TKIP{}}}
	d.PSK = "qwerty"
	var expectation = EnumSecurity{T: SecurityType{WPA2Personal{}}, D: d}
	var result EnumSecurity
	var err = json.Unmarshal(fixture, &result)

	if err != nil {
		panic(err)
	}
	assert.Equal(t, expectation, result)
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
                        "suites": ["aes"]
                    }
                },
                "vlan": 5,
                "radius_acct_servers": [
                    "nas001"
                ]
            }
        `)
	var d = &WPA2PersonalData{}
	d.Suites = []SecuritySuite{SecuritySuite{AES{}}}
	d.PSK = "qwerty"
	var expectation = WLAN{
		Name:        "myhotspot",
		SSID:        "qwertyasdfgh",
		Description: "This is my pretty little honeypot",
		Security: &EnumSecurity{
			T: SecurityType{WPA2Personal{}},
			D: d,
		},
		VLAN: 5,
	}
	expectation.RadiusAcctServers = []UUID{"nas001"}

	var result WLAN
	var err = json.Unmarshal(fixture, &result)

	if err != nil {
		panic(err)
	}
	assert.Equal(t, expectation, result)
}

func TestCPE(t *testing.T) {
	var fixture = []byte(`
        {
            "name": "mycpe",
            "description": "this is my CPE",
            "connected": true,
            "model": "MODELID001",
            "interfaces": {
                "wlan0": {
                    "addr": "macaddr0",
                    "capabilities":{
                        "channels":[
                            {
                                "channel":1,
                                "frequency":2417,
                                "bandwidth":[
                                    "20",
                                    "HT40+"
                                ],
                                "mode":"b/g/n",
                                "radardetection":true,
                                "maxtxpower":23
                            }
                        ]
                    },
                    "type": "wifi",
                    "data": {
                        "name": "WLAN Interface 0",
                        "mac": "macaddr1",
                        "frequency": 2.4,
                        "bandmode": "n",
                        "bandwidth": "20",
                        "channel": 5,
                        "txpower": 100,
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
	d.Channel = 5
	d.TxPower = 100
	d.WLANs = []UUID{"WLAN001", "WLAN002"}

	var i CPEInterface
	i.Addr = "macaddr0"
	i.T = CPEInterfaceType{InterfaceWiFi{}}
	i.D = d

	var channel_caps ChannelCapabilities
	channel_caps.Channel = 1
	channel_caps.Frequency = 2417
	channel_caps.Bandwidth = []BandwidthType{
		BandwidthType{Bandwidth20{}},
		BandwidthType{BandwidthHT40Plus{}},
	}
	channel_caps.Mode = "b/g/n"
	channel_caps.RadarDetection = true
	channel_caps.MaxTxPower = 23

	var caps Capabilities
	caps.Channels = []ChannelCapabilities{channel_caps}
	i.Capabilities = caps

	var expectation CPE
	expectation.Name = "mycpe"
	expectation.Description = "this is my CPE"
	expectation.Connected = true
	expectation.Model = UUID("MODELID001")
	expectation.Interfaces = map[string]CPEInterface{}
	expectation.Interfaces["wlan0"] = i
	expectation.ConfigStatus = ConfigurationStatus{StatusOK{}}

	var result CPE
	var err = json.Unmarshal(fixture, &result)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, expectation, result)
}
