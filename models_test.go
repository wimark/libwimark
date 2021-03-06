package libwimark

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	var fixture1 = SecuritySuiteTKIP
	var fixture2 = SecuritySuiteTKIP
	var fixture3 = SecuritySuiteAES

	assert.Equal(t, fixture1, fixture2)
	assert.NotEqual(t, fixture1, fixture3)

	assert.True(t, fixture1 == fixture2)
	assert.True(t, fixture1 != fixture3)
}

func TestSecuritySuite(t *testing.T) {
	var fixture = []byte(`["tkip"]`)
	var expectation = SecuritySuiteTKIP
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
	d.Suites = []SecuritySuite{SecuritySuiteTKIP}
	d.PSK = "qwerty"
	var expectation = EnumSecurity{Type: SecurityTypeWPA2Personal, Data: d}
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
	d.Suites = []SecuritySuite{SecuritySuiteAES}
	d.PSK = "qwerty"
	var expectation = WLAN{
		Name:        "myhotspot",
		SSID:        "qwertyasdfgh",
		Description: "This is my pretty little honeypot",
		Security: EnumSecurity{
			Type: SecurityTypeWPA2Personal,
			Data: d,
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

// func TestCPE(t *testing.T) {
// 	var fixture = []byte(`
//         {
//             "name": "mycpe",
//             "description": "this is my CPE",
//             "connected": true,
//             "model": {
//                 "id": "MODELID001",
//                 "name": "model 001"
//             },
//             "interfaces": {
//                 "wlan0": {
//                     "addr": "macaddr0",
//                     "capabilities":{
//                         "txpower_offset":0,
//                         "txpwrlist":[
//                             {
//                                 "dbm":0,
//                                 "mw":1
//                             },
//                             {
//                                 "dbm":20,
//                                 "mw":100
//                             }
//                         ],
//                         "htmodelist":{
//                             "HT20":false,
//                             "HT40":true,
//                             "VHT20":true,
//                             "VHT40":true,
//                             "VHT80":true,
//                             "VHT160":false
//                         },
//                         "freqlist":[
//                             {
//                                 "restricted":false,
//                                 "mhz":2412,
//                                 "channel":1,
//                                 "max_txpower":{
//                                     "dbm":20,
//                                     "mw":100
//                                 }
//                             },
//                             {
//                                 "restricted":true,
//                                 "mhz":5825,
//                                 "channel":165,
//                                 "max_txpower":{
//                                     "dbm":18,
//                                     "mw":63
//                                 }
//                             }
//                         ],
//                         "hwmodelist":{
//                             "a":false,
//                             "b":true,
//                             "ac":true,
//                             "g":true,
//                             "n":true
//                         }
//                     },
//                     "type": "wifi",
//                     "data": {
//                         "name": "WLAN Interface 0",
//                         "config": {
//                             "bandmode": "n",
//                             "txpower": "100",
//                             "channels": [1, 3, 5],
//                             "wlans": ["WLAN001", "WLAN002"]
//                         },
//                         "state": {
//                             "frequency": "2.4",
//                             "bandmode": "n",
//                             "bandwidth": "20",
//                             "channel": "5",
//                             "txpower": "100",
//                             "enabled": true
//                         }
//                     }
//                 }
//             },
//             "config_status": "ok"
//         }
//     `)
// 	var d = &WiFiData{}
// 	d.Name = "WLAN Interface 0"

// 	d.State.Frequency = "2.4"
// 	d.State.BandMode = "n"
// 	d.State.Bandwidth = "20"
// 	d.State.Channel = "5"
// 	d.State.TxPower = "100"
// 	d.State.Enabled = true

// 	d.Config.BandMode = "n"
// 	d.Config.TxPower = "100"
// 	d.Config.Channels = []int{1, 3, 5}
// 	d.Config.WLANs = []UUID{"WLAN001", "WLAN002"}

// 	var i CPEInterface
// 	i.Addr = "macaddr0"
// 	i.Type = CPEInterfaceTypeWiFi
// 	i.Data = d

// 	var caps Capabilities
// 	caps.Channels = []CapChannel{
// 		CapChannel{
// 			Restricted: false,
// 			Freq:       2412,
// 			Channel:    1,
// 			MaxPower:   CapTxPower{DBelMw: 20, MilliWatt: 100},
// 		},
// 		CapChannel{
// 			Restricted: true,
// 			Freq:       5825,
// 			Channel:    165,
// 			MaxPower:   CapTxPower{DBelMw: 18, MilliWatt: 63},
// 		},
// 	}
// 	caps.HTModes = map[string]bool{
// 		//		BandwidthTypeHT20:   false,
// 		//		BandwidthTypeHT40:   true,
// 		//		BandwidthTypeVHT20:  true,
// 		//		BandwidthTypeVHT40:  true,
// 		//		BandwidthTypeVHT80:  true,
// 		//		BandwidthTypeVHT160: false,
// 		"HT20":   false,
// 		"HT40":   true,
// 		"VHT20":  true,
// 		"VHT40":  true,
// 		"VHT80":  true,
// 		"VHT160": false,
// 	}
// 	caps.HWModes = map[string]bool{
// 		"a":  false,
// 		"b":  true,
// 		"ac": true,
// 		"g":  true,
// 		"n":  true,
// 	}
// 	caps.TxOffset = 0
// 	caps.TxPowers = []CapTxPower{
// 		CapTxPower{DBelMw: 0, MilliWatt: 1},
// 		CapTxPower{DBelMw: 20, MilliWatt: 100},
// 	}
// 	i.Capabilities = caps

// 	var expectation CPE
// 	expectation.Name = "mycpe"
// 	expectation.Description = "this is my CPE"
// 	expectation.Connected = true
// 	expectation.Model.Id = UUID("MODELID001")
// 	expectation.Model.Name = "model 001"
// 	expectation.Interfaces = map[string]CPEInterface{}
// 	expectation.Interfaces["wlan0"] = i
// 	expectation.ConfigStatus = ConfigurationStatusOK

// 	var result CPE
// 	var err = json.Unmarshal(fixture, &result)

// 	if err != nil {
// 		panic(err)
// 	}

// 	assert.Equal(t, expectation, result)
// }
