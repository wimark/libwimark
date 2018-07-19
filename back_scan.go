package libwimark

import (
	"regexp"
	"strconv"
)

type BackScanModelRaw struct {
	SSID        string          `json:"ssid" bson:"ssid"`
	BSSID       string          `json:"bssid" bson:"bssid"`
	Signal      float32         `json:"signal" bson:"signal"`
	Channel     int             `json:"channel" bson:"channel"`
	Quality     float32         `json:"quality" bson:"quality"`
	QualityMax  int             `json:"quality_max" bson:"quality_max"`
	Mode        string          `json:"mode" bson:"mode"`
	CenterIndex int             `json:"center_idx0" bson:"center_idx0"`
	ChanOffset  int             `json:"sec_channel_offset" bson:"sec_channel_offset"`
	HTModes     map[string]bool `json:"htmodelist" bson:"htmodelist"`
	Encryption  struct {
		Enabled      bool     `json:"enabled" bson:"enabled"`
		AuthAlgs     []string `json:"auth_algs" bson:"auth_algs"`
		Description  string   `json:"description" bson:"description"`
		WEP          bool     `json:"wep" bson:"wep"`
		WPA          int      `json:"wpa" bson:"wpa"`
		AuthSuites   []string `json:"auth_suites" bson:"auth_suites"`
		PairCiphers  []string `json:"pair_ciphers" bson:"pair_ciphers"`
		GroupCiphers []string `json:"group_ciphers" bson:"group_ciphers"`
	} `json:"encryption" bson:"encryption"`
}

type CPEScanData struct {
	CPE       UUID                        `json:"cpe" bson:"cpe"`
	Radio     string                      `json:"radio" bson:"radio"`
	Timestamp int64                       `json:"timestamp" bson:"timestamp"`
	Scanlist  map[string]BackScanModelRaw `json:"scanlist" bson:"scanlist"`
}

type RRMChannel struct {
	Channel int `json:"channel"`
	Mode    int `json:"mode"`   // 0-legacy, 1-HT, 2-VHT
	Width   int `json:"width"`  // 20, 40, 80, 160 (no 80+80 support yet)
	Offset  int `json:"offset"` // -1, 0, 1 - for HT only
	Central int `json:"center"`
}

func RRMChannelFromState(state WiFiState) RRMChannel {
	var res RRMChannel
	res.Width, res.Mode, res.Offset = ParseBandmode(state.Bandwidth)
	res.Channel, _ = strconv.Atoi(state.Channel)
	if res.Offset == 0 && res.Width == 40 {
		if res.Channel <= 9 {
			res.Offset = 1
		} else if res.Channel < 36 {
			res.Offset = -1
		} else if res.Channel <= 144 {
			res.Offset = ((res.Channel-32)%4)/2 - 1
		} else if res.Channel <= 165 {
			res.Offset = ((res.Channel-145)%4)/2 - 1
		}
	}
	res.Central = CalcCentralChannel(res)
	return res
}
func RRMChannelFromScan(scan BackScanModelRaw) RRMChannel {
	var res = RRMChannel{
		Central: scan.Channel,
		Channel: scan.Channel,
		Width:   20,
		Offset:  scan.ChanOffset,
	}
	for htm, present := range scan.HTModes {
		if !present {
			continue
		}
		w, m, o := ParseBandmode(htm)
		if m > res.Mode || w > res.Width {
			res.Mode = m
			res.Width = w
			if o != 0 {
				res.Offset = o
			}
		}
	}
	if res.Mode <= 1 {
		res.Central = CalcCentralChannel(res)
	}

	return res
}
func RRMChannelToReq(settings RRMChannel) CPERRMChannelParams {
	var mhz = settings.Channel * 5
	var cmhz = settings.Central * 5
	if settings.Channel >= 36 {
		mhz += 5000
		cmhz += 5000
	} else {
		mhz += 2407
		cmhz += 2407
	}
	return CPERRMChannelParams{
		Count:       5, // oh, i dunno
		Freq:        mhz,
		Offset:      settings.Offset,
		CenterFreq1: cmhz,
		CenterFreq2: 0,
		Bandwidth:   settings.Width,
		HT:          strconv.FormatBool(settings.Mode > 0),
		VHT:         strconv.FormatBool(settings.Mode == 2),
	}
}

func AvailableChannels(state WiFiState, cfg WiFiConfig, caps WifiCapabilities) []RRMChannel {
	var settings = RRMChannelFromState(state)
	var res []RRMChannel
	var chans []int
	for _, capchan := range caps.Channels {
		if capchan.Restricted {
			continue
		}
		if len(cfg.Channels) == 0 {
			chans = append(chans, capchan.Channel)
		} else {
			for _, cfgchan := range cfg.Channels {
				if cfgchan == capchan.Channel {
					chans = append(chans, capchan.Channel)
				}
			}
		}
	}
	for _, ch := range chans {
		var newset = settings
		newset.Channel = ch
		newset.Central = CalcCentralChannel(newset)
		if newset.Central >= 50 && newset.Central <= 144 {
			// exclude DFS channels for now
			continue
		}
		res = append(res, newset)
	}
	return res
}

func ParseBandmode(bandmode string) (width int, mode int, offset int) {
	var re = regexp.MustCompile("(V?HT)?([0-9]+)?([\\+-])?")
	var modes = re.FindStringSubmatch(bandmode)
	if len(modes) < 3 {
		return 20, 0, 0
	}
	width, _ = strconv.Atoi(modes[2])
	if width == 0 {
		width = 20
	}
	switch modes[1] {
	case "VHT":
		mode = 2
	case "HT":
		mode = 1
	default:
		mode = 0
	}
	switch modes[2] {
	case "+":
		offset = 1
	case "-":
		offset = -1
	default:
		offset = 0
	}
	return
}

var Chan40 = []int{38, 46, 54, 62, 102, 110, 118, 126, 134, 142, 151, 159}
var Chan80 = []int{42, 58, 106, 122, 138, 155}
var Chan160 = []int{50, 114}

func GetChannelWidth(index int) int {
	for _, i := range Chan40 {
		if i == index {
			return 40
		}
	}
	for _, i := range Chan80 {
		if i == index {
			return 80
		}
	}
	for _, i := range Chan160 {
		if i == index {
			return 160
		}
	}
	return 20
}

func CalcCentralChannel(settings RRMChannel) int {
	if settings.Width == 20 || settings.Mode == 0 {
		return settings.Channel
	} else if settings.Width == 40 || settings.Mode == 1 {
		return settings.Channel + settings.Offset*2
	} else {
		var chans []int
		switch settings.Width {
		case 80:
			chans = Chan80
		case 160:
			chans = Chan160
		}
		// choose the closest
		var diff = 0
		var closest = 0
		for _, ch := range chans {
			var d = ch - settings.Channel
			if d < 0 {
				d = -d
			}
			if diff == 0 {
				diff = d
				closest = ch
			} else if d < diff {
				diff = d
				closest = ch
			} else {
				break
			}
		}
		return closest
	}
}
