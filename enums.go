package libwimark

import (
	"encoding/json"
	"errors"
	"github.com/globalsign/mgo/bson"
)

type CPEAgentStatusType string

const CPEAgentStatusTypeException CPEAgentStatusType = "exception"
const CPEAgentStatusTypeSuccess CPEAgentStatusType = "success"
const CPEAgentStatusTypeSyntaxError CPEAgentStatusType = "syntax"
const CPEAgentStatusTypeUndefined CPEAgentStatusType = "undefined"

func (en CPEAgentStatusType) GetPtr() *CPEAgentStatusType { var v = en; return &v }

func (en CPEAgentStatusType) String() string {
	switch en {
	case CPEAgentStatusTypeException:
		return "exception"
	case CPEAgentStatusTypeSuccess:
		return "success"
	case CPEAgentStatusTypeSyntaxError:
		return "syntax"
	case CPEAgentStatusTypeUndefined:
		return "undefined"
	}
	if len(en) == 0 {
		return "undefined"
	}
	panic(errors.New("Invalid value of CPEAgentStatusType: " + string(en)))
}

func (en *CPEAgentStatusType) MarshalJSON() ([]byte, error) {
	switch *en {
	case CPEAgentStatusTypeException:
		return json.Marshal("exception")
	case CPEAgentStatusTypeSuccess:
		return json.Marshal("success")
	case CPEAgentStatusTypeSyntaxError:
		return json.Marshal("syntax")
	case CPEAgentStatusTypeUndefined:
		return json.Marshal("undefined")
	}
	if len(*en) == 0 {
		return json.Marshal("undefined")
	}
	return nil, errors.New("Invalid value of CPEAgentStatusType: " + string(*en))
}

func (en *CPEAgentStatusType) GetBSON() (interface{}, error) {
	switch *en {
	case CPEAgentStatusTypeException:
		return "exception", nil
	case CPEAgentStatusTypeSuccess:
		return "success", nil
	case CPEAgentStatusTypeSyntaxError:
		return "syntax", nil
	case CPEAgentStatusTypeUndefined:
		return "undefined", nil
	}
	if len(*en) == 0 {
		return "undefined", nil
	}
	return nil, errors.New("Invalid value of CPEAgentStatusType: " + string(*en))
}

func (en *CPEAgentStatusType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "exception":
		*en = CPEAgentStatusTypeException
		return nil
	case "success":
		*en = CPEAgentStatusTypeSuccess
		return nil
	case "syntax":
		*en = CPEAgentStatusTypeSyntaxError
		return nil
	case "undefined":
		*en = CPEAgentStatusTypeUndefined
		return nil
	}
	if len(s) == 0 {
		*en = CPEAgentStatusTypeUndefined
		return nil
	}
	return errors.New("Unknown CPEAgentStatusType: " + s)
}

func (en *CPEAgentStatusType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "exception":
		*en = CPEAgentStatusTypeException
		return nil
	case "success":
		*en = CPEAgentStatusTypeSuccess
		return nil
	case "syntax":
		*en = CPEAgentStatusTypeSyntaxError
		return nil
	case "undefined":
		*en = CPEAgentStatusTypeUndefined
		return nil
	}
	if len(s) == 0 {
		*en = CPEAgentStatusTypeUndefined
		return nil
	}
	return errors.New("Unknown CPEAgentStatusType: " + s)
}

type CPEInterfaceState string

const CPEInterfaceStateACS CPEInterfaceState = "ACS"
const CPEInterfaceStateCountryUpdate CPEInterfaceState = "COUNTRY_UPDATE"
const CPEInterfaceStateDFS CPEInterfaceState = "DFS"
const CPEInterfaceStateDisabled CPEInterfaceState = "DISABLED"
const CPEInterfaceStateEnabled CPEInterfaceState = "ENABLED"
const CPEInterfaceStateHtScan CPEInterfaceState = "HT_SCAN"
const CPEInterfaceStateTerminated CPEInterfaceState = "TERMINATED"
const CPEInterfaceStateUninitialized CPEInterfaceState = "UNINITIALIZED"
const CPEInterfaceStateUnknown CPEInterfaceState = "UNKNOWN"

func (en CPEInterfaceState) GetPtr() *CPEInterfaceState { var v = en; return &v }

func (en CPEInterfaceState) String() string {
	switch en {
	case CPEInterfaceStateACS:
		return "ACS"
	case CPEInterfaceStateCountryUpdate:
		return "COUNTRY_UPDATE"
	case CPEInterfaceStateDFS:
		return "DFS"
	case CPEInterfaceStateDisabled:
		return "DISABLED"
	case CPEInterfaceStateEnabled:
		return "ENABLED"
	case CPEInterfaceStateHtScan:
		return "HT_SCAN"
	case CPEInterfaceStateTerminated:
		return "TERMINATED"
	case CPEInterfaceStateUninitialized:
		return "UNINITIALIZED"
	case CPEInterfaceStateUnknown:
		return "UNKNOWN"
	}
	panic(errors.New("Invalid value of CPEInterfaceState: " + string(en)))
}

func (en *CPEInterfaceState) MarshalJSON() ([]byte, error) {
	switch *en {
	case CPEInterfaceStateACS:
		return json.Marshal("ACS")
	case CPEInterfaceStateCountryUpdate:
		return json.Marshal("COUNTRY_UPDATE")
	case CPEInterfaceStateDFS:
		return json.Marshal("DFS")
	case CPEInterfaceStateDisabled:
		return json.Marshal("DISABLED")
	case CPEInterfaceStateEnabled:
		return json.Marshal("ENABLED")
	case CPEInterfaceStateHtScan:
		return json.Marshal("HT_SCAN")
	case CPEInterfaceStateTerminated:
		return json.Marshal("TERMINATED")
	case CPEInterfaceStateUninitialized:
		return json.Marshal("UNINITIALIZED")
	case CPEInterfaceStateUnknown:
		return json.Marshal("UNKNOWN")
	}
	return nil, errors.New("Invalid value of CPEInterfaceState: " + string(*en))
}

func (en *CPEInterfaceState) GetBSON() (interface{}, error) {
	switch *en {
	case CPEInterfaceStateACS:
		return "ACS", nil
	case CPEInterfaceStateCountryUpdate:
		return "COUNTRY_UPDATE", nil
	case CPEInterfaceStateDFS:
		return "DFS", nil
	case CPEInterfaceStateDisabled:
		return "DISABLED", nil
	case CPEInterfaceStateEnabled:
		return "ENABLED", nil
	case CPEInterfaceStateHtScan:
		return "HT_SCAN", nil
	case CPEInterfaceStateTerminated:
		return "TERMINATED", nil
	case CPEInterfaceStateUninitialized:
		return "UNINITIALIZED", nil
	case CPEInterfaceStateUnknown:
		return "UNKNOWN", nil
	}
	return nil, errors.New("Invalid value of CPEInterfaceState: " + string(*en))
}

func (en *CPEInterfaceState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ACS":
		*en = CPEInterfaceStateACS
		return nil
	case "COUNTRY_UPDATE":
		*en = CPEInterfaceStateCountryUpdate
		return nil
	case "DFS":
		*en = CPEInterfaceStateDFS
		return nil
	case "DISABLED":
		*en = CPEInterfaceStateDisabled
		return nil
	case "ENABLED":
		*en = CPEInterfaceStateEnabled
		return nil
	case "HT_SCAN":
		*en = CPEInterfaceStateHtScan
		return nil
	case "TERMINATED":
		*en = CPEInterfaceStateTerminated
		return nil
	case "UNINITIALIZED":
		*en = CPEInterfaceStateUninitialized
		return nil
	case "UNKNOWN":
		*en = CPEInterfaceStateUnknown
		return nil
	}
	return errors.New("Unknown CPEInterfaceState: " + s)
}

func (en *CPEInterfaceState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ACS":
		*en = CPEInterfaceStateACS
		return nil
	case "COUNTRY_UPDATE":
		*en = CPEInterfaceStateCountryUpdate
		return nil
	case "DFS":
		*en = CPEInterfaceStateDFS
		return nil
	case "DISABLED":
		*en = CPEInterfaceStateDisabled
		return nil
	case "ENABLED":
		*en = CPEInterfaceStateEnabled
		return nil
	case "HT_SCAN":
		*en = CPEInterfaceStateHtScan
		return nil
	case "TERMINATED":
		*en = CPEInterfaceStateTerminated
		return nil
	case "UNINITIALIZED":
		*en = CPEInterfaceStateUninitialized
		return nil
	case "UNKNOWN":
		*en = CPEInterfaceStateUnknown
		return nil
	}
	return errors.New("Unknown CPEInterfaceState: " + s)
}

type ClientStatPacketType string

const ClientStatPacketTypeInterim ClientStatPacketType = "interim"
const ClientStatPacketTypeInterimOld ClientStatPacketType = "Interim-Update"
const ClientStatPacketTypeOffOld ClientStatPacketType = "Accounting-Off"
const ClientStatPacketTypeOnOld ClientStatPacketType = "Accounting-On"
const ClientStatPacketTypeStart ClientStatPacketType = "start"
const ClientStatPacketTypeStartOld ClientStatPacketType = "Start"
const ClientStatPacketTypeStop ClientStatPacketType = "stop"
const ClientStatPacketTypeStopOld ClientStatPacketType = "Stop"

func (en ClientStatPacketType) GetPtr() *ClientStatPacketType { var v = en; return &v }

func (en ClientStatPacketType) String() string {
	switch en {
	case ClientStatPacketTypeInterim:
		return "interim"
	case ClientStatPacketTypeInterimOld:
		return "Interim-Update"
	case ClientStatPacketTypeOffOld:
		return "Accounting-Off"
	case ClientStatPacketTypeOnOld:
		return "Accounting-On"
	case ClientStatPacketTypeStart:
		return "start"
	case ClientStatPacketTypeStartOld:
		return "Start"
	case ClientStatPacketTypeStop:
		return "stop"
	case ClientStatPacketTypeStopOld:
		return "Stop"
	}
	panic(errors.New("Invalid value of ClientStatPacketType: " + string(en)))
}

func (en *ClientStatPacketType) MarshalJSON() ([]byte, error) {
	switch *en {
	case ClientStatPacketTypeInterim:
		return json.Marshal("interim")
	case ClientStatPacketTypeInterimOld:
		return json.Marshal("Interim-Update")
	case ClientStatPacketTypeOffOld:
		return json.Marshal("Accounting-Off")
	case ClientStatPacketTypeOnOld:
		return json.Marshal("Accounting-On")
	case ClientStatPacketTypeStart:
		return json.Marshal("start")
	case ClientStatPacketTypeStartOld:
		return json.Marshal("Start")
	case ClientStatPacketTypeStop:
		return json.Marshal("stop")
	case ClientStatPacketTypeStopOld:
		return json.Marshal("Stop")
	}
	return nil, errors.New("Invalid value of ClientStatPacketType: " + string(*en))
}

func (en *ClientStatPacketType) GetBSON() (interface{}, error) {
	switch *en {
	case ClientStatPacketTypeInterim:
		return "interim", nil
	case ClientStatPacketTypeInterimOld:
		return "Interim-Update", nil
	case ClientStatPacketTypeOffOld:
		return "Accounting-Off", nil
	case ClientStatPacketTypeOnOld:
		return "Accounting-On", nil
	case ClientStatPacketTypeStart:
		return "start", nil
	case ClientStatPacketTypeStartOld:
		return "Start", nil
	case ClientStatPacketTypeStop:
		return "stop", nil
	case ClientStatPacketTypeStopOld:
		return "Stop", nil
	}
	return nil, errors.New("Invalid value of ClientStatPacketType: " + string(*en))
}

func (en *ClientStatPacketType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "interim":
		*en = ClientStatPacketTypeInterim
		return nil
	case "Interim-Update":
		*en = ClientStatPacketTypeInterimOld
		return nil
	case "Accounting-Off":
		*en = ClientStatPacketTypeOffOld
		return nil
	case "Accounting-On":
		*en = ClientStatPacketTypeOnOld
		return nil
	case "start":
		*en = ClientStatPacketTypeStart
		return nil
	case "Start":
		*en = ClientStatPacketTypeStartOld
		return nil
	case "stop":
		*en = ClientStatPacketTypeStop
		return nil
	case "Stop":
		*en = ClientStatPacketTypeStopOld
		return nil
	}
	return errors.New("Unknown ClientStatPacketType: " + s)
}

func (en *ClientStatPacketType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "interim":
		*en = ClientStatPacketTypeInterim
		return nil
	case "Interim-Update":
		*en = ClientStatPacketTypeInterimOld
		return nil
	case "Accounting-Off":
		*en = ClientStatPacketTypeOffOld
		return nil
	case "Accounting-On":
		*en = ClientStatPacketTypeOnOld
		return nil
	case "start":
		*en = ClientStatPacketTypeStart
		return nil
	case "Start":
		*en = ClientStatPacketTypeStartOld
		return nil
	case "stop":
		*en = ClientStatPacketTypeStop
		return nil
	case "Stop":
		*en = ClientStatPacketTypeStopOld
		return nil
	}
	return errors.New("Unknown ClientStatPacketType: " + s)
}

type Coll string

const CollBSSStatInfo Coll = "bss_stat_info"
const CollBaseLocation Coll = "base_location"
const CollCPE Coll = "cpes"
const CollCPEConfigTemplate Coll = "config_rule"
const CollCPEMap Coll = "cpe_maps"
const CollCPEModel Coll = "cpe_model"
const CollCPESessionInfo Coll = "cpe_session_info"
const CollCPEStatInfo Coll = "cpe_stat_info"
const CollCPEStats Coll = "stats"
const CollCaptiveRedirect Coll = "captive_redirects"
const CollClientDistance Coll = "client_distance"
const CollClientRF Coll = "client_rf"
const CollClientSession Coll = "client_session_info"
const CollClientStatInfo Coll = "client_stat_info"
const CollClientStats Coll = "client_stats"
const CollClients Coll = "clients"
const CollCollLBSZones Coll = "lbs_zones"
const CollController Coll = "controllers"
const CollExternalAP Coll = "ext_access_points"
const CollHotspotProfiles Coll = "hotspot_profile"
const CollL2Chain Coll = "l2_chains"
const CollLBSCPEInfo Coll = "lbs_cpe_info"
const CollLBSClientCoord Coll = "lbs_client_coords"
const CollLBSClientData Coll = "lbs_client_data"
const CollLBSClientProbes Coll = "lbs_client_probes"
const CollLicenseLogDaily Coll = "license_log_daily"
const CollLocation Coll = "location"
const CollMonitorCPE Coll = "poll_cpe"
const CollMonitorEvents Coll = "events"
const CollMonitorRules Coll = "stat_event_rule"
const CollNTPConfig Coll = "ntp_config"
const CollOperation Coll = "operation"
const CollRADIUS Coll = "radius"
const CollRRMGroups Coll = "rrm_groups"
const CollRadarExport Coll = "radar_export"
const CollRadarExportResult Coll = "radar_export_result"
const CollRadarProbesRaw Coll = "radar_probes_raw"
const CollRadarProbesReal Coll = "radar_probes_real"
const CollRadarVisits Coll = "radar_visits"
const CollRadarVisitsFirst Coll = "radar_visits_first"
const CollRadarVisitsHour Coll = "radar_visits_hour"
const CollReportResult Coll = "report_results"
const CollResamplingConf Coll = "resampling_conf"
const CollSnmpCommunity Coll = "snmp_community"
const CollSnmpGeneral Coll = "snmp_general"
const CollSnmpHosts Coll = "snmp_hosts"
const CollSnmpUsers Coll = "snmp_users"
const CollSnmpUsersGroup Coll = "snmp_user_group"
const CollSnmpWalker Coll = "snmp_walker"
const CollStatReport Coll = "reports"
const CollTags Coll = "tags"
const CollUser Coll = "user"
const CollWLAN Coll = "wlans"
const CollWLANStatInfo Coll = "wlan_stat_info"
const CollWLCConfigs Coll = "wlc_configs"

func (en Coll) GetPtr() *Coll { var v = en; return &v }

func (en Coll) String() string {
	switch en {
	case CollBSSStatInfo:
		return "bss_stat_info"
	case CollBaseLocation:
		return "base_location"
	case CollCPE:
		return "cpes"
	case CollCPEConfigTemplate:
		return "config_rule"
	case CollCPEMap:
		return "cpe_maps"
	case CollCPEModel:
		return "cpe_model"
	case CollCPESessionInfo:
		return "cpe_session_info"
	case CollCPEStatInfo:
		return "cpe_stat_info"
	case CollCPEStats:
		return "stats"
	case CollCaptiveRedirect:
		return "captive_redirects"
	case CollClientDistance:
		return "client_distance"
	case CollClientRF:
		return "client_rf"
	case CollClientSession:
		return "client_session_info"
	case CollClientStatInfo:
		return "client_stat_info"
	case CollClientStats:
		return "client_stats"
	case CollClients:
		return "clients"
	case CollCollLBSZones:
		return "lbs_zones"
	case CollController:
		return "controllers"
	case CollExternalAP:
		return "ext_access_points"
	case CollHotspotProfiles:
		return "hotspot_profile"
	case CollL2Chain:
		return "l2_chains"
	case CollLBSCPEInfo:
		return "lbs_cpe_info"
	case CollLBSClientCoord:
		return "lbs_client_coords"
	case CollLBSClientData:
		return "lbs_client_data"
	case CollLBSClientProbes:
		return "lbs_client_probes"
	case CollLicenseLogDaily:
		return "license_log_daily"
	case CollLocation:
		return "location"
	case CollMonitorCPE:
		return "poll_cpe"
	case CollMonitorEvents:
		return "events"
	case CollMonitorRules:
		return "stat_event_rule"
	case CollNTPConfig:
		return "ntp_config"
	case CollOperation:
		return "operation"
	case CollRADIUS:
		return "radius"
	case CollRRMGroups:
		return "rrm_groups"
	case CollRadarExport:
		return "radar_export"
	case CollRadarExportResult:
		return "radar_export_result"
	case CollRadarProbesRaw:
		return "radar_probes_raw"
	case CollRadarProbesReal:
		return "radar_probes_real"
	case CollRadarVisits:
		return "radar_visits"
	case CollRadarVisitsFirst:
		return "radar_visits_first"
	case CollRadarVisitsHour:
		return "radar_visits_hour"
	case CollReportResult:
		return "report_results"
	case CollResamplingConf:
		return "resampling_conf"
	case CollSnmpCommunity:
		return "snmp_community"
	case CollSnmpGeneral:
		return "snmp_general"
	case CollSnmpHosts:
		return "snmp_hosts"
	case CollSnmpUsers:
		return "snmp_users"
	case CollSnmpUsersGroup:
		return "snmp_user_group"
	case CollSnmpWalker:
		return "snmp_walker"
	case CollStatReport:
		return "reports"
	case CollTags:
		return "tags"
	case CollUser:
		return "user"
	case CollWLAN:
		return "wlans"
	case CollWLANStatInfo:
		return "wlan_stat_info"
	case CollWLCConfigs:
		return "wlc_configs"
	}
	panic(errors.New("Invalid value of Coll: " + string(en)))
}

func (en *Coll) MarshalJSON() ([]byte, error) {
	switch *en {
	case CollBSSStatInfo:
		return json.Marshal("bss_stat_info")
	case CollBaseLocation:
		return json.Marshal("base_location")
	case CollCPE:
		return json.Marshal("cpes")
	case CollCPEConfigTemplate:
		return json.Marshal("config_rule")
	case CollCPEMap:
		return json.Marshal("cpe_maps")
	case CollCPEModel:
		return json.Marshal("cpe_model")
	case CollCPESessionInfo:
		return json.Marshal("cpe_session_info")
	case CollCPEStatInfo:
		return json.Marshal("cpe_stat_info")
	case CollCPEStats:
		return json.Marshal("stats")
	case CollCaptiveRedirect:
		return json.Marshal("captive_redirects")
	case CollClientDistance:
		return json.Marshal("client_distance")
	case CollClientRF:
		return json.Marshal("client_rf")
	case CollClientSession:
		return json.Marshal("client_session_info")
	case CollClientStatInfo:
		return json.Marshal("client_stat_info")
	case CollClientStats:
		return json.Marshal("client_stats")
	case CollClients:
		return json.Marshal("clients")
	case CollCollLBSZones:
		return json.Marshal("lbs_zones")
	case CollController:
		return json.Marshal("controllers")
	case CollExternalAP:
		return json.Marshal("ext_access_points")
	case CollHotspotProfiles:
		return json.Marshal("hotspot_profile")
	case CollL2Chain:
		return json.Marshal("l2_chains")
	case CollLBSCPEInfo:
		return json.Marshal("lbs_cpe_info")
	case CollLBSClientCoord:
		return json.Marshal("lbs_client_coords")
	case CollLBSClientData:
		return json.Marshal("lbs_client_data")
	case CollLBSClientProbes:
		return json.Marshal("lbs_client_probes")
	case CollLicenseLogDaily:
		return json.Marshal("license_log_daily")
	case CollLocation:
		return json.Marshal("location")
	case CollMonitorCPE:
		return json.Marshal("poll_cpe")
	case CollMonitorEvents:
		return json.Marshal("events")
	case CollMonitorRules:
		return json.Marshal("stat_event_rule")
	case CollNTPConfig:
		return json.Marshal("ntp_config")
	case CollOperation:
		return json.Marshal("operation")
	case CollRADIUS:
		return json.Marshal("radius")
	case CollRRMGroups:
		return json.Marshal("rrm_groups")
	case CollRadarExport:
		return json.Marshal("radar_export")
	case CollRadarExportResult:
		return json.Marshal("radar_export_result")
	case CollRadarProbesRaw:
		return json.Marshal("radar_probes_raw")
	case CollRadarProbesReal:
		return json.Marshal("radar_probes_real")
	case CollRadarVisits:
		return json.Marshal("radar_visits")
	case CollRadarVisitsFirst:
		return json.Marshal("radar_visits_first")
	case CollRadarVisitsHour:
		return json.Marshal("radar_visits_hour")
	case CollReportResult:
		return json.Marshal("report_results")
	case CollResamplingConf:
		return json.Marshal("resampling_conf")
	case CollSnmpCommunity:
		return json.Marshal("snmp_community")
	case CollSnmpGeneral:
		return json.Marshal("snmp_general")
	case CollSnmpHosts:
		return json.Marshal("snmp_hosts")
	case CollSnmpUsers:
		return json.Marshal("snmp_users")
	case CollSnmpUsersGroup:
		return json.Marshal("snmp_user_group")
	case CollSnmpWalker:
		return json.Marshal("snmp_walker")
	case CollStatReport:
		return json.Marshal("reports")
	case CollTags:
		return json.Marshal("tags")
	case CollUser:
		return json.Marshal("user")
	case CollWLAN:
		return json.Marshal("wlans")
	case CollWLANStatInfo:
		return json.Marshal("wlan_stat_info")
	case CollWLCConfigs:
		return json.Marshal("wlc_configs")
	}
	return nil, errors.New("Invalid value of Coll: " + string(*en))
}

func (en *Coll) GetBSON() (interface{}, error) {
	switch *en {
	case CollBSSStatInfo:
		return "bss_stat_info", nil
	case CollBaseLocation:
		return "base_location", nil
	case CollCPE:
		return "cpes", nil
	case CollCPEConfigTemplate:
		return "config_rule", nil
	case CollCPEMap:
		return "cpe_maps", nil
	case CollCPEModel:
		return "cpe_model", nil
	case CollCPESessionInfo:
		return "cpe_session_info", nil
	case CollCPEStatInfo:
		return "cpe_stat_info", nil
	case CollCPEStats:
		return "stats", nil
	case CollCaptiveRedirect:
		return "captive_redirects", nil
	case CollClientDistance:
		return "client_distance", nil
	case CollClientRF:
		return "client_rf", nil
	case CollClientSession:
		return "client_session_info", nil
	case CollClientStatInfo:
		return "client_stat_info", nil
	case CollClientStats:
		return "client_stats", nil
	case CollClients:
		return "clients", nil
	case CollCollLBSZones:
		return "lbs_zones", nil
	case CollController:
		return "controllers", nil
	case CollExternalAP:
		return "ext_access_points", nil
	case CollHotspotProfiles:
		return "hotspot_profile", nil
	case CollL2Chain:
		return "l2_chains", nil
	case CollLBSCPEInfo:
		return "lbs_cpe_info", nil
	case CollLBSClientCoord:
		return "lbs_client_coords", nil
	case CollLBSClientData:
		return "lbs_client_data", nil
	case CollLBSClientProbes:
		return "lbs_client_probes", nil
	case CollLicenseLogDaily:
		return "license_log_daily", nil
	case CollLocation:
		return "location", nil
	case CollMonitorCPE:
		return "poll_cpe", nil
	case CollMonitorEvents:
		return "events", nil
	case CollMonitorRules:
		return "stat_event_rule", nil
	case CollNTPConfig:
		return "ntp_config", nil
	case CollOperation:
		return "operation", nil
	case CollRADIUS:
		return "radius", nil
	case CollRRMGroups:
		return "rrm_groups", nil
	case CollRadarExport:
		return "radar_export", nil
	case CollRadarExportResult:
		return "radar_export_result", nil
	case CollRadarProbesRaw:
		return "radar_probes_raw", nil
	case CollRadarProbesReal:
		return "radar_probes_real", nil
	case CollRadarVisits:
		return "radar_visits", nil
	case CollRadarVisitsFirst:
		return "radar_visits_first", nil
	case CollRadarVisitsHour:
		return "radar_visits_hour", nil
	case CollReportResult:
		return "report_results", nil
	case CollResamplingConf:
		return "resampling_conf", nil
	case CollSnmpCommunity:
		return "snmp_community", nil
	case CollSnmpGeneral:
		return "snmp_general", nil
	case CollSnmpHosts:
		return "snmp_hosts", nil
	case CollSnmpUsers:
		return "snmp_users", nil
	case CollSnmpUsersGroup:
		return "snmp_user_group", nil
	case CollSnmpWalker:
		return "snmp_walker", nil
	case CollStatReport:
		return "reports", nil
	case CollTags:
		return "tags", nil
	case CollUser:
		return "user", nil
	case CollWLAN:
		return "wlans", nil
	case CollWLANStatInfo:
		return "wlan_stat_info", nil
	case CollWLCConfigs:
		return "wlc_configs", nil
	}
	return nil, errors.New("Invalid value of Coll: " + string(*en))
}

func (en *Coll) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "bss_stat_info":
		*en = CollBSSStatInfo
		return nil
	case "base_location":
		*en = CollBaseLocation
		return nil
	case "cpes":
		*en = CollCPE
		return nil
	case "config_rule":
		*en = CollCPEConfigTemplate
		return nil
	case "cpe_maps":
		*en = CollCPEMap
		return nil
	case "cpe_model":
		*en = CollCPEModel
		return nil
	case "cpe_session_info":
		*en = CollCPESessionInfo
		return nil
	case "cpe_stat_info":
		*en = CollCPEStatInfo
		return nil
	case "stats":
		*en = CollCPEStats
		return nil
	case "captive_redirects":
		*en = CollCaptiveRedirect
		return nil
	case "client_distance":
		*en = CollClientDistance
		return nil
	case "client_rf":
		*en = CollClientRF
		return nil
	case "client_session_info":
		*en = CollClientSession
		return nil
	case "client_stat_info":
		*en = CollClientStatInfo
		return nil
	case "client_stats":
		*en = CollClientStats
		return nil
	case "clients":
		*en = CollClients
		return nil
	case "lbs_zones":
		*en = CollCollLBSZones
		return nil
	case "controllers":
		*en = CollController
		return nil
	case "ext_access_points":
		*en = CollExternalAP
		return nil
	case "hotspot_profile":
		*en = CollHotspotProfiles
		return nil
	case "l2_chains":
		*en = CollL2Chain
		return nil
	case "lbs_cpe_info":
		*en = CollLBSCPEInfo
		return nil
	case "lbs_client_coords":
		*en = CollLBSClientCoord
		return nil
	case "lbs_client_data":
		*en = CollLBSClientData
		return nil
	case "lbs_client_probes":
		*en = CollLBSClientProbes
		return nil
	case "license_log_daily":
		*en = CollLicenseLogDaily
		return nil
	case "location":
		*en = CollLocation
		return nil
	case "poll_cpe":
		*en = CollMonitorCPE
		return nil
	case "events":
		*en = CollMonitorEvents
		return nil
	case "stat_event_rule":
		*en = CollMonitorRules
		return nil
	case "ntp_config":
		*en = CollNTPConfig
		return nil
	case "operation":
		*en = CollOperation
		return nil
	case "radius":
		*en = CollRADIUS
		return nil
	case "rrm_groups":
		*en = CollRRMGroups
		return nil
	case "radar_export":
		*en = CollRadarExport
		return nil
	case "radar_export_result":
		*en = CollRadarExportResult
		return nil
	case "radar_probes_raw":
		*en = CollRadarProbesRaw
		return nil
	case "radar_probes_real":
		*en = CollRadarProbesReal
		return nil
	case "radar_visits":
		*en = CollRadarVisits
		return nil
	case "radar_visits_first":
		*en = CollRadarVisitsFirst
		return nil
	case "radar_visits_hour":
		*en = CollRadarVisitsHour
		return nil
	case "report_results":
		*en = CollReportResult
		return nil
	case "resampling_conf":
		*en = CollResamplingConf
		return nil
	case "snmp_community":
		*en = CollSnmpCommunity
		return nil
	case "snmp_general":
		*en = CollSnmpGeneral
		return nil
	case "snmp_hosts":
		*en = CollSnmpHosts
		return nil
	case "snmp_users":
		*en = CollSnmpUsers
		return nil
	case "snmp_user_group":
		*en = CollSnmpUsersGroup
		return nil
	case "snmp_walker":
		*en = CollSnmpWalker
		return nil
	case "reports":
		*en = CollStatReport
		return nil
	case "tags":
		*en = CollTags
		return nil
	case "user":
		*en = CollUser
		return nil
	case "wlans":
		*en = CollWLAN
		return nil
	case "wlan_stat_info":
		*en = CollWLANStatInfo
		return nil
	case "wlc_configs":
		*en = CollWLCConfigs
		return nil
	}
	return errors.New("Unknown Coll: " + s)
}

func (en *Coll) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "bss_stat_info":
		*en = CollBSSStatInfo
		return nil
	case "base_location":
		*en = CollBaseLocation
		return nil
	case "cpes":
		*en = CollCPE
		return nil
	case "config_rule":
		*en = CollCPEConfigTemplate
		return nil
	case "cpe_maps":
		*en = CollCPEMap
		return nil
	case "cpe_model":
		*en = CollCPEModel
		return nil
	case "cpe_session_info":
		*en = CollCPESessionInfo
		return nil
	case "cpe_stat_info":
		*en = CollCPEStatInfo
		return nil
	case "stats":
		*en = CollCPEStats
		return nil
	case "captive_redirects":
		*en = CollCaptiveRedirect
		return nil
	case "client_distance":
		*en = CollClientDistance
		return nil
	case "client_rf":
		*en = CollClientRF
		return nil
	case "client_session_info":
		*en = CollClientSession
		return nil
	case "client_stat_info":
		*en = CollClientStatInfo
		return nil
	case "client_stats":
		*en = CollClientStats
		return nil
	case "clients":
		*en = CollClients
		return nil
	case "lbs_zones":
		*en = CollCollLBSZones
		return nil
	case "controllers":
		*en = CollController
		return nil
	case "ext_access_points":
		*en = CollExternalAP
		return nil
	case "hotspot_profile":
		*en = CollHotspotProfiles
		return nil
	case "l2_chains":
		*en = CollL2Chain
		return nil
	case "lbs_cpe_info":
		*en = CollLBSCPEInfo
		return nil
	case "lbs_client_coords":
		*en = CollLBSClientCoord
		return nil
	case "lbs_client_data":
		*en = CollLBSClientData
		return nil
	case "lbs_client_probes":
		*en = CollLBSClientProbes
		return nil
	case "license_log_daily":
		*en = CollLicenseLogDaily
		return nil
	case "location":
		*en = CollLocation
		return nil
	case "poll_cpe":
		*en = CollMonitorCPE
		return nil
	case "events":
		*en = CollMonitorEvents
		return nil
	case "stat_event_rule":
		*en = CollMonitorRules
		return nil
	case "ntp_config":
		*en = CollNTPConfig
		return nil
	case "operation":
		*en = CollOperation
		return nil
	case "radius":
		*en = CollRADIUS
		return nil
	case "rrm_groups":
		*en = CollRRMGroups
		return nil
	case "radar_export":
		*en = CollRadarExport
		return nil
	case "radar_export_result":
		*en = CollRadarExportResult
		return nil
	case "radar_probes_raw":
		*en = CollRadarProbesRaw
		return nil
	case "radar_probes_real":
		*en = CollRadarProbesReal
		return nil
	case "radar_visits":
		*en = CollRadarVisits
		return nil
	case "radar_visits_first":
		*en = CollRadarVisitsFirst
		return nil
	case "radar_visits_hour":
		*en = CollRadarVisitsHour
		return nil
	case "report_results":
		*en = CollReportResult
		return nil
	case "resampling_conf":
		*en = CollResamplingConf
		return nil
	case "snmp_community":
		*en = CollSnmpCommunity
		return nil
	case "snmp_general":
		*en = CollSnmpGeneral
		return nil
	case "snmp_hosts":
		*en = CollSnmpHosts
		return nil
	case "snmp_users":
		*en = CollSnmpUsers
		return nil
	case "snmp_user_group":
		*en = CollSnmpUsersGroup
		return nil
	case "snmp_walker":
		*en = CollSnmpWalker
		return nil
	case "reports":
		*en = CollStatReport
		return nil
	case "tags":
		*en = CollTags
		return nil
	case "user":
		*en = CollUser
		return nil
	case "wlans":
		*en = CollWLAN
		return nil
	case "wlan_stat_info":
		*en = CollWLANStatInfo
		return nil
	case "wlc_configs":
		*en = CollWLCConfigs
		return nil
	}
	return errors.New("Unknown Coll: " + s)
}

type ConfigurationStatus string

const ConfigurationStatusDontUse1 ConfigurationStatus = "pending"
const ConfigurationStatusDontUse2 ConfigurationStatus = "error"
const ConfigurationStatusEmpty ConfigurationStatus = "empty"
const ConfigurationStatusOK ConfigurationStatus = "ok"
const ConfigurationStatusOffline ConfigurationStatus = "offline"
const ConfigurationStatusRebooting ConfigurationStatus = "rebooting"
const ConfigurationStatusUpdating ConfigurationStatus = "updating"
const ConfigurationStatusUpgrading ConfigurationStatus = "upgrading"

func (en ConfigurationStatus) GetPtr() *ConfigurationStatus { var v = en; return &v }

func (en ConfigurationStatus) String() string {
	switch en {
	case ConfigurationStatusDontUse1:
		return "pending"
	case ConfigurationStatusDontUse2:
		return "error"
	case ConfigurationStatusEmpty:
		return "empty"
	case ConfigurationStatusOK:
		return "ok"
	case ConfigurationStatusOffline:
		return "offline"
	case ConfigurationStatusRebooting:
		return "rebooting"
	case ConfigurationStatusUpdating:
		return "updating"
	case ConfigurationStatusUpgrading:
		return "upgrading"
	}
	if len(en) == 0 {
		return "empty"
	}
	panic(errors.New("Invalid value of ConfigurationStatus: " + string(en)))
}

func (en *ConfigurationStatus) MarshalJSON() ([]byte, error) {
	switch *en {
	case ConfigurationStatusDontUse1:
		return json.Marshal("pending")
	case ConfigurationStatusDontUse2:
		return json.Marshal("error")
	case ConfigurationStatusEmpty:
		return json.Marshal("empty")
	case ConfigurationStatusOK:
		return json.Marshal("ok")
	case ConfigurationStatusOffline:
		return json.Marshal("offline")
	case ConfigurationStatusRebooting:
		return json.Marshal("rebooting")
	case ConfigurationStatusUpdating:
		return json.Marshal("updating")
	case ConfigurationStatusUpgrading:
		return json.Marshal("upgrading")
	}
	if len(*en) == 0 {
		return json.Marshal("empty")
	}
	return nil, errors.New("Invalid value of ConfigurationStatus: " + string(*en))
}

func (en *ConfigurationStatus) GetBSON() (interface{}, error) {
	switch *en {
	case ConfigurationStatusDontUse1:
		return "pending", nil
	case ConfigurationStatusDontUse2:
		return "error", nil
	case ConfigurationStatusEmpty:
		return "empty", nil
	case ConfigurationStatusOK:
		return "ok", nil
	case ConfigurationStatusOffline:
		return "offline", nil
	case ConfigurationStatusRebooting:
		return "rebooting", nil
	case ConfigurationStatusUpdating:
		return "updating", nil
	case ConfigurationStatusUpgrading:
		return "upgrading", nil
	}
	if len(*en) == 0 {
		return "empty", nil
	}
	return nil, errors.New("Invalid value of ConfigurationStatus: " + string(*en))
}

func (en *ConfigurationStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "pending":
		*en = ConfigurationStatusDontUse1
		return nil
	case "error":
		*en = ConfigurationStatusDontUse2
		return nil
	case "empty":
		*en = ConfigurationStatusEmpty
		return nil
	case "ok":
		*en = ConfigurationStatusOK
		return nil
	case "offline":
		*en = ConfigurationStatusOffline
		return nil
	case "rebooting":
		*en = ConfigurationStatusRebooting
		return nil
	case "updating":
		*en = ConfigurationStatusUpdating
		return nil
	case "upgrading":
		*en = ConfigurationStatusUpgrading
		return nil
	}
	if len(s) == 0 {
		*en = ConfigurationStatusEmpty
		return nil
	}
	return errors.New("Unknown ConfigurationStatus: " + s)
}

func (en *ConfigurationStatus) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "pending":
		*en = ConfigurationStatusDontUse1
		return nil
	case "error":
		*en = ConfigurationStatusDontUse2
		return nil
	case "empty":
		*en = ConfigurationStatusEmpty
		return nil
	case "ok":
		*en = ConfigurationStatusOK
		return nil
	case "offline":
		*en = ConfigurationStatusOffline
		return nil
	case "rebooting":
		*en = ConfigurationStatusRebooting
		return nil
	case "updating":
		*en = ConfigurationStatusUpdating
		return nil
	case "upgrading":
		*en = ConfigurationStatusUpgrading
		return nil
	}
	if len(s) == 0 {
		*en = ConfigurationStatusEmpty
		return nil
	}
	return errors.New("Unknown ConfigurationStatus: " + s)
}

type ConnectionModeType string

const ConnectionModeTypeModeAC ConnectionModeType = "ac"
const ConnectionModeTypeModeAX ConnectionModeType = "ax"
const ConnectionModeTypeModeLegacy ConnectionModeType = "legacy"
const ConnectionModeTypeModeN ConnectionModeType = "n"

func (en ConnectionModeType) GetPtr() *ConnectionModeType { var v = en; return &v }

func (en ConnectionModeType) String() string {
	switch en {
	case ConnectionModeTypeModeAC:
		return "ac"
	case ConnectionModeTypeModeAX:
		return "ax"
	case ConnectionModeTypeModeLegacy:
		return "legacy"
	case ConnectionModeTypeModeN:
		return "n"
	}
	if len(en) == 0 {
		return "legacy"
	}
	panic(errors.New("Invalid value of ConnectionModeType: " + string(en)))
}

func (en *ConnectionModeType) MarshalJSON() ([]byte, error) {
	switch *en {
	case ConnectionModeTypeModeAC:
		return json.Marshal("ac")
	case ConnectionModeTypeModeAX:
		return json.Marshal("ax")
	case ConnectionModeTypeModeLegacy:
		return json.Marshal("legacy")
	case ConnectionModeTypeModeN:
		return json.Marshal("n")
	}
	if len(*en) == 0 {
		return json.Marshal("legacy")
	}
	return nil, errors.New("Invalid value of ConnectionModeType: " + string(*en))
}

func (en *ConnectionModeType) GetBSON() (interface{}, error) {
	switch *en {
	case ConnectionModeTypeModeAC:
		return "ac", nil
	case ConnectionModeTypeModeAX:
		return "ax", nil
	case ConnectionModeTypeModeLegacy:
		return "legacy", nil
	case ConnectionModeTypeModeN:
		return "n", nil
	}
	if len(*en) == 0 {
		return "legacy", nil
	}
	return nil, errors.New("Invalid value of ConnectionModeType: " + string(*en))
}

func (en *ConnectionModeType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ac":
		*en = ConnectionModeTypeModeAC
		return nil
	case "ax":
		*en = ConnectionModeTypeModeAX
		return nil
	case "legacy":
		*en = ConnectionModeTypeModeLegacy
		return nil
	case "n":
		*en = ConnectionModeTypeModeN
		return nil
	}
	if len(s) == 0 {
		*en = ConnectionModeTypeModeLegacy
		return nil
	}
	return errors.New("Unknown ConnectionModeType: " + s)
}

func (en *ConnectionModeType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ac":
		*en = ConnectionModeTypeModeAC
		return nil
	case "ax":
		*en = ConnectionModeTypeModeAX
		return nil
	case "legacy":
		*en = ConnectionModeTypeModeLegacy
		return nil
	case "n":
		*en = ConnectionModeTypeModeN
		return nil
	}
	if len(s) == 0 {
		*en = ConnectionModeTypeModeLegacy
		return nil
	}
	return errors.New("Unknown ConnectionModeType: " + s)
}

type ControllerStatusType string

const ControllerStatusTypeConnected ControllerStatusType = "connected"
const ControllerStatusTypeDisconnected ControllerStatusType = "disconnected"
const ControllerStatusTypeEmpty ControllerStatusType = "empty"
const ControllerStatusTypeError ControllerStatusType = "error"
const ControllerStatusTypeProvisioning ControllerStatusType = "provision"
const ControllerStatusTypeUpdating ControllerStatusType = "updating"

func (en ControllerStatusType) GetPtr() *ControllerStatusType { var v = en; return &v }

func (en ControllerStatusType) String() string {
	switch en {
	case ControllerStatusTypeConnected:
		return "connected"
	case ControllerStatusTypeDisconnected:
		return "disconnected"
	case ControllerStatusTypeEmpty:
		return "empty"
	case ControllerStatusTypeError:
		return "error"
	case ControllerStatusTypeProvisioning:
		return "provision"
	case ControllerStatusTypeUpdating:
		return "updating"
	}
	if len(en) == 0 {
		return "empty"
	}
	panic(errors.New("Invalid value of ControllerStatusType: " + string(en)))
}

func (en *ControllerStatusType) MarshalJSON() ([]byte, error) {
	switch *en {
	case ControllerStatusTypeConnected:
		return json.Marshal("connected")
	case ControllerStatusTypeDisconnected:
		return json.Marshal("disconnected")
	case ControllerStatusTypeEmpty:
		return json.Marshal("empty")
	case ControllerStatusTypeError:
		return json.Marshal("error")
	case ControllerStatusTypeProvisioning:
		return json.Marshal("provision")
	case ControllerStatusTypeUpdating:
		return json.Marshal("updating")
	}
	if len(*en) == 0 {
		return json.Marshal("empty")
	}
	return nil, errors.New("Invalid value of ControllerStatusType: " + string(*en))
}

func (en *ControllerStatusType) GetBSON() (interface{}, error) {
	switch *en {
	case ControllerStatusTypeConnected:
		return "connected", nil
	case ControllerStatusTypeDisconnected:
		return "disconnected", nil
	case ControllerStatusTypeEmpty:
		return "empty", nil
	case ControllerStatusTypeError:
		return "error", nil
	case ControllerStatusTypeProvisioning:
		return "provision", nil
	case ControllerStatusTypeUpdating:
		return "updating", nil
	}
	if len(*en) == 0 {
		return "empty", nil
	}
	return nil, errors.New("Invalid value of ControllerStatusType: " + string(*en))
}

func (en *ControllerStatusType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "connected":
		*en = ControllerStatusTypeConnected
		return nil
	case "disconnected":
		*en = ControllerStatusTypeDisconnected
		return nil
	case "empty":
		*en = ControllerStatusTypeEmpty
		return nil
	case "error":
		*en = ControllerStatusTypeError
		return nil
	case "provision":
		*en = ControllerStatusTypeProvisioning
		return nil
	case "updating":
		*en = ControllerStatusTypeUpdating
		return nil
	}
	if len(s) == 0 {
		*en = ControllerStatusTypeEmpty
		return nil
	}
	return errors.New("Unknown ControllerStatusType: " + s)
}

func (en *ControllerStatusType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "connected":
		*en = ControllerStatusTypeConnected
		return nil
	case "disconnected":
		*en = ControllerStatusTypeDisconnected
		return nil
	case "empty":
		*en = ControllerStatusTypeEmpty
		return nil
	case "error":
		*en = ControllerStatusTypeError
		return nil
	case "provision":
		*en = ControllerStatusTypeProvisioning
		return nil
	case "updating":
		*en = ControllerStatusTypeUpdating
		return nil
	}
	if len(s) == 0 {
		*en = ControllerStatusTypeEmpty
		return nil
	}
	return errors.New("Unknown ControllerStatusType: " + s)
}

type FirewallDirection string

const FirewallDirectionAny FirewallDirection = "ANY"
const FirewallDirectionIn FirewallDirection = "IN"
const FirewallDirectionOut FirewallDirection = "OUT"

func (en FirewallDirection) GetPtr() *FirewallDirection { var v = en; return &v }

func (en FirewallDirection) String() string {
	switch en {
	case FirewallDirectionAny:
		return "ANY"
	case FirewallDirectionIn:
		return "IN"
	case FirewallDirectionOut:
		return "OUT"
	}
	if len(en) == 0 {
		return "ANY"
	}
	panic(errors.New("Invalid value of FirewallDirection: " + string(en)))
}

func (en *FirewallDirection) MarshalJSON() ([]byte, error) {
	switch *en {
	case FirewallDirectionAny:
		return json.Marshal("ANY")
	case FirewallDirectionIn:
		return json.Marshal("IN")
	case FirewallDirectionOut:
		return json.Marshal("OUT")
	}
	if len(*en) == 0 {
		return json.Marshal("ANY")
	}
	return nil, errors.New("Invalid value of FirewallDirection: " + string(*en))
}

func (en *FirewallDirection) GetBSON() (interface{}, error) {
	switch *en {
	case FirewallDirectionAny:
		return "ANY", nil
	case FirewallDirectionIn:
		return "IN", nil
	case FirewallDirectionOut:
		return "OUT", nil
	}
	if len(*en) == 0 {
		return "ANY", nil
	}
	return nil, errors.New("Invalid value of FirewallDirection: " + string(*en))
}

func (en *FirewallDirection) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ANY":
		*en = FirewallDirectionAny
		return nil
	case "IN":
		*en = FirewallDirectionIn
		return nil
	case "OUT":
		*en = FirewallDirectionOut
		return nil
	}
	if len(s) == 0 {
		*en = FirewallDirectionAny
		return nil
	}
	return errors.New("Unknown FirewallDirection: " + s)
}

func (en *FirewallDirection) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ANY":
		*en = FirewallDirectionAny
		return nil
	case "IN":
		*en = FirewallDirectionIn
		return nil
	case "OUT":
		*en = FirewallDirectionOut
		return nil
	}
	if len(s) == 0 {
		*en = FirewallDirectionAny
		return nil
	}
	return errors.New("Unknown FirewallDirection: " + s)
}

type FirewallPolicy string

const FirewallPolicyAccept FirewallPolicy = "ACCEPT"
const FirewallPolicyDrop FirewallPolicy = "DROP"
const FirewallPolicyEmpty FirewallPolicy = ""
const FirewallPolicyReturn FirewallPolicy = "RETURN"

func (en FirewallPolicy) GetPtr() *FirewallPolicy { var v = en; return &v }

func (en FirewallPolicy) String() string {
	switch en {
	case FirewallPolicyAccept:
		return "ACCEPT"
	case FirewallPolicyDrop:
		return "DROP"
	case FirewallPolicyEmpty:
		return ""
	case FirewallPolicyReturn:
		return "RETURN"
	}
	if len(en) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of FirewallPolicy: " + string(en)))
}

func (en *FirewallPolicy) MarshalJSON() ([]byte, error) {
	switch *en {
	case FirewallPolicyAccept:
		return json.Marshal("ACCEPT")
	case FirewallPolicyDrop:
		return json.Marshal("DROP")
	case FirewallPolicyEmpty:
		return json.Marshal("")
	case FirewallPolicyReturn:
		return json.Marshal("RETURN")
	}
	if len(*en) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of FirewallPolicy: " + string(*en))
}

func (en *FirewallPolicy) GetBSON() (interface{}, error) {
	switch *en {
	case FirewallPolicyAccept:
		return "ACCEPT", nil
	case FirewallPolicyDrop:
		return "DROP", nil
	case FirewallPolicyEmpty:
		return "", nil
	case FirewallPolicyReturn:
		return "RETURN", nil
	}
	if len(*en) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of FirewallPolicy: " + string(*en))
}

func (en *FirewallPolicy) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ACCEPT":
		*en = FirewallPolicyAccept
		return nil
	case "DROP":
		*en = FirewallPolicyDrop
		return nil
	case "":
		*en = FirewallPolicyEmpty
		return nil
	case "RETURN":
		*en = FirewallPolicyReturn
		return nil
	}
	if len(s) == 0 {
		*en = FirewallPolicyEmpty
		return nil
	}
	return errors.New("Unknown FirewallPolicy: " + s)
}

func (en *FirewallPolicy) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ACCEPT":
		*en = FirewallPolicyAccept
		return nil
	case "DROP":
		*en = FirewallPolicyDrop
		return nil
	case "":
		*en = FirewallPolicyEmpty
		return nil
	case "RETURN":
		*en = FirewallPolicyReturn
		return nil
	}
	if len(s) == 0 {
		*en = FirewallPolicyEmpty
		return nil
	}
	return errors.New("Unknown FirewallPolicy: " + s)
}

type FirmwareUpdateMode string

const FirmwareUpdateModeCheck FirmwareUpdateMode = "check"
const FirmwareUpdateModeOff FirmwareUpdateMode = "off"
const FirmwareUpdateModeOn FirmwareUpdateMode = "on"

func (en FirmwareUpdateMode) GetPtr() *FirmwareUpdateMode { var v = en; return &v }

func (en FirmwareUpdateMode) String() string {
	switch en {
	case FirmwareUpdateModeCheck:
		return "check"
	case FirmwareUpdateModeOff:
		return "off"
	case FirmwareUpdateModeOn:
		return "on"
	}
	if len(en) == 0 {
		return "check"
	}
	panic(errors.New("Invalid value of FirmwareUpdateMode: " + string(en)))
}

func (en *FirmwareUpdateMode) MarshalJSON() ([]byte, error) {
	switch *en {
	case FirmwareUpdateModeCheck:
		return json.Marshal("check")
	case FirmwareUpdateModeOff:
		return json.Marshal("off")
	case FirmwareUpdateModeOn:
		return json.Marshal("on")
	}
	if len(*en) == 0 {
		return json.Marshal("check")
	}
	return nil, errors.New("Invalid value of FirmwareUpdateMode: " + string(*en))
}

func (en *FirmwareUpdateMode) GetBSON() (interface{}, error) {
	switch *en {
	case FirmwareUpdateModeCheck:
		return "check", nil
	case FirmwareUpdateModeOff:
		return "off", nil
	case FirmwareUpdateModeOn:
		return "on", nil
	}
	if len(*en) == 0 {
		return "check", nil
	}
	return nil, errors.New("Invalid value of FirmwareUpdateMode: " + string(*en))
}

func (en *FirmwareUpdateMode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "check":
		*en = FirmwareUpdateModeCheck
		return nil
	case "off":
		*en = FirmwareUpdateModeOff
		return nil
	case "on":
		*en = FirmwareUpdateModeOn
		return nil
	}
	if len(s) == 0 {
		*en = FirmwareUpdateModeCheck
		return nil
	}
	return errors.New("Unknown FirmwareUpdateMode: " + s)
}

func (en *FirmwareUpdateMode) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "check":
		*en = FirmwareUpdateModeCheck
		return nil
	case "off":
		*en = FirmwareUpdateModeOff
		return nil
	case "on":
		*en = FirmwareUpdateModeOn
		return nil
	}
	if len(s) == 0 {
		*en = FirmwareUpdateModeCheck
		return nil
	}
	return errors.New("Unknown FirmwareUpdateMode: " + s)
}

type L3Protocol string

const L3ProtocolEmpty L3Protocol = ""
const L3ProtocolIP L3Protocol = "ip"
const L3ProtocolIPv4 L3Protocol = "ipv4"
const L3ProtocolIPv6 L3Protocol = "ipv6"

func (en L3Protocol) GetPtr() *L3Protocol { var v = en; return &v }

func (en L3Protocol) String() string {
	switch en {
	case L3ProtocolEmpty:
		return ""
	case L3ProtocolIP:
		return "ip"
	case L3ProtocolIPv4:
		return "ipv4"
	case L3ProtocolIPv6:
		return "ipv6"
	}
	if len(en) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of L3Protocol: " + string(en)))
}

func (en *L3Protocol) MarshalJSON() ([]byte, error) {
	switch *en {
	case L3ProtocolEmpty:
		return json.Marshal("")
	case L3ProtocolIP:
		return json.Marshal("ip")
	case L3ProtocolIPv4:
		return json.Marshal("ipv4")
	case L3ProtocolIPv6:
		return json.Marshal("ipv6")
	}
	if len(*en) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of L3Protocol: " + string(*en))
}

func (en *L3Protocol) GetBSON() (interface{}, error) {
	switch *en {
	case L3ProtocolEmpty:
		return "", nil
	case L3ProtocolIP:
		return "ip", nil
	case L3ProtocolIPv4:
		return "ipv4", nil
	case L3ProtocolIPv6:
		return "ipv6", nil
	}
	if len(*en) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of L3Protocol: " + string(*en))
}

func (en *L3Protocol) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "":
		*en = L3ProtocolEmpty
		return nil
	case "ip":
		*en = L3ProtocolIP
		return nil
	case "ipv4":
		*en = L3ProtocolIPv4
		return nil
	case "ipv6":
		*en = L3ProtocolIPv6
		return nil
	}
	if len(s) == 0 {
		*en = L3ProtocolEmpty
		return nil
	}
	return errors.New("Unknown L3Protocol: " + s)
}

func (en *L3Protocol) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "":
		*en = L3ProtocolEmpty
		return nil
	case "ip":
		*en = L3ProtocolIP
		return nil
	case "ipv4":
		*en = L3ProtocolIPv4
		return nil
	case "ipv6":
		*en = L3ProtocolIPv6
		return nil
	}
	if len(s) == 0 {
		*en = L3ProtocolEmpty
		return nil
	}
	return errors.New("Unknown L3Protocol: " + s)
}

type L4Protocol string

const L4ProtocolEmpty L4Protocol = ""
const L4ProtocolTCP L4Protocol = "TCP"
const L4ProtocolUDP L4Protocol = "UDP"

func (en L4Protocol) GetPtr() *L4Protocol { var v = en; return &v }

func (en L4Protocol) String() string {
	switch en {
	case L4ProtocolEmpty:
		return ""
	case L4ProtocolTCP:
		return "TCP"
	case L4ProtocolUDP:
		return "UDP"
	}
	if len(en) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of L4Protocol: " + string(en)))
}

func (en *L4Protocol) MarshalJSON() ([]byte, error) {
	switch *en {
	case L4ProtocolEmpty:
		return json.Marshal("")
	case L4ProtocolTCP:
		return json.Marshal("TCP")
	case L4ProtocolUDP:
		return json.Marshal("UDP")
	}
	if len(*en) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of L4Protocol: " + string(*en))
}

func (en *L4Protocol) GetBSON() (interface{}, error) {
	switch *en {
	case L4ProtocolEmpty:
		return "", nil
	case L4ProtocolTCP:
		return "TCP", nil
	case L4ProtocolUDP:
		return "UDP", nil
	}
	if len(*en) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of L4Protocol: " + string(*en))
}

func (en *L4Protocol) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "":
		*en = L4ProtocolEmpty
		return nil
	case "TCP":
		*en = L4ProtocolTCP
		return nil
	case "UDP":
		*en = L4ProtocolUDP
		return nil
	}
	if len(s) == 0 {
		*en = L4ProtocolEmpty
		return nil
	}
	return errors.New("Unknown L4Protocol: " + s)
}

func (en *L4Protocol) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "":
		*en = L4ProtocolEmpty
		return nil
	case "TCP":
		*en = L4ProtocolTCP
		return nil
	case "UDP":
		*en = L4ProtocolUDP
		return nil
	}
	if len(s) == 0 {
		*en = L4ProtocolEmpty
		return nil
	}
	return errors.New("Unknown L4Protocol: " + s)
}

type MCSRequire string

const MCSRequireHT MCSRequire = "ht"
const MCSRequireOff MCSRequire = "off"
const MCSRequireVHT MCSRequire = "vht"

func (en MCSRequire) GetPtr() *MCSRequire { var v = en; return &v }

func (en MCSRequire) String() string {
	switch en {
	case MCSRequireHT:
		return "ht"
	case MCSRequireOff:
		return "off"
	case MCSRequireVHT:
		return "vht"
	}
	if len(en) == 0 {
		return "off"
	}
	panic(errors.New("Invalid value of MCSRequire: " + string(en)))
}

func (en *MCSRequire) MarshalJSON() ([]byte, error) {
	switch *en {
	case MCSRequireHT:
		return json.Marshal("ht")
	case MCSRequireOff:
		return json.Marshal("off")
	case MCSRequireVHT:
		return json.Marshal("vht")
	}
	if len(*en) == 0 {
		return json.Marshal("off")
	}
	return nil, errors.New("Invalid value of MCSRequire: " + string(*en))
}

func (en *MCSRequire) GetBSON() (interface{}, error) {
	switch *en {
	case MCSRequireHT:
		return "ht", nil
	case MCSRequireOff:
		return "off", nil
	case MCSRequireVHT:
		return "vht", nil
	}
	if len(*en) == 0 {
		return "off", nil
	}
	return nil, errors.New("Invalid value of MCSRequire: " + string(*en))
}

func (en *MCSRequire) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ht":
		*en = MCSRequireHT
		return nil
	case "off":
		*en = MCSRequireOff
		return nil
	case "vht":
		*en = MCSRequireVHT
		return nil
	}
	if len(s) == 0 {
		*en = MCSRequireOff
		return nil
	}
	return errors.New("Unknown MCSRequire: " + s)
}

func (en *MCSRequire) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ht":
		*en = MCSRequireHT
		return nil
	case "off":
		*en = MCSRequireOff
		return nil
	case "vht":
		*en = MCSRequireVHT
		return nil
	}
	if len(s) == 0 {
		*en = MCSRequireOff
		return nil
	}
	return errors.New("Unknown MCSRequire: " + s)
}

type MacFilterType string

const MacFilterTypeBlackList MacFilterType = "BlackList"
const MacFilterTypeNone MacFilterType = "None"
const MacFilterTypeWhiteList MacFilterType = "WhiteList"

func (en MacFilterType) GetPtr() *MacFilterType { var v = en; return &v }

func (en MacFilterType) String() string {
	switch en {
	case MacFilterTypeBlackList:
		return "BlackList"
	case MacFilterTypeNone:
		return "None"
	case MacFilterTypeWhiteList:
		return "WhiteList"
	}
	if len(en) == 0 {
		return "None"
	}
	panic(errors.New("Invalid value of MacFilterType: " + string(en)))
}

func (en *MacFilterType) MarshalJSON() ([]byte, error) {
	switch *en {
	case MacFilterTypeBlackList:
		return json.Marshal("BlackList")
	case MacFilterTypeNone:
		return json.Marshal("None")
	case MacFilterTypeWhiteList:
		return json.Marshal("WhiteList")
	}
	if len(*en) == 0 {
		return json.Marshal("None")
	}
	return nil, errors.New("Invalid value of MacFilterType: " + string(*en))
}

func (en *MacFilterType) GetBSON() (interface{}, error) {
	switch *en {
	case MacFilterTypeBlackList:
		return "BlackList", nil
	case MacFilterTypeNone:
		return "None", nil
	case MacFilterTypeWhiteList:
		return "WhiteList", nil
	}
	if len(*en) == 0 {
		return "None", nil
	}
	return nil, errors.New("Invalid value of MacFilterType: " + string(*en))
}

func (en *MacFilterType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "BlackList":
		*en = MacFilterTypeBlackList
		return nil
	case "None":
		*en = MacFilterTypeNone
		return nil
	case "WhiteList":
		*en = MacFilterTypeWhiteList
		return nil
	}
	if len(s) == 0 {
		*en = MacFilterTypeNone
		return nil
	}
	return errors.New("Unknown MacFilterType: " + s)
}

func (en *MacFilterType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "BlackList":
		*en = MacFilterTypeBlackList
		return nil
	case "None":
		*en = MacFilterTypeNone
		return nil
	case "WhiteList":
		*en = MacFilterTypeWhiteList
		return nil
	}
	if len(s) == 0 {
		*en = MacFilterTypeNone
		return nil
	}
	return errors.New("Unknown MacFilterType: " + s)
}

type Model string

const ModelBSSRating Model = "bss_rating"
const ModelBaseLocation Model = "base_location"
const ModelCPE Model = "cpe"
const ModelCPEAggStat Model = "cpe_agg_stats"
const ModelCPEConfigTemplates Model = "cpe_config_templates"
const ModelCPECurrentState Model = "cpe_current_state"
const ModelCPEDashboard Model = "cpe_dashboard"
const ModelCPEEvents Model = "cpe_events"
const ModelCPEFWUpgrade Model = "cpe_fw_upgrade"
const ModelCPEMap Model = "mapobjects"
const ModelCPEMapCPEList Model = "cpe_map_cpe_list"
const ModelCPEMapPosition Model = "cpe_map_position"
const ModelCPEMapScale Model = "cpe_map_scale"
const ModelCPEModels Model = "cpe_models"
const ModelCPERating Model = "cpe_rating"
const ModelCPERemote Model = "cpe_remote"
const ModelCPEReport Model = "cpe_report"
const ModelCPEScanLast Model = "cpe_scan_last"
const ModelCPESession Model = "cpe_session"
const ModelCPESessionAvg Model = "cpe_session_avg"
const ModelCPEStats Model = "stat"
const ModelCPETags Model = "cpe_tags"
const ModelCPETimeline Model = "cpe_timeline"
const ModelCPEWLANAggStat Model = "cpe_wlan_agg_stats"
const ModelCaptiveRedirect Model = "captive-redirect"
const ModelClient Model = "client"
const ModelClientCurrentState Model = "client_current_state"
const ModelClientDashboard Model = "client_dashboard"
const ModelClientDistance Model = "client_distance"
const ModelClientEvents Model = "client_events"
const ModelClientList Model = "client-list"
const ModelClientRF Model = "client_rf"
const ModelClientRating Model = "client_rating"
const ModelClientSession Model = "client_session"
const ModelClientSessionAvg Model = "client_session_avg"
const ModelClientStatExt Model = "client-stat-ext"
const ModelClientStats Model = "client-stat"
const ModelClientTimeline Model = "client_timeline"
const ModelCommonDashboard Model = "common_dashboard"
const ModelCommonEvents Model = "common_events"
const ModelCompactDashboard Model = "compact_dashboard"
const ModelConfiguration Model = "configuration"
const ModelController Model = "controller"
const ModelDitExport Model = "dit_export"
const ModelExternalAP Model = "ext_access_point"
const ModelHeatMapClients Model = "heatmap_clients"
const ModelHeatMapLBSClients Model = "heatmap_lbs_clients"
const ModelHeatMapTraffic Model = "heatmap_traffic"
const ModelHost Model = "host"
const ModelHotspotProfile Model = "hotspot-profile"
const ModelL2Chain Model = "l2-chain"
const ModelLBSCPEInfo Model = "lbs-cpe-info"
const ModelLBSClientCoord Model = "lbs-client-coords"
const ModelLBSClientData Model = "lbs-client-data"
const ModelLBSClientProbes Model = "lbs-client-probes"
const ModelLBSClientTrack Model = "lbs-client-track"
const ModelLBSZones Model = "lbs_zones"
const ModelLicense Model = "license"
const ModelLicenseLogDaily Model = "license_log_daily"
const ModelLoaderSettings Model = "settings_loader"
const ModelLoaderTask Model = "periodic_task"
const ModelLocation Model = "location"
const ModelMoniEvents Model = "monitor_events"
const ModelMonitorCPE Model = "poll-cpe"
const ModelMonitorEvents Model = "event-stat-rule-violation"
const ModelMonitorRules Model = "stat-event-rule"
const ModelOperation Model = "operation"
const ModelRADIUS Model = "radius"
const ModelRRMGroups Model = "rrm-groups"
const ModelRadarExport Model = "radar_exports"
const ModelResampling Model = "resampling"
const ModelSnmpWalker Model = "snmp_walker"
const ModelStatReport Model = "report"
const ModelSystemCPERateStats Model = "system_cpe_rate_stats"
const ModelSystemCPEStats Model = "system_cpe_stats"
const ModelSystemDashboard Model = "system_dashboard"
const ModelSystemEvents Model = "system_events"
const ModelSystemRating Model = "system_rating"
const ModelSystemTimeline Model = "system_timeline"
const ModelUser Model = "user"
const ModelWLAN Model = "wlan"
const ModelWLANAggStat Model = "wlan_agg_stats"
const ModelWLANCPEAggStat Model = "wlan_cpe_agg_stats"
const ModelWLANCPETimeline Model = "wlan_cpe_timeline"
const ModelWLANClientRating Model = "wlan_client_rating"
const ModelWLANClientTimeline Model = "wlan_client_timeline"
const ModelWLANCurrentState Model = "wlan_current_state"
const ModelWLANRateStats Model = "wlans_rate_stats"
const ModelWLANRating Model = "wlan_rating"
const ModelWLANSessionAvg Model = "wlan_session_avg"
const ModelWLANStats Model = "wlans_stats"
const ModelWLANTags Model = "wlan_tags"
const ModelWLANTimeline Model = "wlan_timeline"
const ModelWlanDashboard Model = "wlan_dashboard"

func (en Model) GetPtr() *Model { var v = en; return &v }

func (en Model) String() string {
	switch en {
	case ModelBSSRating:
		return "bss_rating"
	case ModelBaseLocation:
		return "base_location"
	case ModelCPE:
		return "cpe"
	case ModelCPEAggStat:
		return "cpe_agg_stats"
	case ModelCPEConfigTemplates:
		return "cpe_config_templates"
	case ModelCPECurrentState:
		return "cpe_current_state"
	case ModelCPEDashboard:
		return "cpe_dashboard"
	case ModelCPEEvents:
		return "cpe_events"
	case ModelCPEFWUpgrade:
		return "cpe_fw_upgrade"
	case ModelCPEMap:
		return "mapobjects"
	case ModelCPEMapCPEList:
		return "cpe_map_cpe_list"
	case ModelCPEMapPosition:
		return "cpe_map_position"
	case ModelCPEMapScale:
		return "cpe_map_scale"
	case ModelCPEModels:
		return "cpe_models"
	case ModelCPERating:
		return "cpe_rating"
	case ModelCPERemote:
		return "cpe_remote"
	case ModelCPEReport:
		return "cpe_report"
	case ModelCPEScanLast:
		return "cpe_scan_last"
	case ModelCPESession:
		return "cpe_session"
	case ModelCPESessionAvg:
		return "cpe_session_avg"
	case ModelCPEStats:
		return "stat"
	case ModelCPETags:
		return "cpe_tags"
	case ModelCPETimeline:
		return "cpe_timeline"
	case ModelCPEWLANAggStat:
		return "cpe_wlan_agg_stats"
	case ModelCaptiveRedirect:
		return "captive-redirect"
	case ModelClient:
		return "client"
	case ModelClientCurrentState:
		return "client_current_state"
	case ModelClientDashboard:
		return "client_dashboard"
	case ModelClientDistance:
		return "client_distance"
	case ModelClientEvents:
		return "client_events"
	case ModelClientList:
		return "client-list"
	case ModelClientRF:
		return "client_rf"
	case ModelClientRating:
		return "client_rating"
	case ModelClientSession:
		return "client_session"
	case ModelClientSessionAvg:
		return "client_session_avg"
	case ModelClientStatExt:
		return "client-stat-ext"
	case ModelClientStats:
		return "client-stat"
	case ModelClientTimeline:
		return "client_timeline"
	case ModelCommonDashboard:
		return "common_dashboard"
	case ModelCommonEvents:
		return "common_events"
	case ModelCompactDashboard:
		return "compact_dashboard"
	case ModelConfiguration:
		return "configuration"
	case ModelController:
		return "controller"
	case ModelDitExport:
		return "dit_export"
	case ModelExternalAP:
		return "ext_access_point"
	case ModelHeatMapClients:
		return "heatmap_clients"
	case ModelHeatMapLBSClients:
		return "heatmap_lbs_clients"
	case ModelHeatMapTraffic:
		return "heatmap_traffic"
	case ModelHost:
		return "host"
	case ModelHotspotProfile:
		return "hotspot-profile"
	case ModelL2Chain:
		return "l2-chain"
	case ModelLBSCPEInfo:
		return "lbs-cpe-info"
	case ModelLBSClientCoord:
		return "lbs-client-coords"
	case ModelLBSClientData:
		return "lbs-client-data"
	case ModelLBSClientProbes:
		return "lbs-client-probes"
	case ModelLBSClientTrack:
		return "lbs-client-track"
	case ModelLBSZones:
		return "lbs_zones"
	case ModelLicense:
		return "license"
	case ModelLicenseLogDaily:
		return "license_log_daily"
	case ModelLoaderSettings:
		return "settings_loader"
	case ModelLoaderTask:
		return "periodic_task"
	case ModelLocation:
		return "location"
	case ModelMoniEvents:
		return "monitor_events"
	case ModelMonitorCPE:
		return "poll-cpe"
	case ModelMonitorEvents:
		return "event-stat-rule-violation"
	case ModelMonitorRules:
		return "stat-event-rule"
	case ModelOperation:
		return "operation"
	case ModelRADIUS:
		return "radius"
	case ModelRRMGroups:
		return "rrm-groups"
	case ModelRadarExport:
		return "radar_exports"
	case ModelResampling:
		return "resampling"
	case ModelSnmpWalker:
		return "snmp_walker"
	case ModelStatReport:
		return "report"
	case ModelSystemCPERateStats:
		return "system_cpe_rate_stats"
	case ModelSystemCPEStats:
		return "system_cpe_stats"
	case ModelSystemDashboard:
		return "system_dashboard"
	case ModelSystemEvents:
		return "system_events"
	case ModelSystemRating:
		return "system_rating"
	case ModelSystemTimeline:
		return "system_timeline"
	case ModelUser:
		return "user"
	case ModelWLAN:
		return "wlan"
	case ModelWLANAggStat:
		return "wlan_agg_stats"
	case ModelWLANCPEAggStat:
		return "wlan_cpe_agg_stats"
	case ModelWLANCPETimeline:
		return "wlan_cpe_timeline"
	case ModelWLANClientRating:
		return "wlan_client_rating"
	case ModelWLANClientTimeline:
		return "wlan_client_timeline"
	case ModelWLANCurrentState:
		return "wlan_current_state"
	case ModelWLANRateStats:
		return "wlans_rate_stats"
	case ModelWLANRating:
		return "wlan_rating"
	case ModelWLANSessionAvg:
		return "wlan_session_avg"
	case ModelWLANStats:
		return "wlans_stats"
	case ModelWLANTags:
		return "wlan_tags"
	case ModelWLANTimeline:
		return "wlan_timeline"
	case ModelWlanDashboard:
		return "wlan_dashboard"
	}
	panic(errors.New("Invalid value of Model: " + string(en)))
}

func (en *Model) MarshalJSON() ([]byte, error) {
	switch *en {
	case ModelBSSRating:
		return json.Marshal("bss_rating")
	case ModelBaseLocation:
		return json.Marshal("base_location")
	case ModelCPE:
		return json.Marshal("cpe")
	case ModelCPEAggStat:
		return json.Marshal("cpe_agg_stats")
	case ModelCPEConfigTemplates:
		return json.Marshal("cpe_config_templates")
	case ModelCPECurrentState:
		return json.Marshal("cpe_current_state")
	case ModelCPEDashboard:
		return json.Marshal("cpe_dashboard")
	case ModelCPEEvents:
		return json.Marshal("cpe_events")
	case ModelCPEFWUpgrade:
		return json.Marshal("cpe_fw_upgrade")
	case ModelCPEMap:
		return json.Marshal("mapobjects")
	case ModelCPEMapCPEList:
		return json.Marshal("cpe_map_cpe_list")
	case ModelCPEMapPosition:
		return json.Marshal("cpe_map_position")
	case ModelCPEMapScale:
		return json.Marshal("cpe_map_scale")
	case ModelCPEModels:
		return json.Marshal("cpe_models")
	case ModelCPERating:
		return json.Marshal("cpe_rating")
	case ModelCPERemote:
		return json.Marshal("cpe_remote")
	case ModelCPEReport:
		return json.Marshal("cpe_report")
	case ModelCPEScanLast:
		return json.Marshal("cpe_scan_last")
	case ModelCPESession:
		return json.Marshal("cpe_session")
	case ModelCPESessionAvg:
		return json.Marshal("cpe_session_avg")
	case ModelCPEStats:
		return json.Marshal("stat")
	case ModelCPETags:
		return json.Marshal("cpe_tags")
	case ModelCPETimeline:
		return json.Marshal("cpe_timeline")
	case ModelCPEWLANAggStat:
		return json.Marshal("cpe_wlan_agg_stats")
	case ModelCaptiveRedirect:
		return json.Marshal("captive-redirect")
	case ModelClient:
		return json.Marshal("client")
	case ModelClientCurrentState:
		return json.Marshal("client_current_state")
	case ModelClientDashboard:
		return json.Marshal("client_dashboard")
	case ModelClientDistance:
		return json.Marshal("client_distance")
	case ModelClientEvents:
		return json.Marshal("client_events")
	case ModelClientList:
		return json.Marshal("client-list")
	case ModelClientRF:
		return json.Marshal("client_rf")
	case ModelClientRating:
		return json.Marshal("client_rating")
	case ModelClientSession:
		return json.Marshal("client_session")
	case ModelClientSessionAvg:
		return json.Marshal("client_session_avg")
	case ModelClientStatExt:
		return json.Marshal("client-stat-ext")
	case ModelClientStats:
		return json.Marshal("client-stat")
	case ModelClientTimeline:
		return json.Marshal("client_timeline")
	case ModelCommonDashboard:
		return json.Marshal("common_dashboard")
	case ModelCommonEvents:
		return json.Marshal("common_events")
	case ModelCompactDashboard:
		return json.Marshal("compact_dashboard")
	case ModelConfiguration:
		return json.Marshal("configuration")
	case ModelController:
		return json.Marshal("controller")
	case ModelDitExport:
		return json.Marshal("dit_export")
	case ModelExternalAP:
		return json.Marshal("ext_access_point")
	case ModelHeatMapClients:
		return json.Marshal("heatmap_clients")
	case ModelHeatMapLBSClients:
		return json.Marshal("heatmap_lbs_clients")
	case ModelHeatMapTraffic:
		return json.Marshal("heatmap_traffic")
	case ModelHost:
		return json.Marshal("host")
	case ModelHotspotProfile:
		return json.Marshal("hotspot-profile")
	case ModelL2Chain:
		return json.Marshal("l2-chain")
	case ModelLBSCPEInfo:
		return json.Marshal("lbs-cpe-info")
	case ModelLBSClientCoord:
		return json.Marshal("lbs-client-coords")
	case ModelLBSClientData:
		return json.Marshal("lbs-client-data")
	case ModelLBSClientProbes:
		return json.Marshal("lbs-client-probes")
	case ModelLBSClientTrack:
		return json.Marshal("lbs-client-track")
	case ModelLBSZones:
		return json.Marshal("lbs_zones")
	case ModelLicense:
		return json.Marshal("license")
	case ModelLicenseLogDaily:
		return json.Marshal("license_log_daily")
	case ModelLoaderSettings:
		return json.Marshal("settings_loader")
	case ModelLoaderTask:
		return json.Marshal("periodic_task")
	case ModelLocation:
		return json.Marshal("location")
	case ModelMoniEvents:
		return json.Marshal("monitor_events")
	case ModelMonitorCPE:
		return json.Marshal("poll-cpe")
	case ModelMonitorEvents:
		return json.Marshal("event-stat-rule-violation")
	case ModelMonitorRules:
		return json.Marshal("stat-event-rule")
	case ModelOperation:
		return json.Marshal("operation")
	case ModelRADIUS:
		return json.Marshal("radius")
	case ModelRRMGroups:
		return json.Marshal("rrm-groups")
	case ModelRadarExport:
		return json.Marshal("radar_exports")
	case ModelResampling:
		return json.Marshal("resampling")
	case ModelSnmpWalker:
		return json.Marshal("snmp_walker")
	case ModelStatReport:
		return json.Marshal("report")
	case ModelSystemCPERateStats:
		return json.Marshal("system_cpe_rate_stats")
	case ModelSystemCPEStats:
		return json.Marshal("system_cpe_stats")
	case ModelSystemDashboard:
		return json.Marshal("system_dashboard")
	case ModelSystemEvents:
		return json.Marshal("system_events")
	case ModelSystemRating:
		return json.Marshal("system_rating")
	case ModelSystemTimeline:
		return json.Marshal("system_timeline")
	case ModelUser:
		return json.Marshal("user")
	case ModelWLAN:
		return json.Marshal("wlan")
	case ModelWLANAggStat:
		return json.Marshal("wlan_agg_stats")
	case ModelWLANCPEAggStat:
		return json.Marshal("wlan_cpe_agg_stats")
	case ModelWLANCPETimeline:
		return json.Marshal("wlan_cpe_timeline")
	case ModelWLANClientRating:
		return json.Marshal("wlan_client_rating")
	case ModelWLANClientTimeline:
		return json.Marshal("wlan_client_timeline")
	case ModelWLANCurrentState:
		return json.Marshal("wlan_current_state")
	case ModelWLANRateStats:
		return json.Marshal("wlans_rate_stats")
	case ModelWLANRating:
		return json.Marshal("wlan_rating")
	case ModelWLANSessionAvg:
		return json.Marshal("wlan_session_avg")
	case ModelWLANStats:
		return json.Marshal("wlans_stats")
	case ModelWLANTags:
		return json.Marshal("wlan_tags")
	case ModelWLANTimeline:
		return json.Marshal("wlan_timeline")
	case ModelWlanDashboard:
		return json.Marshal("wlan_dashboard")
	}
	return nil, errors.New("Invalid value of Model: " + string(*en))
}

func (en *Model) GetBSON() (interface{}, error) {
	switch *en {
	case ModelBSSRating:
		return "bss_rating", nil
	case ModelBaseLocation:
		return "base_location", nil
	case ModelCPE:
		return "cpe", nil
	case ModelCPEAggStat:
		return "cpe_agg_stats", nil
	case ModelCPEConfigTemplates:
		return "cpe_config_templates", nil
	case ModelCPECurrentState:
		return "cpe_current_state", nil
	case ModelCPEDashboard:
		return "cpe_dashboard", nil
	case ModelCPEEvents:
		return "cpe_events", nil
	case ModelCPEFWUpgrade:
		return "cpe_fw_upgrade", nil
	case ModelCPEMap:
		return "mapobjects", nil
	case ModelCPEMapCPEList:
		return "cpe_map_cpe_list", nil
	case ModelCPEMapPosition:
		return "cpe_map_position", nil
	case ModelCPEMapScale:
		return "cpe_map_scale", nil
	case ModelCPEModels:
		return "cpe_models", nil
	case ModelCPERating:
		return "cpe_rating", nil
	case ModelCPERemote:
		return "cpe_remote", nil
	case ModelCPEReport:
		return "cpe_report", nil
	case ModelCPEScanLast:
		return "cpe_scan_last", nil
	case ModelCPESession:
		return "cpe_session", nil
	case ModelCPESessionAvg:
		return "cpe_session_avg", nil
	case ModelCPEStats:
		return "stat", nil
	case ModelCPETags:
		return "cpe_tags", nil
	case ModelCPETimeline:
		return "cpe_timeline", nil
	case ModelCPEWLANAggStat:
		return "cpe_wlan_agg_stats", nil
	case ModelCaptiveRedirect:
		return "captive-redirect", nil
	case ModelClient:
		return "client", nil
	case ModelClientCurrentState:
		return "client_current_state", nil
	case ModelClientDashboard:
		return "client_dashboard", nil
	case ModelClientDistance:
		return "client_distance", nil
	case ModelClientEvents:
		return "client_events", nil
	case ModelClientList:
		return "client-list", nil
	case ModelClientRF:
		return "client_rf", nil
	case ModelClientRating:
		return "client_rating", nil
	case ModelClientSession:
		return "client_session", nil
	case ModelClientSessionAvg:
		return "client_session_avg", nil
	case ModelClientStatExt:
		return "client-stat-ext", nil
	case ModelClientStats:
		return "client-stat", nil
	case ModelClientTimeline:
		return "client_timeline", nil
	case ModelCommonDashboard:
		return "common_dashboard", nil
	case ModelCommonEvents:
		return "common_events", nil
	case ModelCompactDashboard:
		return "compact_dashboard", nil
	case ModelConfiguration:
		return "configuration", nil
	case ModelController:
		return "controller", nil
	case ModelDitExport:
		return "dit_export", nil
	case ModelExternalAP:
		return "ext_access_point", nil
	case ModelHeatMapClients:
		return "heatmap_clients", nil
	case ModelHeatMapLBSClients:
		return "heatmap_lbs_clients", nil
	case ModelHeatMapTraffic:
		return "heatmap_traffic", nil
	case ModelHost:
		return "host", nil
	case ModelHotspotProfile:
		return "hotspot-profile", nil
	case ModelL2Chain:
		return "l2-chain", nil
	case ModelLBSCPEInfo:
		return "lbs-cpe-info", nil
	case ModelLBSClientCoord:
		return "lbs-client-coords", nil
	case ModelLBSClientData:
		return "lbs-client-data", nil
	case ModelLBSClientProbes:
		return "lbs-client-probes", nil
	case ModelLBSClientTrack:
		return "lbs-client-track", nil
	case ModelLBSZones:
		return "lbs_zones", nil
	case ModelLicense:
		return "license", nil
	case ModelLicenseLogDaily:
		return "license_log_daily", nil
	case ModelLoaderSettings:
		return "settings_loader", nil
	case ModelLoaderTask:
		return "periodic_task", nil
	case ModelLocation:
		return "location", nil
	case ModelMoniEvents:
		return "monitor_events", nil
	case ModelMonitorCPE:
		return "poll-cpe", nil
	case ModelMonitorEvents:
		return "event-stat-rule-violation", nil
	case ModelMonitorRules:
		return "stat-event-rule", nil
	case ModelOperation:
		return "operation", nil
	case ModelRADIUS:
		return "radius", nil
	case ModelRRMGroups:
		return "rrm-groups", nil
	case ModelRadarExport:
		return "radar_exports", nil
	case ModelResampling:
		return "resampling", nil
	case ModelSnmpWalker:
		return "snmp_walker", nil
	case ModelStatReport:
		return "report", nil
	case ModelSystemCPERateStats:
		return "system_cpe_rate_stats", nil
	case ModelSystemCPEStats:
		return "system_cpe_stats", nil
	case ModelSystemDashboard:
		return "system_dashboard", nil
	case ModelSystemEvents:
		return "system_events", nil
	case ModelSystemRating:
		return "system_rating", nil
	case ModelSystemTimeline:
		return "system_timeline", nil
	case ModelUser:
		return "user", nil
	case ModelWLAN:
		return "wlan", nil
	case ModelWLANAggStat:
		return "wlan_agg_stats", nil
	case ModelWLANCPEAggStat:
		return "wlan_cpe_agg_stats", nil
	case ModelWLANCPETimeline:
		return "wlan_cpe_timeline", nil
	case ModelWLANClientRating:
		return "wlan_client_rating", nil
	case ModelWLANClientTimeline:
		return "wlan_client_timeline", nil
	case ModelWLANCurrentState:
		return "wlan_current_state", nil
	case ModelWLANRateStats:
		return "wlans_rate_stats", nil
	case ModelWLANRating:
		return "wlan_rating", nil
	case ModelWLANSessionAvg:
		return "wlan_session_avg", nil
	case ModelWLANStats:
		return "wlans_stats", nil
	case ModelWLANTags:
		return "wlan_tags", nil
	case ModelWLANTimeline:
		return "wlan_timeline", nil
	case ModelWlanDashboard:
		return "wlan_dashboard", nil
	}
	return nil, errors.New("Invalid value of Model: " + string(*en))
}

func (en *Model) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "bss_rating":
		*en = ModelBSSRating
		return nil
	case "base_location":
		*en = ModelBaseLocation
		return nil
	case "cpe":
		*en = ModelCPE
		return nil
	case "cpe_agg_stats":
		*en = ModelCPEAggStat
		return nil
	case "cpe_config_templates":
		*en = ModelCPEConfigTemplates
		return nil
	case "cpe_current_state":
		*en = ModelCPECurrentState
		return nil
	case "cpe_dashboard":
		*en = ModelCPEDashboard
		return nil
	case "cpe_events":
		*en = ModelCPEEvents
		return nil
	case "cpe_fw_upgrade":
		*en = ModelCPEFWUpgrade
		return nil
	case "mapobjects":
		*en = ModelCPEMap
		return nil
	case "cpe_map_cpe_list":
		*en = ModelCPEMapCPEList
		return nil
	case "cpe_map_position":
		*en = ModelCPEMapPosition
		return nil
	case "cpe_map_scale":
		*en = ModelCPEMapScale
		return nil
	case "cpe_models":
		*en = ModelCPEModels
		return nil
	case "cpe_rating":
		*en = ModelCPERating
		return nil
	case "cpe_remote":
		*en = ModelCPERemote
		return nil
	case "cpe_report":
		*en = ModelCPEReport
		return nil
	case "cpe_scan_last":
		*en = ModelCPEScanLast
		return nil
	case "cpe_session":
		*en = ModelCPESession
		return nil
	case "cpe_session_avg":
		*en = ModelCPESessionAvg
		return nil
	case "stat":
		*en = ModelCPEStats
		return nil
	case "cpe_tags":
		*en = ModelCPETags
		return nil
	case "cpe_timeline":
		*en = ModelCPETimeline
		return nil
	case "cpe_wlan_agg_stats":
		*en = ModelCPEWLANAggStat
		return nil
	case "captive-redirect":
		*en = ModelCaptiveRedirect
		return nil
	case "client":
		*en = ModelClient
		return nil
	case "client_current_state":
		*en = ModelClientCurrentState
		return nil
	case "client_dashboard":
		*en = ModelClientDashboard
		return nil
	case "client_distance":
		*en = ModelClientDistance
		return nil
	case "client_events":
		*en = ModelClientEvents
		return nil
	case "client-list":
		*en = ModelClientList
		return nil
	case "client_rf":
		*en = ModelClientRF
		return nil
	case "client_rating":
		*en = ModelClientRating
		return nil
	case "client_session":
		*en = ModelClientSession
		return nil
	case "client_session_avg":
		*en = ModelClientSessionAvg
		return nil
	case "client-stat-ext":
		*en = ModelClientStatExt
		return nil
	case "client-stat":
		*en = ModelClientStats
		return nil
	case "client_timeline":
		*en = ModelClientTimeline
		return nil
	case "common_dashboard":
		*en = ModelCommonDashboard
		return nil
	case "common_events":
		*en = ModelCommonEvents
		return nil
	case "compact_dashboard":
		*en = ModelCompactDashboard
		return nil
	case "configuration":
		*en = ModelConfiguration
		return nil
	case "controller":
		*en = ModelController
		return nil
	case "dit_export":
		*en = ModelDitExport
		return nil
	case "ext_access_point":
		*en = ModelExternalAP
		return nil
	case "heatmap_clients":
		*en = ModelHeatMapClients
		return nil
	case "heatmap_lbs_clients":
		*en = ModelHeatMapLBSClients
		return nil
	case "heatmap_traffic":
		*en = ModelHeatMapTraffic
		return nil
	case "host":
		*en = ModelHost
		return nil
	case "hotspot-profile":
		*en = ModelHotspotProfile
		return nil
	case "l2-chain":
		*en = ModelL2Chain
		return nil
	case "lbs-cpe-info":
		*en = ModelLBSCPEInfo
		return nil
	case "lbs-client-coords":
		*en = ModelLBSClientCoord
		return nil
	case "lbs-client-data":
		*en = ModelLBSClientData
		return nil
	case "lbs-client-probes":
		*en = ModelLBSClientProbes
		return nil
	case "lbs-client-track":
		*en = ModelLBSClientTrack
		return nil
	case "lbs_zones":
		*en = ModelLBSZones
		return nil
	case "license":
		*en = ModelLicense
		return nil
	case "license_log_daily":
		*en = ModelLicenseLogDaily
		return nil
	case "settings_loader":
		*en = ModelLoaderSettings
		return nil
	case "periodic_task":
		*en = ModelLoaderTask
		return nil
	case "location":
		*en = ModelLocation
		return nil
	case "monitor_events":
		*en = ModelMoniEvents
		return nil
	case "poll-cpe":
		*en = ModelMonitorCPE
		return nil
	case "event-stat-rule-violation":
		*en = ModelMonitorEvents
		return nil
	case "stat-event-rule":
		*en = ModelMonitorRules
		return nil
	case "operation":
		*en = ModelOperation
		return nil
	case "radius":
		*en = ModelRADIUS
		return nil
	case "rrm-groups":
		*en = ModelRRMGroups
		return nil
	case "radar_exports":
		*en = ModelRadarExport
		return nil
	case "resampling":
		*en = ModelResampling
		return nil
	case "snmp_walker":
		*en = ModelSnmpWalker
		return nil
	case "report":
		*en = ModelStatReport
		return nil
	case "system_cpe_rate_stats":
		*en = ModelSystemCPERateStats
		return nil
	case "system_cpe_stats":
		*en = ModelSystemCPEStats
		return nil
	case "system_dashboard":
		*en = ModelSystemDashboard
		return nil
	case "system_events":
		*en = ModelSystemEvents
		return nil
	case "system_rating":
		*en = ModelSystemRating
		return nil
	case "system_timeline":
		*en = ModelSystemTimeline
		return nil
	case "user":
		*en = ModelUser
		return nil
	case "wlan":
		*en = ModelWLAN
		return nil
	case "wlan_agg_stats":
		*en = ModelWLANAggStat
		return nil
	case "wlan_cpe_agg_stats":
		*en = ModelWLANCPEAggStat
		return nil
	case "wlan_cpe_timeline":
		*en = ModelWLANCPETimeline
		return nil
	case "wlan_client_rating":
		*en = ModelWLANClientRating
		return nil
	case "wlan_client_timeline":
		*en = ModelWLANClientTimeline
		return nil
	case "wlan_current_state":
		*en = ModelWLANCurrentState
		return nil
	case "wlans_rate_stats":
		*en = ModelWLANRateStats
		return nil
	case "wlan_rating":
		*en = ModelWLANRating
		return nil
	case "wlan_session_avg":
		*en = ModelWLANSessionAvg
		return nil
	case "wlans_stats":
		*en = ModelWLANStats
		return nil
	case "wlan_tags":
		*en = ModelWLANTags
		return nil
	case "wlan_timeline":
		*en = ModelWLANTimeline
		return nil
	case "wlan_dashboard":
		*en = ModelWlanDashboard
		return nil
	}
	return errors.New("Unknown Model: " + s)
}

func (en *Model) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "bss_rating":
		*en = ModelBSSRating
		return nil
	case "base_location":
		*en = ModelBaseLocation
		return nil
	case "cpe":
		*en = ModelCPE
		return nil
	case "cpe_agg_stats":
		*en = ModelCPEAggStat
		return nil
	case "cpe_config_templates":
		*en = ModelCPEConfigTemplates
		return nil
	case "cpe_current_state":
		*en = ModelCPECurrentState
		return nil
	case "cpe_dashboard":
		*en = ModelCPEDashboard
		return nil
	case "cpe_events":
		*en = ModelCPEEvents
		return nil
	case "cpe_fw_upgrade":
		*en = ModelCPEFWUpgrade
		return nil
	case "mapobjects":
		*en = ModelCPEMap
		return nil
	case "cpe_map_cpe_list":
		*en = ModelCPEMapCPEList
		return nil
	case "cpe_map_position":
		*en = ModelCPEMapPosition
		return nil
	case "cpe_map_scale":
		*en = ModelCPEMapScale
		return nil
	case "cpe_models":
		*en = ModelCPEModels
		return nil
	case "cpe_rating":
		*en = ModelCPERating
		return nil
	case "cpe_remote":
		*en = ModelCPERemote
		return nil
	case "cpe_report":
		*en = ModelCPEReport
		return nil
	case "cpe_scan_last":
		*en = ModelCPEScanLast
		return nil
	case "cpe_session":
		*en = ModelCPESession
		return nil
	case "cpe_session_avg":
		*en = ModelCPESessionAvg
		return nil
	case "stat":
		*en = ModelCPEStats
		return nil
	case "cpe_tags":
		*en = ModelCPETags
		return nil
	case "cpe_timeline":
		*en = ModelCPETimeline
		return nil
	case "cpe_wlan_agg_stats":
		*en = ModelCPEWLANAggStat
		return nil
	case "captive-redirect":
		*en = ModelCaptiveRedirect
		return nil
	case "client":
		*en = ModelClient
		return nil
	case "client_current_state":
		*en = ModelClientCurrentState
		return nil
	case "client_dashboard":
		*en = ModelClientDashboard
		return nil
	case "client_distance":
		*en = ModelClientDistance
		return nil
	case "client_events":
		*en = ModelClientEvents
		return nil
	case "client-list":
		*en = ModelClientList
		return nil
	case "client_rf":
		*en = ModelClientRF
		return nil
	case "client_rating":
		*en = ModelClientRating
		return nil
	case "client_session":
		*en = ModelClientSession
		return nil
	case "client_session_avg":
		*en = ModelClientSessionAvg
		return nil
	case "client-stat-ext":
		*en = ModelClientStatExt
		return nil
	case "client-stat":
		*en = ModelClientStats
		return nil
	case "client_timeline":
		*en = ModelClientTimeline
		return nil
	case "common_dashboard":
		*en = ModelCommonDashboard
		return nil
	case "common_events":
		*en = ModelCommonEvents
		return nil
	case "compact_dashboard":
		*en = ModelCompactDashboard
		return nil
	case "configuration":
		*en = ModelConfiguration
		return nil
	case "controller":
		*en = ModelController
		return nil
	case "dit_export":
		*en = ModelDitExport
		return nil
	case "ext_access_point":
		*en = ModelExternalAP
		return nil
	case "heatmap_clients":
		*en = ModelHeatMapClients
		return nil
	case "heatmap_lbs_clients":
		*en = ModelHeatMapLBSClients
		return nil
	case "heatmap_traffic":
		*en = ModelHeatMapTraffic
		return nil
	case "host":
		*en = ModelHost
		return nil
	case "hotspot-profile":
		*en = ModelHotspotProfile
		return nil
	case "l2-chain":
		*en = ModelL2Chain
		return nil
	case "lbs-cpe-info":
		*en = ModelLBSCPEInfo
		return nil
	case "lbs-client-coords":
		*en = ModelLBSClientCoord
		return nil
	case "lbs-client-data":
		*en = ModelLBSClientData
		return nil
	case "lbs-client-probes":
		*en = ModelLBSClientProbes
		return nil
	case "lbs-client-track":
		*en = ModelLBSClientTrack
		return nil
	case "lbs_zones":
		*en = ModelLBSZones
		return nil
	case "license":
		*en = ModelLicense
		return nil
	case "license_log_daily":
		*en = ModelLicenseLogDaily
		return nil
	case "settings_loader":
		*en = ModelLoaderSettings
		return nil
	case "periodic_task":
		*en = ModelLoaderTask
		return nil
	case "location":
		*en = ModelLocation
		return nil
	case "monitor_events":
		*en = ModelMoniEvents
		return nil
	case "poll-cpe":
		*en = ModelMonitorCPE
		return nil
	case "event-stat-rule-violation":
		*en = ModelMonitorEvents
		return nil
	case "stat-event-rule":
		*en = ModelMonitorRules
		return nil
	case "operation":
		*en = ModelOperation
		return nil
	case "radius":
		*en = ModelRADIUS
		return nil
	case "rrm-groups":
		*en = ModelRRMGroups
		return nil
	case "radar_exports":
		*en = ModelRadarExport
		return nil
	case "resampling":
		*en = ModelResampling
		return nil
	case "snmp_walker":
		*en = ModelSnmpWalker
		return nil
	case "report":
		*en = ModelStatReport
		return nil
	case "system_cpe_rate_stats":
		*en = ModelSystemCPERateStats
		return nil
	case "system_cpe_stats":
		*en = ModelSystemCPEStats
		return nil
	case "system_dashboard":
		*en = ModelSystemDashboard
		return nil
	case "system_events":
		*en = ModelSystemEvents
		return nil
	case "system_rating":
		*en = ModelSystemRating
		return nil
	case "system_timeline":
		*en = ModelSystemTimeline
		return nil
	case "user":
		*en = ModelUser
		return nil
	case "wlan":
		*en = ModelWLAN
		return nil
	case "wlan_agg_stats":
		*en = ModelWLANAggStat
		return nil
	case "wlan_cpe_agg_stats":
		*en = ModelWLANCPEAggStat
		return nil
	case "wlan_cpe_timeline":
		*en = ModelWLANCPETimeline
		return nil
	case "wlan_client_rating":
		*en = ModelWLANClientRating
		return nil
	case "wlan_client_timeline":
		*en = ModelWLANClientTimeline
		return nil
	case "wlan_current_state":
		*en = ModelWLANCurrentState
		return nil
	case "wlans_rate_stats":
		*en = ModelWLANRateStats
		return nil
	case "wlan_rating":
		*en = ModelWLANRating
		return nil
	case "wlan_session_avg":
		*en = ModelWLANSessionAvg
		return nil
	case "wlans_stats":
		*en = ModelWLANStats
		return nil
	case "wlan_tags":
		*en = ModelWLANTags
		return nil
	case "wlan_timeline":
		*en = ModelWLANTimeline
		return nil
	case "wlan_dashboard":
		*en = ModelWlanDashboard
		return nil
	}
	return errors.New("Unknown Model: " + s)
}

type Module string

const ModuleAC Module = "AC"
const ModuleAnalMW Module = "ANAL-MW"
const ModuleAny Module = "+"
const ModuleBackend Module = "BACKEND"
const ModuleCLI Module = "CLI"
const ModuleCPE Module = "CPE"
const ModuleCPEStat Module = "CPE_STAT"
const ModuleClientDistance Module = "CLIENT-DISTANCE"
const ModuleClientStat Module = "CLIENT_STAT"
const ModuleConfig Module = "CONFIG"
const ModuleDB Module = "DB"
const ModuleDummy Module = "DUMMY"
const ModuleFW Module = "FW"
const ModuleGSNMP Module = "GSNMP"
const ModuleLBS Module = "LBS"
const ModuleMQTTLog Module = "MQTT_LOG"
const ModuleMediator Module = "MEDIATOR"
const ModuleMonitor Module = "MONITOR"
const ModuleNone Module = ""
const ModulePortalBack Module = "PORTAL_BACKEND"
const ModuleRRM Module = "RRM"
const ModuleRadarExportMW Module = "RADAR-EXPORT-MW"
const ModuleRadarMW Module = "RADAR-MW"
const ModuleRadiusGw Module = "RADIUS_GATEWAY"
const ModuleRedirect Module = "REDIRECT"
const ModuleResampling Module = "RESAMPLING"
const ModuleSnmpWalker Module = "SNMP_WALKER"
const ModuleStat Module = "STAT"
const ModuleStatLBS Module = "STAT-LBS"
const ModuleTunManager Module = "TUN_MANAGER"

func (en Module) GetPtr() *Module { var v = en; return &v }

func (en Module) String() string {
	switch en {
	case ModuleAC:
		return "AC"
	case ModuleAnalMW:
		return "ANAL-MW"
	case ModuleAny:
		return "+"
	case ModuleBackend:
		return "BACKEND"
	case ModuleCLI:
		return "CLI"
	case ModuleCPE:
		return "CPE"
	case ModuleCPEStat:
		return "CPE_STAT"
	case ModuleClientDistance:
		return "CLIENT-DISTANCE"
	case ModuleClientStat:
		return "CLIENT_STAT"
	case ModuleConfig:
		return "CONFIG"
	case ModuleDB:
		return "DB"
	case ModuleDummy:
		return "DUMMY"
	case ModuleFW:
		return "FW"
	case ModuleGSNMP:
		return "GSNMP"
	case ModuleLBS:
		return "LBS"
	case ModuleMQTTLog:
		return "MQTT_LOG"
	case ModuleMediator:
		return "MEDIATOR"
	case ModuleMonitor:
		return "MONITOR"
	case ModuleNone:
		return ""
	case ModulePortalBack:
		return "PORTAL_BACKEND"
	case ModuleRRM:
		return "RRM"
	case ModuleRadarExportMW:
		return "RADAR-EXPORT-MW"
	case ModuleRadarMW:
		return "RADAR-MW"
	case ModuleRadiusGw:
		return "RADIUS_GATEWAY"
	case ModuleRedirect:
		return "REDIRECT"
	case ModuleResampling:
		return "RESAMPLING"
	case ModuleSnmpWalker:
		return "SNMP_WALKER"
	case ModuleStat:
		return "STAT"
	case ModuleStatLBS:
		return "STAT-LBS"
	case ModuleTunManager:
		return "TUN_MANAGER"
	}
	if len(en) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of Module: " + string(en)))
}

func (en *Module) MarshalJSON() ([]byte, error) {
	switch *en {
	case ModuleAC:
		return json.Marshal("AC")
	case ModuleAnalMW:
		return json.Marshal("ANAL-MW")
	case ModuleAny:
		return json.Marshal("+")
	case ModuleBackend:
		return json.Marshal("BACKEND")
	case ModuleCLI:
		return json.Marshal("CLI")
	case ModuleCPE:
		return json.Marshal("CPE")
	case ModuleCPEStat:
		return json.Marshal("CPE_STAT")
	case ModuleClientDistance:
		return json.Marshal("CLIENT-DISTANCE")
	case ModuleClientStat:
		return json.Marshal("CLIENT_STAT")
	case ModuleConfig:
		return json.Marshal("CONFIG")
	case ModuleDB:
		return json.Marshal("DB")
	case ModuleDummy:
		return json.Marshal("DUMMY")
	case ModuleFW:
		return json.Marshal("FW")
	case ModuleGSNMP:
		return json.Marshal("GSNMP")
	case ModuleLBS:
		return json.Marshal("LBS")
	case ModuleMQTTLog:
		return json.Marshal("MQTT_LOG")
	case ModuleMediator:
		return json.Marshal("MEDIATOR")
	case ModuleMonitor:
		return json.Marshal("MONITOR")
	case ModuleNone:
		return json.Marshal("")
	case ModulePortalBack:
		return json.Marshal("PORTAL_BACKEND")
	case ModuleRRM:
		return json.Marshal("RRM")
	case ModuleRadarExportMW:
		return json.Marshal("RADAR-EXPORT-MW")
	case ModuleRadarMW:
		return json.Marshal("RADAR-MW")
	case ModuleRadiusGw:
		return json.Marshal("RADIUS_GATEWAY")
	case ModuleRedirect:
		return json.Marshal("REDIRECT")
	case ModuleResampling:
		return json.Marshal("RESAMPLING")
	case ModuleSnmpWalker:
		return json.Marshal("SNMP_WALKER")
	case ModuleStat:
		return json.Marshal("STAT")
	case ModuleStatLBS:
		return json.Marshal("STAT-LBS")
	case ModuleTunManager:
		return json.Marshal("TUN_MANAGER")
	}
	if len(*en) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of Module: " + string(*en))
}

func (en *Module) GetBSON() (interface{}, error) {
	switch *en {
	case ModuleAC:
		return "AC", nil
	case ModuleAnalMW:
		return "ANAL-MW", nil
	case ModuleAny:
		return "+", nil
	case ModuleBackend:
		return "BACKEND", nil
	case ModuleCLI:
		return "CLI", nil
	case ModuleCPE:
		return "CPE", nil
	case ModuleCPEStat:
		return "CPE_STAT", nil
	case ModuleClientDistance:
		return "CLIENT-DISTANCE", nil
	case ModuleClientStat:
		return "CLIENT_STAT", nil
	case ModuleConfig:
		return "CONFIG", nil
	case ModuleDB:
		return "DB", nil
	case ModuleDummy:
		return "DUMMY", nil
	case ModuleFW:
		return "FW", nil
	case ModuleGSNMP:
		return "GSNMP", nil
	case ModuleLBS:
		return "LBS", nil
	case ModuleMQTTLog:
		return "MQTT_LOG", nil
	case ModuleMediator:
		return "MEDIATOR", nil
	case ModuleMonitor:
		return "MONITOR", nil
	case ModuleNone:
		return "", nil
	case ModulePortalBack:
		return "PORTAL_BACKEND", nil
	case ModuleRRM:
		return "RRM", nil
	case ModuleRadarExportMW:
		return "RADAR-EXPORT-MW", nil
	case ModuleRadarMW:
		return "RADAR-MW", nil
	case ModuleRadiusGw:
		return "RADIUS_GATEWAY", nil
	case ModuleRedirect:
		return "REDIRECT", nil
	case ModuleResampling:
		return "RESAMPLING", nil
	case ModuleSnmpWalker:
		return "SNMP_WALKER", nil
	case ModuleStat:
		return "STAT", nil
	case ModuleStatLBS:
		return "STAT-LBS", nil
	case ModuleTunManager:
		return "TUN_MANAGER", nil
	}
	if len(*en) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of Module: " + string(*en))
}

func (en *Module) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "AC":
		*en = ModuleAC
		return nil
	case "ANAL-MW":
		*en = ModuleAnalMW
		return nil
	case "+":
		*en = ModuleAny
		return nil
	case "BACKEND":
		*en = ModuleBackend
		return nil
	case "CLI":
		*en = ModuleCLI
		return nil
	case "CPE":
		*en = ModuleCPE
		return nil
	case "CPE_STAT":
		*en = ModuleCPEStat
		return nil
	case "CLIENT-DISTANCE":
		*en = ModuleClientDistance
		return nil
	case "CLIENT_STAT":
		*en = ModuleClientStat
		return nil
	case "CONFIG":
		*en = ModuleConfig
		return nil
	case "DB":
		*en = ModuleDB
		return nil
	case "DUMMY":
		*en = ModuleDummy
		return nil
	case "FW":
		*en = ModuleFW
		return nil
	case "GSNMP":
		*en = ModuleGSNMP
		return nil
	case "LBS":
		*en = ModuleLBS
		return nil
	case "MQTT_LOG":
		*en = ModuleMQTTLog
		return nil
	case "MEDIATOR":
		*en = ModuleMediator
		return nil
	case "MONITOR":
		*en = ModuleMonitor
		return nil
	case "":
		*en = ModuleNone
		return nil
	case "PORTAL_BACKEND":
		*en = ModulePortalBack
		return nil
	case "RRM":
		*en = ModuleRRM
		return nil
	case "RADAR-EXPORT-MW":
		*en = ModuleRadarExportMW
		return nil
	case "RADAR-MW":
		*en = ModuleRadarMW
		return nil
	case "RADIUS_GATEWAY":
		*en = ModuleRadiusGw
		return nil
	case "REDIRECT":
		*en = ModuleRedirect
		return nil
	case "RESAMPLING":
		*en = ModuleResampling
		return nil
	case "SNMP_WALKER":
		*en = ModuleSnmpWalker
		return nil
	case "STAT":
		*en = ModuleStat
		return nil
	case "STAT-LBS":
		*en = ModuleStatLBS
		return nil
	case "TUN_MANAGER":
		*en = ModuleTunManager
		return nil
	}
	if len(s) == 0 {
		*en = ModuleNone
		return nil
	}
	return errors.New("Unknown Module: " + s)
}

func (en *Module) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "AC":
		*en = ModuleAC
		return nil
	case "ANAL-MW":
		*en = ModuleAnalMW
		return nil
	case "+":
		*en = ModuleAny
		return nil
	case "BACKEND":
		*en = ModuleBackend
		return nil
	case "CLI":
		*en = ModuleCLI
		return nil
	case "CPE":
		*en = ModuleCPE
		return nil
	case "CPE_STAT":
		*en = ModuleCPEStat
		return nil
	case "CLIENT-DISTANCE":
		*en = ModuleClientDistance
		return nil
	case "CLIENT_STAT":
		*en = ModuleClientStat
		return nil
	case "CONFIG":
		*en = ModuleConfig
		return nil
	case "DB":
		*en = ModuleDB
		return nil
	case "DUMMY":
		*en = ModuleDummy
		return nil
	case "FW":
		*en = ModuleFW
		return nil
	case "GSNMP":
		*en = ModuleGSNMP
		return nil
	case "LBS":
		*en = ModuleLBS
		return nil
	case "MQTT_LOG":
		*en = ModuleMQTTLog
		return nil
	case "MEDIATOR":
		*en = ModuleMediator
		return nil
	case "MONITOR":
		*en = ModuleMonitor
		return nil
	case "":
		*en = ModuleNone
		return nil
	case "PORTAL_BACKEND":
		*en = ModulePortalBack
		return nil
	case "RRM":
		*en = ModuleRRM
		return nil
	case "RADAR-EXPORT-MW":
		*en = ModuleRadarExportMW
		return nil
	case "RADAR-MW":
		*en = ModuleRadarMW
		return nil
	case "RADIUS_GATEWAY":
		*en = ModuleRadiusGw
		return nil
	case "REDIRECT":
		*en = ModuleRedirect
		return nil
	case "RESAMPLING":
		*en = ModuleResampling
		return nil
	case "SNMP_WALKER":
		*en = ModuleSnmpWalker
		return nil
	case "STAT":
		*en = ModuleStat
		return nil
	case "STAT-LBS":
		*en = ModuleStatLBS
		return nil
	case "TUN_MANAGER":
		*en = ModuleTunManager
		return nil
	}
	if len(s) == 0 {
		*en = ModuleNone
		return nil
	}
	return errors.New("Unknown Module: " + s)
}

type NotifyType string

const NotifyTypeEmail NotifyType = "email"
const NotifyTypeEmpty NotifyType = ""
const NotifyTypeMqtt NotifyType = "mqtt"
const NotifyTypeTelegram NotifyType = "telegram"

func (en NotifyType) GetPtr() *NotifyType { var v = en; return &v }

func (en NotifyType) String() string {
	switch en {
	case NotifyTypeEmail:
		return "email"
	case NotifyTypeEmpty:
		return ""
	case NotifyTypeMqtt:
		return "mqtt"
	case NotifyTypeTelegram:
		return "telegram"
	}
	if len(en) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of NotifyType: " + string(en)))
}

func (en *NotifyType) MarshalJSON() ([]byte, error) {
	switch *en {
	case NotifyTypeEmail:
		return json.Marshal("email")
	case NotifyTypeEmpty:
		return json.Marshal("")
	case NotifyTypeMqtt:
		return json.Marshal("mqtt")
	case NotifyTypeTelegram:
		return json.Marshal("telegram")
	}
	if len(*en) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of NotifyType: " + string(*en))
}

func (en *NotifyType) GetBSON() (interface{}, error) {
	switch *en {
	case NotifyTypeEmail:
		return "email", nil
	case NotifyTypeEmpty:
		return "", nil
	case NotifyTypeMqtt:
		return "mqtt", nil
	case NotifyTypeTelegram:
		return "telegram", nil
	}
	if len(*en) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of NotifyType: " + string(*en))
}

func (en *NotifyType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "email":
		*en = NotifyTypeEmail
		return nil
	case "":
		*en = NotifyTypeEmpty
		return nil
	case "mqtt":
		*en = NotifyTypeMqtt
		return nil
	case "telegram":
		*en = NotifyTypeTelegram
		return nil
	}
	if len(s) == 0 {
		*en = NotifyTypeEmpty
		return nil
	}
	return errors.New("Unknown NotifyType: " + s)
}

func (en *NotifyType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "email":
		*en = NotifyTypeEmail
		return nil
	case "":
		*en = NotifyTypeEmpty
		return nil
	case "mqtt":
		*en = NotifyTypeMqtt
		return nil
	case "telegram":
		*en = NotifyTypeTelegram
		return nil
	}
	if len(s) == 0 {
		*en = NotifyTypeEmpty
		return nil
	}
	return errors.New("Unknown NotifyType: " + s)
}

type Operation string

const OperationAny Operation = "+"
const OperationCPEStatus Operation = "STATUS"
const OperationCreate Operation = "C"
const OperationDelete Operation = "D"
const OperationJSONRPC Operation = "JSONRPC"
const OperationLuaScript Operation = "LUA"
const OperationRead Operation = "R"
const OperationSHScript Operation = "SH"
const OperationUpdate Operation = "U"

func (en Operation) GetPtr() *Operation { var v = en; return &v }

func (en Operation) String() string {
	switch en {
	case OperationAny:
		return "+"
	case OperationCPEStatus:
		return "STATUS"
	case OperationCreate:
		return "C"
	case OperationDelete:
		return "D"
	case OperationJSONRPC:
		return "JSONRPC"
	case OperationLuaScript:
		return "LUA"
	case OperationRead:
		return "R"
	case OperationSHScript:
		return "SH"
	case OperationUpdate:
		return "U"
	}
	panic(errors.New("Invalid value of Operation: " + string(en)))
}

func (en *Operation) MarshalJSON() ([]byte, error) {
	switch *en {
	case OperationAny:
		return json.Marshal("+")
	case OperationCPEStatus:
		return json.Marshal("STATUS")
	case OperationCreate:
		return json.Marshal("C")
	case OperationDelete:
		return json.Marshal("D")
	case OperationJSONRPC:
		return json.Marshal("JSONRPC")
	case OperationLuaScript:
		return json.Marshal("LUA")
	case OperationRead:
		return json.Marshal("R")
	case OperationSHScript:
		return json.Marshal("SH")
	case OperationUpdate:
		return json.Marshal("U")
	}
	return nil, errors.New("Invalid value of Operation: " + string(*en))
}

func (en *Operation) GetBSON() (interface{}, error) {
	switch *en {
	case OperationAny:
		return "+", nil
	case OperationCPEStatus:
		return "STATUS", nil
	case OperationCreate:
		return "C", nil
	case OperationDelete:
		return "D", nil
	case OperationJSONRPC:
		return "JSONRPC", nil
	case OperationLuaScript:
		return "LUA", nil
	case OperationRead:
		return "R", nil
	case OperationSHScript:
		return "SH", nil
	case OperationUpdate:
		return "U", nil
	}
	return nil, errors.New("Invalid value of Operation: " + string(*en))
}

func (en *Operation) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "+":
		*en = OperationAny
		return nil
	case "STATUS":
		*en = OperationCPEStatus
		return nil
	case "C":
		*en = OperationCreate
		return nil
	case "D":
		*en = OperationDelete
		return nil
	case "JSONRPC":
		*en = OperationJSONRPC
		return nil
	case "LUA":
		*en = OperationLuaScript
		return nil
	case "R":
		*en = OperationRead
		return nil
	case "SH":
		*en = OperationSHScript
		return nil
	case "U":
		*en = OperationUpdate
		return nil
	}
	return errors.New("Unknown Operation: " + s)
}

func (en *Operation) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "+":
		*en = OperationAny
		return nil
	case "STATUS":
		*en = OperationCPEStatus
		return nil
	case "C":
		*en = OperationCreate
		return nil
	case "D":
		*en = OperationDelete
		return nil
	case "JSONRPC":
		*en = OperationJSONRPC
		return nil
	case "LUA":
		*en = OperationLuaScript
		return nil
	case "R":
		*en = OperationRead
		return nil
	case "SH":
		*en = OperationSHScript
		return nil
	case "U":
		*en = OperationUpdate
		return nil
	}
	return errors.New("Unknown Operation: " + s)
}

type OperationStatus string

const OperationStatusError OperationStatus = "error"
const OperationStatusPending OperationStatus = "pending"
const OperationStatusSuccess OperationStatus = "success"

func (en OperationStatus) GetPtr() *OperationStatus { var v = en; return &v }

func (en OperationStatus) String() string {
	switch en {
	case OperationStatusError:
		return "error"
	case OperationStatusPending:
		return "pending"
	case OperationStatusSuccess:
		return "success"
	}
	panic(errors.New("Invalid value of OperationStatus: " + string(en)))
}

func (en *OperationStatus) MarshalJSON() ([]byte, error) {
	switch *en {
	case OperationStatusError:
		return json.Marshal("error")
	case OperationStatusPending:
		return json.Marshal("pending")
	case OperationStatusSuccess:
		return json.Marshal("success")
	}
	return nil, errors.New("Invalid value of OperationStatus: " + string(*en))
}

func (en *OperationStatus) GetBSON() (interface{}, error) {
	switch *en {
	case OperationStatusError:
		return "error", nil
	case OperationStatusPending:
		return "pending", nil
	case OperationStatusSuccess:
		return "success", nil
	}
	return nil, errors.New("Invalid value of OperationStatus: " + string(*en))
}

func (en *OperationStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "error":
		*en = OperationStatusError
		return nil
	case "pending":
		*en = OperationStatusPending
		return nil
	case "success":
		*en = OperationStatusSuccess
		return nil
	}
	return errors.New("Unknown OperationStatus: " + s)
}

func (en *OperationStatus) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "error":
		*en = OperationStatusError
		return nil
	case "pending":
		*en = OperationStatusPending
		return nil
	case "success":
		*en = OperationStatusSuccess
		return nil
	}
	return errors.New("Unknown OperationStatus: " + s)
}

type Option82CircuitIDType string

const Option82CircuitIDTypeEmpty Option82CircuitIDType = ""
const Option82CircuitIDTypeIfname Option82CircuitIDType = "IFNAME"
const Option82CircuitIDTypeSsid Option82CircuitIDType = "SSID"

func (en Option82CircuitIDType) GetPtr() *Option82CircuitIDType { var v = en; return &v }

func (en Option82CircuitIDType) String() string {
	switch en {
	case Option82CircuitIDTypeEmpty:
		return ""
	case Option82CircuitIDTypeIfname:
		return "IFNAME"
	case Option82CircuitIDTypeSsid:
		return "SSID"
	}
	if len(en) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of Option82CircuitIDType: " + string(en)))
}

func (en *Option82CircuitIDType) MarshalJSON() ([]byte, error) {
	switch *en {
	case Option82CircuitIDTypeEmpty:
		return json.Marshal("")
	case Option82CircuitIDTypeIfname:
		return json.Marshal("IFNAME")
	case Option82CircuitIDTypeSsid:
		return json.Marshal("SSID")
	}
	if len(*en) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of Option82CircuitIDType: " + string(*en))
}

func (en *Option82CircuitIDType) GetBSON() (interface{}, error) {
	switch *en {
	case Option82CircuitIDTypeEmpty:
		return "", nil
	case Option82CircuitIDTypeIfname:
		return "IFNAME", nil
	case Option82CircuitIDTypeSsid:
		return "SSID", nil
	}
	if len(*en) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of Option82CircuitIDType: " + string(*en))
}

func (en *Option82CircuitIDType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "":
		*en = Option82CircuitIDTypeEmpty
		return nil
	case "IFNAME":
		*en = Option82CircuitIDTypeIfname
		return nil
	case "SSID":
		*en = Option82CircuitIDTypeSsid
		return nil
	}
	if len(s) == 0 {
		*en = Option82CircuitIDTypeEmpty
		return nil
	}
	return errors.New("Unknown Option82CircuitIDType: " + s)
}

func (en *Option82CircuitIDType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "":
		*en = Option82CircuitIDTypeEmpty
		return nil
	case "IFNAME":
		*en = Option82CircuitIDTypeIfname
		return nil
	case "SSID":
		*en = Option82CircuitIDTypeSsid
		return nil
	}
	if len(s) == 0 {
		*en = Option82CircuitIDTypeEmpty
		return nil
	}
	return errors.New("Unknown Option82CircuitIDType: " + s)
}

type Option82RemoteIDType string

const Option82RemoteIDTypeApMac Option82RemoteIDType = "APMAC"
const Option82RemoteIDTypeApMacSSID Option82RemoteIDType = "APMAC:SSID"
const Option82RemoteIDTypeApMacSiteID Option82RemoteIDType = "APMAC:SITEID"
const Option82RemoteIDTypeBSSIDHostname Option82RemoteIDType = "BSSID:HOSTNAME"
const Option82RemoteIDTypeVlanSSID Option82RemoteIDType = "VLAN:SSID"
const Option82RemoteIDTypeWlanIFName Option82RemoteIDType = "WLAN:IFNAME"

func (en Option82RemoteIDType) GetPtr() *Option82RemoteIDType { var v = en; return &v }

func (en Option82RemoteIDType) String() string {
	switch en {
	case Option82RemoteIDTypeApMac:
		return "APMAC"
	case Option82RemoteIDTypeApMacSSID:
		return "APMAC:SSID"
	case Option82RemoteIDTypeApMacSiteID:
		return "APMAC:SITEID"
	case Option82RemoteIDTypeBSSIDHostname:
		return "BSSID:HOSTNAME"
	case Option82RemoteIDTypeVlanSSID:
		return "VLAN:SSID"
	case Option82RemoteIDTypeWlanIFName:
		return "WLAN:IFNAME"
	}
	if len(en) == 0 {
		return "APMAC:SSID"
	}
	panic(errors.New("Invalid value of Option82RemoteIDType: " + string(en)))
}

func (en *Option82RemoteIDType) MarshalJSON() ([]byte, error) {
	switch *en {
	case Option82RemoteIDTypeApMac:
		return json.Marshal("APMAC")
	case Option82RemoteIDTypeApMacSSID:
		return json.Marshal("APMAC:SSID")
	case Option82RemoteIDTypeApMacSiteID:
		return json.Marshal("APMAC:SITEID")
	case Option82RemoteIDTypeBSSIDHostname:
		return json.Marshal("BSSID:HOSTNAME")
	case Option82RemoteIDTypeVlanSSID:
		return json.Marshal("VLAN:SSID")
	case Option82RemoteIDTypeWlanIFName:
		return json.Marshal("WLAN:IFNAME")
	}
	if len(*en) == 0 {
		return json.Marshal("APMAC:SSID")
	}
	return nil, errors.New("Invalid value of Option82RemoteIDType: " + string(*en))
}

func (en *Option82RemoteIDType) GetBSON() (interface{}, error) {
	switch *en {
	case Option82RemoteIDTypeApMac:
		return "APMAC", nil
	case Option82RemoteIDTypeApMacSSID:
		return "APMAC:SSID", nil
	case Option82RemoteIDTypeApMacSiteID:
		return "APMAC:SITEID", nil
	case Option82RemoteIDTypeBSSIDHostname:
		return "BSSID:HOSTNAME", nil
	case Option82RemoteIDTypeVlanSSID:
		return "VLAN:SSID", nil
	case Option82RemoteIDTypeWlanIFName:
		return "WLAN:IFNAME", nil
	}
	if len(*en) == 0 {
		return "APMAC:SSID", nil
	}
	return nil, errors.New("Invalid value of Option82RemoteIDType: " + string(*en))
}

func (en *Option82RemoteIDType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "APMAC":
		*en = Option82RemoteIDTypeApMac
		return nil
	case "APMAC:SSID":
		*en = Option82RemoteIDTypeApMacSSID
		return nil
	case "APMAC:SITEID":
		*en = Option82RemoteIDTypeApMacSiteID
		return nil
	case "BSSID:HOSTNAME":
		*en = Option82RemoteIDTypeBSSIDHostname
		return nil
	case "VLAN:SSID":
		*en = Option82RemoteIDTypeVlanSSID
		return nil
	case "WLAN:IFNAME":
		*en = Option82RemoteIDTypeWlanIFName
		return nil
	}
	if len(s) == 0 {
		*en = Option82RemoteIDTypeApMacSSID
		return nil
	}
	return errors.New("Unknown Option82RemoteIDType: " + s)
}

func (en *Option82RemoteIDType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "APMAC":
		*en = Option82RemoteIDTypeApMac
		return nil
	case "APMAC:SSID":
		*en = Option82RemoteIDTypeApMacSSID
		return nil
	case "APMAC:SITEID":
		*en = Option82RemoteIDTypeApMacSiteID
		return nil
	case "BSSID:HOSTNAME":
		*en = Option82RemoteIDTypeBSSIDHostname
		return nil
	case "VLAN:SSID":
		*en = Option82RemoteIDTypeVlanSSID
		return nil
	case "WLAN:IFNAME":
		*en = Option82RemoteIDTypeWlanIFName
		return nil
	}
	if len(s) == 0 {
		*en = Option82RemoteIDTypeApMacSSID
		return nil
	}
	return errors.New("Unknown Option82RemoteIDType: " + s)
}

type PortalAuthType string

const PortalAuthTypeExternal PortalAuthType = "external"
const PortalAuthTypeNone PortalAuthType = ""
const PortalAuthTypeOAuth2 PortalAuthType = "oauth2"
const PortalAuthTypeRADIUS PortalAuthType = "radius"
const PortalAuthTypeSMS PortalAuthType = "sms"

func (en PortalAuthType) GetPtr() *PortalAuthType { var v = en; return &v }

func (en PortalAuthType) String() string {
	switch en {
	case PortalAuthTypeExternal:
		return "external"
	case PortalAuthTypeNone:
		return ""
	case PortalAuthTypeOAuth2:
		return "oauth2"
	case PortalAuthTypeRADIUS:
		return "radius"
	case PortalAuthTypeSMS:
		return "sms"
	}
	if len(en) == 0 {
		return ""
	}
	panic(errors.New("Invalid value of PortalAuthType: " + string(en)))
}

func (en *PortalAuthType) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalAuthTypeExternal:
		return json.Marshal("external")
	case PortalAuthTypeNone:
		return json.Marshal("")
	case PortalAuthTypeOAuth2:
		return json.Marshal("oauth2")
	case PortalAuthTypeRADIUS:
		return json.Marshal("radius")
	case PortalAuthTypeSMS:
		return json.Marshal("sms")
	}
	if len(*en) == 0 {
		return json.Marshal("")
	}
	return nil, errors.New("Invalid value of PortalAuthType: " + string(*en))
}

func (en *PortalAuthType) GetBSON() (interface{}, error) {
	switch *en {
	case PortalAuthTypeExternal:
		return "external", nil
	case PortalAuthTypeNone:
		return "", nil
	case PortalAuthTypeOAuth2:
		return "oauth2", nil
	case PortalAuthTypeRADIUS:
		return "radius", nil
	case PortalAuthTypeSMS:
		return "sms", nil
	}
	if len(*en) == 0 {
		return "", nil
	}
	return nil, errors.New("Invalid value of PortalAuthType: " + string(*en))
}

func (en *PortalAuthType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "external":
		*en = PortalAuthTypeExternal
		return nil
	case "":
		*en = PortalAuthTypeNone
		return nil
	case "oauth2":
		*en = PortalAuthTypeOAuth2
		return nil
	case "radius":
		*en = PortalAuthTypeRADIUS
		return nil
	case "sms":
		*en = PortalAuthTypeSMS
		return nil
	}
	if len(s) == 0 {
		*en = PortalAuthTypeNone
		return nil
	}
	return errors.New("Unknown PortalAuthType: " + s)
}

func (en *PortalAuthType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "external":
		*en = PortalAuthTypeExternal
		return nil
	case "":
		*en = PortalAuthTypeNone
		return nil
	case "oauth2":
		*en = PortalAuthTypeOAuth2
		return nil
	case "radius":
		*en = PortalAuthTypeRADIUS
		return nil
	case "sms":
		*en = PortalAuthTypeSMS
		return nil
	}
	if len(s) == 0 {
		*en = PortalAuthTypeNone
		return nil
	}
	return errors.New("Unknown PortalAuthType: " + s)
}

type PortalProfileType string

const PortalProfileTypeFree PortalProfileType = "free"
const PortalProfileTypePremium PortalProfileType = "premium"
const PortalProfileTypeSponsor PortalProfileType = "sponsor"

func (en PortalProfileType) GetPtr() *PortalProfileType { var v = en; return &v }

func (en PortalProfileType) String() string {
	switch en {
	case PortalProfileTypeFree:
		return "free"
	case PortalProfileTypePremium:
		return "premium"
	case PortalProfileTypeSponsor:
		return "sponsor"
	}
	if len(en) == 0 {
		return "free"
	}
	panic(errors.New("Invalid value of PortalProfileType: " + string(en)))
}

func (en *PortalProfileType) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalProfileTypeFree:
		return json.Marshal("free")
	case PortalProfileTypePremium:
		return json.Marshal("premium")
	case PortalProfileTypeSponsor:
		return json.Marshal("sponsor")
	}
	if len(*en) == 0 {
		return json.Marshal("free")
	}
	return nil, errors.New("Invalid value of PortalProfileType: " + string(*en))
}

func (en *PortalProfileType) GetBSON() (interface{}, error) {
	switch *en {
	case PortalProfileTypeFree:
		return "free", nil
	case PortalProfileTypePremium:
		return "premium", nil
	case PortalProfileTypeSponsor:
		return "sponsor", nil
	}
	if len(*en) == 0 {
		return "free", nil
	}
	return nil, errors.New("Invalid value of PortalProfileType: " + string(*en))
}

func (en *PortalProfileType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "free":
		*en = PortalProfileTypeFree
		return nil
	case "premium":
		*en = PortalProfileTypePremium
		return nil
	case "sponsor":
		*en = PortalProfileTypeSponsor
		return nil
	}
	if len(s) == 0 {
		*en = PortalProfileTypeFree
		return nil
	}
	return errors.New("Unknown PortalProfileType: " + s)
}

func (en *PortalProfileType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "free":
		*en = PortalProfileTypeFree
		return nil
	case "premium":
		*en = PortalProfileTypePremium
		return nil
	case "sponsor":
		*en = PortalProfileTypeSponsor
		return nil
	}
	if len(s) == 0 {
		*en = PortalProfileTypeFree
		return nil
	}
	return errors.New("Unknown PortalProfileType: " + s)
}

type RRMAlgoType string

const RRMAlgoTypeBlind RRMAlgoType = "Blind"
const RRMAlgoTypeDummy RRMAlgoType = "Dummy"
const RRMAlgoTypeGreed RRMAlgoType = "Greed"

func (en RRMAlgoType) GetPtr() *RRMAlgoType { var v = en; return &v }

func (en RRMAlgoType) String() string {
	switch en {
	case RRMAlgoTypeBlind:
		return "Blind"
	case RRMAlgoTypeDummy:
		return "Dummy"
	case RRMAlgoTypeGreed:
		return "Greed"
	}
	if len(en) == 0 {
		return "Greed"
	}
	panic(errors.New("Invalid value of RRMAlgoType: " + string(en)))
}

func (en *RRMAlgoType) MarshalJSON() ([]byte, error) {
	switch *en {
	case RRMAlgoTypeBlind:
		return json.Marshal("Blind")
	case RRMAlgoTypeDummy:
		return json.Marshal("Dummy")
	case RRMAlgoTypeGreed:
		return json.Marshal("Greed")
	}
	if len(*en) == 0 {
		return json.Marshal("Greed")
	}
	return nil, errors.New("Invalid value of RRMAlgoType: " + string(*en))
}

func (en *RRMAlgoType) GetBSON() (interface{}, error) {
	switch *en {
	case RRMAlgoTypeBlind:
		return "Blind", nil
	case RRMAlgoTypeDummy:
		return "Dummy", nil
	case RRMAlgoTypeGreed:
		return "Greed", nil
	}
	if len(*en) == 0 {
		return "Greed", nil
	}
	return nil, errors.New("Invalid value of RRMAlgoType: " + string(*en))
}

func (en *RRMAlgoType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "Blind":
		*en = RRMAlgoTypeBlind
		return nil
	case "Dummy":
		*en = RRMAlgoTypeDummy
		return nil
	case "Greed":
		*en = RRMAlgoTypeGreed
		return nil
	}
	if len(s) == 0 {
		*en = RRMAlgoTypeGreed
		return nil
	}
	return errors.New("Unknown RRMAlgoType: " + s)
}

func (en *RRMAlgoType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "Blind":
		*en = RRMAlgoTypeBlind
		return nil
	case "Dummy":
		*en = RRMAlgoTypeDummy
		return nil
	case "Greed":
		*en = RRMAlgoTypeGreed
		return nil
	}
	if len(s) == 0 {
		*en = RRMAlgoTypeGreed
		return nil
	}
	return errors.New("Unknown RRMAlgoType: " + s)
}

type RadarExportFilter string

const RadarExportFilterAll RadarExportFilter = "all"
const RadarExportFilterNew RadarExportFilter = "new"
const RadarExportFilterReturn RadarExportFilter = "return"

func (en RadarExportFilter) GetPtr() *RadarExportFilter { var v = en; return &v }

func (en RadarExportFilter) String() string {
	switch en {
	case RadarExportFilterAll:
		return "all"
	case RadarExportFilterNew:
		return "new"
	case RadarExportFilterReturn:
		return "return"
	}
	if len(en) == 0 {
		return "all"
	}
	panic(errors.New("Invalid value of RadarExportFilter: " + string(en)))
}

func (en *RadarExportFilter) MarshalJSON() ([]byte, error) {
	switch *en {
	case RadarExportFilterAll:
		return json.Marshal("all")
	case RadarExportFilterNew:
		return json.Marshal("new")
	case RadarExportFilterReturn:
		return json.Marshal("return")
	}
	if len(*en) == 0 {
		return json.Marshal("all")
	}
	return nil, errors.New("Invalid value of RadarExportFilter: " + string(*en))
}

func (en *RadarExportFilter) GetBSON() (interface{}, error) {
	switch *en {
	case RadarExportFilterAll:
		return "all", nil
	case RadarExportFilterNew:
		return "new", nil
	case RadarExportFilterReturn:
		return "return", nil
	}
	if len(*en) == 0 {
		return "all", nil
	}
	return nil, errors.New("Invalid value of RadarExportFilter: " + string(*en))
}

func (en *RadarExportFilter) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "all":
		*en = RadarExportFilterAll
		return nil
	case "new":
		*en = RadarExportFilterNew
		return nil
	case "return":
		*en = RadarExportFilterReturn
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportFilterAll
		return nil
	}
	return errors.New("Unknown RadarExportFilter: " + s)
}

func (en *RadarExportFilter) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "all":
		*en = RadarExportFilterAll
		return nil
	case "new":
		*en = RadarExportFilterNew
		return nil
	case "return":
		*en = RadarExportFilterReturn
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportFilterAll
		return nil
	}
	return errors.New("Unknown RadarExportFilter: " + s)
}

type RadarExportFormat string

const RadarExportFormatCSV RadarExportFormat = "csv"
const RadarExportFormatJson RadarExportFormat = "json"
const RadarExportFormatTxt RadarExportFormat = "txt"

func (en RadarExportFormat) GetPtr() *RadarExportFormat { var v = en; return &v }

func (en RadarExportFormat) String() string {
	switch en {
	case RadarExportFormatCSV:
		return "csv"
	case RadarExportFormatJson:
		return "json"
	case RadarExportFormatTxt:
		return "txt"
	}
	if len(en) == 0 {
		return "csv"
	}
	panic(errors.New("Invalid value of RadarExportFormat: " + string(en)))
}

func (en *RadarExportFormat) MarshalJSON() ([]byte, error) {
	switch *en {
	case RadarExportFormatCSV:
		return json.Marshal("csv")
	case RadarExportFormatJson:
		return json.Marshal("json")
	case RadarExportFormatTxt:
		return json.Marshal("txt")
	}
	if len(*en) == 0 {
		return json.Marshal("csv")
	}
	return nil, errors.New("Invalid value of RadarExportFormat: " + string(*en))
}

func (en *RadarExportFormat) GetBSON() (interface{}, error) {
	switch *en {
	case RadarExportFormatCSV:
		return "csv", nil
	case RadarExportFormatJson:
		return "json", nil
	case RadarExportFormatTxt:
		return "txt", nil
	}
	if len(*en) == 0 {
		return "csv", nil
	}
	return nil, errors.New("Invalid value of RadarExportFormat: " + string(*en))
}

func (en *RadarExportFormat) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "csv":
		*en = RadarExportFormatCSV
		return nil
	case "json":
		*en = RadarExportFormatJson
		return nil
	case "txt":
		*en = RadarExportFormatTxt
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportFormatCSV
		return nil
	}
	return errors.New("Unknown RadarExportFormat: " + s)
}

func (en *RadarExportFormat) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "csv":
		*en = RadarExportFormatCSV
		return nil
	case "json":
		*en = RadarExportFormatJson
		return nil
	case "txt":
		*en = RadarExportFormatTxt
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportFormatCSV
		return nil
	}
	return errors.New("Unknown RadarExportFormat: " + s)
}

type RadarExportMacs string

const RadarExportMacsAll RadarExportMacs = "all"
const RadarExportMacsFake RadarExportMacs = "fake"
const RadarExportMacsReal RadarExportMacs = "real"

func (en RadarExportMacs) GetPtr() *RadarExportMacs { var v = en; return &v }

func (en RadarExportMacs) String() string {
	switch en {
	case RadarExportMacsAll:
		return "all"
	case RadarExportMacsFake:
		return "fake"
	case RadarExportMacsReal:
		return "real"
	}
	if len(en) == 0 {
		return "all"
	}
	panic(errors.New("Invalid value of RadarExportMacs: " + string(en)))
}

func (en *RadarExportMacs) MarshalJSON() ([]byte, error) {
	switch *en {
	case RadarExportMacsAll:
		return json.Marshal("all")
	case RadarExportMacsFake:
		return json.Marshal("fake")
	case RadarExportMacsReal:
		return json.Marshal("real")
	}
	if len(*en) == 0 {
		return json.Marshal("all")
	}
	return nil, errors.New("Invalid value of RadarExportMacs: " + string(*en))
}

func (en *RadarExportMacs) GetBSON() (interface{}, error) {
	switch *en {
	case RadarExportMacsAll:
		return "all", nil
	case RadarExportMacsFake:
		return "fake", nil
	case RadarExportMacsReal:
		return "real", nil
	}
	if len(*en) == 0 {
		return "all", nil
	}
	return nil, errors.New("Invalid value of RadarExportMacs: " + string(*en))
}

func (en *RadarExportMacs) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "all":
		*en = RadarExportMacsAll
		return nil
	case "fake":
		*en = RadarExportMacsFake
		return nil
	case "real":
		*en = RadarExportMacsReal
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportMacsAll
		return nil
	}
	return errors.New("Unknown RadarExportMacs: " + s)
}

func (en *RadarExportMacs) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "all":
		*en = RadarExportMacsAll
		return nil
	case "fake":
		*en = RadarExportMacsFake
		return nil
	case "real":
		*en = RadarExportMacsReal
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportMacsAll
		return nil
	}
	return errors.New("Unknown RadarExportMacs: " + s)
}

type RadarExportStatus string

const RadarExportStatusCreated RadarExportStatus = "created"
const RadarExportStatusFinished RadarExportStatus = "finished"
const RadarExportStatusRunning RadarExportStatus = "running"
const RadarExportStatusUpdated RadarExportStatus = "updated"

func (en RadarExportStatus) GetPtr() *RadarExportStatus { var v = en; return &v }

func (en RadarExportStatus) String() string {
	switch en {
	case RadarExportStatusCreated:
		return "created"
	case RadarExportStatusFinished:
		return "finished"
	case RadarExportStatusRunning:
		return "running"
	case RadarExportStatusUpdated:
		return "updated"
	}
	if len(en) == 0 {
		return "created"
	}
	panic(errors.New("Invalid value of RadarExportStatus: " + string(en)))
}

func (en *RadarExportStatus) MarshalJSON() ([]byte, error) {
	switch *en {
	case RadarExportStatusCreated:
		return json.Marshal("created")
	case RadarExportStatusFinished:
		return json.Marshal("finished")
	case RadarExportStatusRunning:
		return json.Marshal("running")
	case RadarExportStatusUpdated:
		return json.Marshal("updated")
	}
	if len(*en) == 0 {
		return json.Marshal("created")
	}
	return nil, errors.New("Invalid value of RadarExportStatus: " + string(*en))
}

func (en *RadarExportStatus) GetBSON() (interface{}, error) {
	switch *en {
	case RadarExportStatusCreated:
		return "created", nil
	case RadarExportStatusFinished:
		return "finished", nil
	case RadarExportStatusRunning:
		return "running", nil
	case RadarExportStatusUpdated:
		return "updated", nil
	}
	if len(*en) == 0 {
		return "created", nil
	}
	return nil, errors.New("Invalid value of RadarExportStatus: " + string(*en))
}

func (en *RadarExportStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "created":
		*en = RadarExportStatusCreated
		return nil
	case "finished":
		*en = RadarExportStatusFinished
		return nil
	case "running":
		*en = RadarExportStatusRunning
		return nil
	case "updated":
		*en = RadarExportStatusUpdated
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportStatusCreated
		return nil
	}
	return errors.New("Unknown RadarExportStatus: " + s)
}

func (en *RadarExportStatus) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "created":
		*en = RadarExportStatusCreated
		return nil
	case "finished":
		*en = RadarExportStatusFinished
		return nil
	case "running":
		*en = RadarExportStatusRunning
		return nil
	case "updated":
		*en = RadarExportStatusUpdated
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportStatusCreated
		return nil
	}
	return errors.New("Unknown RadarExportStatus: " + s)
}

type RadarExportType string

const RadarExportTypeBeePro RadarExportType = "beepro"
const RadarExportTypeEmail RadarExportType = "email"
const RadarExportTypeExternal RadarExportType = "external"
const RadarExportTypeMytarget RadarExportType = "mytarget"
const RadarExportTypeYandex RadarExportType = "yandex"

func (en RadarExportType) GetPtr() *RadarExportType { var v = en; return &v }

func (en RadarExportType) String() string {
	switch en {
	case RadarExportTypeBeePro:
		return "beepro"
	case RadarExportTypeEmail:
		return "email"
	case RadarExportTypeExternal:
		return "external"
	case RadarExportTypeMytarget:
		return "mytarget"
	case RadarExportTypeYandex:
		return "yandex"
	}
	if len(en) == 0 {
		return "email"
	}
	panic(errors.New("Invalid value of RadarExportType: " + string(en)))
}

func (en *RadarExportType) MarshalJSON() ([]byte, error) {
	switch *en {
	case RadarExportTypeBeePro:
		return json.Marshal("beepro")
	case RadarExportTypeEmail:
		return json.Marshal("email")
	case RadarExportTypeExternal:
		return json.Marshal("external")
	case RadarExportTypeMytarget:
		return json.Marshal("mytarget")
	case RadarExportTypeYandex:
		return json.Marshal("yandex")
	}
	if len(*en) == 0 {
		return json.Marshal("email")
	}
	return nil, errors.New("Invalid value of RadarExportType: " + string(*en))
}

func (en *RadarExportType) GetBSON() (interface{}, error) {
	switch *en {
	case RadarExportTypeBeePro:
		return "beepro", nil
	case RadarExportTypeEmail:
		return "email", nil
	case RadarExportTypeExternal:
		return "external", nil
	case RadarExportTypeMytarget:
		return "mytarget", nil
	case RadarExportTypeYandex:
		return "yandex", nil
	}
	if len(*en) == 0 {
		return "email", nil
	}
	return nil, errors.New("Invalid value of RadarExportType: " + string(*en))
}

func (en *RadarExportType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "beepro":
		*en = RadarExportTypeBeePro
		return nil
	case "email":
		*en = RadarExportTypeEmail
		return nil
	case "external":
		*en = RadarExportTypeExternal
		return nil
	case "mytarget":
		*en = RadarExportTypeMytarget
		return nil
	case "yandex":
		*en = RadarExportTypeYandex
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportTypeEmail
		return nil
	}
	return errors.New("Unknown RadarExportType: " + s)
}

func (en *RadarExportType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "beepro":
		*en = RadarExportTypeBeePro
		return nil
	case "email":
		*en = RadarExportTypeEmail
		return nil
	case "external":
		*en = RadarExportTypeExternal
		return nil
	case "mytarget":
		*en = RadarExportTypeMytarget
		return nil
	case "yandex":
		*en = RadarExportTypeYandex
		return nil
	}
	if len(s) == 0 {
		*en = RadarExportTypeEmail
		return nil
	}
	return errors.New("Unknown RadarExportType: " + s)
}

type RadiusMessageType string

const RadiusMessageTypeAccessAccept RadiusMessageType = "access-accept"
const RadiusMessageTypeAccessReject RadiusMessageType = "access-reject"
const RadiusMessageTypeAccessRequest RadiusMessageType = "access-request"
const RadiusMessageTypeAccountingRequest RadiusMessageType = "accounting"

func (en RadiusMessageType) GetPtr() *RadiusMessageType { var v = en; return &v }

func (en RadiusMessageType) String() string {
	switch en {
	case RadiusMessageTypeAccessAccept:
		return "access-accept"
	case RadiusMessageTypeAccessReject:
		return "access-reject"
	case RadiusMessageTypeAccessRequest:
		return "access-request"
	case RadiusMessageTypeAccountingRequest:
		return "accounting"
	}
	if len(en) == 0 {
		return "accounting"
	}
	panic(errors.New("Invalid value of RadiusMessageType: " + string(en)))
}

func (en *RadiusMessageType) MarshalJSON() ([]byte, error) {
	switch *en {
	case RadiusMessageTypeAccessAccept:
		return json.Marshal("access-accept")
	case RadiusMessageTypeAccessReject:
		return json.Marshal("access-reject")
	case RadiusMessageTypeAccessRequest:
		return json.Marshal("access-request")
	case RadiusMessageTypeAccountingRequest:
		return json.Marshal("accounting")
	}
	if len(*en) == 0 {
		return json.Marshal("accounting")
	}
	return nil, errors.New("Invalid value of RadiusMessageType: " + string(*en))
}

func (en *RadiusMessageType) GetBSON() (interface{}, error) {
	switch *en {
	case RadiusMessageTypeAccessAccept:
		return "access-accept", nil
	case RadiusMessageTypeAccessReject:
		return "access-reject", nil
	case RadiusMessageTypeAccessRequest:
		return "access-request", nil
	case RadiusMessageTypeAccountingRequest:
		return "accounting", nil
	}
	if len(*en) == 0 {
		return "accounting", nil
	}
	return nil, errors.New("Invalid value of RadiusMessageType: " + string(*en))
}

func (en *RadiusMessageType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "access-accept":
		*en = RadiusMessageTypeAccessAccept
		return nil
	case "access-reject":
		*en = RadiusMessageTypeAccessReject
		return nil
	case "access-request":
		*en = RadiusMessageTypeAccessRequest
		return nil
	case "accounting":
		*en = RadiusMessageTypeAccountingRequest
		return nil
	}
	if len(s) == 0 {
		*en = RadiusMessageTypeAccountingRequest
		return nil
	}
	return errors.New("Unknown RadiusMessageType: " + s)
}

func (en *RadiusMessageType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "access-accept":
		*en = RadiusMessageTypeAccessAccept
		return nil
	case "access-reject":
		*en = RadiusMessageTypeAccessReject
		return nil
	case "access-request":
		*en = RadiusMessageTypeAccessRequest
		return nil
	case "accounting":
		*en = RadiusMessageTypeAccountingRequest
		return nil
	}
	if len(s) == 0 {
		*en = RadiusMessageTypeAccountingRequest
		return nil
	}
	return errors.New("Unknown RadiusMessageType: " + s)
}

type ReportActionType string

const ReportActionTypeEmail ReportActionType = "email"
const ReportActionTypeScript ReportActionType = "script"

func (en ReportActionType) GetPtr() *ReportActionType { var v = en; return &v }

func (en ReportActionType) String() string {
	switch en {
	case ReportActionTypeEmail:
		return "email"
	case ReportActionTypeScript:
		return "script"
	}
	if len(en) == 0 {
		return "email"
	}
	panic(errors.New("Invalid value of ReportActionType: " + string(en)))
}

func (en *ReportActionType) MarshalJSON() ([]byte, error) {
	switch *en {
	case ReportActionTypeEmail:
		return json.Marshal("email")
	case ReportActionTypeScript:
		return json.Marshal("script")
	}
	if len(*en) == 0 {
		return json.Marshal("email")
	}
	return nil, errors.New("Invalid value of ReportActionType: " + string(*en))
}

func (en *ReportActionType) GetBSON() (interface{}, error) {
	switch *en {
	case ReportActionTypeEmail:
		return "email", nil
	case ReportActionTypeScript:
		return "script", nil
	}
	if len(*en) == 0 {
		return "email", nil
	}
	return nil, errors.New("Invalid value of ReportActionType: " + string(*en))
}

func (en *ReportActionType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "email":
		*en = ReportActionTypeEmail
		return nil
	case "script":
		*en = ReportActionTypeScript
		return nil
	}
	if len(s) == 0 {
		*en = ReportActionTypeEmail
		return nil
	}
	return errors.New("Unknown ReportActionType: " + s)
}

func (en *ReportActionType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "email":
		*en = ReportActionTypeEmail
		return nil
	case "script":
		*en = ReportActionTypeScript
		return nil
	}
	if len(s) == 0 {
		*en = ReportActionTypeEmail
		return nil
	}
	return errors.New("Unknown ReportActionType: " + s)
}

type ReportFormat string

const ReportFormatCSV ReportFormat = "csv"
const ReportFormatJson ReportFormat = "json"
const ReportFormatTxt ReportFormat = "txt"
const ReportFormatXLSX ReportFormat = "xlsx"

func (en ReportFormat) GetPtr() *ReportFormat { var v = en; return &v }

func (en ReportFormat) String() string {
	switch en {
	case ReportFormatCSV:
		return "csv"
	case ReportFormatJson:
		return "json"
	case ReportFormatTxt:
		return "txt"
	case ReportFormatXLSX:
		return "xlsx"
	}
	if len(en) == 0 {
		return "json"
	}
	panic(errors.New("Invalid value of ReportFormat: " + string(en)))
}

func (en *ReportFormat) MarshalJSON() ([]byte, error) {
	switch *en {
	case ReportFormatCSV:
		return json.Marshal("csv")
	case ReportFormatJson:
		return json.Marshal("json")
	case ReportFormatTxt:
		return json.Marshal("txt")
	case ReportFormatXLSX:
		return json.Marshal("xlsx")
	}
	if len(*en) == 0 {
		return json.Marshal("json")
	}
	return nil, errors.New("Invalid value of ReportFormat: " + string(*en))
}

func (en *ReportFormat) GetBSON() (interface{}, error) {
	switch *en {
	case ReportFormatCSV:
		return "csv", nil
	case ReportFormatJson:
		return "json", nil
	case ReportFormatTxt:
		return "txt", nil
	case ReportFormatXLSX:
		return "xlsx", nil
	}
	if len(*en) == 0 {
		return "json", nil
	}
	return nil, errors.New("Invalid value of ReportFormat: " + string(*en))
}

func (en *ReportFormat) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "csv":
		*en = ReportFormatCSV
		return nil
	case "json":
		*en = ReportFormatJson
		return nil
	case "txt":
		*en = ReportFormatTxt
		return nil
	case "xlsx":
		*en = ReportFormatXLSX
		return nil
	}
	if len(s) == 0 {
		*en = ReportFormatJson
		return nil
	}
	return errors.New("Unknown ReportFormat: " + s)
}

func (en *ReportFormat) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "csv":
		*en = ReportFormatCSV
		return nil
	case "json":
		*en = ReportFormatJson
		return nil
	case "txt":
		*en = ReportFormatTxt
		return nil
	case "xlsx":
		*en = ReportFormatXLSX
		return nil
	}
	if len(s) == 0 {
		*en = ReportFormatJson
		return nil
	}
	return errors.New("Unknown ReportFormat: " + s)
}

type ReportPeriod string

const ReportPeriodDay ReportPeriod = "day"
const ReportPeriodMonth ReportPeriod = "month"
const ReportPeriodNow ReportPeriod = "now"
const ReportPeriodWeek ReportPeriod = "week"

func (en ReportPeriod) GetPtr() *ReportPeriod { var v = en; return &v }

func (en ReportPeriod) String() string {
	switch en {
	case ReportPeriodDay:
		return "day"
	case ReportPeriodMonth:
		return "month"
	case ReportPeriodNow:
		return "now"
	case ReportPeriodWeek:
		return "week"
	}
	if len(en) == 0 {
		return "now"
	}
	panic(errors.New("Invalid value of ReportPeriod: " + string(en)))
}

func (en *ReportPeriod) MarshalJSON() ([]byte, error) {
	switch *en {
	case ReportPeriodDay:
		return json.Marshal("day")
	case ReportPeriodMonth:
		return json.Marshal("month")
	case ReportPeriodNow:
		return json.Marshal("now")
	case ReportPeriodWeek:
		return json.Marshal("week")
	}
	if len(*en) == 0 {
		return json.Marshal("now")
	}
	return nil, errors.New("Invalid value of ReportPeriod: " + string(*en))
}

func (en *ReportPeriod) GetBSON() (interface{}, error) {
	switch *en {
	case ReportPeriodDay:
		return "day", nil
	case ReportPeriodMonth:
		return "month", nil
	case ReportPeriodNow:
		return "now", nil
	case ReportPeriodWeek:
		return "week", nil
	}
	if len(*en) == 0 {
		return "now", nil
	}
	return nil, errors.New("Invalid value of ReportPeriod: " + string(*en))
}

func (en *ReportPeriod) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "day":
		*en = ReportPeriodDay
		return nil
	case "month":
		*en = ReportPeriodMonth
		return nil
	case "now":
		*en = ReportPeriodNow
		return nil
	case "week":
		*en = ReportPeriodWeek
		return nil
	}
	if len(s) == 0 {
		*en = ReportPeriodNow
		return nil
	}
	return errors.New("Unknown ReportPeriod: " + s)
}

func (en *ReportPeriod) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "day":
		*en = ReportPeriodDay
		return nil
	case "month":
		*en = ReportPeriodMonth
		return nil
	case "now":
		*en = ReportPeriodNow
		return nil
	case "week":
		*en = ReportPeriodWeek
		return nil
	}
	if len(s) == 0 {
		*en = ReportPeriodNow
		return nil
	}
	return errors.New("Unknown ReportPeriod: " + s)
}

type ReportSubject string

const ReportSubjectCPECommon ReportSubject = "cpes_common"
const ReportSubjectCPEs ReportSubject = "cpes"
const ReportSubjectClients ReportSubject = "clients"
const ReportSubjectEvents ReportSubject = "events"

func (en ReportSubject) GetPtr() *ReportSubject { var v = en; return &v }

func (en ReportSubject) String() string {
	switch en {
	case ReportSubjectCPECommon:
		return "cpes_common"
	case ReportSubjectCPEs:
		return "cpes"
	case ReportSubjectClients:
		return "clients"
	case ReportSubjectEvents:
		return "events"
	}
	if len(en) == 0 {
		return "clients"
	}
	panic(errors.New("Invalid value of ReportSubject: " + string(en)))
}

func (en *ReportSubject) MarshalJSON() ([]byte, error) {
	switch *en {
	case ReportSubjectCPECommon:
		return json.Marshal("cpes_common")
	case ReportSubjectCPEs:
		return json.Marshal("cpes")
	case ReportSubjectClients:
		return json.Marshal("clients")
	case ReportSubjectEvents:
		return json.Marshal("events")
	}
	if len(*en) == 0 {
		return json.Marshal("clients")
	}
	return nil, errors.New("Invalid value of ReportSubject: " + string(*en))
}

func (en *ReportSubject) GetBSON() (interface{}, error) {
	switch *en {
	case ReportSubjectCPECommon:
		return "cpes_common", nil
	case ReportSubjectCPEs:
		return "cpes", nil
	case ReportSubjectClients:
		return "clients", nil
	case ReportSubjectEvents:
		return "events", nil
	}
	if len(*en) == 0 {
		return "clients", nil
	}
	return nil, errors.New("Invalid value of ReportSubject: " + string(*en))
}

func (en *ReportSubject) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "cpes_common":
		*en = ReportSubjectCPECommon
		return nil
	case "cpes":
		*en = ReportSubjectCPEs
		return nil
	case "clients":
		*en = ReportSubjectClients
		return nil
	case "events":
		*en = ReportSubjectEvents
		return nil
	}
	if len(s) == 0 {
		*en = ReportSubjectClients
		return nil
	}
	return errors.New("Unknown ReportSubject: " + s)
}

func (en *ReportSubject) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "cpes_common":
		*en = ReportSubjectCPECommon
		return nil
	case "cpes":
		*en = ReportSubjectCPEs
		return nil
	case "clients":
		*en = ReportSubjectClients
		return nil
	case "events":
		*en = ReportSubjectEvents
		return nil
	}
	if len(s) == 0 {
		*en = ReportSubjectClients
		return nil
	}
	return errors.New("Unknown ReportSubject: " + s)
}

type ReportType string

const ReportTypeClientsAll ReportType = "clients"
const ReportTypeClientsAuthorized ReportType = "clients_auth"
const ReportTypeCurrent ReportType = "current"
const ReportTypeSummary ReportType = "summary"

func (en ReportType) GetPtr() *ReportType { var v = en; return &v }

func (en ReportType) String() string {
	switch en {
	case ReportTypeClientsAll:
		return "clients"
	case ReportTypeClientsAuthorized:
		return "clients_auth"
	case ReportTypeCurrent:
		return "current"
	case ReportTypeSummary:
		return "summary"
	}
	if len(en) == 0 {
		return "current"
	}
	panic(errors.New("Invalid value of ReportType: " + string(en)))
}

func (en *ReportType) MarshalJSON() ([]byte, error) {
	switch *en {
	case ReportTypeClientsAll:
		return json.Marshal("clients")
	case ReportTypeClientsAuthorized:
		return json.Marshal("clients_auth")
	case ReportTypeCurrent:
		return json.Marshal("current")
	case ReportTypeSummary:
		return json.Marshal("summary")
	}
	if len(*en) == 0 {
		return json.Marshal("current")
	}
	return nil, errors.New("Invalid value of ReportType: " + string(*en))
}

func (en *ReportType) GetBSON() (interface{}, error) {
	switch *en {
	case ReportTypeClientsAll:
		return "clients", nil
	case ReportTypeClientsAuthorized:
		return "clients_auth", nil
	case ReportTypeCurrent:
		return "current", nil
	case ReportTypeSummary:
		return "summary", nil
	}
	if len(*en) == 0 {
		return "current", nil
	}
	return nil, errors.New("Invalid value of ReportType: " + string(*en))
}

func (en *ReportType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "clients":
		*en = ReportTypeClientsAll
		return nil
	case "clients_auth":
		*en = ReportTypeClientsAuthorized
		return nil
	case "current":
		*en = ReportTypeCurrent
		return nil
	case "summary":
		*en = ReportTypeSummary
		return nil
	}
	if len(s) == 0 {
		*en = ReportTypeCurrent
		return nil
	}
	return errors.New("Unknown ReportType: " + s)
}

func (en *ReportType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "clients":
		*en = ReportTypeClientsAll
		return nil
	case "clients_auth":
		*en = ReportTypeClientsAuthorized
		return nil
	case "current":
		*en = ReportTypeCurrent
		return nil
	case "summary":
		*en = ReportTypeSummary
		return nil
	}
	if len(s) == 0 {
		*en = ReportTypeCurrent
		return nil
	}
	return errors.New("Unknown ReportType: " + s)
}

type SNMPAccessMode string

const SNMPAccessModeRead SNMPAccessMode = "read"
const SNMPAccessModeReadWrite SNMPAccessMode = "read/write"

func (en SNMPAccessMode) GetPtr() *SNMPAccessMode { var v = en; return &v }

func (en SNMPAccessMode) String() string {
	switch en {
	case SNMPAccessModeRead:
		return "read"
	case SNMPAccessModeReadWrite:
		return "read/write"
	}
	if len(en) == 0 {
		return "read"
	}
	panic(errors.New("Invalid value of SNMPAccessMode: " + string(en)))
}

func (en *SNMPAccessMode) MarshalJSON() ([]byte, error) {
	switch *en {
	case SNMPAccessModeRead:
		return json.Marshal("read")
	case SNMPAccessModeReadWrite:
		return json.Marshal("read/write")
	}
	if len(*en) == 0 {
		return json.Marshal("read")
	}
	return nil, errors.New("Invalid value of SNMPAccessMode: " + string(*en))
}

func (en *SNMPAccessMode) GetBSON() (interface{}, error) {
	switch *en {
	case SNMPAccessModeRead:
		return "read", nil
	case SNMPAccessModeReadWrite:
		return "read/write", nil
	}
	if len(*en) == 0 {
		return "read", nil
	}
	return nil, errors.New("Invalid value of SNMPAccessMode: " + string(*en))
}

func (en *SNMPAccessMode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "read":
		*en = SNMPAccessModeRead
		return nil
	case "read/write":
		*en = SNMPAccessModeReadWrite
		return nil
	}
	if len(s) == 0 {
		*en = SNMPAccessModeRead
		return nil
	}
	return errors.New("Unknown SNMPAccessMode: " + s)
}

func (en *SNMPAccessMode) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "read":
		*en = SNMPAccessModeRead
		return nil
	case "read/write":
		*en = SNMPAccessModeReadWrite
		return nil
	}
	if len(s) == 0 {
		*en = SNMPAccessModeRead
		return nil
	}
	return errors.New("Unknown SNMPAccessMode: " + s)
}

type SNMPAuthProtocol string

const SNMPAuthProtocolMD5 SNMPAuthProtocol = "MD5"
const SNMPAuthProtocolSHA SNMPAuthProtocol = "SHA"
const SNMPAuthProtocolSHA224 SNMPAuthProtocol = "SHA-224"
const SNMPAuthProtocolSHA256 SNMPAuthProtocol = "SHA-256"
const SNMPAuthProtocolSHA384 SNMPAuthProtocol = "SHA-384"
const SNMPAuthProtocolSHA512 SNMPAuthProtocol = "SHA-512"

func (en SNMPAuthProtocol) GetPtr() *SNMPAuthProtocol { var v = en; return &v }

func (en SNMPAuthProtocol) String() string {
	switch en {
	case SNMPAuthProtocolMD5:
		return "MD5"
	case SNMPAuthProtocolSHA:
		return "SHA"
	case SNMPAuthProtocolSHA224:
		return "SHA-224"
	case SNMPAuthProtocolSHA256:
		return "SHA-256"
	case SNMPAuthProtocolSHA384:
		return "SHA-384"
	case SNMPAuthProtocolSHA512:
		return "SHA-512"
	}
	if len(en) == 0 {
		return "MD5"
	}
	panic(errors.New("Invalid value of SNMPAuthProtocol: " + string(en)))
}

func (en *SNMPAuthProtocol) MarshalJSON() ([]byte, error) {
	switch *en {
	case SNMPAuthProtocolMD5:
		return json.Marshal("MD5")
	case SNMPAuthProtocolSHA:
		return json.Marshal("SHA")
	case SNMPAuthProtocolSHA224:
		return json.Marshal("SHA-224")
	case SNMPAuthProtocolSHA256:
		return json.Marshal("SHA-256")
	case SNMPAuthProtocolSHA384:
		return json.Marshal("SHA-384")
	case SNMPAuthProtocolSHA512:
		return json.Marshal("SHA-512")
	}
	if len(*en) == 0 {
		return json.Marshal("MD5")
	}
	return nil, errors.New("Invalid value of SNMPAuthProtocol: " + string(*en))
}

func (en *SNMPAuthProtocol) GetBSON() (interface{}, error) {
	switch *en {
	case SNMPAuthProtocolMD5:
		return "MD5", nil
	case SNMPAuthProtocolSHA:
		return "SHA", nil
	case SNMPAuthProtocolSHA224:
		return "SHA-224", nil
	case SNMPAuthProtocolSHA256:
		return "SHA-256", nil
	case SNMPAuthProtocolSHA384:
		return "SHA-384", nil
	case SNMPAuthProtocolSHA512:
		return "SHA-512", nil
	}
	if len(*en) == 0 {
		return "MD5", nil
	}
	return nil, errors.New("Invalid value of SNMPAuthProtocol: " + string(*en))
}

func (en *SNMPAuthProtocol) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "MD5":
		*en = SNMPAuthProtocolMD5
		return nil
	case "SHA":
		*en = SNMPAuthProtocolSHA
		return nil
	case "SHA-224":
		*en = SNMPAuthProtocolSHA224
		return nil
	case "SHA-256":
		*en = SNMPAuthProtocolSHA256
		return nil
	case "SHA-384":
		*en = SNMPAuthProtocolSHA384
		return nil
	case "SHA-512":
		*en = SNMPAuthProtocolSHA512
		return nil
	}
	if len(s) == 0 {
		*en = SNMPAuthProtocolMD5
		return nil
	}
	return errors.New("Unknown SNMPAuthProtocol: " + s)
}

func (en *SNMPAuthProtocol) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "MD5":
		*en = SNMPAuthProtocolMD5
		return nil
	case "SHA":
		*en = SNMPAuthProtocolSHA
		return nil
	case "SHA-224":
		*en = SNMPAuthProtocolSHA224
		return nil
	case "SHA-256":
		*en = SNMPAuthProtocolSHA256
		return nil
	case "SHA-384":
		*en = SNMPAuthProtocolSHA384
		return nil
	case "SHA-512":
		*en = SNMPAuthProtocolSHA512
		return nil
	}
	if len(s) == 0 {
		*en = SNMPAuthProtocolMD5
		return nil
	}
	return errors.New("Unknown SNMPAuthProtocol: " + s)
}

type SNMPPrivacyProtocol string

const SNMPPrivacyProtocolAES SNMPPrivacyProtocol = "AES"
const SNMPPrivacyProtocolAES192 SNMPPrivacyProtocol = "AES-192"
const SNMPPrivacyProtocolAES256 SNMPPrivacyProtocol = "AES-256"
const SNMPPrivacyProtocolDES SNMPPrivacyProtocol = "DES"

func (en SNMPPrivacyProtocol) GetPtr() *SNMPPrivacyProtocol { var v = en; return &v }

func (en SNMPPrivacyProtocol) String() string {
	switch en {
	case SNMPPrivacyProtocolAES:
		return "AES"
	case SNMPPrivacyProtocolAES192:
		return "AES-192"
	case SNMPPrivacyProtocolAES256:
		return "AES-256"
	case SNMPPrivacyProtocolDES:
		return "DES"
	}
	if len(en) == 0 {
		return "DES"
	}
	panic(errors.New("Invalid value of SNMPPrivacyProtocol: " + string(en)))
}

func (en *SNMPPrivacyProtocol) MarshalJSON() ([]byte, error) {
	switch *en {
	case SNMPPrivacyProtocolAES:
		return json.Marshal("AES")
	case SNMPPrivacyProtocolAES192:
		return json.Marshal("AES-192")
	case SNMPPrivacyProtocolAES256:
		return json.Marshal("AES-256")
	case SNMPPrivacyProtocolDES:
		return json.Marshal("DES")
	}
	if len(*en) == 0 {
		return json.Marshal("DES")
	}
	return nil, errors.New("Invalid value of SNMPPrivacyProtocol: " + string(*en))
}

func (en *SNMPPrivacyProtocol) GetBSON() (interface{}, error) {
	switch *en {
	case SNMPPrivacyProtocolAES:
		return "AES", nil
	case SNMPPrivacyProtocolAES192:
		return "AES-192", nil
	case SNMPPrivacyProtocolAES256:
		return "AES-256", nil
	case SNMPPrivacyProtocolDES:
		return "DES", nil
	}
	if len(*en) == 0 {
		return "DES", nil
	}
	return nil, errors.New("Invalid value of SNMPPrivacyProtocol: " + string(*en))
}

func (en *SNMPPrivacyProtocol) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "AES":
		*en = SNMPPrivacyProtocolAES
		return nil
	case "AES-192":
		*en = SNMPPrivacyProtocolAES192
		return nil
	case "AES-256":
		*en = SNMPPrivacyProtocolAES256
		return nil
	case "DES":
		*en = SNMPPrivacyProtocolDES
		return nil
	}
	if len(s) == 0 {
		*en = SNMPPrivacyProtocolDES
		return nil
	}
	return errors.New("Unknown SNMPPrivacyProtocol: " + s)
}

func (en *SNMPPrivacyProtocol) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "AES":
		*en = SNMPPrivacyProtocolAES
		return nil
	case "AES-192":
		*en = SNMPPrivacyProtocolAES192
		return nil
	case "AES-256":
		*en = SNMPPrivacyProtocolAES256
		return nil
	case "DES":
		*en = SNMPPrivacyProtocolDES
		return nil
	}
	if len(s) == 0 {
		*en = SNMPPrivacyProtocolDES
		return nil
	}
	return errors.New("Unknown SNMPPrivacyProtocol: " + s)
}

type SNMPSecurityLevelType string

const SNMPSecurityLevelTypeAuthNoPriv SNMPSecurityLevelType = "authNoPriv"
const SNMPSecurityLevelTypeAuthPriv SNMPSecurityLevelType = "authPriv"
const SNMPSecurityLevelTypeNoAuthNoPriv SNMPSecurityLevelType = "noAuthNoPriv"

func (en SNMPSecurityLevelType) GetPtr() *SNMPSecurityLevelType { var v = en; return &v }

func (en SNMPSecurityLevelType) String() string {
	switch en {
	case SNMPSecurityLevelTypeAuthNoPriv:
		return "authNoPriv"
	case SNMPSecurityLevelTypeAuthPriv:
		return "authPriv"
	case SNMPSecurityLevelTypeNoAuthNoPriv:
		return "noAuthNoPriv"
	}
	if len(en) == 0 {
		return "noAuthNoPriv"
	}
	panic(errors.New("Invalid value of SNMPSecurityLevelType: " + string(en)))
}

func (en *SNMPSecurityLevelType) MarshalJSON() ([]byte, error) {
	switch *en {
	case SNMPSecurityLevelTypeAuthNoPriv:
		return json.Marshal("authNoPriv")
	case SNMPSecurityLevelTypeAuthPriv:
		return json.Marshal("authPriv")
	case SNMPSecurityLevelTypeNoAuthNoPriv:
		return json.Marshal("noAuthNoPriv")
	}
	if len(*en) == 0 {
		return json.Marshal("noAuthNoPriv")
	}
	return nil, errors.New("Invalid value of SNMPSecurityLevelType: " + string(*en))
}

func (en *SNMPSecurityLevelType) GetBSON() (interface{}, error) {
	switch *en {
	case SNMPSecurityLevelTypeAuthNoPriv:
		return "authNoPriv", nil
	case SNMPSecurityLevelTypeAuthPriv:
		return "authPriv", nil
	case SNMPSecurityLevelTypeNoAuthNoPriv:
		return "noAuthNoPriv", nil
	}
	if len(*en) == 0 {
		return "noAuthNoPriv", nil
	}
	return nil, errors.New("Invalid value of SNMPSecurityLevelType: " + string(*en))
}

func (en *SNMPSecurityLevelType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "authNoPriv":
		*en = SNMPSecurityLevelTypeAuthNoPriv
		return nil
	case "authPriv":
		*en = SNMPSecurityLevelTypeAuthPriv
		return nil
	case "noAuthNoPriv":
		*en = SNMPSecurityLevelTypeNoAuthNoPriv
		return nil
	}
	if len(s) == 0 {
		*en = SNMPSecurityLevelTypeNoAuthNoPriv
		return nil
	}
	return errors.New("Unknown SNMPSecurityLevelType: " + s)
}

func (en *SNMPSecurityLevelType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "authNoPriv":
		*en = SNMPSecurityLevelTypeAuthNoPriv
		return nil
	case "authPriv":
		*en = SNMPSecurityLevelTypeAuthPriv
		return nil
	case "noAuthNoPriv":
		*en = SNMPSecurityLevelTypeNoAuthNoPriv
		return nil
	}
	if len(s) == 0 {
		*en = SNMPSecurityLevelTypeNoAuthNoPriv
		return nil
	}
	return errors.New("Unknown SNMPSecurityLevelType: " + s)
}

type SNMPVersion string

const SNMPVersionV1 SNMPVersion = "1"
const SNMPVersionV2 SNMPVersion = "2c"
const SNMPVersionV3 SNMPVersion = "3"

func (en SNMPVersion) GetPtr() *SNMPVersion { var v = en; return &v }

func (en SNMPVersion) String() string {
	switch en {
	case SNMPVersionV1:
		return "1"
	case SNMPVersionV2:
		return "2c"
	case SNMPVersionV3:
		return "3"
	}
	if len(en) == 0 {
		return "1"
	}
	panic(errors.New("Invalid value of SNMPVersion: " + string(en)))
}

func (en *SNMPVersion) MarshalJSON() ([]byte, error) {
	switch *en {
	case SNMPVersionV1:
		return json.Marshal("1")
	case SNMPVersionV2:
		return json.Marshal("2c")
	case SNMPVersionV3:
		return json.Marshal("3")
	}
	if len(*en) == 0 {
		return json.Marshal("1")
	}
	return nil, errors.New("Invalid value of SNMPVersion: " + string(*en))
}

func (en *SNMPVersion) GetBSON() (interface{}, error) {
	switch *en {
	case SNMPVersionV1:
		return "1", nil
	case SNMPVersionV2:
		return "2c", nil
	case SNMPVersionV3:
		return "3", nil
	}
	if len(*en) == 0 {
		return "1", nil
	}
	return nil, errors.New("Invalid value of SNMPVersion: " + string(*en))
}

func (en *SNMPVersion) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "1":
		*en = SNMPVersionV1
		return nil
	case "2c":
		*en = SNMPVersionV2
		return nil
	case "3":
		*en = SNMPVersionV3
		return nil
	}
	if len(s) == 0 {
		*en = SNMPVersionV1
		return nil
	}
	return errors.New("Unknown SNMPVersion: " + s)
}

func (en *SNMPVersion) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "1":
		*en = SNMPVersionV1
		return nil
	case "2c":
		*en = SNMPVersionV2
		return nil
	case "3":
		*en = SNMPVersionV3
		return nil
	}
	if len(s) == 0 {
		*en = SNMPVersionV1
		return nil
	}
	return errors.New("Unknown SNMPVersion: " + s)
}

type SecuritySuite string

const SecuritySuiteAES SecuritySuite = "aes"
const SecuritySuiteCCMP SecuritySuite = "ccmp"
const SecuritySuiteTKIP SecuritySuite = "tkip"

func (en SecuritySuite) GetPtr() *SecuritySuite { var v = en; return &v }

func (en SecuritySuite) String() string {
	switch en {
	case SecuritySuiteAES:
		return "aes"
	case SecuritySuiteCCMP:
		return "ccmp"
	case SecuritySuiteTKIP:
		return "tkip"
	}
	panic(errors.New("Invalid value of SecuritySuite: " + string(en)))
}

func (en *SecuritySuite) MarshalJSON() ([]byte, error) {
	switch *en {
	case SecuritySuiteAES:
		return json.Marshal("aes")
	case SecuritySuiteCCMP:
		return json.Marshal("ccmp")
	case SecuritySuiteTKIP:
		return json.Marshal("tkip")
	}
	return nil, errors.New("Invalid value of SecuritySuite: " + string(*en))
}

func (en *SecuritySuite) GetBSON() (interface{}, error) {
	switch *en {
	case SecuritySuiteAES:
		return "aes", nil
	case SecuritySuiteCCMP:
		return "ccmp", nil
	case SecuritySuiteTKIP:
		return "tkip", nil
	}
	return nil, errors.New("Invalid value of SecuritySuite: " + string(*en))
}

func (en *SecuritySuite) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "aes":
		*en = SecuritySuiteAES
		return nil
	case "ccmp":
		*en = SecuritySuiteCCMP
		return nil
	case "tkip":
		*en = SecuritySuiteTKIP
		return nil
	}
	return errors.New("Unknown SecuritySuite: " + s)
}

func (en *SecuritySuite) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "aes":
		*en = SecuritySuiteAES
		return nil
	case "ccmp":
		*en = SecuritySuiteCCMP
		return nil
	case "tkip":
		*en = SecuritySuiteTKIP
		return nil
	}
	return errors.New("Unknown SecuritySuite: " + s)
}

type SecurityType string

const SecurityTypeNone SecurityType = "open"
const SecurityTypeWPA23Enterprise SecurityType = "wpa23enterprise"
const SecurityTypeWPA23Personal SecurityType = "wpa23personal"
const SecurityTypeWPA2Enterprise SecurityType = "wpa2enterprise"
const SecurityTypeWPA2Personal SecurityType = "wpa2personal"
const SecurityTypeWPA3Enterprise SecurityType = "wpa3enterprise"
const SecurityTypeWPA3Personal SecurityType = "wpa3personal"
const SecurityTypeWPAEnterprise SecurityType = "wpaenterprise"
const SecurityTypeWPAPersonal SecurityType = "wpapersonal"

func (en SecurityType) GetPtr() *SecurityType { var v = en; return &v }

func (en SecurityType) String() string {
	switch en {
	case SecurityTypeNone:
		return "open"
	case SecurityTypeWPA23Enterprise:
		return "wpa23enterprise"
	case SecurityTypeWPA23Personal:
		return "wpa23personal"
	case SecurityTypeWPA2Enterprise:
		return "wpa2enterprise"
	case SecurityTypeWPA2Personal:
		return "wpa2personal"
	case SecurityTypeWPA3Enterprise:
		return "wpa3enterprise"
	case SecurityTypeWPA3Personal:
		return "wpa3personal"
	case SecurityTypeWPAEnterprise:
		return "wpaenterprise"
	case SecurityTypeWPAPersonal:
		return "wpapersonal"
	}
	if len(en) == 0 {
		return "open"
	}
	panic(errors.New("Invalid value of SecurityType: " + string(en)))
}

func (en *SecurityType) MarshalJSON() ([]byte, error) {
	switch *en {
	case SecurityTypeNone:
		return json.Marshal("open")
	case SecurityTypeWPA23Enterprise:
		return json.Marshal("wpa23enterprise")
	case SecurityTypeWPA23Personal:
		return json.Marshal("wpa23personal")
	case SecurityTypeWPA2Enterprise:
		return json.Marshal("wpa2enterprise")
	case SecurityTypeWPA2Personal:
		return json.Marshal("wpa2personal")
	case SecurityTypeWPA3Enterprise:
		return json.Marshal("wpa3enterprise")
	case SecurityTypeWPA3Personal:
		return json.Marshal("wpa3personal")
	case SecurityTypeWPAEnterprise:
		return json.Marshal("wpaenterprise")
	case SecurityTypeWPAPersonal:
		return json.Marshal("wpapersonal")
	}
	if len(*en) == 0 {
		return json.Marshal("open")
	}
	return nil, errors.New("Invalid value of SecurityType: " + string(*en))
}

func (en *SecurityType) GetBSON() (interface{}, error) {
	switch *en {
	case SecurityTypeNone:
		return "open", nil
	case SecurityTypeWPA23Enterprise:
		return "wpa23enterprise", nil
	case SecurityTypeWPA23Personal:
		return "wpa23personal", nil
	case SecurityTypeWPA2Enterprise:
		return "wpa2enterprise", nil
	case SecurityTypeWPA2Personal:
		return "wpa2personal", nil
	case SecurityTypeWPA3Enterprise:
		return "wpa3enterprise", nil
	case SecurityTypeWPA3Personal:
		return "wpa3personal", nil
	case SecurityTypeWPAEnterprise:
		return "wpaenterprise", nil
	case SecurityTypeWPAPersonal:
		return "wpapersonal", nil
	}
	if len(*en) == 0 {
		return "open", nil
	}
	return nil, errors.New("Invalid value of SecurityType: " + string(*en))
}

func (en *SecurityType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "open":
		*en = SecurityTypeNone
		return nil
	case "wpa23enterprise":
		*en = SecurityTypeWPA23Enterprise
		return nil
	case "wpa23personal":
		*en = SecurityTypeWPA23Personal
		return nil
	case "wpa2enterprise":
		*en = SecurityTypeWPA2Enterprise
		return nil
	case "wpa2personal":
		*en = SecurityTypeWPA2Personal
		return nil
	case "wpa3enterprise":
		*en = SecurityTypeWPA3Enterprise
		return nil
	case "wpa3personal":
		*en = SecurityTypeWPA3Personal
		return nil
	case "wpaenterprise":
		*en = SecurityTypeWPAEnterprise
		return nil
	case "wpapersonal":
		*en = SecurityTypeWPAPersonal
		return nil
	}
	if len(s) == 0 {
		*en = SecurityTypeNone
		return nil
	}
	return errors.New("Unknown SecurityType: " + s)
}

func (en *SecurityType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "open":
		*en = SecurityTypeNone
		return nil
	case "wpa23enterprise":
		*en = SecurityTypeWPA23Enterprise
		return nil
	case "wpa23personal":
		*en = SecurityTypeWPA23Personal
		return nil
	case "wpa2enterprise":
		*en = SecurityTypeWPA2Enterprise
		return nil
	case "wpa2personal":
		*en = SecurityTypeWPA2Personal
		return nil
	case "wpa3enterprise":
		*en = SecurityTypeWPA3Enterprise
		return nil
	case "wpa3personal":
		*en = SecurityTypeWPA3Personal
		return nil
	case "wpaenterprise":
		*en = SecurityTypeWPAEnterprise
		return nil
	case "wpapersonal":
		*en = SecurityTypeWPAPersonal
		return nil
	}
	if len(s) == 0 {
		*en = SecurityTypeNone
		return nil
	}
	return errors.New("Unknown SecurityType: " + s)
}

type ServiceState string

const ServiceStateConnected ServiceState = "CONNECTED"
const ServiceStateDisconnected ServiceState = "DISCONNECTED"
const ServiceStatePending ServiceState = "PENDING"

func (en ServiceState) GetPtr() *ServiceState { var v = en; return &v }

func (en ServiceState) String() string {
	switch en {
	case ServiceStateConnected:
		return "CONNECTED"
	case ServiceStateDisconnected:
		return "DISCONNECTED"
	case ServiceStatePending:
		return "PENDING"
	}
	panic(errors.New("Invalid value of ServiceState: " + string(en)))
}

func (en *ServiceState) MarshalJSON() ([]byte, error) {
	switch *en {
	case ServiceStateConnected:
		return json.Marshal("CONNECTED")
	case ServiceStateDisconnected:
		return json.Marshal("DISCONNECTED")
	case ServiceStatePending:
		return json.Marshal("PENDING")
	}
	return nil, errors.New("Invalid value of ServiceState: " + string(*en))
}

func (en *ServiceState) GetBSON() (interface{}, error) {
	switch *en {
	case ServiceStateConnected:
		return "CONNECTED", nil
	case ServiceStateDisconnected:
		return "DISCONNECTED", nil
	case ServiceStatePending:
		return "PENDING", nil
	}
	return nil, errors.New("Invalid value of ServiceState: " + string(*en))
}

func (en *ServiceState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		*en = ServiceStateConnected
		return nil
	case "DISCONNECTED":
		*en = ServiceStateDisconnected
		return nil
	case "PENDING":
		*en = ServiceStatePending
		return nil
	}
	return errors.New("Unknown ServiceState: " + s)
}

func (en *ServiceState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		*en = ServiceStateConnected
		return nil
	case "DISCONNECTED":
		*en = ServiceStateDisconnected
		return nil
	case "PENDING":
		*en = ServiceStatePending
		return nil
	}
	return errors.New("Unknown ServiceState: " + s)
}

type SpeedType string

const SpeedTypeGbit SpeedType = "gbit"
const SpeedTypeKbit SpeedType = "kbit"
const SpeedTypeMbit SpeedType = "mbit"
const SpeedTypeTbit SpeedType = "tbit"

func (en SpeedType) GetPtr() *SpeedType { var v = en; return &v }

func (en SpeedType) String() string {
	switch en {
	case SpeedTypeGbit:
		return "gbit"
	case SpeedTypeKbit:
		return "kbit"
	case SpeedTypeMbit:
		return "mbit"
	case SpeedTypeTbit:
		return "tbit"
	}
	if len(en) == 0 {
		return "kbit"
	}
	panic(errors.New("Invalid value of SpeedType: " + string(en)))
}

func (en *SpeedType) MarshalJSON() ([]byte, error) {
	switch *en {
	case SpeedTypeGbit:
		return json.Marshal("gbit")
	case SpeedTypeKbit:
		return json.Marshal("kbit")
	case SpeedTypeMbit:
		return json.Marshal("mbit")
	case SpeedTypeTbit:
		return json.Marshal("tbit")
	}
	if len(*en) == 0 {
		return json.Marshal("kbit")
	}
	return nil, errors.New("Invalid value of SpeedType: " + string(*en))
}

func (en *SpeedType) GetBSON() (interface{}, error) {
	switch *en {
	case SpeedTypeGbit:
		return "gbit", nil
	case SpeedTypeKbit:
		return "kbit", nil
	case SpeedTypeMbit:
		return "mbit", nil
	case SpeedTypeTbit:
		return "tbit", nil
	}
	if len(*en) == 0 {
		return "kbit", nil
	}
	return nil, errors.New("Invalid value of SpeedType: " + string(*en))
}

func (en *SpeedType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "gbit":
		*en = SpeedTypeGbit
		return nil
	case "kbit":
		*en = SpeedTypeKbit
		return nil
	case "mbit":
		*en = SpeedTypeMbit
		return nil
	case "tbit":
		*en = SpeedTypeTbit
		return nil
	}
	if len(s) == 0 {
		*en = SpeedTypeKbit
		return nil
	}
	return errors.New("Unknown SpeedType: " + s)
}

func (en *SpeedType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "gbit":
		*en = SpeedTypeGbit
		return nil
	case "kbit":
		*en = SpeedTypeKbit
		return nil
	case "mbit":
		*en = SpeedTypeMbit
		return nil
	case "tbit":
		*en = SpeedTypeTbit
		return nil
	}
	if len(s) == 0 {
		*en = SpeedTypeKbit
		return nil
	}
	return errors.New("Unknown SpeedType: " + s)
}

type StatEventRuleType string

const StatEventRuleTypeCPUload StatEventRuleType = "cpu_load"
const StatEventRuleTypeClientCon StatEventRuleType = "client_con"
const StatEventRuleTypeClientDis StatEventRuleType = "client_dis"
const StatEventRuleTypeClientFar StatEventRuleType = "client_far"
const StatEventRuleTypeConfigError StatEventRuleType = "config_error"
const StatEventRuleTypeConnected StatEventRuleType = "connected"
const StatEventRuleTypeCustomerActivity StatEventRuleType = "customer_activity"
const StatEventRuleTypeDisconnected StatEventRuleType = "disconnected"
const StatEventRuleTypeFreeRAM StatEventRuleType = "free_ram"
const StatEventRuleTypeIfaceError StatEventRuleType = "iface_error"

func (en StatEventRuleType) GetPtr() *StatEventRuleType { var v = en; return &v }

func (en StatEventRuleType) String() string {
	switch en {
	case StatEventRuleTypeCPUload:
		return "cpu_load"
	case StatEventRuleTypeClientCon:
		return "client_con"
	case StatEventRuleTypeClientDis:
		return "client_dis"
	case StatEventRuleTypeClientFar:
		return "client_far"
	case StatEventRuleTypeConfigError:
		return "config_error"
	case StatEventRuleTypeConnected:
		return "connected"
	case StatEventRuleTypeCustomerActivity:
		return "customer_activity"
	case StatEventRuleTypeDisconnected:
		return "disconnected"
	case StatEventRuleTypeFreeRAM:
		return "free_ram"
	case StatEventRuleTypeIfaceError:
		return "iface_error"
	}
	panic(errors.New("Invalid value of StatEventRuleType: " + string(en)))
}

func (en *StatEventRuleType) MarshalJSON() ([]byte, error) {
	switch *en {
	case StatEventRuleTypeCPUload:
		return json.Marshal("cpu_load")
	case StatEventRuleTypeClientCon:
		return json.Marshal("client_con")
	case StatEventRuleTypeClientDis:
		return json.Marshal("client_dis")
	case StatEventRuleTypeClientFar:
		return json.Marshal("client_far")
	case StatEventRuleTypeConfigError:
		return json.Marshal("config_error")
	case StatEventRuleTypeConnected:
		return json.Marshal("connected")
	case StatEventRuleTypeCustomerActivity:
		return json.Marshal("customer_activity")
	case StatEventRuleTypeDisconnected:
		return json.Marshal("disconnected")
	case StatEventRuleTypeFreeRAM:
		return json.Marshal("free_ram")
	case StatEventRuleTypeIfaceError:
		return json.Marshal("iface_error")
	}
	return nil, errors.New("Invalid value of StatEventRuleType: " + string(*en))
}

func (en *StatEventRuleType) GetBSON() (interface{}, error) {
	switch *en {
	case StatEventRuleTypeCPUload:
		return "cpu_load", nil
	case StatEventRuleTypeClientCon:
		return "client_con", nil
	case StatEventRuleTypeClientDis:
		return "client_dis", nil
	case StatEventRuleTypeClientFar:
		return "client_far", nil
	case StatEventRuleTypeConfigError:
		return "config_error", nil
	case StatEventRuleTypeConnected:
		return "connected", nil
	case StatEventRuleTypeCustomerActivity:
		return "customer_activity", nil
	case StatEventRuleTypeDisconnected:
		return "disconnected", nil
	case StatEventRuleTypeFreeRAM:
		return "free_ram", nil
	case StatEventRuleTypeIfaceError:
		return "iface_error", nil
	}
	return nil, errors.New("Invalid value of StatEventRuleType: " + string(*en))
}

func (en *StatEventRuleType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "cpu_load":
		*en = StatEventRuleTypeCPUload
		return nil
	case "client_con":
		*en = StatEventRuleTypeClientCon
		return nil
	case "client_dis":
		*en = StatEventRuleTypeClientDis
		return nil
	case "client_far":
		*en = StatEventRuleTypeClientFar
		return nil
	case "config_error":
		*en = StatEventRuleTypeConfigError
		return nil
	case "connected":
		*en = StatEventRuleTypeConnected
		return nil
	case "customer_activity":
		*en = StatEventRuleTypeCustomerActivity
		return nil
	case "disconnected":
		*en = StatEventRuleTypeDisconnected
		return nil
	case "free_ram":
		*en = StatEventRuleTypeFreeRAM
		return nil
	case "iface_error":
		*en = StatEventRuleTypeIfaceError
		return nil
	}
	return errors.New("Unknown StatEventRuleType: " + s)
}

func (en *StatEventRuleType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "cpu_load":
		*en = StatEventRuleTypeCPUload
		return nil
	case "client_con":
		*en = StatEventRuleTypeClientCon
		return nil
	case "client_dis":
		*en = StatEventRuleTypeClientDis
		return nil
	case "client_far":
		*en = StatEventRuleTypeClientFar
		return nil
	case "config_error":
		*en = StatEventRuleTypeConfigError
		return nil
	case "connected":
		*en = StatEventRuleTypeConnected
		return nil
	case "customer_activity":
		*en = StatEventRuleTypeCustomerActivity
		return nil
	case "disconnected":
		*en = StatEventRuleTypeDisconnected
		return nil
	case "free_ram":
		*en = StatEventRuleTypeFreeRAM
		return nil
	case "iface_error":
		*en = StatEventRuleTypeIfaceError
		return nil
	}
	return errors.New("Unknown StatEventRuleType: " + s)
}

type SystemEventLevel string

const SystemEventLevelDEBUG SystemEventLevel = "DEBUG"
const SystemEventLevelERROR SystemEventLevel = "ERROR"
const SystemEventLevelINFO SystemEventLevel = "INFO"
const SystemEventLevelWARNING SystemEventLevel = "WARNING"

func (en SystemEventLevel) GetPtr() *SystemEventLevel { var v = en; return &v }

func (en SystemEventLevel) String() string {
	switch en {
	case SystemEventLevelDEBUG:
		return "DEBUG"
	case SystemEventLevelERROR:
		return "ERROR"
	case SystemEventLevelINFO:
		return "INFO"
	case SystemEventLevelWARNING:
		return "WARNING"
	}
	if len(en) == 0 {
		return "DEBUG"
	}
	panic(errors.New("Invalid value of SystemEventLevel: " + string(en)))
}

func (en *SystemEventLevel) MarshalJSON() ([]byte, error) {
	switch *en {
	case SystemEventLevelDEBUG:
		return json.Marshal("DEBUG")
	case SystemEventLevelERROR:
		return json.Marshal("ERROR")
	case SystemEventLevelINFO:
		return json.Marshal("INFO")
	case SystemEventLevelWARNING:
		return json.Marshal("WARNING")
	}
	if len(*en) == 0 {
		return json.Marshal("DEBUG")
	}
	return nil, errors.New("Invalid value of SystemEventLevel: " + string(*en))
}

func (en *SystemEventLevel) GetBSON() (interface{}, error) {
	switch *en {
	case SystemEventLevelDEBUG:
		return "DEBUG", nil
	case SystemEventLevelERROR:
		return "ERROR", nil
	case SystemEventLevelINFO:
		return "INFO", nil
	case SystemEventLevelWARNING:
		return "WARNING", nil
	}
	if len(*en) == 0 {
		return "DEBUG", nil
	}
	return nil, errors.New("Invalid value of SystemEventLevel: " + string(*en))
}

func (en *SystemEventLevel) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "DEBUG":
		*en = SystemEventLevelDEBUG
		return nil
	case "ERROR":
		*en = SystemEventLevelERROR
		return nil
	case "INFO":
		*en = SystemEventLevelINFO
		return nil
	case "WARNING":
		*en = SystemEventLevelWARNING
		return nil
	}
	if len(s) == 0 {
		*en = SystemEventLevelDEBUG
		return nil
	}
	return errors.New("Unknown SystemEventLevel: " + s)
}

func (en *SystemEventLevel) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "DEBUG":
		*en = SystemEventLevelDEBUG
		return nil
	case "ERROR":
		*en = SystemEventLevelERROR
		return nil
	case "INFO":
		*en = SystemEventLevelINFO
		return nil
	case "WARNING":
		*en = SystemEventLevelWARNING
		return nil
	}
	if len(s) == 0 {
		*en = SystemEventLevelDEBUG
		return nil
	}
	return errors.New("Unknown SystemEventLevel: " + s)
}

type SystemEventType string

const SystemEventTypeAny SystemEventType = "+"
const SystemEventTypeCPEConnected SystemEventType = "CPE_CONNECTED"
const SystemEventTypeCPEDisconnected SystemEventType = "CPE_DISCONNECTED"
const SystemEventTypeCPEInterfaceState SystemEventType = "CPE_INTERFACE_STATE"
const SystemEventTypeClientAuthorization SystemEventType = "CLIENT_AUTHORIZATION"
const SystemEventTypeClientConnected SystemEventType = "CLIENT_CONNECTED"
const SystemEventTypeClientDisconnected SystemEventType = "CLIENT_DISCONNECTED"
const SystemEventTypeCpeFirmwareAvailable SystemEventType = "CPE_FIRMWARE_AVAILABLE"
const SystemEventTypeDHCPAck SystemEventType = "DHCP_ACK"
const SystemEventTypeDaemonSettingsChanged SystemEventType = "DAEMON_SETTINGS_CHANGE"
const SystemEventTypeFirmwareUploaded SystemEventType = "FIRMWARE_UPLOADED"
const SystemEventTypeLocationCacheReload SystemEventType = "LOCATION_CACHE_RELOAD"
const SystemEventTypeLoggedError SystemEventType = "LOGGED_ERROR"
const SystemEventTypeMonitorRuleViolation SystemEventType = "MONITOR_RULE_VIOLATION"
const SystemEventTypeRRMGroupApplyAlgo SystemEventType = "RRM_GROUP_APPLY_ALGO"
const SystemEventTypeRRMStatus SystemEventType = "RRM_STATUS_DATA"
const SystemEventTypeRadarExportUpdate SystemEventType = "RADAR_EXPORT_UPDATE"
const SystemEventTypeRadiusAccountingSend SystemEventType = "RADIUS_ACCOUNTING_SEND"
const SystemEventTypeServiceConnected SystemEventType = "SERVICE_CONNECTED"
const SystemEventTypeServiceDisconnected SystemEventType = "SERVICE_DISCONNECTED"
const SystemEventTypeServiceFatalError SystemEventType = "SERVICE_FATAL_ERROR"
const SystemEventTypeSystemTimeChanged SystemEventType = "SYSTEM_TIME_CHANGE"
const SystemEventTypeUserAuthSuccess SystemEventType = "USER_AUTHORIZATION_SUCCESS"
const SystemEventTypeWLANCentrAccChanged SystemEventType = "WLAN_CENTR_ACC_CHANGE"

func (en SystemEventType) GetPtr() *SystemEventType { var v = en; return &v }

func (en SystemEventType) String() string {
	switch en {
	case SystemEventTypeAny:
		return "+"
	case SystemEventTypeCPEConnected:
		return "CPE_CONNECTED"
	case SystemEventTypeCPEDisconnected:
		return "CPE_DISCONNECTED"
	case SystemEventTypeCPEInterfaceState:
		return "CPE_INTERFACE_STATE"
	case SystemEventTypeClientAuthorization:
		return "CLIENT_AUTHORIZATION"
	case SystemEventTypeClientConnected:
		return "CLIENT_CONNECTED"
	case SystemEventTypeClientDisconnected:
		return "CLIENT_DISCONNECTED"
	case SystemEventTypeCpeFirmwareAvailable:
		return "CPE_FIRMWARE_AVAILABLE"
	case SystemEventTypeDHCPAck:
		return "DHCP_ACK"
	case SystemEventTypeDaemonSettingsChanged:
		return "DAEMON_SETTINGS_CHANGE"
	case SystemEventTypeFirmwareUploaded:
		return "FIRMWARE_UPLOADED"
	case SystemEventTypeLocationCacheReload:
		return "LOCATION_CACHE_RELOAD"
	case SystemEventTypeLoggedError:
		return "LOGGED_ERROR"
	case SystemEventTypeMonitorRuleViolation:
		return "MONITOR_RULE_VIOLATION"
	case SystemEventTypeRRMGroupApplyAlgo:
		return "RRM_GROUP_APPLY_ALGO"
	case SystemEventTypeRRMStatus:
		return "RRM_STATUS_DATA"
	case SystemEventTypeRadarExportUpdate:
		return "RADAR_EXPORT_UPDATE"
	case SystemEventTypeRadiusAccountingSend:
		return "RADIUS_ACCOUNTING_SEND"
	case SystemEventTypeServiceConnected:
		return "SERVICE_CONNECTED"
	case SystemEventTypeServiceDisconnected:
		return "SERVICE_DISCONNECTED"
	case SystemEventTypeServiceFatalError:
		return "SERVICE_FATAL_ERROR"
	case SystemEventTypeSystemTimeChanged:
		return "SYSTEM_TIME_CHANGE"
	case SystemEventTypeUserAuthSuccess:
		return "USER_AUTHORIZATION_SUCCESS"
	case SystemEventTypeWLANCentrAccChanged:
		return "WLAN_CENTR_ACC_CHANGE"
	}
	panic(errors.New("Invalid value of SystemEventType: " + string(en)))
}

func (en *SystemEventType) MarshalJSON() ([]byte, error) {
	switch *en {
	case SystemEventTypeAny:
		return json.Marshal("+")
	case SystemEventTypeCPEConnected:
		return json.Marshal("CPE_CONNECTED")
	case SystemEventTypeCPEDisconnected:
		return json.Marshal("CPE_DISCONNECTED")
	case SystemEventTypeCPEInterfaceState:
		return json.Marshal("CPE_INTERFACE_STATE")
	case SystemEventTypeClientAuthorization:
		return json.Marshal("CLIENT_AUTHORIZATION")
	case SystemEventTypeClientConnected:
		return json.Marshal("CLIENT_CONNECTED")
	case SystemEventTypeClientDisconnected:
		return json.Marshal("CLIENT_DISCONNECTED")
	case SystemEventTypeCpeFirmwareAvailable:
		return json.Marshal("CPE_FIRMWARE_AVAILABLE")
	case SystemEventTypeDHCPAck:
		return json.Marshal("DHCP_ACK")
	case SystemEventTypeDaemonSettingsChanged:
		return json.Marshal("DAEMON_SETTINGS_CHANGE")
	case SystemEventTypeFirmwareUploaded:
		return json.Marshal("FIRMWARE_UPLOADED")
	case SystemEventTypeLocationCacheReload:
		return json.Marshal("LOCATION_CACHE_RELOAD")
	case SystemEventTypeLoggedError:
		return json.Marshal("LOGGED_ERROR")
	case SystemEventTypeMonitorRuleViolation:
		return json.Marshal("MONITOR_RULE_VIOLATION")
	case SystemEventTypeRRMGroupApplyAlgo:
		return json.Marshal("RRM_GROUP_APPLY_ALGO")
	case SystemEventTypeRRMStatus:
		return json.Marshal("RRM_STATUS_DATA")
	case SystemEventTypeRadarExportUpdate:
		return json.Marshal("RADAR_EXPORT_UPDATE")
	case SystemEventTypeRadiusAccountingSend:
		return json.Marshal("RADIUS_ACCOUNTING_SEND")
	case SystemEventTypeServiceConnected:
		return json.Marshal("SERVICE_CONNECTED")
	case SystemEventTypeServiceDisconnected:
		return json.Marshal("SERVICE_DISCONNECTED")
	case SystemEventTypeServiceFatalError:
		return json.Marshal("SERVICE_FATAL_ERROR")
	case SystemEventTypeSystemTimeChanged:
		return json.Marshal("SYSTEM_TIME_CHANGE")
	case SystemEventTypeUserAuthSuccess:
		return json.Marshal("USER_AUTHORIZATION_SUCCESS")
	case SystemEventTypeWLANCentrAccChanged:
		return json.Marshal("WLAN_CENTR_ACC_CHANGE")
	}
	return nil, errors.New("Invalid value of SystemEventType: " + string(*en))
}

func (en *SystemEventType) GetBSON() (interface{}, error) {
	switch *en {
	case SystemEventTypeAny:
		return "+", nil
	case SystemEventTypeCPEConnected:
		return "CPE_CONNECTED", nil
	case SystemEventTypeCPEDisconnected:
		return "CPE_DISCONNECTED", nil
	case SystemEventTypeCPEInterfaceState:
		return "CPE_INTERFACE_STATE", nil
	case SystemEventTypeClientAuthorization:
		return "CLIENT_AUTHORIZATION", nil
	case SystemEventTypeClientConnected:
		return "CLIENT_CONNECTED", nil
	case SystemEventTypeClientDisconnected:
		return "CLIENT_DISCONNECTED", nil
	case SystemEventTypeCpeFirmwareAvailable:
		return "CPE_FIRMWARE_AVAILABLE", nil
	case SystemEventTypeDHCPAck:
		return "DHCP_ACK", nil
	case SystemEventTypeDaemonSettingsChanged:
		return "DAEMON_SETTINGS_CHANGE", nil
	case SystemEventTypeFirmwareUploaded:
		return "FIRMWARE_UPLOADED", nil
	case SystemEventTypeLocationCacheReload:
		return "LOCATION_CACHE_RELOAD", nil
	case SystemEventTypeLoggedError:
		return "LOGGED_ERROR", nil
	case SystemEventTypeMonitorRuleViolation:
		return "MONITOR_RULE_VIOLATION", nil
	case SystemEventTypeRRMGroupApplyAlgo:
		return "RRM_GROUP_APPLY_ALGO", nil
	case SystemEventTypeRRMStatus:
		return "RRM_STATUS_DATA", nil
	case SystemEventTypeRadarExportUpdate:
		return "RADAR_EXPORT_UPDATE", nil
	case SystemEventTypeRadiusAccountingSend:
		return "RADIUS_ACCOUNTING_SEND", nil
	case SystemEventTypeServiceConnected:
		return "SERVICE_CONNECTED", nil
	case SystemEventTypeServiceDisconnected:
		return "SERVICE_DISCONNECTED", nil
	case SystemEventTypeServiceFatalError:
		return "SERVICE_FATAL_ERROR", nil
	case SystemEventTypeSystemTimeChanged:
		return "SYSTEM_TIME_CHANGE", nil
	case SystemEventTypeUserAuthSuccess:
		return "USER_AUTHORIZATION_SUCCESS", nil
	case SystemEventTypeWLANCentrAccChanged:
		return "WLAN_CENTR_ACC_CHANGE", nil
	}
	return nil, errors.New("Invalid value of SystemEventType: " + string(*en))
}

func (en *SystemEventType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "+":
		*en = SystemEventTypeAny
		return nil
	case "CPE_CONNECTED":
		*en = SystemEventTypeCPEConnected
		return nil
	case "CPE_DISCONNECTED":
		*en = SystemEventTypeCPEDisconnected
		return nil
	case "CPE_INTERFACE_STATE":
		*en = SystemEventTypeCPEInterfaceState
		return nil
	case "CLIENT_AUTHORIZATION":
		*en = SystemEventTypeClientAuthorization
		return nil
	case "CLIENT_CONNECTED":
		*en = SystemEventTypeClientConnected
		return nil
	case "CLIENT_DISCONNECTED":
		*en = SystemEventTypeClientDisconnected
		return nil
	case "CPE_FIRMWARE_AVAILABLE":
		*en = SystemEventTypeCpeFirmwareAvailable
		return nil
	case "DHCP_ACK":
		*en = SystemEventTypeDHCPAck
		return nil
	case "DAEMON_SETTINGS_CHANGE":
		*en = SystemEventTypeDaemonSettingsChanged
		return nil
	case "FIRMWARE_UPLOADED":
		*en = SystemEventTypeFirmwareUploaded
		return nil
	case "LOCATION_CACHE_RELOAD":
		*en = SystemEventTypeLocationCacheReload
		return nil
	case "LOGGED_ERROR":
		*en = SystemEventTypeLoggedError
		return nil
	case "MONITOR_RULE_VIOLATION":
		*en = SystemEventTypeMonitorRuleViolation
		return nil
	case "RRM_GROUP_APPLY_ALGO":
		*en = SystemEventTypeRRMGroupApplyAlgo
		return nil
	case "RRM_STATUS_DATA":
		*en = SystemEventTypeRRMStatus
		return nil
	case "RADAR_EXPORT_UPDATE":
		*en = SystemEventTypeRadarExportUpdate
		return nil
	case "RADIUS_ACCOUNTING_SEND":
		*en = SystemEventTypeRadiusAccountingSend
		return nil
	case "SERVICE_CONNECTED":
		*en = SystemEventTypeServiceConnected
		return nil
	case "SERVICE_DISCONNECTED":
		*en = SystemEventTypeServiceDisconnected
		return nil
	case "SERVICE_FATAL_ERROR":
		*en = SystemEventTypeServiceFatalError
		return nil
	case "SYSTEM_TIME_CHANGE":
		*en = SystemEventTypeSystemTimeChanged
		return nil
	case "USER_AUTHORIZATION_SUCCESS":
		*en = SystemEventTypeUserAuthSuccess
		return nil
	case "WLAN_CENTR_ACC_CHANGE":
		*en = SystemEventTypeWLANCentrAccChanged
		return nil
	}
	return errors.New("Unknown SystemEventType: " + s)
}

func (en *SystemEventType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "+":
		*en = SystemEventTypeAny
		return nil
	case "CPE_CONNECTED":
		*en = SystemEventTypeCPEConnected
		return nil
	case "CPE_DISCONNECTED":
		*en = SystemEventTypeCPEDisconnected
		return nil
	case "CPE_INTERFACE_STATE":
		*en = SystemEventTypeCPEInterfaceState
		return nil
	case "CLIENT_AUTHORIZATION":
		*en = SystemEventTypeClientAuthorization
		return nil
	case "CLIENT_CONNECTED":
		*en = SystemEventTypeClientConnected
		return nil
	case "CLIENT_DISCONNECTED":
		*en = SystemEventTypeClientDisconnected
		return nil
	case "CPE_FIRMWARE_AVAILABLE":
		*en = SystemEventTypeCpeFirmwareAvailable
		return nil
	case "DHCP_ACK":
		*en = SystemEventTypeDHCPAck
		return nil
	case "DAEMON_SETTINGS_CHANGE":
		*en = SystemEventTypeDaemonSettingsChanged
		return nil
	case "FIRMWARE_UPLOADED":
		*en = SystemEventTypeFirmwareUploaded
		return nil
	case "LOCATION_CACHE_RELOAD":
		*en = SystemEventTypeLocationCacheReload
		return nil
	case "LOGGED_ERROR":
		*en = SystemEventTypeLoggedError
		return nil
	case "MONITOR_RULE_VIOLATION":
		*en = SystemEventTypeMonitorRuleViolation
		return nil
	case "RRM_GROUP_APPLY_ALGO":
		*en = SystemEventTypeRRMGroupApplyAlgo
		return nil
	case "RRM_STATUS_DATA":
		*en = SystemEventTypeRRMStatus
		return nil
	case "RADAR_EXPORT_UPDATE":
		*en = SystemEventTypeRadarExportUpdate
		return nil
	case "RADIUS_ACCOUNTING_SEND":
		*en = SystemEventTypeRadiusAccountingSend
		return nil
	case "SERVICE_CONNECTED":
		*en = SystemEventTypeServiceConnected
		return nil
	case "SERVICE_DISCONNECTED":
		*en = SystemEventTypeServiceDisconnected
		return nil
	case "SERVICE_FATAL_ERROR":
		*en = SystemEventTypeServiceFatalError
		return nil
	case "SYSTEM_TIME_CHANGE":
		*en = SystemEventTypeSystemTimeChanged
		return nil
	case "USER_AUTHORIZATION_SUCCESS":
		*en = SystemEventTypeUserAuthSuccess
		return nil
	case "WLAN_CENTR_ACC_CHANGE":
		*en = SystemEventTypeWLANCentrAccChanged
		return nil
	}
	return errors.New("Unknown SystemEventType: " + s)
}

type TunManagerRPC string

const TunManagerRPCCreateL2TunnelSession TunManagerRPC = "CreateL2TunnelSession"
const TunManagerRPCDeleteL2TunnelSession TunManagerRPC = "DeleteL2TunnelSession"

func (en TunManagerRPC) GetPtr() *TunManagerRPC { var v = en; return &v }

func (en TunManagerRPC) String() string {
	switch en {
	case TunManagerRPCCreateL2TunnelSession:
		return "CreateL2TunnelSession"
	case TunManagerRPCDeleteL2TunnelSession:
		return "DeleteL2TunnelSession"
	}
	panic(errors.New("Invalid value of TunManagerRPC: " + string(en)))
}

func (en *TunManagerRPC) MarshalJSON() ([]byte, error) {
	switch *en {
	case TunManagerRPCCreateL2TunnelSession:
		return json.Marshal("CreateL2TunnelSession")
	case TunManagerRPCDeleteL2TunnelSession:
		return json.Marshal("DeleteL2TunnelSession")
	}
	return nil, errors.New("Invalid value of TunManagerRPC: " + string(*en))
}

func (en *TunManagerRPC) GetBSON() (interface{}, error) {
	switch *en {
	case TunManagerRPCCreateL2TunnelSession:
		return "CreateL2TunnelSession", nil
	case TunManagerRPCDeleteL2TunnelSession:
		return "DeleteL2TunnelSession", nil
	}
	return nil, errors.New("Invalid value of TunManagerRPC: " + string(*en))
}

func (en *TunManagerRPC) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "CreateL2TunnelSession":
		*en = TunManagerRPCCreateL2TunnelSession
		return nil
	case "DeleteL2TunnelSession":
		*en = TunManagerRPCDeleteL2TunnelSession
		return nil
	}
	return errors.New("Unknown TunManagerRPC: " + s)
}

func (en *TunManagerRPC) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "CreateL2TunnelSession":
		*en = TunManagerRPCCreateL2TunnelSession
		return nil
	case "DeleteL2TunnelSession":
		*en = TunManagerRPCDeleteL2TunnelSession
		return nil
	}
	return errors.New("Unknown TunManagerRPC: " + s)
}

type TunnelType string

const TunnelTypeEoGRE TunnelType = "gretap"
const TunnelTypeGRE TunnelType = "gre"
const TunnelTypeL2TP TunnelType = "l2tpv3"

func (en TunnelType) GetPtr() *TunnelType { var v = en; return &v }

func (en TunnelType) String() string {
	switch en {
	case TunnelTypeEoGRE:
		return "gretap"
	case TunnelTypeGRE:
		return "gre"
	case TunnelTypeL2TP:
		return "l2tpv3"
	}
	if len(en) == 0 {
		return "l2tpv3"
	}
	panic(errors.New("Invalid value of TunnelType: " + string(en)))
}

func (en *TunnelType) MarshalJSON() ([]byte, error) {
	switch *en {
	case TunnelTypeEoGRE:
		return json.Marshal("gretap")
	case TunnelTypeGRE:
		return json.Marshal("gre")
	case TunnelTypeL2TP:
		return json.Marshal("l2tpv3")
	}
	if len(*en) == 0 {
		return json.Marshal("l2tpv3")
	}
	return nil, errors.New("Invalid value of TunnelType: " + string(*en))
}

func (en *TunnelType) GetBSON() (interface{}, error) {
	switch *en {
	case TunnelTypeEoGRE:
		return "gretap", nil
	case TunnelTypeGRE:
		return "gre", nil
	case TunnelTypeL2TP:
		return "l2tpv3", nil
	}
	if len(*en) == 0 {
		return "l2tpv3", nil
	}
	return nil, errors.New("Invalid value of TunnelType: " + string(*en))
}

func (en *TunnelType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "gretap":
		*en = TunnelTypeEoGRE
		return nil
	case "gre":
		*en = TunnelTypeGRE
		return nil
	case "l2tpv3":
		*en = TunnelTypeL2TP
		return nil
	}
	if len(s) == 0 {
		*en = TunnelTypeL2TP
		return nil
	}
	return errors.New("Unknown TunnelType: " + s)
}

func (en *TunnelType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "gretap":
		*en = TunnelTypeEoGRE
		return nil
	case "gre":
		*en = TunnelTypeGRE
		return nil
	case "l2tpv3":
		*en = TunnelTypeL2TP
		return nil
	}
	if len(s) == 0 {
		*en = TunnelTypeL2TP
		return nil
	}
	return errors.New("Unknown TunnelType: " + s)
}

type WMMAccessCategory string

const WMMAccessCategoryBackground WMMAccessCategory = "BK"
const WMMAccessCategoryBestEffort WMMAccessCategory = "BE"
const WMMAccessCategoryVideo WMMAccessCategory = "VI"
const WMMAccessCategoryVoice WMMAccessCategory = "VO"

func (en WMMAccessCategory) GetPtr() *WMMAccessCategory { var v = en; return &v }

func (en WMMAccessCategory) String() string {
	switch en {
	case WMMAccessCategoryBackground:
		return "BK"
	case WMMAccessCategoryBestEffort:
		return "BE"
	case WMMAccessCategoryVideo:
		return "VI"
	case WMMAccessCategoryVoice:
		return "VO"
	}
	if len(en) == 0 {
		return "BK"
	}
	panic(errors.New("Invalid value of WMMAccessCategory: " + string(en)))
}

func (en *WMMAccessCategory) MarshalJSON() ([]byte, error) {
	switch *en {
	case WMMAccessCategoryBackground:
		return json.Marshal("BK")
	case WMMAccessCategoryBestEffort:
		return json.Marshal("BE")
	case WMMAccessCategoryVideo:
		return json.Marshal("VI")
	case WMMAccessCategoryVoice:
		return json.Marshal("VO")
	}
	if len(*en) == 0 {
		return json.Marshal("BK")
	}
	return nil, errors.New("Invalid value of WMMAccessCategory: " + string(*en))
}

func (en *WMMAccessCategory) GetBSON() (interface{}, error) {
	switch *en {
	case WMMAccessCategoryBackground:
		return "BK", nil
	case WMMAccessCategoryBestEffort:
		return "BE", nil
	case WMMAccessCategoryVideo:
		return "VI", nil
	case WMMAccessCategoryVoice:
		return "VO", nil
	}
	if len(*en) == 0 {
		return "BK", nil
	}
	return nil, errors.New("Invalid value of WMMAccessCategory: " + string(*en))
}

func (en *WMMAccessCategory) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "BK":
		*en = WMMAccessCategoryBackground
		return nil
	case "BE":
		*en = WMMAccessCategoryBestEffort
		return nil
	case "VI":
		*en = WMMAccessCategoryVideo
		return nil
	case "VO":
		*en = WMMAccessCategoryVoice
		return nil
	}
	if len(s) == 0 {
		*en = WMMAccessCategoryBackground
		return nil
	}
	return errors.New("Unknown WMMAccessCategory: " + s)
}

func (en *WMMAccessCategory) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "BK":
		*en = WMMAccessCategoryBackground
		return nil
	case "BE":
		*en = WMMAccessCategoryBestEffort
		return nil
	case "VI":
		*en = WMMAccessCategoryVideo
		return nil
	case "VO":
		*en = WMMAccessCategoryVoice
		return nil
	}
	if len(s) == 0 {
		*en = WMMAccessCategoryBackground
		return nil
	}
	return errors.New("Unknown WMMAccessCategory: " + s)
}

type WirelessClientState string

const WirelessClientStateConnected WirelessClientState = "CONNECTED"
const WirelessClientStateDisconnected WirelessClientState = "DISCONNECTED"

func (en WirelessClientState) GetPtr() *WirelessClientState { var v = en; return &v }

func (en WirelessClientState) String() string {
	switch en {
	case WirelessClientStateConnected:
		return "CONNECTED"
	case WirelessClientStateDisconnected:
		return "DISCONNECTED"
	}
	panic(errors.New("Invalid value of WirelessClientState: " + string(en)))
}

func (en *WirelessClientState) MarshalJSON() ([]byte, error) {
	switch *en {
	case WirelessClientStateConnected:
		return json.Marshal("CONNECTED")
	case WirelessClientStateDisconnected:
		return json.Marshal("DISCONNECTED")
	}
	return nil, errors.New("Invalid value of WirelessClientState: " + string(*en))
}

func (en *WirelessClientState) GetBSON() (interface{}, error) {
	switch *en {
	case WirelessClientStateConnected:
		return "CONNECTED", nil
	case WirelessClientStateDisconnected:
		return "DISCONNECTED", nil
	}
	return nil, errors.New("Invalid value of WirelessClientState: " + string(*en))
}

func (en *WirelessClientState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		*en = WirelessClientStateConnected
		return nil
	case "DISCONNECTED":
		*en = WirelessClientStateDisconnected
		return nil
	}
	return errors.New("Unknown WirelessClientState: " + s)
}

func (en *WirelessClientState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "CONNECTED":
		*en = WirelessClientStateConnected
		return nil
	case "DISCONNECTED":
		*en = WirelessClientStateDisconnected
		return nil
	}
	return errors.New("Unknown WirelessClientState: " + s)
}

type WirelessClientType string

const WirelessClientTypeCamera WirelessClientType = "camera"
const WirelessClientTypeOther WirelessClientType = "other"
const WirelessClientTypeWired WirelessClientType = "wired"

func (en WirelessClientType) GetPtr() *WirelessClientType { var v = en; return &v }

func (en WirelessClientType) String() string {
	switch en {
	case WirelessClientTypeCamera:
		return "camera"
	case WirelessClientTypeOther:
		return "other"
	case WirelessClientTypeWired:
		return "wired"
	}
	if len(en) == 0 {
		return "other"
	}
	panic(errors.New("Invalid value of WirelessClientType: " + string(en)))
}

func (en *WirelessClientType) MarshalJSON() ([]byte, error) {
	switch *en {
	case WirelessClientTypeCamera:
		return json.Marshal("camera")
	case WirelessClientTypeOther:
		return json.Marshal("other")
	case WirelessClientTypeWired:
		return json.Marshal("wired")
	}
	if len(*en) == 0 {
		return json.Marshal("other")
	}
	return nil, errors.New("Invalid value of WirelessClientType: " + string(*en))
}

func (en *WirelessClientType) GetBSON() (interface{}, error) {
	switch *en {
	case WirelessClientTypeCamera:
		return "camera", nil
	case WirelessClientTypeOther:
		return "other", nil
	case WirelessClientTypeWired:
		return "wired", nil
	}
	if len(*en) == 0 {
		return "other", nil
	}
	return nil, errors.New("Invalid value of WirelessClientType: " + string(*en))
}

func (en *WirelessClientType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "camera":
		*en = WirelessClientTypeCamera
		return nil
	case "other":
		*en = WirelessClientTypeOther
		return nil
	case "wired":
		*en = WirelessClientTypeWired
		return nil
	}
	if len(s) == 0 {
		*en = WirelessClientTypeOther
		return nil
	}
	return errors.New("Unknown WirelessClientType: " + s)
}

func (en *WirelessClientType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "camera":
		*en = WirelessClientTypeCamera
		return nil
	case "other":
		*en = WirelessClientTypeOther
		return nil
	case "wired":
		*en = WirelessClientTypeWired
		return nil
	}
	if len(s) == 0 {
		*en = WirelessClientTypeOther
		return nil
	}
	return errors.New("Unknown WirelessClientType: " + s)
}
