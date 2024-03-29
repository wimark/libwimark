//go:generate easyjson -all stat.go
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
		Type    string  `json:"type"`
		Uptime  int64   `json:"uptime"`
		Tx      float64 `json:"tx"`
		Rx      float64 `json:"rx"`
		TxSpeed float64 `json:"tx_speed"`
		RxSpeed float64 `json:"rx_speed"`
		TxDelta float64 `json:"tx_delta"`
		RxDelta float64 `json:"rx_delta"`
	} `json:"interfaces"`
	Eth0State struct {
		Carrier       int16  `json:"carrier"`
		CarrierChange int16  `json:"carrier_changes"`
		Speed         int16  `json:"speed"`
		OperState     string `json:"operstate"`
		Duplex        string `json:"duplex"`
	} `json:"eth0_state"`
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
	Tx_rate   int  `json:"tx_rate,omitempty"`
	Rx_rate   int  `json:"rx_rate,omitempty"`
	Rx_vht    bool `json:"rx_vht"`
	Tx_vht    bool `json:"tx_vht"`
	Rx_mhz    int  `json:"rx_mhz,omitempty"`
	Tx_mhz    int  `json:"tx_mhz,omitempty"`
	Signal    int  `json:"signal,omitempty"`
	Noise     int  `json:"noise,omitempty"`
	Frequency int  `json:"frequency,omitempty"`
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
	RxRate int  `json:"rx_rate"`
	TxRate int  `json:"tx_rate"`
	RxHt   bool `json:"rx_ht"`
	TxHt   bool `json:"tx_ht"`
	RxVht  bool `json:"rx_vht"`
	TxVht  bool `json:"tx_vht"`
	RxHE   bool `json:"rx_he"`
	TxHE   bool `json:"tx_he"`

	TxMcs     int  `json:"tx_mcs"`
	RxMcs     int  `json:"rx_mcs"`
	TxNSS     int  `json:"tx_nss"`
	RxNSS     int  `json:"rx_nss"`
	Tx40Mhz   bool `json:"tx_40mhz"`
	Rx40Mhz   bool `json:"rx_40mhz"`
	TxShortGi bool `json:"tx_short_gi"`
	RxShortGi bool `json:"rx_short_gi"`

	ExpThroughput int `json:"expected_throughput"`
}

type AccountingDPI struct {
	DestIPs []string `json:"dest_ips" bson:"dest_ips"`
}

type ClientStat struct {
	Type ClientStatPacketType `json:"type"`

	MAC       string `json:"macaddr"`
	SSID      string `json:"ssid"`
	WLAN      UUID   `json:"wlan_id"`
	CPE       UUID   `json:"cpe_id"`
	RadioId   string `json:"radio_id"`
	BSSID     string `json:"bssid"`
	Frequency int    `json:"frequency,omitempty"`

	Username    string `json:"identity"`
	SessionId   string `json:"session_id"`
	SessionTime int    `json:"session_time"`

	Timestamp  int64           `json:"timestamp"`
	Accounting AccountingData  `json:"accounting"`
	Radio      AccountingRadio `json:"rf"`

	DPI AccountingDPI `json:"dpi"`
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
	Name           string      `json:"name"`
	PostScript     string      `json:"post_script"`
	NotifyType     NotifyType  `json:"notify_type"`
	NotifySettings interface{} `json:"notify_settings"`
}

func (eventRule *StatEventRule) MarshalJSON() ([]byte, error) {
	var b []byte
	{
		var err error
		b, err = json.Marshal(eventRule.StatEventRuleObject)
		if err != nil {
			return nil, err
		}
	}

	var doc Document
	{
		var err = json.Unmarshal(b, &doc)
		if err != nil {
			return nil, err
		}
	}

	doc["name"] = eventRule.Name
	doc["post_script"] = eventRule.PostScript
	doc["notify_type"] = eventRule.NotifyType
	doc["notify_settings"] = eventRule.NotifySettings

	return json.Marshal(doc)
}

func (eventRule *StatEventRule) UnmarshalJSON(b []byte) error {
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
			eventRule.Name = name
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
			eventRule.PostScript = postScript
		} else {
			return tsErr
		}
	}

	delete(doc, "post_script")

	var notifyTypeRaw, notifyTypeExists = doc["notify_type"]
	if notifyTypeExists {
		var notifyType NotifyType
		var tsErr = json.Unmarshal(notifyTypeRaw, &notifyType)
		if tsErr == nil {
			eventRule.NotifyType = notifyType
		} else {
			return tsErr
		}
	}

	delete(doc, "notify_type")

	var notifySettingsRaw, notifySettingsExists = doc["notify_settings"]
	if notifySettingsExists {
		var notifySettings interface{}
		var tsErr = json.Unmarshal(notifySettingsRaw, &notifySettings)
		if tsErr == nil {
			eventRule.NotifySettings = notifySettings
		} else {
			return tsErr
		}
	}

	delete(doc, "notify_settings")

	var v, _ = json.Marshal(doc)

	return eventRule.StatEventRuleObject.UnmarshalJSON(v)
}

func (eventRule *StatEventRule) GetBSON() (interface{}, error) {
	var out bson.M

	var obj_b []byte
	{
		var err error
		obj_b, err = bson.Marshal(eventRule.StatEventRuleObject)
		if err != nil {
			return nil, err
		}
	}

	var obj bson.M
	{
		var err = bson.Unmarshal(obj_b, &obj)
		if err != nil {
			return nil, err
		}
	}
	out = obj
	out["name"] = eventRule.Name
	out["post_script"] = eventRule.PostScript
	out["notify_type"] = eventRule.NotifyType
	out["notify_settings"] = eventRule.NotifySettings

	return out, nil
}

func (eventRule *StatEventRule) SetBSON(raw bson.Raw) error {
	var in = bson.M{}
	{
		var err = raw.Unmarshal(&in)
		if err != nil {
			return err
		}
	}

	//name
	var v_name, k_found = in["name"]
	if !k_found {
		return errors.New("no name found")
	}

	eventRule.Name = v_name.(string)
	delete(in, "name")

	//post_script
	var ps_script, ps_found = in["post_script"]
	if ps_found {
		eventRule.PostScript = ps_script.(string)
		delete(in, "post_script")
	}

	//notify_type
	var notify_type, notify_type_found = in["notify_type"]
	if notify_type_found {
		eventRule.NotifyType = NotifyType(notify_type.(string))
		delete(in, "notify_type")
	}

	//notify_settings
	var notify_settings, notify_settings_found = in["notify_settings"]
	if notify_settings_found {
		eventRule.NotifySettings = notify_settings
		delete(in, "notify_settings")
	}

	var err error
	obj_b, err := bson.Marshal(in)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(obj_b, &eventRule.StatEventRuleObject)
	if err != nil {
		return err
	}

	return nil
}

type LBSClientSignal struct {
	RSSI      int   `json:"rssi" bson:"rssi"`
	Timestamp int64 `json:"timestamp"`
}

type LBSClientData struct {
	Timestamp int64             `json:"timestamp"`
	CPE       UUID              `json:"cpe"`
	Radio     string            `json:"radio"`
	ClientMac string            `json:"client_mac"`
	RSSI      float64           `json:"rssi"`
	Frequency int               `json:"freq"`
	SSIDs     []string          `json:"ssids"`
	Signals   []LBSClientSignal `json:"signals"`
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
	Timestamp int64     `json:"timestamp"`
	Group     UUID      `json:"group"`
	Mac       string    `json:"mac"`
	BestCPE   UUID      `json:"bestcpe"`
	Zone      string    `json:"zone"`
	ExpireAt  time.Time `json:"expire_at"`
	X         float64   `json:"x"`
	Y         float64   `json:"y"`
	Z         float64   `json:"z"`
}

// CPEStatInfo struct for store CPE accumulated stats
type CPEStatInfo struct {
	ID           string    `json:"id" bson:"_id"`
	CPE          string    `json:"cpe_id" bson:"cpe_id"`
	Timestamp    time.Time `json:"timestamp" bson:"timestamp"` // as create_at
	CPULoad      float64   `json:"cpu_load" bson:"cpu_load"`
	MemoryFree   int       `json:"memory_free" bson:"memory_free"`
	MemoryTotal  int       `json:"memory_total" bson:"memory_total"`
	TotalRxBytes int64     `json:"total_rx_bytes" bson:"total_rx_bytes"`
	TotalTxBytes int64     `json:"total_tx_bytes" bson:"total_tx_bytes"`
	LastRxBytes  int64     `json:"last_rx_bytes" bson:"last_rx_bytes"`
	LastTxBytes  int64     `json:"last_tx_bytes" bson:"last_tx_bytes"`
	DeltaTxBytes int64     `json:"delta_tx_bytes" bson:"delta_tx_bytes"`
	DeltaRxBytes int64     `json:"delta_rx_bytes" bson:"delta_rx_bytes"`
	TxSpeed      float64   `json:"tx_speed" bson:"tx_speed"`
	RxSpeed      float64   `json:"rx_speed" bson:"rx_speed"`
	Uptime       int64     `json:"uptime" bson:"uptime"`

	ConnectedClients []string `json:"connected_clients" bson:"connected_clients"`

	// // new clients
	// Clients map[string]bool `json:"-" bson:"-"`
	// // hits to count average cpu / mem load
	// Hits int `json:"-" bson:"-"`
}

// WLANStatInfo struct for store WLAN accumulated stats
type WLANStatInfo struct {
	ID        string `json:"id" bson:"_id"`
	WLAN      string `json:"wlan_id" bson:"wlan_id"`
	SSID      string `json:"ssid" bson:"ssid"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`

	TotalRxBytes     int64    `json:"total_rx_bytes" bson:"total_rx_bytes"`
	TotalTxBytes     int64    `json:"total_tx_bytes" bson:"total_tx_bytes"`
	DeltaTxBytes     int64    `json:"delta_tx_bytes" bson:"delta_tx_bytes"`
	DeltaRxBytes     int64    `json:"delta_rx_bytes" bson:"delta_rx_bytes"`
	ConnectedClients []string `json:"connected_clients" bson:"connected_clients"`

	CreateAt time.Time `json:"create_at" bson:"create_at"`

	// new clients
	// Clients map[string]bool `json:"-" bson:"-"`
}

// BSSStatInfo struct for store BSS (CPE+WLAN) accumulated stats
type BSSStatInfo struct {
	ID        string `json:"id" bson:"_id"`
	WLAN      string `json:"wlan_id" bson:"wlan_id"`
	SSID      string `json:"ssid" bson:"ssid"`
	CPE       string `json:"cpe_id" bson:"cpe_id"`
	Radio     string `json:"radio" bson:"radio"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`

	TotalRxBytes     int64    `json:"total_rx_bytes" bson:"total_rx_bytes"`
	TotalTxBytes     int64    `json:"total_tx_bytes" bson:"total_tx_bytes"`
	DeltaTxBytes     int64    `json:"delta_tx_bytes" bson:"delta_tx_bytes"`
	DeltaRxBytes     int64    `json:"delta_rx_bytes" bson:"delta_rx_bytes"`
	ConnectedClients []string `json:"connected_clients" bson:"connected_clients"`

	CreateAt time.Time `json:"create_at" bson:"create_at"`

	// new clients
	// Clients map[string]bool `json:"-" bson:"-"`
}

type ClientStatInfo struct {
	ID    string `json:"id" bson:"_id"`
	MAC   string `json:"mac" bson:"mac"`
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

	CreateAt time.Time `json:"create_at" bson:"create_at"`
}

type WirelessClientUpdate struct {
	MAC     string `json:"mac" bson:"_id"`
	MacAddr string `json:"mac_addr" bson:"mac_addr"`
	CPE     string `json:"cpe_id" bson:"cpe_id"`

	Type WirelessClientType
	Data interface{} `bson:"data" json:"data"`
}

type WirelessClient struct {
	MAC       string `json:"mac" bson:"_id"`
	Timestamp int64  `json:"timestamp"`

	// general client data
	Type         WirelessClientType
	Data         interface{} `bson:"data" json:"data"`
	Manufacturer string      `json:"manufacturer" bson:"manufacturer"`
	MacAddr      string      `json:"mac_addr" bson:"mac_addr"`

	// last connection data
	State   WirelessClientState
	WLAN    string             `json:"wlan_id" bson:"wlan_id"`
	SSID    string             `json:"wlan_ssid" bson:"wlan_ssid"`
	CPE     string             `json:"cpe_id" bson:"cpe_id"`
	Radio   string             `json:"radio_id" bson:"radio_id"`
	Freq    string             `json:"freq"`
	Channel string             `json:"channel"`
	Mode    ConnectionModeType `json:"mode"`
	Ip      string             `json:"ip"`

	// radio state
	Rssi  int `json:"rssi"`
	Noise int `json:"noise"`
	SNR   int `json:"snr" bson:"snr"`
	// heath score as Cisco DNA assurance - connected score (1 || 4) + best over 5 minute SNR
	HealthScore int `json:"health_score" bson:"health_score"`

	// DEPRECATED -- last session traffic data
	InPackets  int64 `json:"in_packets" bson:"in_packets"`
	OutPackets int64 `json:"out_packets" bson:"out_packets"`
	InKBytes   int64 `json:"in_kbytes" bson:"in_kbytes"`
	OutKBytes  int64 `json:"out_kbytes" bson:"out_kbytes"`

	// session overall data
	FirstConnect  int64 `json:"first_connect" bson:"first_connect"`
	LastConnect   int64 `json:"last_connect" bson:"last_connect"`
	LastAuthorise int64 `json:"last_auth" bson:"last_auth"`

	// auth data
	UserAgent string `json:"useragent" bson:"useragent"`
	UserName  string `json:"username" bson:"username"`

	OS        string `json:"os" bson:"os"`
	OSVersion string `json:"os_version" bson:"os_version"`
	UADevice  string `json:"ua_device" bson:"ua_device"`
	UAType    string `json:"ua_type" bson:"ua_type"`

	CreateAt time.Time `json:"create_at" bson:"create_at"`
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
	ID       string `json:"id" bson:"_id"`
	CPE      string `json:"cpe_id" bson:"cpe_id"`
	Start    int64  `json:"start" bson:"start"`
	Stop     int64  `json:"stop" bson:"stop"`
	Duration int64  `json:"duration" bson:"duration"`

	CreateAt time.Time `json:"create_at" bson:"create_at"`
}

type ClientSessionInfo struct {
	ID   string `json:"id" bson:"_id"`
	MAC  string `json:"mac" bson:"mac"`
	WLAN string `json:"wlan_id" bson:"wlan_id"`
	SSID string `json:"ssid" bson:"ssid"`
	CPE  string `json:"cpe_id" bson:"cpe_id"`

	Radio string `json:"radio_id" bson:"radio_id"`
	Freq  string `json:"freq" bson:"freq"`
	Mode  string `json:"mode" bson:"mode"`

	Start     int64 `json:"start" bson:"start"`
	Stop      int64 `json:"stop" bson:"stop"`
	Duration  int64 `json:"duration" bson:"duration"`
	Timestamp int64 `json:"timestamp" bson:"timestamp"`

	StartNoise int `json:"start_noise" bson:"start_noise"`
	StopNoise  int `json:"stop_noise" bson:"stop_noise"`

	StartRSSI int `json:"start_rssi" bson:"start_rssi"`
	StopRSSI  int `json:"stop_rssi" bson:"stop_rssi"`

	RxBytes int64 `json:"rx_bytes" bson:"rx_bytes"`
	TxBytes int64 `json:"tx_bytes" bson:"tx_bytes"`

	// auth data
	UserAgent  string `json:"useragent" bson:"useragent"`
	UserName   string `json:"username" bson:"username"`
	AuthenType string `json:"authen_type,omitempty" bson:"authen_type"`
	AuthType   string `json:"auth_type,omitempty" bson:"auth_type"`

	CreateAt time.Time `json:"create_at" bson:"create_at"`

	// list of clients IPS
	DPI AccountingDPI `json:"dpi" bson:"dpi"`
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

//LBSZone зоны на картах как Cisco CMX
type LBSZone struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	//Identity поля для определения зоны
	Identity string         `json:"identity" bson:"identity"`
	Group    UUID           `json:"group"`
	Corners  []CornerCoords `json:"corners" bson:"corners"`
}

//CornerCoords координаты углов зоны
type CornerCoords struct {
	X float64 `json:"x" bson:"x"`
	Y float64 `json:"y" bson:"y"`
	Z float64 `json:"z" bson:"z"`
}
