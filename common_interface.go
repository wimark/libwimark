package libwimark

import (
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
