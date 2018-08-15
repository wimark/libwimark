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

func (self *ModuleStatus) UnmarshalJSON(b []byte) error {
	var doc map[string]json.RawMessage
	if err := json.Unmarshal(b, &doc); err != nil {
		return err
	}
	type MS ModuleStatus
	var s MS
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*self = ModuleStatus(s)
	var meta, has_meta = doc["meta"]
	has_meta = has_meta && !bytes.Equal(meta, []byte("null"))

	switch self.Service {
	case ModuleCPE:
		var m CpeStatusMeta
		if has_meta {
			if err := json.Unmarshal(meta, &m); err != nil {
				return err
			}
		}
		self.Meta = m
	case ModuleTunManager:
		var m TunManagerBroadcastMeta
		if has_meta {
			if err := json.Unmarshal(meta, &m); err != nil {
				return err
			}
		}
		self.Meta = m
	}
	return nil
}

func (self ModuleStatus) Connected() ModuleStatus {
	var v = self
	v.State = ServiceStateConnected

	return v
}

func (self ModuleStatus) Disconnected() ModuleStatus {
	var v = self
	v.State = ServiceStateConnected

	return v
}

func (self ModuleStatus) String() string {
	var s, sErr = json.Marshal(self)
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
	var self Version

	self.Version = version
	self.Commit = commit
	build_num, _ := strconv.Atoi(build)
	self.Build = build_num

	return self
}
