package libwimark

import (
	"encoding/json"
)

type DBResponseBase struct {
	Errors []ModelError `json:"errors,omitempty"`
}

type DBDataObj struct {
	WLANs           map[UUID]WLAN             `json:"wlan,omitempty"`
	CPEs            map[UUID]CPE              `json:"cpe,omitempty"`
	Stats           map[UUID]Stat             `json:"stat,omitempty"`
	ClientStats     map[UUID]ClientStat       `json:"client-stat,omitempty"`
	Events          map[UUID]SystemEvent      `json:"event,omitempty"`
	StatEventRules  map[UUID]StatEventRule    `json:"stat-event-rule,omitempty"`
	PollCPE         map[UUID]CPEPollSettings  `json:"poll-cpe,omitempty"`
	Radius          map[UUID]Radius           `json:"radius,omitempty"`
	LBSCPEInfo      map[UUID]LBSCPEInfo       `json:"lbs-cpe-info,omitempty"`
	LBSClientData   map[UUID]LBSClientData    `json:"lbs-client-data,omitempty"`
	LBSClientCoords map[UUID]LBSClientCoords  `json:"lbs-client-coords,omitempty"`
	VPNHosts        map[UUID]VPNHost          `json:"vpn-host,omitempty"`
	CPEScanData     map[UUID]CPEScanData      `json:"cpe-scan-data,omitempty"`
	CPEModel        map[UUID]CPEModel         `json:"cpe-model,omitempty"`
	ConfigRule      map[UUID]ConfigRule       `json:"config-rule,omitempty"`
	L2Chain         map[UUID]L2Chain          `json:"l2-chain,omitempty"`
	CaptiveRedirect map[UUID]CaptiveRedirect  `json:"captive-redirect,omitempty"`
	HotspotProfile  map[UUID]Hotspot20Profile `json:"hotspot-profile,omitempty"`
	Controller      map[UUID]Controller       `json:"controller,omitempty"`
	LBSZones        map[UUID]LBSZone          `json:"lbs_zones,omitempty"`
}

func (dbd *DBDataObj) Reset() {
	*dbd = DBDataObj{}
}

type DBDataUUID struct {
	WLANs           []UUID `json:"wlan,omitempty"`
	CPEs            []UUID `json:"cpe,omitempty"`
	Stats           []UUID `json:"stat,omitempty"`
	ClientStats     []UUID `json:"client-stat,omitempty"`
	Events          []UUID `json:"event,omitempty"`
	StatEventRules  []UUID `json:"stat-event-rule,omitempty"`
	PollCPE         []UUID `json:"poll-cpe,omitempty"`
	Radius          []UUID `json:"radius,omitempty"`
	LBSCPEInfo      []UUID `json:"lbs-cpe-info,omitempty"`
	LBSClientData   []UUID `json:"lbs-client-data,omitempty"`
	LBSClientCoords []UUID `json:"lbs-client-coords,omitempty"`
	VPNHosts        []UUID `json:"vpn-host,omitempty"`
	CPEScanData     []UUID `json:"cpe-scan-data,omitempty"`
	CPEModel        []UUID `json:"cpe-model,omitempty"`
	ConfigRule      []UUID `json:"config-rule,omitempty"`
	L2Chain         []UUID `json:"l2-chain,omitempty"`
	CaptiveRedirect []UUID `json:"captive-redirect,omitempty"`
	HotspotProfile  []UUID `json:"hotspot-profile,omitempty"`
	Controller      []UUID `json:"controller,omitempty"`
	LBSZones        []UUID `json:"lbs_zones,omitempty"`
}

func (dbd *DBDataUUID) Reset() {
	*dbd = DBDataUUID{}
}

type DBDataMasks struct {
	WLANs           *WLANMask            `json:"wlan,omitempty"`
	CPEs            *CPEMask             `json:"cpe,omitempty"`
	Stats           *StatsMask           `json:"stat,omitempty"`
	ClientStats     *SimpleMask          `json:"client-stat,omitempty"`
	Events          *EventMask           `json:"event,omitempty"`
	StatEventRules  *SimpleMask          `json:"stat-event-rule,omitempty"`
	PollCPE         *SimpleMask          `json:"poll-cpe,omitempty"`
	Radius          *SimpleMask          `json:"radius,omitempty"`
	LBSCPEInfo      *LBSCPEInfoMask      `json:"lbs-cpe-info,omitempty"`
	LBSClientData   *LBSClientDataMask   `json:"lbs-client-data,omitempty"`
	LBSClientCoords *LBSClientCoordsMask `json:"lbs-client-coords,omitempty"`
	VPNHosts        *SimpleMask          `json:"vpn-host,omitempty"`
	CPEScanData     *SimpleMask          `json:"cpe-scan-data,omitempty"`
	CPEModel        *CPEModelMask        `json:"cpe-model,omitempty"`
	ConfigRule      *ConfigRuleMask      `json:"config-rule,omitempty"`
	L2Chain         *SimpleMask          `json:"l2-chain,omitempty"`
	CaptiveRedirect *SimpleMask          `json:"captive-redirect,omitempty"`
	HotspotProfile  *SimpleMask          `json:"hotspot-profile,omitempty"`
	Controller      *ControllerMask      `json:"controller,omitempty"`
	LBSZones        *SimpleMask          `json:"lbs_zones,omitempty"`
	BaseLocation    *SimpleMask          `json:",omitempty"`
}

func (dbd *DBDataMasks) Reset() {
	*dbd = DBDataMasks{}
}

type DBRequestC DBDataObj
type DBRequestU DBDataObj
type DBRequestR DBDataMasks
type DBRequestD DBDataMasks

func (dbd *DBRequestC) UnmarshalJSON(b []byte) error {
	(*DBDataObj)(dbd).Reset()

	return json.Unmarshal(b, (*DBDataObj)(dbd))
}

func (dbd *DBRequestR) UnmarshalJSON(b []byte) error {
	(*DBDataMasks)(dbd).Reset()

	return json.Unmarshal(b, (*DBDataMasks)(dbd))
}

func (dbd *DBRequestU) UnmarshalJSON(b []byte) error {
	(*DBDataObj)(dbd).Reset()

	return json.Unmarshal(b, (*DBDataObj)(dbd))
}

func (dbd *DBRequestD) UnmarshalJSON(b []byte) error {
	(*DBDataMasks)(dbd).Reset()

	return json.Unmarshal(b, (*DBDataMasks)(dbd))
}

type DBResponseObj struct {
	DBResponseBase `json:",inline"`
	DBDataObj      `json:"data,inline"`
}

type DBResponseUUID struct {
	DBResponseBase `json:",inline"`
	DBDataUUID     `json:"data,inline"`
}

type DBResponseC DBResponseUUID
type DBResponseR DBResponseObj
type DBResponseU DBResponseUUID
type DBResponseD DBResponseUUID

func (dbd *DBResponseC) UnmarshalJSON(b []byte) error {
	dbd.DBDataUUID.Reset()

	return json.Unmarshal(b, (*DBResponseUUID)(dbd))
}

func (dbd *DBResponseR) UnmarshalJSON(b []byte) error {
	dbd.DBDataObj.Reset()

	return json.Unmarshal(b, (*DBResponseObj)(dbd))
}

func (dbd *DBResponseU) UnmarshalJSON(b []byte) error {
	dbd.DBDataUUID.Reset()

	return json.Unmarshal(b, (*DBResponseUUID)(dbd))
}

func (dbd *DBResponseD) UnmarshalJSON(b []byte) error {
	dbd.DBDataUUID.Reset()

	return json.Unmarshal(b, (*DBResponseUUID)(dbd))
}

type ConnectorInfo struct {
	DbType    string   `json:"db_type"`
	DbServers []string `json:"db_servers"`
	Models    []string `json:"models"`
}
