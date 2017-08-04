package libwimark

import (
	"encoding/json"
	"errors"
	"sort"
	"strconv"

	"github.com/vorot93/goutil"
	"gopkg.in/mgo.v2/bson"
)

type Document goutil.Document

type UUID string

type Radius struct {
	Name      string `json:"name"`
	Hostname  string `json:"hostname"`
	Auth_port string `json:"auth_port"`
	Acc_port  string `json:"acc_port"`
	Secret    string `json:"secret"`
}

type InterfaceType string

const (
	I2_4 = InterfaceType("2.4")
	I5_0 = InterfaceType("5.0")
)

type WPA2Common struct {
	Suites []SecuritySuite `json:"suites"`
}

type WPA2PersonalData struct {
	WPA2Common `bson:",inline"`
	PSK        string `json:"psk"`
}

type WPA2EnterpriseData struct {
	WPA2Common           `bson:",inline"`
	NasID                string `json:"nasid"`
	PMKCaching           bool   `json:"pmkcaching"`
	RadiusAuthentication []UUID `json:"radiusauthentication"`
}

type SecuritySettings interface {
	is_security_settings()
}

type WLAN struct {
	Name              string        `json:"name"`
	SSID              string        `json:"ssid"`
	Description       string        `json:"description"`
	Security          *EnumSecurity `json:"security"`
	VLAN              int           `json:"vlan"`
	Hidden            bool          `json:"hidden"`
	NasID             *string       `json:"nas_id"`
	RadiusAcctServers []UUID        `json:"radius_acct_servers"`
	WhiteList         []string      `json:"whitelist"`
	BlackList         []string      `json:"blacklist"`
	FilterMode        MacFilterType `json:"filtermode"`
	L2Isolate         bool          `json:"l2isolate"`
}

type InterfaceConfiguration struct {
	WLANs []UUID
}

type Configuration struct {
	IfaceConfigs map[InterfaceType]InterfaceConfiguration
}

type WiredData struct {
	Name string `json:"name"`
	Mac  string `json:"mac"`
	VLAN int    `json:"vlan"`
}

type WiFiData struct {
	Name      string `json:"name,omitempty"`
	Mac       string `json:"mac,omitempty"`
	Frequency string `json:"frequency,omitempty"`
	BandMode  string `json:"bandmode,omitempty"`
	Bandwidth string `json:"bandwidth,omitempty"`
	Channel   string `json:"channel,omitempty"`
	TxPower   string `json:"txpower,omitempty"`
	WLANs     []UUID `json:"wlans,omitempty"`
	ChanList  []int  `json:"chanlist,omitempty"`
}

type CPEInterface struct {
	CPEInterfaceInfo
	Addr         string
	Capabilities Capabilities
}

func (self CPEInterface) MarshalJSON() ([]byte, error) {
	var b []byte
	{
		var err error
		b, err = json.Marshal(self.CPEInterfaceInfo)
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

	doc["addr"] = self.Addr
	doc["capabilities"] = self.Capabilities

	return json.Marshal(doc)
}

func (self *CPEInterface) UnmarshalJSON(b []byte) error {
	var doc map[string]json.RawMessage
	var err = json.Unmarshal(b, &doc)
	if err != nil {
		return err
	}

	if doc == nil {
		return nil
	}

	var addrRaw, addrExists = doc["addr"]
	if addrExists {
		var addr string
		var addrErr = json.Unmarshal(addrRaw, &addr)
		if addrErr == nil {
			self.Addr = addr
		} else {
			return addrErr
		}
	}

	delete(doc, "addr")

	var capsRaw, capsExists = doc["capabilities"]
	if capsExists {
		var caps Capabilities
		var capsErr = json.Unmarshal(capsRaw, &caps)
		if capsErr == nil {
			self.Capabilities = caps
		} else {
			return capsErr
		}
	}

	delete(doc, "capabilities")

	var v, _ = json.Marshal(doc)

	return self.CPEInterfaceInfo.UnmarshalJSON(v)
}

type CPEInterfaces map[string]CPEInterface

func (self CPEInterfaces) GetBSON() (interface{}, error) {
	var out = []bson.M{}
	var keys = []string{}
	for k, _ := range self {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := self[k]
		var b []byte
		{
			var err error
			b, err = bson.Marshal(v)
			if err != nil {
				return nil, err
			}
		}

		var obj_m bson.M
		{
			var err error
			err = bson.Unmarshal(b, &obj_m)
			if err != nil {
				return nil, err
			}
		}

		obj_m["_id"] = k
		out = append(out, obj_m)
	}

	return out, nil
}

func (self *CPEInterfaces) SetBSON(raw bson.Raw) error {
	var in = []bson.M{}
	{
		var err error
		err = raw.Unmarshal(&in)
		if err != nil {
			return err
		}
	}

	var out = CPEInterfaces{}

	for _, v := range in {
		var obj_b []byte
		{
			var err error
			obj_b, err = bson.Marshal(v)
			if err != nil {
				return err
			}
		}

		var obj CPEInterface
		{
			var err error
			err = bson.Unmarshal(obj_b, &obj)
			if err != nil {
				return err
			}
		}
		var k_any, k_found = v["_id"]
		if !k_found {
			return errors.New("No ID found")
		}

		var k, k_correct = k_any.(string)
		if !k_correct {
			return errors.New("Invalid key type")
		}

		out[k] = obj
	}

	*self = out

	return nil
}

type CPE struct {
	Name        string `json:"name"`
	Connected   bool   `json:"connected"`
	Description string `json:"description"`
	Model       struct {
		Id   UUID   `json:"id"`
		Name string `json:"name"`
	} `json:"model"`
	Interfaces       CPEInterfaces       `json:"interfaces"`
	ConfigStatus     ConfigurationStatus `json:"config_status"`
	LbsConfig        LBSConfig           `json:"lbs_config"`
	StatisticsConfig StatisticsConfig    `json:"stats_config"`
	LogConfig        LogConfig           `json:"log_config"`
}

type Stat struct {
	CPE       UUID    `json:"cpeid"`
	Timestamp int64   `json:"timestamp"`
	CPU       float64 `json:"cpu"`
	RAM       struct {
		Total    float64 `json:"total"`
		Buffered float64 `json:"buffered"`
		Shared   float64 `json:"shared"`
		Free     float64 `json:"free"`
	} `json:"memory"`
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
}

type CapTxPower struct {
	DBelMw    int `json:"dbm"`
	MilliWatt int `json:"mw"`
}

type CapChannel struct {
	Restricted bool       `json:"restricted"`
	Freq       int        `json:"mhz"`
	Channel    int        `json:"channel"`
	MaxPower   CapTxPower `json:"max_txpower"`
}

type Capabilities struct {
	TxPowers []CapTxPower    `json:"txpwrlist"`
	HTModes  map[string]bool `json:"htmodelist"`
	HWModes  map[string]bool `json:"hwmodelist"`
	Channels []CapChannel    `json:"freqlist"`
	TxOffset int             `json:"txpower_offset"`
}

type CPEAgentResponse struct {
	Result json.RawMessage `json:"result"`
	Status CPEAgentError   `json:"error"`
}

type LimitBetween struct {
	Upper float64 `json:"upper"`
	Lower float64 `json:"lower"`
}

type LimitBetweenOptional struct {
	Upper *float64 `json:"upper"`
	Lower *float64 `json:"lower"`
}

type CameraClient struct {
	Rtsp_stream string `json:"rtsp_stream"`
}

type OtherClient struct {
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

// Events
type SystemEvent struct {
	SystemEventObject
	Timestamp  int64            `json:"timestamp"`
	Subject_id string           `json:"subject_id"`
	Level      SystemEventLevel `json:"level"`
}

func (self *SystemEvent) MarshalJSON() ([]byte, error) {
	var b []byte
	{
		var err error
		b, err = json.Marshal(self.SystemEventObject)
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

	doc["timestamp"] = self.Timestamp
	doc["subject_id"] = self.Subject_id
	doc["level"] = self.Level

	return json.Marshal(doc)
}

func (self *SystemEvent) UnmarshalJSON(b []byte) error {
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

	var tsRaw, tsExists = doc["timestamp"]
	if tsExists {
		var ts int64
		var tsErr = json.Unmarshal(tsRaw, &ts)
		if tsErr == nil {
			self.Timestamp = ts
		} else {
			return tsErr
		}
	}

	delete(doc, "timestamp")

	var subRaw, subExists = doc["subject_id"]
	if subExists {
		var subject_id string
		var subErr = json.Unmarshal(subRaw, &subject_id)
		if subErr == nil {
			self.Subject_id = subject_id
		} else {
			return subErr
		}
	}

	delete(doc, "subject_id")

	var levelRaw, levelExists = doc["level"]
	if levelExists {
		var level SystemEventLevel
		var levErr = json.Unmarshal(levelRaw, &level)
		if levErr == nil {
			self.Level = level
		} else {
			return levErr
		}
	}

	delete(doc, "level")

	var v, _ = json.Marshal(doc)

	return self.SystemEventObject.UnmarshalJSON(v)
}

func (self *SystemEvent) GetBSON() (interface{}, error) {
	var out = bson.M{}

	levelBson, err := self.Level.GetBSON()
	if err != nil {
		return nil, err
	}

	var obj_b []byte
	{
		var err error
		obj_b, err = bson.Marshal(self.SystemEventObject)
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
	out["timestamp"] = self.Timestamp
	out["subject_id"] = self.Subject_id
	out["level"] = levelBson

	return out, nil
}

func (self *SystemEvent) SetBSON(raw bson.Raw) error {
	var in = bson.M{}
	{
		var err error
		err = raw.Unmarshal(&in)
		if err != nil {
			return err
		}
	}

	//timestamp
	var v_timestamp, k_found = in["timestamp"]
	if !k_found {
		return errors.New("No timestamp found")
	}
	self.Timestamp = v_timestamp.(int64)

	delete(in, "timestamp")

	//subject_id
	v_subj, k_found := in["subject_id"]
	if !k_found {
		return errors.New("No subject found")
	}

	self.Subject_id = v_subj.(string)

	delete(in, "subject_id")

	//subject_id
	v_level, k_found := in["level"]
	if !k_found {
		return errors.New("No subject found")
	}

	var err error
	obj_b, err := bson.Marshal(v_level)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(obj_b, &self.Level)
	if err != nil {
		return err
	}
	delete(in, "level")

	{
		var err error
		obj_b, err = bson.Marshal(in)
		if err != nil {
			return err
		}
	}
	err = bson.Unmarshal(obj_b, &self.SystemEventObject)
	if err != nil {
		return err
	}

	return nil
}

type CPEConnectedData struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Build   string `json:"build"`
}

type MonitorRuleViolationData struct {
	Rule_name string `json:"rule_name"`
	Cpe_id    UUID   `json:"cpe_id"`
}

type ClientConnectedData struct {
	Session_id string `json:"session_id"`
	Freq       string `json:"freq"`
	Cpe_id     UUID   `json:"cpe_id"`
	Wlan_id    UUID   `json:"wlan_id"`
	Radio_id   string `json:"radio_id"`
}

type ClientDisconnectedData struct {
	Session_id string `json:"session_id"`
	Freq       string `json:"freq"`
	Cpe_id     UUID   `json:"cpe_id"`
	Wlan_id    UUID   `json:"wlan_id"`
	Radio_id   string `json:"radio_id"`
}

type CPEConfigurationErrorData struct {
	Wifi        WifiConfigurationError `json:"wifi,omitempty"`
	Radius      string                 `json:"radius,omitempty"`
	Description string                 `json:"description,omitempty"`
}

type WifiConfigurationError struct {
	Wlan_id  UUID `json:"wlan_id"`
	Radio_id UUID `json:"radio_id"`
}

type ServiceFatalErrorData struct {
	Sw_version  string `json:"sw_version"`
	Description string `json:"description"`
}

type ServiceConnectedData struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Build   int    `json:"build"`
}

type LBSClientData struct {
	Timestamp int64  `json:"timestamp"`
	CPE       UUID   `json:"cpe"`
	Radio     string `json:"radio"`
	ClientMac string `json:"client_mac"`
	RSSI      int    `json:"rssi"`
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
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Z         float64 `json:"z"`
}

type LBSConfig struct {
	Enabled bool `json:"enabled"`
	// in seconds
	ReportPeriod  int `json:"reportperiod,omitempty"`
	ClientTimeout int `json:"clienttimeout,omitempty"`
}

type StatisticsConfig struct {
	Enabled bool `json:"enabled"`
	// in seconds
	ReportPeriod int `json:"reportperiod,omitempty"`
}

type LogConfig struct {
	Enabled bool `json:"enabled"`
	// in seconds
	ReportPeriod int `json:"reportperiod,omitempty"`
}

type SimpleMask struct {
	UUID []UUID `json:"uuid"`
}

type TimestampMask struct {
	UUID  []UUID `json:"uuid"`
	Start *int64 `json:"start"`
	Stop  *int64 `json:"stop"`
}

type EventMask struct {
	TimestampMask
	Type       []SystemEventObjectType `json:"type"`
	Subject_id []string                `json:"subject_id"`
	Level      []SystemEventLevel      `json:"level"`
}

type ClientStatMask struct {
	TimestampMask
	CPE              []UUID
	CallingStationId []string
}

type CPEMask struct {
	UUID      []UUID `json:"uuid"`
	HasWLANs  []UUID `json:"has_wlans"`
	Connected *bool  `json:"connected"`
}

type WLANMask struct {
	UUID      []UUID `json:"uuid"`
	HasRadius []UUID `json:"has_radius"`
}

type StatsMask struct {
	UUID    []UUID `json:"uuid"`
	CPEUUID []UUID `json:"cpe"`
	Start   *int64 `json:"start"`
	Stop    *int64 `json:"stop"`
}

type LBSClientDataMask struct {
	TimestampMask
	CPE       []UUID   `json:"cpe"`
	Radio     []string `json:"radio"`
	ClientMac []string `json:"client_mac"`
	RSSI      []int    `json:"rssi"`
}

type LBSCPEInfoMask struct {
	SimpleMask
	Group []UUID               `json:"group"`
	CPE   []UUID               `json:"cpe"`
	Name  []string             `json:"name"`
	X     LimitBetweenOptional `json:"x"`
	Y     LimitBetweenOptional `json:"y"`
	Z     LimitBetweenOptional `json:"z"`
}

type LBSClientCoordsMask struct {
	TimestampMask
	Group []UUID               `json:"group"`
	Mac   []string             `json:"mac"`
	X     LimitBetweenOptional `json:"x"`
	Y     LimitBetweenOptional `json:"y"`
	Z     LimitBetweenOptional `json:"z"`
}

type WirelessClient struct {
	Type        WirelessClientType
	Data        interface{} `bson:"data" json:"data"`
	Mac         string      `json:"mac" bson:"_id"`
	Ip          string      `json:"ip"`
	Net_mask    string      `json:"net_mask"`
	Wlan_id     string      `json:"wlan_id"`
	Cpe_id      string      `json:"cpe_id"`
	Freq        string      `json:"freq"`
	Radio_id    string      `json:"radio_id"`
	Rssi        int         `json:"rssi"`
	Timestamp   int64       `json:"timestamp"`
	In_packets  int64       `json:"in_packets"`
	Out_packets int64       `json:"out_packets"`
	In_kbytes   int64       `json:"in_kbytes"`
	Out_kbytes  int64       `json:"out_kbytes"`
	State       WirelessClientState
}

func (wc *WirelessClient) GetSpecificWCInfo() *WirelessClientObject {
	wco := &WirelessClientObject{Data: wc.Data, Type: wc.Type}
	return wco
}

func (wc *WirelessClient) SetSpecificWCInfo(wco *WirelessClientObject) {
	wc.Type = wco.Type
	wc.Data = wco.Data
}

type Version struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Build   int    `json:"build"`
}

func MakeVersion(version string, commit string, build string) Version {
	var self Version

	self.Version = version
	self.Commit = commit
	build_num, _ := strconv.Atoi(build)
	self.Build = build_num

	return self
}
