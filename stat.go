package libwimark

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
)

type Stat struct {
	CPE       UUID    `json:"cpe_id"`
	Timestamp int64   `json:"timestamp"`
	CPU       float64 `json:"cpu"`
	RAM       struct {
		Total    float64 `json:"total"`
		Buffered float64 `json:"buffered"`
		Shared   float64 `json:"shared"`
		Free     float64 `json:"free"`
	} `json:"memory"`
	Uptime       int64   `json:"uptime"`
	Storage      float64 `json:"storage"`
	ProcActive   uint64  `json:"processes_active"`
	ProcSleeping uint64  `json:"processes_sleeping"`
	Interfaces   map[string]struct {
		Type string  `json:"type"`
		Tx   float64 `json:"tx"`
		Rx   float64 `json:"rx"`
	} `json:"interfaces"`
}

type ClientStatOld struct {
	AcctStatusType      ClientStatPacketType `json:"Acct-Status-Type"`
	CPE                 UUID                 `json:"cpe_id"`
	WLAN                UUID                 `json:"wlan_id"`
	RadioId             string               `json:"radio_id"`
	CallingStationId    string               `json:"Calling-Station-Id"`
	CalledStationId     string               `json:"Called-Station-Id"`
	UserName            string               `json:"User-Name"`
	AcctDelayTime       int                  `json:"Acct-Delay-Time"`
	AcctSessionId       string               `json:"Acct-Session-Id"`
	AcctInputGigawords  *int                 `json:"Acct-Input-Gigawords"`
	AcctOutputGigawords *int                 `json:"Acct-Output-Gigawords"`
	AcctOutputOctets    *int                 `json:"Acct-Output-Octets"`
	AcctInputOctets     *int                 `json:"Acct-Input-Octets"`
	AcctInputPackets    *int                 `json:"Acct-Input-Packets"`
	AcctOutputPackets   *int                 `json:"Acct-Output-Packets"`
	AcctSessionTime     *int                 `json:"Acct-Session-Time"`
	Timestamp           int                  `json:"Timestamp"`
	ConnectInfo         string               `json:"Connect-Info"`
	NasIPAddress        string               `json:"NAS-IP-Address"`
	NasPortType         string               `json:"NAS-Port-Type"`
	NasPort             string               `json:"NAS-Port"`

	// newly added fields
	Inactive  int  `json:"inactive"`
	Tx_ht     bool `json:"tx_ht"`
	Rx_ht     bool `json:"rx_ht"`
	Tx_rate   int  `json:"tx_rate, omitempty"`
	Rx_rate   int  `json:"rx_rate, omitempty"`
	Rx_vht    bool `json:"rx_vht"`
	Tx_vht    bool `json:"tx_vht"`
	Rx_mhz    int  `json:"rx_mhz, omitempty"`
	Tx_mhz    int  `json:"tx_mhz, omitempty"`
	Signal    int  `json:"signal, omitempty"`
	Noise     int  `json:"noise, omitempty"`
	Frequency int  `json:"frequency, omitempty"`
}

type AccountingData struct {
	RxBytes   int64 `json:"rx_bytes" `
	TxBytes   int64 `json:"tx_bytes"`
	RxPackets int64 `json:"rx_packets"`
	TxPackets int64 `json:"tx_packets"`
}

type AccountingRadio struct {
	Signal   int `json:"signal"`
	Noise    int `json:"noise"`
	Inactive int `json:"inactive"`

	RxMhz  int  `json:"rx_mhz"`
	TxMhz  int  `json:"tx_mhz"`
	TxRate int  `json:"tx_rate"`
	RxRate int  `json:"rx_rate"`
	TxHt   bool `json:"tx_ht"`
	RxHt   bool `json:"rx_ht"`
	RxVht  bool `json:"rx_vht"`
	TxVht  bool `json:"tx_vht"`

	TxMcs     int  `json:"tx_mcs"`
	Tx40Mhz   bool `json:"tx_40mhz"`
	TxShortGi bool `json:"tx_short_gi"`
}

type ClientStat struct {
	Type ClientStatPacketType `json:"type"`

	MAC       string `json:"macaddr"`
	SSID      string `json:"ssid"`
	WLAN      UUID   `json:"wlan_id"`
	CPE       UUID   `json:"cpe_id"`
	RadioId   string `json:"radio_id"`
	BSSID     string `json:"bssid"`
	Frequency int    `json:"frequency, omitempty"`

	Username    string `json:"identity"`
	SessionId   string `json:"session_id"`
	SessionTime int    `json:"session_time"`

	Timestamp  int64           `json:"timestamp"`
	Accounting AccountingData  `json:"accounting"`
	Radio      AccountingRadio `json:"rf"`
}

type CameraClientData struct {
	Rtsp_stream string `json:"rtsp_stream"`
	Description string `json:"description"`
}

type OtherClientData struct {
	Description string `json:"description"`
}

type CPEPollSettings struct {
	Rules []UUID `json:"rules"`
}

type StatEventRule struct {
	StatEventRuleObject
	Name       string `json:"name"`
	PostScript string `json:"post_script"`
}

func (self *StatEventRule) MarshalJSON() ([]byte, error) {
	var b []byte
	{
		var err error
		b, err = json.Marshal(self.StatEventRuleObject)
		if err != nil {
			return nil, err
		}
	}

	var doc Document
	{
		var err error
		err = json.Unmarshal(b, &doc)
		if err != nil {
			return nil, err
		}
	}

	doc["name"] = self.Name
	doc["post_script"] = self.PostScript

	return json.Marshal(doc)
}

func (self *StatEventRule) UnmarshalJSON(b []byte) error {
	var doc map[string]json.RawMessage
	var err = json.Unmarshal(b, &doc)
	if err != nil {
		return err
	}

	if doc == nil {
		return nil
	}

	// may be there is a way how to make code more reusable
	// don't have time to figure out

	var nameRaw, nameExists = doc["name"]
	if nameExists {
		var name string
		var tsErr = json.Unmarshal(nameRaw, &name)
		if tsErr == nil {
			self.Name = name
		} else {
			return tsErr
		}
	}

	delete(doc, "name")

	var postScriptRaw, postScriptExists = doc["post_script"]
	if postScriptExists {
		var postScript string
		var tsErr = json.Unmarshal(postScriptRaw, &postScript)
		if tsErr == nil {
			self.PostScript = postScript
		} else {
			return tsErr
		}
	}

	delete(doc, "post_script")

	var v, _ = json.Marshal(doc)

	return self.StatEventRuleObject.UnmarshalJSON(v)
}

func (self *StatEventRule) GetBSON() (interface{}, error) {
	var out = bson.M{}

	var obj_b []byte
	{
		var err error
		obj_b, err = bson.Marshal(self.StatEventRuleObject)
		if err != nil {
			return nil, err
		}
	}

	var obj bson.M
	{
		var err error
		err = bson.Unmarshal(obj_b, &obj)
		if err != nil {
			return nil, err
		}
	}
	out = obj
	out["name"] = self.Name
	out["post_script"] = self.PostScript

	return out, nil
}

func (self *StatEventRule) SetBSON(raw bson.Raw) error {
	var in = bson.M{}
	{
		var err error
		err = raw.Unmarshal(&in)
		if err != nil {
			return err
		}
	}

	//name
	var v_name, k_found = in["name"]

	if !k_found {
		return errors.New("No name found")
	}

	self.Name = v_name.(string)
	delete(in, "name")

	//post_script
	var ps_script, ps_found = in["post_script"]

	if ps_found {
		self.PostScript = ps_script.(string)
		delete(in, "post_script")
	}

	var err error
	obj_b, err := bson.Marshal(in)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(obj_b, &self.StatEventRuleObject)
	if err != nil {
		return err
	}

	return nil
}

type LBSClientData struct {
	Timestamp int64   `json:"timestamp"`
	CPE       UUID    `json:"cpe"`
	Radio     string  `json:"radio"`
	ClientMac string  `json:"client_mac"`
	RSSI      float64 `json:"rssi"`
	Frequency int     `json:"freq"`
}

type LBSCPEInfo struct {
	Group UUID    `json:"group"`
	CPE   UUID    `json:"cpe"`
	Name  string  `json:"name"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Z     float64 `json:"z"`
}

type LBSClientCoords struct {
	Timestamp int64   `json:"timestamp"`
	Group     UUID    `json:"group"`
	Mac       string  `json:"mac"`
	BestCPE   UUID    `json:"bestcpe"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Z         float64 `json:"z"`
}

type CPEStatInfo struct {
	ID               string    `json:"id" bson:"_id"`
	CPE              string    `json:"cpe_id" bson:"cpe_id"`
	Timestamp        time.Time `json:"timestamp" bson:"timestamp"`
	CPULoad          float64   `json:"cpu_load" bson:"cpu_load"`
	MemoryFree       int       `json:"memory_free" bson:"memory_free"`
	MemoryTotal      int       `json:"memory_total" bson:"memory_total"`
	TotalRxBytes     int64     `json:"total_rx_bytes" bson:"total_rx_bytes"`
	TotalTxBytes     int64     `json:"total_tx_bytes" bson:"total_tx_bytes"`
	LastRxBytes      int64     `json:"last_rx_bytes" bson:"last_rx_bytes"`
	LastTxBytes      int64     `json:"last_tx_bytes" bson:"last_tx_bytes"`
	DeltaTxBytes     int64     `json:"delta_tx_bytes" bson:"delta_tx_bytes"`
	DeltaRxBytes     int64     `json:"delta_rx_bytes" bson:"delta_rx_bytes"`
	ConnectedClients []string  `json:"connected_clients" bson:"connected_clients"`
}

type WLANStatInfo struct {
	ID               string   `json:"id" bson:"_id"`
	WLAN             string   `json:"wlan_id" bson:"wlan_id"`
	SSID             string   `json:"ssid" bson:"ssid"`
	Timestamp        int64    `json:"timestamp" bson:"timestamp"`
	TotalRxBytes     int64    `json:"total_rx_bytes" bson:"total_rx_bytes"`
	TotalTxBytes     int64    `json:"total_tx_bytes" bson:"total_tx_bytes"`
	DeltaTxBytes     int64    `json:"delta_tx_bytes" bson:"delta_tx_bytes"`
	DeltaRxBytes     int64    `json:"delta_rx_bytes" bson:"delta_rx_bytes"`
	ConnectedClients []string `json:"connected_clients" bson:"connected_clients"`
}

type BSSStatInfo struct {
	ID               string   `json:"id" bson:"_id"`
	WLAN             string   `json:"wlan_id" bson:"wlan_id"`
	SSID             string   `json:"ssid" bson:"ssid"`
	CPE              string   `json:"cpe_id" bson:"cpe_id"`
	Radio            string   `json:"radio" bson:"radio"`
	Timestamp        int64    `json:"timestamp" bson:"timestamp"`
	TotalRxBytes     int64    `json:"total_rx_bytes" bson:"total_rx_bytes"`
	TotalTxBytes     int64    `json:"total_tx_bytes" bson:"total_tx_bytes"`
	DeltaTxBytes     int64    `json:"delta_tx_bytes" bson:"delta_tx_bytes"`
	DeltaRxBytes     int64    `json:"delta_rx_bytes" bson:"delta_rx_bytes"`
	ConnectedClients []string `json:"connected_clients" bson:"connected_clients"`
}

type ClientStatInfo struct {
	ID    string `json:"id" bson:"_id"`
	MAC   string `json:"mac" bson: "mac"`
	WLAN  string `json:"wlan_id" bson:"wlan_id"`
	SSID  string `json:"ssid" bson:"ssid"`
	CPE   string `json:"cpe_id" bson:"cpe_id"`
	Radio string `json:"radio" bson:"radio"`

	Channel   string             `json:"channel" bson:"channel"`
	Frequence string             `json:"freq" bson:"freq"`
	Noise     int                `json:"noise" bson:"noise"`
	RSSI      int                `json:"rssi" bson:"rssi"`
	Mode      ConnectionModeType `json:"mode" bson:"mode"`

	Timestamp    int64 `json:"timestamp" bson:"timestamp"`
	TotalRxBytes int64 `json:"total_rx_bytes" bson:"total_rx_bytes"`
	TotalTxBytes int64 `json:"total_tx_bytes" bson:"total_tx_bytes"`
	DeltaTxBytes int64 `json:"delta_tx_bytes" bson:"delta_tx_bytes"`
	DeltaRxBytes int64 `json:"delta_rx_bytes" bson:"delta_rx_bytes"`
	LastRxBytes  int64 `json:"last_rx_bytes" bson:"last_rx_bytes"`
	LastTxBytes  int64 `json:"last_tx_bytes" bson:"last_tx_bytes"`
}

type WirelessClient struct {
	MAC       string `json:"mac" bson:"_id"`
	Timestamp int64  `json:"timestamp"`

	Type         WirelessClientType
	State        WirelessClientState
	Data         interface{} `bson:"data" json:"data"`
	Manufacturer string      `json:"manufacturer" bson:"manufacturer"`

	WLAN    string             `json:"wlan_id" bson:"wlan_id"`
	SSID    string             `json:"wlan_ssid" bson:"wlan_ssid"`
	CPE     string             `json:"cpe_id" bson:"cpe_id"`
	Radio   string             `json:"radio_id" bson:"radio_id"`
	Freq    string             `json:"freq"`
	Channel string             `json:"channel"`
	Rssi    int                `json:"rssi"`
	Noise   int                `json:"noise"`
	Mode    ConnectionModeType `json:"mode"`

	InPackets  int64 `json:"in_packets" bson:"in_packets"`
	OutPackets int64 `json:"out_packets" bson:"out_packets"`
	InKBytes   int64 `json:"in_kbytes" bson:"in_kbytes"`
	OutKBytes  int64 `json:"out_kbytes" bson:"out_kbytes"`

	Ip string `json:"ip"`

	FirstConnect int64 `json:"first_connect" bson:"first_connect"`
	LastConnect  int64 `json:"last_connect" bson:"last_connect"`
}

func (wc *WirelessClient) GetSpecificWCInfo() *WirelessClientObject {
	wco := &WirelessClientObject{Data: wc.Data, Type: wc.Type}
	return wco
}

func (wc *WirelessClient) SetSpecificWCInfo(wco *WirelessClientObject) {
	wc.Type = wco.Type
	wc.Data = wco.Data
}

type CPESessionInfo struct {
	ID    string `json:"id" bson:"_id"`
	CPE   string `json:"cpe_id" bson:"cpe_id"`
	Start int64  `json:"start" bson:"start"`
	Stop  int64  `json:"stop" bson:"stop"`
}

type ClientSessionInfo struct {
	ID    string `json:"id" bson:"_id"`
	MAC   string `json:"mac" bson:"mac"`
	WLAN  string `json:"wlan_id" bson:"wlan_id"`
	SSID  string `json:"ssid" bson:"ssid"`
	CPE   string `json:"cpe_id" bson:"cpe_id"`
	Radio string `json:"radio_id" bson:"radio_id"`
	Freq  string `json:"freq" bson:"freq"`
	Mode  string `json:"mode" bson:"mode"`

	Start int64 `json:"start" bson:"start"`
	Stop  int64 `json:"stop" bson:"stop"`

	StartNoise int `json:"start_noise" bson:"start_noise"`
	StopNoise  int `json:"stop_noise" bson:"stop_noise"`

	StartRSSI int `json:"start_rssi" bson:"start_rssi"`
	StopRSSI  int `json:"stop_rssi" bson:"stop_rssi"`

	RxBytes int64 `json:"rx_bytes" bson:"rx_bytes"`
	TxBytes int64 `json:"tx_bytes" bson:"tx_bytes"`
}

// for new storing lbs probe data collection
type ClientAssocData struct {
	WLAN    string `json:"wlan_id" bson:"wlan_id"`
	SSID    string `json:"wlan_ssid" bson:"wlan_ssid"`
	Desc    string `json:"wlan_desc" bson:"wlan_desc"`
	CPE     string `json:"cpe_id" bson:"cpe_id"`
	Radio   string `json:"radio_id" bson:"radio_id"`
	Freq    string `json:"freq" bson:"freq"`
	Channel string `json:"channel" bson:"channel"`
}

type ClientProbeData struct {
	Id           string          `json:"id" bson:"_id"`
	Timestamp    int64           `json:"timestamp" bson:"timestamp"`
	MAC          string          `json:"mac" bson:"mac"`
	CPE          string          `json:"cpe" bson:"cpe"`
	CPEName      string          `json:"cpe_name" bson:"cpe_name"`
	CPEMAC       string          `json:"cpe_mac" bson:"cpe_mac"`
	Radio        string          `json:"radio" bson:"radio"`
	Frequency    int             `json:"freq" bson:"freq"`
	RSSI         int             `json:"rssi" bson:"rssi"`
	Manufacturer string          `json:"manufacturer" bson:"manufacturer"`
	AssocData    ClientAssocData `json:"assoc_data,omitempty" bson:"assoc_data"`
}

type StationDumpData struct {
	Radio     string                     `json:"radio" bson:"radio"`
	AssocList map[string]AccountingRadio `json:"assoclist" bson:"assoclist"`
}
