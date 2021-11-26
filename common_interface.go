package libwimark

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type ModuleStatus struct {
	Service Module             `json:"service"`
	Id      string             `json:"id"`
	Version string             `json:"version"`
	Commit  string             `json:"commit"`
	Build   int                `json:"build"`
	State   ServiceState       `json:"state"`
	Meta    interface{}        `json:"meta,omitempty"`
	Statics map[string]Version `json:"statics,omitempty"`
}

type CpeStatusMeta struct {
	Model   string             `json:"model"`
	Statics map[string]Version `json:"statics"`
}

type TunManagerBroadcastMeta struct {
	Hostname       string                 `json:"hostname"`
	HostUUID       string                 `json:"host_uuid"`
	HostInterfaces []LinkDescriptor       `json:"active_out_interfaces"`
	HostTunnels    []CPETunnelDescription `json:"active_cpe_tunnels"`
}

func (ms *ModuleStatus) UnmarshalJSON(b []byte) error {
	var doc map[string]json.RawMessage
	if err := json.Unmarshal(b, &doc); err != nil {
		return err
	}
	type MS ModuleStatus
	var s MS
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*ms = ModuleStatus(s)
	var meta, has_meta = doc["meta"]
	has_meta = has_meta && !bytes.Equal(meta, []byte("null"))

	switch ms.Service {
	case ModuleCPE:
		var m CpeStatusMeta
		if has_meta {
			if err := json.Unmarshal(meta, &m); err != nil {
				return err
			}
		}
		ms.Meta = m
	case ModuleTunManager:
		var m TunManagerBroadcastMeta
		if has_meta {
			if err := json.Unmarshal(meta, &m); err != nil {
				return err
			}
		}
		ms.Meta = m
	}
	return nil
}

func (ms ModuleStatus) Connected() ModuleStatus {
	var v = ms
	v.State = ServiceStateConnected

	return v
}

func (ms ModuleStatus) Disconnected() ModuleStatus {
	var v = ms
	v.State = ServiceStateConnected

	return v
}

func (ms ModuleStatus) String() string {
	var s, sErr = json.Marshal(ms)
	if sErr != nil {
		panic(sErr)
	}

	return string(s)
}

type Version struct {
	Version string `json:"version" bson:"version"`
	Commit  string `json:"commit" bson:"commit"`
	Build   int    `json:"build" bson:"build"`
}

func MakeVersion(version string, commit string, build string) Version {
	var ms Version

	ms.Version = version
	ms.Commit = commit
	build_num, _ := strconv.Atoi(build)
	ms.Build = build_num

	return ms
}
