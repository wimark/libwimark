package libwimark

import (
	"encoding/json"
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"
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

type ClientStat struct {
	AcctStatusType      ClientStatPacketType `json:"Acct-Status-Type"`
	CPE                 UUID                 `json:"cpe_id"`
	WLAN                UUID                 `json:"wlan_id"`
	RadioId             string               `json:"radio_id"`
	CallingStationId    string               `json:"Calling-Station-Id"`
	UserName            string               `json:"User-Name"`
	AcctDelayTime       int                  `json:"Acct-Delay-Time"`
	AcctSessionId       string               `json:"Acct-Session-Id"`
	AcctInputGigawords  *int                 `json:"Acct-Input-Gigawords"`
	CalledStationId     string               `json:"Called-Station-Id"`
	AcctOutputGigawords *int                 `json:"Acct-Output-Gigawords"`
	AcctOutputOctets    *int                 `json:"Acct-Output-Octets"`
	AcctInputOctets     *int                 `json:"Acct-Input-Octets"`
	AcctSessionTime     *int                 `json:"Acct-Session-Time"`
	AcctInputPackets    *int                 `json:"Acct-Input-Packets"`
	AcctOutputPackets   *int                 `json:"Acct-Output-Packets"`
	Timestamp           int                  `json:"Timestamp"`

	// newly added fields
	Tx_ht     bool   `json:"tx_ht"`
	Rx_ht     bool   `json:"rx_ht"`
	Noise     int    `json:"noise, omitempty"`
	Tx_rate   int    `json:"tx_rate, omitempty"`
	Inactive  int    `json:"inactive"`
	Rx_rate   int    `json:"rx_rate, omitempty"`
	Rx_vht    bool   `json:"rx_vht"`
	Rx_mhz    bool   `json:"rx_mhz, omitempty"`
	Tx_mhz    bool   `json:"tx_mhz, omitempty"`
	Tx_vht    bool   `json:"tx_vht"`
	Signal    int    `json:"signal, omitempty"`
	Frequency string `json:"frecuency, omitempty"`
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
	Name string `json:"name"`
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
	ID      string `json:"id" bson:"_id"`
	MAC     string `json:"mac" bson: "mac"`
	WLAN    string `json:"wlan_id" bson:"wlan_id"`
	CPE     string `json:"cpe_id" bson:"cpe_id"`
	Radio   string `json:"radio" bson:"radio"`
	Channel string `json:"channel" bson:"channel"`
	// wrong spelling
	Frequence string             `json:"freq" bson:"freq"`
	Noise     int                `json:"noise" bson:"noise"`
	RSSI      int                `json:"rssi" bson:"rssi"`
	Inactive  int                `json:"inactive" bson:"inactive"`
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

	Type  WirelessClientType
	State WirelessClientState
	Data  interface{} `bson:"data" json:"data"`

	WLAN     string             `json:"wlan_id" bson:"wlan_id"`
	CPE      string             `json:"cpe_id" bson:"cpe_id"`
	Radio    string             `json:"radio_id" bson:"radio_id"`
	Freq     string             `json:"freq"`
	Channel  string             `json:"channel"`
	Rssi     int                `json:"rssi"`
	Noise    int                `json:"noise"`
	Mode     ConnectionModeType `json:"mode"`
	Inactive int                `json:"inactive"`

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
	MAC   string `json:"mac" baon:"mac"`
	WLAN  string `json:"wlan_id" bson:"wlan_id"`
	CPE   string `json:"cpe_id" bson:"cpe_id"`
	Radio string `json:"radio_id" bson:"radio_id"`
	Start int64  `json:"start" bson:"start"`
	Stop  int64  `json:"stop" bson:"stop"`
}
