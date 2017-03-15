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
