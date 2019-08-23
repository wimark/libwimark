package libwimark

import (
	"encoding/json"
	"errors"
	"github.com/globalsign/mgo/bson"
)

type PortalAuthenticationState string

const PortalAuthenticationStateChecked PortalAuthenticationState = "checked"
const PortalAuthenticationStateNone PortalAuthenticationState = "none"
const PortalAuthenticationStateSent PortalAuthenticationState = "sent"

func (self PortalAuthenticationState) GetPtr() *PortalAuthenticationState { var v = self; return &v }

func (self PortalAuthenticationState) String() string {
	switch self {
	case PortalAuthenticationStateChecked:
		return "checked"
	case PortalAuthenticationStateNone:
		return "none"
	case PortalAuthenticationStateSent:
		return "sent"
	}
	if len(self) == 0 {
		return "none"
	}
	panic(errors.New("Invalid value of PortalAuthenticationState: " + string(self)))
}

func (self *PortalAuthenticationState) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalAuthenticationStateChecked:
		return json.Marshal("checked")
	case PortalAuthenticationStateNone:
		return json.Marshal("none")
	case PortalAuthenticationStateSent:
		return json.Marshal("sent")
	}
	if len(*self) == 0 {
		return json.Marshal("none")
	}
	return nil, errors.New("Invalid value of PortalAuthenticationState: " + string(*self))
}

func (self *PortalAuthenticationState) GetBSON() (interface{}, error) {
	switch *self {
	case PortalAuthenticationStateChecked:
		return "checked", nil
	case PortalAuthenticationStateNone:
		return "none", nil
	case PortalAuthenticationStateSent:
		return "sent", nil
	}
	if len(*self) == 0 {
		return "none", nil
	}
	return nil, errors.New("Invalid value of PortalAuthenticationState: " + string(*self))
}

func (self *PortalAuthenticationState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "checked":
		*self = PortalAuthenticationStateChecked
		return nil
	case "none":
		*self = PortalAuthenticationStateNone
		return nil
	case "sent":
		*self = PortalAuthenticationStateSent
		return nil
	}
	if len(s) == 0 {
		*self = PortalAuthenticationStateNone
		return nil
	}
	return errors.New("Unknown PortalAuthenticationState: " + s)
}

func (self *PortalAuthenticationState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "checked":
		*self = PortalAuthenticationStateChecked
		return nil
	case "none":
		*self = PortalAuthenticationStateNone
		return nil
	case "sent":
		*self = PortalAuthenticationStateSent
		return nil
	}
	if len(s) == 0 {
		*self = PortalAuthenticationStateNone
		return nil
	}
	return errors.New("Unknown PortalAuthenticationState: " + s)
}

type PortalAuthenticationType string

const PortalAuthenticationTypeCallback PortalAuthenticationType = "callback"
const PortalAuthenticationTypeESIA PortalAuthenticationType = "esia"
const PortalAuthenticationTypeNone PortalAuthenticationType = "none"
const PortalAuthenticationTypeSMS PortalAuthenticationType = "sms"
const PortalAuthenticationTypeUserPass PortalAuthenticationType = "userpass"
const PortalAuthenticationTypeVouncher PortalAuthenticationType = "voucher"

func (self PortalAuthenticationType) GetPtr() *PortalAuthenticationType { var v = self; return &v }

func (self PortalAuthenticationType) String() string {
	switch self {
	case PortalAuthenticationTypeCallback:
		return "callback"
	case PortalAuthenticationTypeESIA:
		return "esia"
	case PortalAuthenticationTypeNone:
		return "none"
	case PortalAuthenticationTypeSMS:
		return "sms"
	case PortalAuthenticationTypeUserPass:
		return "userpass"
	case PortalAuthenticationTypeVouncher:
		return "voucher"
	}
	if len(self) == 0 {
		return "sms"
	}
	panic(errors.New("Invalid value of PortalAuthenticationType: " + string(self)))
}

func (self *PortalAuthenticationType) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalAuthenticationTypeCallback:
		return json.Marshal("callback")
	case PortalAuthenticationTypeESIA:
		return json.Marshal("esia")
	case PortalAuthenticationTypeNone:
		return json.Marshal("none")
	case PortalAuthenticationTypeSMS:
		return json.Marshal("sms")
	case PortalAuthenticationTypeUserPass:
		return json.Marshal("userpass")
	case PortalAuthenticationTypeVouncher:
		return json.Marshal("voucher")
	}
	if len(*self) == 0 {
		return json.Marshal("sms")
	}
	return nil, errors.New("Invalid value of PortalAuthenticationType: " + string(*self))
}

func (self *PortalAuthenticationType) GetBSON() (interface{}, error) {
	switch *self {
	case PortalAuthenticationTypeCallback:
		return "callback", nil
	case PortalAuthenticationTypeESIA:
		return "esia", nil
	case PortalAuthenticationTypeNone:
		return "none", nil
	case PortalAuthenticationTypeSMS:
		return "sms", nil
	case PortalAuthenticationTypeUserPass:
		return "userpass", nil
	case PortalAuthenticationTypeVouncher:
		return "voucher", nil
	}
	if len(*self) == 0 {
		return "sms", nil
	}
	return nil, errors.New("Invalid value of PortalAuthenticationType: " + string(*self))
}

func (self *PortalAuthenticationType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "callback":
		*self = PortalAuthenticationTypeCallback
		return nil
	case "esia":
		*self = PortalAuthenticationTypeESIA
		return nil
	case "none":
		*self = PortalAuthenticationTypeNone
		return nil
	case "sms":
		*self = PortalAuthenticationTypeSMS
		return nil
	case "userpass":
		*self = PortalAuthenticationTypeUserPass
		return nil
	case "voucher":
		*self = PortalAuthenticationTypeVouncher
		return nil
	}
	if len(s) == 0 {
		*self = PortalAuthenticationTypeSMS
		return nil
	}
	return errors.New("Unknown PortalAuthenticationType: " + s)
}

func (self *PortalAuthenticationType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "callback":
		*self = PortalAuthenticationTypeCallback
		return nil
	case "esia":
		*self = PortalAuthenticationTypeESIA
		return nil
	case "none":
		*self = PortalAuthenticationTypeNone
		return nil
	case "sms":
		*self = PortalAuthenticationTypeSMS
		return nil
	case "userpass":
		*self = PortalAuthenticationTypeUserPass
		return nil
	case "voucher":
		*self = PortalAuthenticationTypeVouncher
		return nil
	}
	if len(s) == 0 {
		*self = PortalAuthenticationTypeSMS
		return nil
	}
	return errors.New("Unknown PortalAuthenticationType: " + s)
}

type PortalAuthorizationState string

const PortalAuthorizationStateChecked PortalAuthorizationState = "checked"
const PortalAuthorizationStateNone PortalAuthorizationState = "none"

func (self PortalAuthorizationState) GetPtr() *PortalAuthorizationState { var v = self; return &v }

func (self PortalAuthorizationState) String() string {
	switch self {
	case PortalAuthorizationStateChecked:
		return "checked"
	case PortalAuthorizationStateNone:
		return "none"
	}
	if len(self) == 0 {
		return "none"
	}
	panic(errors.New("Invalid value of PortalAuthorizationState: " + string(self)))
}

func (self *PortalAuthorizationState) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalAuthorizationStateChecked:
		return json.Marshal("checked")
	case PortalAuthorizationStateNone:
		return json.Marshal("none")
	}
	if len(*self) == 0 {
		return json.Marshal("none")
	}
	return nil, errors.New("Invalid value of PortalAuthorizationState: " + string(*self))
}

func (self *PortalAuthorizationState) GetBSON() (interface{}, error) {
	switch *self {
	case PortalAuthorizationStateChecked:
		return "checked", nil
	case PortalAuthorizationStateNone:
		return "none", nil
	}
	if len(*self) == 0 {
		return "none", nil
	}
	return nil, errors.New("Invalid value of PortalAuthorizationState: " + string(*self))
}

func (self *PortalAuthorizationState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "checked":
		*self = PortalAuthorizationStateChecked
		return nil
	case "none":
		*self = PortalAuthorizationStateNone
		return nil
	}
	if len(s) == 0 {
		*self = PortalAuthorizationStateNone
		return nil
	}
	return errors.New("Unknown PortalAuthorizationState: " + s)
}

func (self *PortalAuthorizationState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "checked":
		*self = PortalAuthorizationStateChecked
		return nil
	case "none":
		*self = PortalAuthorizationStateNone
		return nil
	}
	if len(s) == 0 {
		*self = PortalAuthorizationStateNone
		return nil
	}
	return errors.New("Unknown PortalAuthorizationState: " + s)
}

type PortalAuthorizationType string

const PortalAuthorizationTypeExtVoucher PortalAuthorizationType = "ext_voucher"
const PortalAuthorizationTypeNone PortalAuthorizationType = "none"
const PortalAuthorizationTypeSponsor PortalAuthorizationType = "sponsor"
const PortalAuthorizationTypeVouncher PortalAuthorizationType = "voucher"

func (self PortalAuthorizationType) GetPtr() *PortalAuthorizationType { var v = self; return &v }

func (self PortalAuthorizationType) String() string {
	switch self {
	case PortalAuthorizationTypeExtVoucher:
		return "ext_voucher"
	case PortalAuthorizationTypeNone:
		return "none"
	case PortalAuthorizationTypeSponsor:
		return "sponsor"
	case PortalAuthorizationTypeVouncher:
		return "voucher"
	}
	if len(self) == 0 {
		return "none"
	}
	panic(errors.New("Invalid value of PortalAuthorizationType: " + string(self)))
}

func (self *PortalAuthorizationType) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalAuthorizationTypeExtVoucher:
		return json.Marshal("ext_voucher")
	case PortalAuthorizationTypeNone:
		return json.Marshal("none")
	case PortalAuthorizationTypeSponsor:
		return json.Marshal("sponsor")
	case PortalAuthorizationTypeVouncher:
		return json.Marshal("voucher")
	}
	if len(*self) == 0 {
		return json.Marshal("none")
	}
	return nil, errors.New("Invalid value of PortalAuthorizationType: " + string(*self))
}

func (self *PortalAuthorizationType) GetBSON() (interface{}, error) {
	switch *self {
	case PortalAuthorizationTypeExtVoucher:
		return "ext_voucher", nil
	case PortalAuthorizationTypeNone:
		return "none", nil
	case PortalAuthorizationTypeSponsor:
		return "sponsor", nil
	case PortalAuthorizationTypeVouncher:
		return "voucher", nil
	}
	if len(*self) == 0 {
		return "none", nil
	}
	return nil, errors.New("Invalid value of PortalAuthorizationType: " + string(*self))
}

func (self *PortalAuthorizationType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ext_voucher":
		*self = PortalAuthorizationTypeExtVoucher
		return nil
	case "none":
		*self = PortalAuthorizationTypeNone
		return nil
	case "sponsor":
		*self = PortalAuthorizationTypeSponsor
		return nil
	case "voucher":
		*self = PortalAuthorizationTypeVouncher
		return nil
	}
	if len(s) == 0 {
		*self = PortalAuthorizationTypeNone
		return nil
	}
	return errors.New("Unknown PortalAuthorizationType: " + s)
}

func (self *PortalAuthorizationType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ext_voucher":
		*self = PortalAuthorizationTypeExtVoucher
		return nil
	case "none":
		*self = PortalAuthorizationTypeNone
		return nil
	case "sponsor":
		*self = PortalAuthorizationTypeSponsor
		return nil
	case "voucher":
		*self = PortalAuthorizationTypeVouncher
		return nil
	}
	if len(s) == 0 {
		*self = PortalAuthorizationTypeNone
		return nil
	}
	return errors.New("Unknown PortalAuthorizationType: " + s)
}

type PortalResponseStatus string

const PortalResponseStatusError PortalResponseStatus = "error"
const PortalResponseStatusSuccess PortalResponseStatus = "success"

func (self PortalResponseStatus) GetPtr() *PortalResponseStatus { var v = self; return &v }

func (self PortalResponseStatus) String() string {
	switch self {
	case PortalResponseStatusError:
		return "error"
	case PortalResponseStatusSuccess:
		return "success"
	}
	if len(self) == 0 {
		return "success"
	}
	panic(errors.New("Invalid value of PortalResponseStatus: " + string(self)))
}

func (self *PortalResponseStatus) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalResponseStatusError:
		return json.Marshal("error")
	case PortalResponseStatusSuccess:
		return json.Marshal("success")
	}
	if len(*self) == 0 {
		return json.Marshal("success")
	}
	return nil, errors.New("Invalid value of PortalResponseStatus: " + string(*self))
}

func (self *PortalResponseStatus) GetBSON() (interface{}, error) {
	switch *self {
	case PortalResponseStatusError:
		return "error", nil
	case PortalResponseStatusSuccess:
		return "success", nil
	}
	if len(*self) == 0 {
		return "success", nil
	}
	return nil, errors.New("Invalid value of PortalResponseStatus: " + string(*self))
}

func (self *PortalResponseStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "error":
		*self = PortalResponseStatusError
		return nil
	case "success":
		*self = PortalResponseStatusSuccess
		return nil
	}
	if len(s) == 0 {
		*self = PortalResponseStatusSuccess
		return nil
	}
	return errors.New("Unknown PortalResponseStatus: " + s)
}

func (self *PortalResponseStatus) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "error":
		*self = PortalResponseStatusError
		return nil
	case "success":
		*self = PortalResponseStatusSuccess
		return nil
	}
	if len(s) == 0 {
		*self = PortalResponseStatusSuccess
		return nil
	}
	return errors.New("Unknown PortalResponseStatus: " + s)
}

type PortalUserState string

const PortalUserStateAdvertisement PortalUserState = "advertisement"
const PortalUserStateAuthentication PortalUserState = "authentication"
const PortalUserStateAuthorization PortalUserState = "authorization"
const PortalUserStatePassing PortalUserState = "passing"

func (self PortalUserState) GetPtr() *PortalUserState { var v = self; return &v }

func (self PortalUserState) String() string {
	switch self {
	case PortalUserStateAdvertisement:
		return "advertisement"
	case PortalUserStateAuthentication:
		return "authentication"
	case PortalUserStateAuthorization:
		return "authorization"
	case PortalUserStatePassing:
		return "passing"
	}
	if len(self) == 0 {
		return "authentication"
	}
	panic(errors.New("Invalid value of PortalUserState: " + string(self)))
}

func (self *PortalUserState) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalUserStateAdvertisement:
		return json.Marshal("advertisement")
	case PortalUserStateAuthentication:
		return json.Marshal("authentication")
	case PortalUserStateAuthorization:
		return json.Marshal("authorization")
	case PortalUserStatePassing:
		return json.Marshal("passing")
	}
	if len(*self) == 0 {
		return json.Marshal("authentication")
	}
	return nil, errors.New("Invalid value of PortalUserState: " + string(*self))
}

func (self *PortalUserState) GetBSON() (interface{}, error) {
	switch *self {
	case PortalUserStateAdvertisement:
		return "advertisement", nil
	case PortalUserStateAuthentication:
		return "authentication", nil
	case PortalUserStateAuthorization:
		return "authorization", nil
	case PortalUserStatePassing:
		return "passing", nil
	}
	if len(*self) == 0 {
		return "authentication", nil
	}
	return nil, errors.New("Invalid value of PortalUserState: " + string(*self))
}

func (self *PortalUserState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "advertisement":
		*self = PortalUserStateAdvertisement
		return nil
	case "authentication":
		*self = PortalUserStateAuthentication
		return nil
	case "authorization":
		*self = PortalUserStateAuthorization
		return nil
	case "passing":
		*self = PortalUserStatePassing
		return nil
	}
	if len(s) == 0 {
		*self = PortalUserStateAuthentication
		return nil
	}
	return errors.New("Unknown PortalUserState: " + s)
}

func (self *PortalUserState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "advertisement":
		*self = PortalUserStateAdvertisement
		return nil
	case "authentication":
		*self = PortalUserStateAuthentication
		return nil
	case "authorization":
		*self = PortalUserStateAuthorization
		return nil
	case "passing":
		*self = PortalUserStatePassing
		return nil
	}
	if len(s) == 0 {
		*self = PortalUserStateAuthentication
		return nil
	}
	return errors.New("Unknown PortalUserState: " + s)
}
