package libwimark

import (
	"encoding/json"
)

type DBResponseBase struct {
	Errors []ModelError `json:"errors,omitempty"`
}

type DBDataObj struct {
	WLANs             map[UUID]WLAN              `json:"wlan,omitempty"`
	CPEs              map[UUID]CPE               `json:"cpe,omitempty"`
	Stats             map[UUID]Stat              `json:"stat,omitempty"`
	ClientStats       map[UUID]ClientStat        `json:"client-stat,omitempty"`
	Events            map[UUID]SystemEvent       `json:"event,omitempty"`
	StatEventRules    map[UUID]StatEventRule     `json:"stat-event-rule,omitempty"`
	PollCPE           map[UUID]CPEPollSettings   `json:"poll-cpe,omitempty"`
	Radius            map[UUID]Radius            `json:"radius,omitempty"`
	LBSCPEInfo        map[UUID]LBSCPEInfo        `json:"lbs-cpe-info,omitempty"`
	LBSClientData     map[UUID]LBSClientData     `json:"lbs-client-data,omitempty"`
	LBSClientCoords   map[UUID]LBSClientCoords   `json:"lbs-client-coords,omitempty"`
	VPNHosts          map[UUID]VPNHost           `json:"vpn-host,omitempty"`
	L2TPTunnelSession map[UUID]L2TPTunnelSession `json:"tunnels,omitempty"`
	CPEScanData       map[UUID]CPEScanData       `json:"cpe-scan-data,omitempty"`
	CPEModel          map[UUID]CPEModel          `json:"cpe-model,omitempty"`
	ConfigRule        map[UUID]ConfigRule        `json:"config-rule,omitempty"`
	L2Chain           map[UUID]L2Chain           `json:"l2-chain,omitempty"`
	CaptiveRedirect   map[UUID]CaptiveRedirect   `json:"captive-redirect,omitempty"`
}

func (self *DBDataObj) Reset() {
	*self = DBDataObj{}
}

type DBDataUUID struct {
	WLANs             []UUID `json:"wlan,omitempty"`
	CPEs              []UUID `json:"cpe,omitempty"`
	Stats             []UUID `json:"stat,omitempty"`
	ClientStats       []UUID `json:"client-stat,omitempty"`
	Events            []UUID `json:"event,omitempty"`
	StatEventRules    []UUID `json:"stat-event-rule,omitempty"`
	PollCPE           []UUID `json:"poll-cpe,omitempty"`
	Radius            []UUID `json:"radius,omitempty"`
	LBSCPEInfo        []UUID `json:"lbs-cpe-info,omitempty"`
	LBSClientData     []UUID `json:"lbs-client-data,omitempty"`
	LBSClientCoords   []UUID `json:"lbs-client-coords,omitempty"`
	VPNHosts          []UUID `json:"vpn-host,omitempty"`
	L2TPTunnelSession []UUID `json:"tunnels,omitempty"`
	CPEScanData       []UUID `json:"cpe-scan-data,omitempty"`
	CPEModel          []UUID `json:"cpe-model,omitempty"`
	ConfigRule        []UUID `json:"config-rule,omitempty"`
	L2Chain           []UUID `json:"l2-chain,omitempty"`
	CaptiveRedirect   []UUID `json:"captive-redirect,omitempty"`
}

func (self *DBDataUUID) Reset() {
	*self = DBDataUUID{}
}

type DBDataMasks struct {
	WLANs             *WLANMask            `json:"wlan,omitempty"`
	CPEs              *CPEMask             `json:"cpe,omitempty"`
	Stats             *StatsMask           `json:"stat,omitempty"`
	ClientStats       *SimpleMask          `json:"client-stat,omitempty"`
	Events            *EventMask           `json:"event,omitempty"`
	StatEventRules    *SimpleMask          `json:"stat-event-rule,omitempty"`
	PollCPE           *SimpleMask          `json:"poll-cpe,omitempty"`
	Radius            *SimpleMask          `json:"radius,omitempty"`
	LBSCPEInfo        *LBSCPEInfoMask      `json:"lbs-cpe-info,omitempty"`
	LBSClientData     *LBSClientDataMask   `json:"lbs-client-data,omitempty"`
	LBSClientCoords   *LBSClientCoordsMask `json:"lbs-client-coords,omitempty"`
	VPNHosts          *SimpleMask          `json:"vpn-host,omitempty"`
	L2TPTunnelSession *TunnelMask          `json:"tunnels,omitempty"`
	CPEScanData       *SimpleMask          `json:"cpe-scan-data,omitempty"`
	CPEModel          *CPEModelMask        `json:"cpe-model,omitempty"`
	ConfigRule        *ConfigRuleMask      `json:"config-rule,omitempty"`
	L2Chain           *SimpleMask          `json:"l2-chain,omitempty"`
	CaptiveRedirect   *SimpleMask          `json:"captive-redirect,omitempty"`
}

func (self *DBDataMasks) Reset() {
	*self = DBDataMasks{}
}

type DBRequestC DBDataObj
type DBRequestU DBDataObj
type DBRequestR DBDataMasks
type DBRequestD DBDataMasks

func (self *DBRequestC) UnmarshalJSON(b []byte) error {
	(*DBDataObj)(self).Reset()

	return json.Unmarshal(b, (*DBDataObj)(self))
}

func (self *DBRequestR) UnmarshalJSON(b []byte) error {
	(*DBDataMasks)(self).Reset()

	return json.Unmarshal(b, (*DBDataMasks)(self))
}

func (self *DBRequestU) UnmarshalJSON(b []byte) error {
	(*DBDataObj)(self).Reset()

	return json.Unmarshal(b, (*DBDataObj)(self))
}

func (self *DBRequestD) UnmarshalJSON(b []byte) error {
	(*DBDataMasks)(self).Reset()

	return json.Unmarshal(b, (*DBDataMasks)(self))
}

type DBResponseObj struct {
	DBResponseBase `json:",inline"`
	DBDataObj      `json:"data",inline`
}

type DBResponseUUID struct {
	DBResponseBase `json:",inline"`
	DBDataUUID     `json:"data",inline`
}

type DBResponseC DBResponseUUID
type DBResponseR DBResponseObj
type DBResponseU DBResponseUUID
type DBResponseD DBResponseUUID

func (self *DBResponseC) UnmarshalJSON(b []byte) error {
	self.DBDataUUID.Reset()

	return json.Unmarshal(b, (*DBResponseUUID)(self))
}

func (self *DBResponseR) UnmarshalJSON(b []byte) error {
	self.DBDataObj.Reset()

	return json.Unmarshal(b, (*DBResponseObj)(self))
}

func (self *DBResponseU) UnmarshalJSON(b []byte) error {
	self.DBDataUUID.Reset()

	return json.Unmarshal(b, (*DBResponseUUID)(self))
}

func (self *DBResponseD) UnmarshalJSON(b []byte) error {
	self.DBDataUUID.Reset()

	return json.Unmarshal(b, (*DBResponseUUID)(self))
}

type ConnectorInfo struct {
	DbType    string   `json:"db_type"`
	DbServers []string `json:"db_servers"`
	Models    []string `json:"models"`
}
