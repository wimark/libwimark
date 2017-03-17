package libwimark

import (
	"encoding/json"
)
import (
	"errors"
)

type CPEInterfaceTypeIface interface {
	CPEInterfaceTypeIfaceFunc()
}
type CPEInterfaceType struct{ CPEInterfaceTypeIface }

func (self *CPEInterfaceType) Value() CPEInterfaceTypeIface { return self.CPEInterfaceTypeIface }

type InterfaceWiFi struct{}

func (InterfaceWiFi) CPEInterfaceTypeIfaceFunc() {}

type InterfaceWired struct{}

func (InterfaceWired) CPEInterfaceTypeIfaceFunc() {}
func (self *CPEInterfaceType) MarshalJSON() ([]byte, error) {
	switch self.CPEInterfaceTypeIface.(type) {
	case InterfaceWiFi:
		return json.Marshal("wifi")
	case InterfaceWired:
		return json.Marshal("wired")

	}
	return nil, errors.New("Not implemented")

}
func (self *CPEInterfaceType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "wifi":
		self.CPEInterfaceTypeIface = InterfaceWiFi{}
		return nil
	case "wired":
		self.CPEInterfaceTypeIface = InterfaceWired{}
		return nil

	}
	return errors.New("Unknown CPEInterfaceType")

}

type ConfigurationStatusIface interface {
	ConfigurationStatusIfaceFunc()
}
type ConfigurationStatus struct{ ConfigurationStatusIface }

func (self *ConfigurationStatus) Value() ConfigurationStatusIface {
	return self.ConfigurationStatusIface
}

type StatusEmpty struct{}

func (StatusEmpty) ConfigurationStatusIfaceFunc() {}

type StatusError struct{}

func (StatusError) ConfigurationStatusIfaceFunc() {}

type StatusOK struct{}

func (StatusOK) ConfigurationStatusIfaceFunc() {}

type StatusPending struct{}

func (StatusPending) ConfigurationStatusIfaceFunc() {}
func (self *ConfigurationStatus) MarshalJSON() ([]byte, error) {
	switch self.ConfigurationStatusIface.(type) {
	case StatusEmpty:
		return json.Marshal("empty")
	case StatusError:
		return json.Marshal("error")
	case StatusOK:
		return json.Marshal("ok")
	case StatusPending:
		return json.Marshal("pending")

	}
	return nil, errors.New("Not implemented")

}
func (self *ConfigurationStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "empty":
		self.ConfigurationStatusIface = StatusEmpty{}
		return nil
	case "error":
		self.ConfigurationStatusIface = StatusError{}
		return nil
	case "ok":
		self.ConfigurationStatusIface = StatusOK{}
		return nil
	case "pending":
		self.ConfigurationStatusIface = StatusPending{}
		return nil

	}
	return errors.New("Unknown ConfigurationStatus")

}

type ModuleIface interface {
	ModuleIfaceFunc()
}
type Module struct{ ModuleIface }

func (self *Module) Value() ModuleIface { return self.ModuleIface }

type ModuleAny struct{}

func (ModuleAny) ModuleIfaceFunc() {}

type ModuleBackend struct{}

func (ModuleBackend) ModuleIfaceFunc() {}

type ModuleCPE struct{}

func (ModuleCPE) ModuleIfaceFunc() {}

type ModuleConfig struct{}

func (ModuleConfig) ModuleIfaceFunc() {}

type ModuleDB struct{}

func (ModuleDB) ModuleIfaceFunc() {}

type ModuleStat struct{}

func (ModuleStat) ModuleIfaceFunc() {}
func (self *Module) MarshalJSON() ([]byte, error) {
	switch self.ModuleIface.(type) {
	case ModuleAny:
		return json.Marshal("+")
	case ModuleBackend:
		return json.Marshal("BACKEND")
	case ModuleCPE:
		return json.Marshal("CPE")
	case ModuleConfig:
		return json.Marshal("CONFIG")
	case ModuleDB:
		return json.Marshal("DB")
	case ModuleStat:
		return json.Marshal("STAT")

	}
	return nil, errors.New("Not implemented")

}
func (self *Module) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "+":
		self.ModuleIface = ModuleAny{}
		return nil
	case "BACKEND":
		self.ModuleIface = ModuleBackend{}
		return nil
	case "CPE":
		self.ModuleIface = ModuleCPE{}
		return nil
	case "CONFIG":
		self.ModuleIface = ModuleConfig{}
		return nil
	case "DB":
		self.ModuleIface = ModuleDB{}
		return nil
	case "STAT":
		self.ModuleIface = ModuleStat{}
		return nil

	}
	return errors.New("Unknown Module")

}

type OperationIface interface {
	OperationIfaceFunc()
}
type Operation struct{ OperationIface }

func (self *Operation) Value() OperationIface { return self.OperationIface }

type OperationAny struct{}

func (OperationAny) OperationIfaceFunc() {}

type OperationCreate struct{}

func (OperationCreate) OperationIfaceFunc() {}

type OperationDelete struct{}

func (OperationDelete) OperationIfaceFunc() {}

type OperationLuaScript struct{}

func (OperationLuaScript) OperationIfaceFunc() {}

type OperationRead struct{}

func (OperationRead) OperationIfaceFunc() {}

type OperationSHScript struct{}

func (OperationSHScript) OperationIfaceFunc() {}

type OperationUpdate struct{}

func (OperationUpdate) OperationIfaceFunc() {}
func (self *Operation) MarshalJSON() ([]byte, error) {
	switch self.OperationIface.(type) {
	case OperationAny:
		return json.Marshal("+")
	case OperationCreate:
		return json.Marshal("C")
	case OperationDelete:
		return json.Marshal("D")
	case OperationLuaScript:
		return json.Marshal("LUA")
	case OperationRead:
		return json.Marshal("R")
	case OperationSHScript:
		return json.Marshal("SH")
	case OperationUpdate:
		return json.Marshal("U")

	}
	return nil, errors.New("Not implemented")

}
func (self *Operation) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "+":
		self.OperationIface = OperationAny{}
		return nil
	case "C":
		self.OperationIface = OperationCreate{}
		return nil
	case "D":
		self.OperationIface = OperationDelete{}
		return nil
	case "LUA":
		self.OperationIface = OperationLuaScript{}
		return nil
	case "R":
		self.OperationIface = OperationRead{}
		return nil
	case "SH":
		self.OperationIface = OperationSHScript{}
		return nil
	case "U":
		self.OperationIface = OperationUpdate{}
		return nil

	}
	return errors.New("Unknown Operation")

}

type RadiusTypeIface interface {
	RadiusTypeIfaceFunc()
}
type RadiusType struct{ RadiusTypeIface }

func (self *RadiusType) Value() RadiusTypeIface { return self.RadiusTypeIface }

type RadiusAccounting struct{}

func (RadiusAccounting) RadiusTypeIfaceFunc() {}

type RadiusAuthentication struct{}

func (RadiusAuthentication) RadiusTypeIfaceFunc() {}

type RadiusBoth struct{}

func (RadiusBoth) RadiusTypeIfaceFunc() {}
func (self *RadiusType) MarshalJSON() ([]byte, error) {
	switch self.RadiusTypeIface.(type) {
	case RadiusAccounting:
		return json.Marshal("acc")
	case RadiusAuthentication:
		return json.Marshal("auth")
	case RadiusBoth:
		return json.Marshal("both")

	}
	return nil, errors.New("Not implemented")

}
func (self *RadiusType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "acc":
		self.RadiusTypeIface = RadiusAccounting{}
		return nil
	case "auth":
		self.RadiusTypeIface = RadiusAuthentication{}
		return nil
	case "both":
		self.RadiusTypeIface = RadiusBoth{}
		return nil

	}
	return errors.New("Unknown RadiusType")

}

type SecuritySuiteIface interface {
	SecuritySuiteIfaceFunc()
}
type SecuritySuite struct{ SecuritySuiteIface }

func (self *SecuritySuite) Value() SecuritySuiteIface { return self.SecuritySuiteIface }

type AES struct{}

func (AES) SecuritySuiteIfaceFunc() {}

type TKIP struct{}

func (TKIP) SecuritySuiteIfaceFunc() {}
func (self *SecuritySuite) MarshalJSON() ([]byte, error) {
	switch self.SecuritySuiteIface.(type) {
	case AES:
		return json.Marshal("aes")
	case TKIP:
		return json.Marshal("tkip")

	}
	return nil, errors.New("Not implemented")

}
func (self *SecuritySuite) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "aes":
		self.SecuritySuiteIface = AES{}
		return nil
	case "tkip":
		self.SecuritySuiteIface = TKIP{}
		return nil

	}
	return errors.New("Unknown SecuritySuite")

}

type SecurityTypeIface interface {
	SecurityTypeIfaceFunc()
}
type SecurityType struct{ SecurityTypeIface }

func (self *SecurityType) Value() SecurityTypeIface { return self.SecurityTypeIface }

type WPA2Enterprise struct{}

func (WPA2Enterprise) SecurityTypeIfaceFunc() {}

type WPA2Personal struct{}

func (WPA2Personal) SecurityTypeIfaceFunc() {}
func (self *SecurityType) MarshalJSON() ([]byte, error) {
	switch self.SecurityTypeIface.(type) {
	case WPA2Enterprise:
		return json.Marshal("wpa2enterprise")
	case WPA2Personal:
		return json.Marshal("wpa2personal")

	}
	return nil, errors.New("Not implemented")

}
func (self *SecurityType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "wpa2enterprise":
		self.SecurityTypeIface = WPA2Enterprise{}
		return nil
	case "wpa2personal":
		self.SecurityTypeIface = WPA2Personal{}
		return nil

	}
	return errors.New("Unknown SecurityType")

}
