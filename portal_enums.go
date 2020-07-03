package libwimark

import (
	"encoding/json"
	"errors"
	"github.com/globalsign/mgo/bson"
)

type PortalActionListType string

const PortalActionListTypeAddToAccessList PortalActionListType = "add_al"
const PortalActionListTypeAddToBlackList PortalActionListType = "add_bl"
const PortalActionListTypeDelFromAccessList PortalActionListType = "del_al"
const PortalActionListTypeDelFromBlackList PortalActionListType = "del_bl"
const PortalActionListTypeNone PortalActionListType = "none"

func (self PortalActionListType) GetPtr() *PortalActionListType { var v = self; return &v }

func (self PortalActionListType) String() string {
	switch self {
	case PortalActionListTypeAddToAccessList:
		return "add_al"
	case PortalActionListTypeAddToBlackList:
		return "add_bl"
	case PortalActionListTypeDelFromAccessList:
		return "del_al"
	case PortalActionListTypeDelFromBlackList:
		return "del_bl"
	case PortalActionListTypeNone:
		return "none"
	}
	if len(self) == 0 {
		return "none"
	}
	panic(errors.New("Invalid value of PortalActionListType: " + string(self)))
}

func (self *PortalActionListType) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalActionListTypeAddToAccessList:
		return json.Marshal("add_al")
	case PortalActionListTypeAddToBlackList:
		return json.Marshal("add_bl")
	case PortalActionListTypeDelFromAccessList:
		return json.Marshal("del_al")
	case PortalActionListTypeDelFromBlackList:
		return json.Marshal("del_bl")
	case PortalActionListTypeNone:
		return json.Marshal("none")
	}
	if len(*self) == 0 {
		return json.Marshal("none")
	}
	return nil, errors.New("Invalid value of PortalActionListType: " + string(*self))
}

func (self *PortalActionListType) GetBSON() (interface{}, error) {
	switch *self {
	case PortalActionListTypeAddToAccessList:
		return "add_al", nil
	case PortalActionListTypeAddToBlackList:
		return "add_bl", nil
	case PortalActionListTypeDelFromAccessList:
		return "del_al", nil
	case PortalActionListTypeDelFromBlackList:
		return "del_bl", nil
	case PortalActionListTypeNone:
		return "none", nil
	}
	if len(*self) == 0 {
		return "none", nil
	}
	return nil, errors.New("Invalid value of PortalActionListType: " + string(*self))
}

func (self *PortalActionListType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "add_al":
		*self = PortalActionListTypeAddToAccessList
		return nil
	case "add_bl":
		*self = PortalActionListTypeAddToBlackList
		return nil
	case "del_al":
		*self = PortalActionListTypeDelFromAccessList
		return nil
	case "del_bl":
		*self = PortalActionListTypeDelFromBlackList
		return nil
	case "none":
		*self = PortalActionListTypeNone
		return nil
	}
	if len(s) == 0 {
		*self = PortalActionListTypeNone
		return nil
	}
	return errors.New("Unknown PortalActionListType: " + s)
}

func (self *PortalActionListType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "add_al":
		*self = PortalActionListTypeAddToAccessList
		return nil
	case "add_bl":
		*self = PortalActionListTypeAddToBlackList
		return nil
	case "del_al":
		*self = PortalActionListTypeDelFromAccessList
		return nil
	case "del_bl":
		*self = PortalActionListTypeDelFromBlackList
		return nil
	case "none":
		*self = PortalActionListTypeNone
		return nil
	}
	if len(s) == 0 {
		*self = PortalActionListTypeNone
		return nil
	}
	return errors.New("Unknown PortalActionListType: " + s)
}

type PortalAdvertisementState string

const PortalAdvertisementStateChecked PortalAdvertisementState = "checked"
const PortalAdvertisementStateNeed PortalAdvertisementState = "none"

func (self PortalAdvertisementState) GetPtr() *PortalAdvertisementState { var v = self; return &v }

func (self PortalAdvertisementState) String() string {
	switch self {
	case PortalAdvertisementStateChecked:
		return "checked"
	case PortalAdvertisementStateNeed:
		return "none"
	}
	if len(self) == 0 {
		return "none"
	}
	panic(errors.New("Invalid value of PortalAdvertisementState: " + string(self)))
}

func (self *PortalAdvertisementState) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalAdvertisementStateChecked:
		return json.Marshal("checked")
	case PortalAdvertisementStateNeed:
		return json.Marshal("none")
	}
	if len(*self) == 0 {
		return json.Marshal("none")
	}
	return nil, errors.New("Invalid value of PortalAdvertisementState: " + string(*self))
}

func (self *PortalAdvertisementState) GetBSON() (interface{}, error) {
	switch *self {
	case PortalAdvertisementStateChecked:
		return "checked", nil
	case PortalAdvertisementStateNeed:
		return "none", nil
	}
	if len(*self) == 0 {
		return "none", nil
	}
	return nil, errors.New("Invalid value of PortalAdvertisementState: " + string(*self))
}

func (self *PortalAdvertisementState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "checked":
		*self = PortalAdvertisementStateChecked
		return nil
	case "none":
		*self = PortalAdvertisementStateNeed
		return nil
	}
	if len(s) == 0 {
		*self = PortalAdvertisementStateNeed
		return nil
	}
	return errors.New("Unknown PortalAdvertisementState: " + s)
}

func (self *PortalAdvertisementState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "checked":
		*self = PortalAdvertisementStateChecked
		return nil
	case "none":
		*self = PortalAdvertisementStateNeed
		return nil
	}
	if len(s) == 0 {
		*self = PortalAdvertisementStateNeed
		return nil
	}
	return errors.New("Unknown PortalAdvertisementState: " + s)
}

type PortalAdvertisementType string

const PortalAdvertisementTypeFlash PortalAdvertisementType = "flash"
const PortalAdvertisementTypeIframe PortalAdvertisementType = "iframe"
const PortalAdvertisementTypeImage PortalAdvertisementType = "image"
const PortalAdvertisementTypePoll PortalAdvertisementType = "poll"
const PortalAdvertisementTypePollIm PortalAdvertisementType = "poll_image"
const PortalAdvertisementTypePollQuality PortalAdvertisementType = "poll_quality"
const PortalAdvertisementTypePollUserData PortalAdvertisementType = "poll_user_data"
const PortalAdvertisementTypeVideo PortalAdvertisementType = "video"

func (self PortalAdvertisementType) GetPtr() *PortalAdvertisementType { var v = self; return &v }

func (self PortalAdvertisementType) String() string {
	switch self {
	case PortalAdvertisementTypeFlash:
		return "flash"
	case PortalAdvertisementTypeIframe:
		return "iframe"
	case PortalAdvertisementTypeImage:
		return "image"
	case PortalAdvertisementTypePoll:
		return "poll"
	case PortalAdvertisementTypePollIm:
		return "poll_image"
	case PortalAdvertisementTypePollQuality:
		return "poll_quality"
	case PortalAdvertisementTypePollUserData:
		return "poll_user_data"
	case PortalAdvertisementTypeVideo:
		return "video"
	}
	if len(self) == 0 {
		return "image"
	}
	panic(errors.New("Invalid value of PortalAdvertisementType: " + string(self)))
}

func (self *PortalAdvertisementType) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalAdvertisementTypeFlash:
		return json.Marshal("flash")
	case PortalAdvertisementTypeIframe:
		return json.Marshal("iframe")
	case PortalAdvertisementTypeImage:
		return json.Marshal("image")
	case PortalAdvertisementTypePoll:
		return json.Marshal("poll")
	case PortalAdvertisementTypePollIm:
		return json.Marshal("poll_image")
	case PortalAdvertisementTypePollQuality:
		return json.Marshal("poll_quality")
	case PortalAdvertisementTypePollUserData:
		return json.Marshal("poll_user_data")
	case PortalAdvertisementTypeVideo:
		return json.Marshal("video")
	}
	if len(*self) == 0 {
		return json.Marshal("image")
	}
	return nil, errors.New("Invalid value of PortalAdvertisementType: " + string(*self))
}

func (self *PortalAdvertisementType) GetBSON() (interface{}, error) {
	switch *self {
	case PortalAdvertisementTypeFlash:
		return "flash", nil
	case PortalAdvertisementTypeIframe:
		return "iframe", nil
	case PortalAdvertisementTypeImage:
		return "image", nil
	case PortalAdvertisementTypePoll:
		return "poll", nil
	case PortalAdvertisementTypePollIm:
		return "poll_image", nil
	case PortalAdvertisementTypePollQuality:
		return "poll_quality", nil
	case PortalAdvertisementTypePollUserData:
		return "poll_user_data", nil
	case PortalAdvertisementTypeVideo:
		return "video", nil
	}
	if len(*self) == 0 {
		return "image", nil
	}
	return nil, errors.New("Invalid value of PortalAdvertisementType: " + string(*self))
}

func (self *PortalAdvertisementType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "flash":
		*self = PortalAdvertisementTypeFlash
		return nil
	case "iframe":
		*self = PortalAdvertisementTypeIframe
		return nil
	case "image":
		*self = PortalAdvertisementTypeImage
		return nil
	case "poll":
		*self = PortalAdvertisementTypePoll
		return nil
	case "poll_image":
		*self = PortalAdvertisementTypePollIm
		return nil
	case "poll_quality":
		*self = PortalAdvertisementTypePollQuality
		return nil
	case "poll_user_data":
		*self = PortalAdvertisementTypePollUserData
		return nil
	case "video":
		*self = PortalAdvertisementTypeVideo
		return nil
	}
	if len(s) == 0 {
		*self = PortalAdvertisementTypeImage
		return nil
	}
	return errors.New("Unknown PortalAdvertisementType: " + s)
}

func (self *PortalAdvertisementType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "flash":
		*self = PortalAdvertisementTypeFlash
		return nil
	case "iframe":
		*self = PortalAdvertisementTypeIframe
		return nil
	case "image":
		*self = PortalAdvertisementTypeImage
		return nil
	case "poll":
		*self = PortalAdvertisementTypePoll
		return nil
	case "poll_image":
		*self = PortalAdvertisementTypePollIm
		return nil
	case "poll_quality":
		*self = PortalAdvertisementTypePollQuality
		return nil
	case "poll_user_data":
		*self = PortalAdvertisementTypePollUserData
		return nil
	case "video":
		*self = PortalAdvertisementTypeVideo
		return nil
	}
	if len(s) == 0 {
		*self = PortalAdvertisementTypeImage
		return nil
	}
	return errors.New("Unknown PortalAdvertisementType: " + s)
}

type PortalAuthenticationState string

const PortalAuthenticationStateChecked PortalAuthenticationState = "checked"
const PortalAuthenticationStateNeed PortalAuthenticationState = "need"
const PortalAuthenticationStateSent PortalAuthenticationState = "sent"

func (self PortalAuthenticationState) GetPtr() *PortalAuthenticationState { var v = self; return &v }

func (self PortalAuthenticationState) String() string {
	switch self {
	case PortalAuthenticationStateChecked:
		return "checked"
	case PortalAuthenticationStateNeed:
		return "need"
	case PortalAuthenticationStateSent:
		return "sent"
	}
	if len(self) == 0 {
		return "need"
	}
	panic(errors.New("Invalid value of PortalAuthenticationState: " + string(self)))
}

func (self *PortalAuthenticationState) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalAuthenticationStateChecked:
		return json.Marshal("checked")
	case PortalAuthenticationStateNeed:
		return json.Marshal("need")
	case PortalAuthenticationStateSent:
		return json.Marshal("sent")
	}
	if len(*self) == 0 {
		return json.Marshal("need")
	}
	return nil, errors.New("Invalid value of PortalAuthenticationState: " + string(*self))
}

func (self *PortalAuthenticationState) GetBSON() (interface{}, error) {
	switch *self {
	case PortalAuthenticationStateChecked:
		return "checked", nil
	case PortalAuthenticationStateNeed:
		return "need", nil
	case PortalAuthenticationStateSent:
		return "sent", nil
	}
	if len(*self) == 0 {
		return "need", nil
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
	case "need":
		*self = PortalAuthenticationStateNeed
		return nil
	case "sent":
		*self = PortalAuthenticationStateSent
		return nil
	}
	if len(s) == 0 {
		*self = PortalAuthenticationStateNeed
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
	case "need":
		*self = PortalAuthenticationStateNeed
		return nil
	case "sent":
		*self = PortalAuthenticationStateSent
		return nil
	}
	if len(s) == 0 {
		*self = PortalAuthenticationStateNeed
		return nil
	}
	return errors.New("Unknown PortalAuthenticationState: " + s)
}

type PortalAuthenticationType string

const PortalAuthenticationTypeCallback PortalAuthenticationType = "callback"
const PortalAuthenticationTypeESIA PortalAuthenticationType = "esia"
const PortalAuthenticationTypeEmail PortalAuthenticationType = "email"
const PortalAuthenticationTypeHotelVoucher PortalAuthenticationType = "hotel_voucher"
const PortalAuthenticationTypeNone PortalAuthenticationType = "none"
const PortalAuthenticationTypeSMS PortalAuthenticationType = "sms"
const PortalAuthenticationTypeUserPass PortalAuthenticationType = "userpass"
const PortalAuthenticationTypeVoucher PortalAuthenticationType = "voucher"

func (self PortalAuthenticationType) GetPtr() *PortalAuthenticationType { var v = self; return &v }

func (self PortalAuthenticationType) String() string {
	switch self {
	case PortalAuthenticationTypeCallback:
		return "callback"
	case PortalAuthenticationTypeESIA:
		return "esia"
	case PortalAuthenticationTypeEmail:
		return "email"
	case PortalAuthenticationTypeHotelVoucher:
		return "hotel_voucher"
	case PortalAuthenticationTypeNone:
		return "none"
	case PortalAuthenticationTypeSMS:
		return "sms"
	case PortalAuthenticationTypeUserPass:
		return "userpass"
	case PortalAuthenticationTypeVoucher:
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
	case PortalAuthenticationTypeEmail:
		return json.Marshal("email")
	case PortalAuthenticationTypeHotelVoucher:
		return json.Marshal("hotel_voucher")
	case PortalAuthenticationTypeNone:
		return json.Marshal("none")
	case PortalAuthenticationTypeSMS:
		return json.Marshal("sms")
	case PortalAuthenticationTypeUserPass:
		return json.Marshal("userpass")
	case PortalAuthenticationTypeVoucher:
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
	case PortalAuthenticationTypeEmail:
		return "email", nil
	case PortalAuthenticationTypeHotelVoucher:
		return "hotel_voucher", nil
	case PortalAuthenticationTypeNone:
		return "none", nil
	case PortalAuthenticationTypeSMS:
		return "sms", nil
	case PortalAuthenticationTypeUserPass:
		return "userpass", nil
	case PortalAuthenticationTypeVoucher:
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
	case "email":
		*self = PortalAuthenticationTypeEmail
		return nil
	case "hotel_voucher":
		*self = PortalAuthenticationTypeHotelVoucher
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
		*self = PortalAuthenticationTypeVoucher
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
	case "email":
		*self = PortalAuthenticationTypeEmail
		return nil
	case "hotel_voucher":
		*self = PortalAuthenticationTypeHotelVoucher
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
		*self = PortalAuthenticationTypeVoucher
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
const PortalAuthorizationStateNeed PortalAuthorizationState = "none"

func (self PortalAuthorizationState) GetPtr() *PortalAuthorizationState { var v = self; return &v }

func (self PortalAuthorizationState) String() string {
	switch self {
	case PortalAuthorizationStateChecked:
		return "checked"
	case PortalAuthorizationStateNeed:
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
	case PortalAuthorizationStateNeed:
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
	case PortalAuthorizationStateNeed:
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
		*self = PortalAuthorizationStateNeed
		return nil
	}
	if len(s) == 0 {
		*self = PortalAuthorizationStateNeed
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
		*self = PortalAuthorizationStateNeed
		return nil
	}
	if len(s) == 0 {
		*self = PortalAuthorizationStateNeed
		return nil
	}
	return errors.New("Unknown PortalAuthorizationState: " + s)
}

type PortalAuthorizationType string

const PortalAuthorizationTypeExtVoucher PortalAuthorizationType = "ext_voucher"
const PortalAuthorizationTypeFB PortalAuthorizationType = "facebook"
const PortalAuthorizationTypeFree PortalAuthorizationType = "free"
const PortalAuthorizationTypeHotelVoucher PortalAuthorizationType = "hotel_voucher"
const PortalAuthorizationTypeIG PortalAuthorizationType = "instagram"
const PortalAuthorizationTypeNone PortalAuthorizationType = "none"
const PortalAuthorizationTypeSkip PortalAuthorizationType = "skip"
const PortalAuthorizationTypeSponsor PortalAuthorizationType = "sponsor"
const PortalAuthorizationTypeStaff PortalAuthorizationType = "staff"
const PortalAuthorizationTypeVK PortalAuthorizationType = "vk"
const PortalAuthorizationTypeVoucher PortalAuthorizationType = "voucher"

func (self PortalAuthorizationType) GetPtr() *PortalAuthorizationType { var v = self; return &v }

func (self PortalAuthorizationType) String() string {
	switch self {
	case PortalAuthorizationTypeExtVoucher:
		return "ext_voucher"
	case PortalAuthorizationTypeFB:
		return "facebook"
	case PortalAuthorizationTypeFree:
		return "free"
	case PortalAuthorizationTypeHotelVoucher:
		return "hotel_voucher"
	case PortalAuthorizationTypeIG:
		return "instagram"
	case PortalAuthorizationTypeNone:
		return "none"
	case PortalAuthorizationTypeSkip:
		return "skip"
	case PortalAuthorizationTypeSponsor:
		return "sponsor"
	case PortalAuthorizationTypeStaff:
		return "staff"
	case PortalAuthorizationTypeVK:
		return "vk"
	case PortalAuthorizationTypeVoucher:
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
	case PortalAuthorizationTypeFB:
		return json.Marshal("facebook")
	case PortalAuthorizationTypeFree:
		return json.Marshal("free")
	case PortalAuthorizationTypeHotelVoucher:
		return json.Marshal("hotel_voucher")
	case PortalAuthorizationTypeIG:
		return json.Marshal("instagram")
	case PortalAuthorizationTypeNone:
		return json.Marshal("none")
	case PortalAuthorizationTypeSkip:
		return json.Marshal("skip")
	case PortalAuthorizationTypeSponsor:
		return json.Marshal("sponsor")
	case PortalAuthorizationTypeStaff:
		return json.Marshal("staff")
	case PortalAuthorizationTypeVK:
		return json.Marshal("vk")
	case PortalAuthorizationTypeVoucher:
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
	case PortalAuthorizationTypeFB:
		return "facebook", nil
	case PortalAuthorizationTypeFree:
		return "free", nil
	case PortalAuthorizationTypeHotelVoucher:
		return "hotel_voucher", nil
	case PortalAuthorizationTypeIG:
		return "instagram", nil
	case PortalAuthorizationTypeNone:
		return "none", nil
	case PortalAuthorizationTypeSkip:
		return "skip", nil
	case PortalAuthorizationTypeSponsor:
		return "sponsor", nil
	case PortalAuthorizationTypeStaff:
		return "staff", nil
	case PortalAuthorizationTypeVK:
		return "vk", nil
	case PortalAuthorizationTypeVoucher:
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
	case "facebook":
		*self = PortalAuthorizationTypeFB
		return nil
	case "free":
		*self = PortalAuthorizationTypeFree
		return nil
	case "hotel_voucher":
		*self = PortalAuthorizationTypeHotelVoucher
		return nil
	case "instagram":
		*self = PortalAuthorizationTypeIG
		return nil
	case "none":
		*self = PortalAuthorizationTypeNone
		return nil
	case "skip":
		*self = PortalAuthorizationTypeSkip
		return nil
	case "sponsor":
		*self = PortalAuthorizationTypeSponsor
		return nil
	case "staff":
		*self = PortalAuthorizationTypeStaff
		return nil
	case "vk":
		*self = PortalAuthorizationTypeVK
		return nil
	case "voucher":
		*self = PortalAuthorizationTypeVoucher
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
	case "facebook":
		*self = PortalAuthorizationTypeFB
		return nil
	case "free":
		*self = PortalAuthorizationTypeFree
		return nil
	case "hotel_voucher":
		*self = PortalAuthorizationTypeHotelVoucher
		return nil
	case "instagram":
		*self = PortalAuthorizationTypeIG
		return nil
	case "none":
		*self = PortalAuthorizationTypeNone
		return nil
	case "skip":
		*self = PortalAuthorizationTypeSkip
		return nil
	case "sponsor":
		*self = PortalAuthorizationTypeSponsor
		return nil
	case "staff":
		*self = PortalAuthorizationTypeStaff
		return nil
	case "vk":
		*self = PortalAuthorizationTypeVK
		return nil
	case "voucher":
		*self = PortalAuthorizationTypeVoucher
		return nil
	}
	if len(s) == 0 {
		*self = PortalAuthorizationTypeNone
		return nil
	}
	return errors.New("Unknown PortalAuthorizationType: " + s)
}

type PortalOSType string

const PortalOSTypeAndroid PortalOSType = "Android"
const PortalOSTypeIOS PortalOSType = "iPhone OS"
const PortalOSTypeLinux PortalOSType = "Linux"
const PortalOSTypeMacOs PortalOSType = "Mac Os"
const PortalOSTypeNone PortalOSType = "none"
const PortalOSTypeWindows PortalOSType = "Windows"

func (self PortalOSType) GetPtr() *PortalOSType { var v = self; return &v }

func (self PortalOSType) String() string {
	switch self {
	case PortalOSTypeAndroid:
		return "Android"
	case PortalOSTypeIOS:
		return "iPhone OS"
	case PortalOSTypeLinux:
		return "Linux"
	case PortalOSTypeMacOs:
		return "Mac Os"
	case PortalOSTypeNone:
		return "none"
	case PortalOSTypeWindows:
		return "Windows"
	}
	if len(self) == 0 {
		return "none"
	}
	panic(errors.New("Invalid value of PortalOSType: " + string(self)))
}

func (self *PortalOSType) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalOSTypeAndroid:
		return json.Marshal("Android")
	case PortalOSTypeIOS:
		return json.Marshal("iPhone OS")
	case PortalOSTypeLinux:
		return json.Marshal("Linux")
	case PortalOSTypeMacOs:
		return json.Marshal("Mac Os")
	case PortalOSTypeNone:
		return json.Marshal("none")
	case PortalOSTypeWindows:
		return json.Marshal("Windows")
	}
	if len(*self) == 0 {
		return json.Marshal("none")
	}
	return nil, errors.New("Invalid value of PortalOSType: " + string(*self))
}

func (self *PortalOSType) GetBSON() (interface{}, error) {
	switch *self {
	case PortalOSTypeAndroid:
		return "Android", nil
	case PortalOSTypeIOS:
		return "iPhone OS", nil
	case PortalOSTypeLinux:
		return "Linux", nil
	case PortalOSTypeMacOs:
		return "Mac Os", nil
	case PortalOSTypeNone:
		return "none", nil
	case PortalOSTypeWindows:
		return "Windows", nil
	}
	if len(*self) == 0 {
		return "none", nil
	}
	return nil, errors.New("Invalid value of PortalOSType: " + string(*self))
}

func (self *PortalOSType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "Android":
		*self = PortalOSTypeAndroid
		return nil
	case "iPhone OS":
		*self = PortalOSTypeIOS
		return nil
	case "Linux":
		*self = PortalOSTypeLinux
		return nil
	case "Mac Os":
		*self = PortalOSTypeMacOs
		return nil
	case "none":
		*self = PortalOSTypeNone
		return nil
	case "Windows":
		*self = PortalOSTypeWindows
		return nil
	}
	if len(s) == 0 {
		*self = PortalOSTypeNone
		return nil
	}
	return errors.New("Unknown PortalOSType: " + s)
}

func (self *PortalOSType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "Android":
		*self = PortalOSTypeAndroid
		return nil
	case "iPhone OS":
		*self = PortalOSTypeIOS
		return nil
	case "Linux":
		*self = PortalOSTypeLinux
		return nil
	case "Mac Os":
		*self = PortalOSTypeMacOs
		return nil
	case "none":
		*self = PortalOSTypeNone
		return nil
	case "Windows":
		*self = PortalOSTypeWindows
		return nil
	}
	if len(s) == 0 {
		*self = PortalOSTypeNone
		return nil
	}
	return errors.New("Unknown PortalOSType: " + s)
}

type PortalPaymentSystemType string

const PortalPaymentSystemTypeExternal PortalPaymentSystemType = "ext"
const PortalPaymentSystemTypeExternalURL PortalPaymentSystemType = "ext_url"
const PortalPaymentSystemTypeGuide PortalPaymentSystemType = "guide"

func (self PortalPaymentSystemType) GetPtr() *PortalPaymentSystemType { var v = self; return &v }

func (self PortalPaymentSystemType) String() string {
	switch self {
	case PortalPaymentSystemTypeExternal:
		return "ext"
	case PortalPaymentSystemTypeExternalURL:
		return "ext_url"
	case PortalPaymentSystemTypeGuide:
		return "guide"
	}
	if len(self) == 0 {
		return "guide"
	}
	panic(errors.New("Invalid value of PortalPaymentSystemType: " + string(self)))
}

func (self *PortalPaymentSystemType) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalPaymentSystemTypeExternal:
		return json.Marshal("ext")
	case PortalPaymentSystemTypeExternalURL:
		return json.Marshal("ext_url")
	case PortalPaymentSystemTypeGuide:
		return json.Marshal("guide")
	}
	if len(*self) == 0 {
		return json.Marshal("guide")
	}
	return nil, errors.New("Invalid value of PortalPaymentSystemType: " + string(*self))
}

func (self *PortalPaymentSystemType) GetBSON() (interface{}, error) {
	switch *self {
	case PortalPaymentSystemTypeExternal:
		return "ext", nil
	case PortalPaymentSystemTypeExternalURL:
		return "ext_url", nil
	case PortalPaymentSystemTypeGuide:
		return "guide", nil
	}
	if len(*self) == 0 {
		return "guide", nil
	}
	return nil, errors.New("Invalid value of PortalPaymentSystemType: " + string(*self))
}

func (self *PortalPaymentSystemType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ext":
		*self = PortalPaymentSystemTypeExternal
		return nil
	case "ext_url":
		*self = PortalPaymentSystemTypeExternalURL
		return nil
	case "guide":
		*self = PortalPaymentSystemTypeGuide
		return nil
	}
	if len(s) == 0 {
		*self = PortalPaymentSystemTypeGuide
		return nil
	}
	return errors.New("Unknown PortalPaymentSystemType: " + s)
}

func (self *PortalPaymentSystemType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ext":
		*self = PortalPaymentSystemTypeExternal
		return nil
	case "ext_url":
		*self = PortalPaymentSystemTypeExternalURL
		return nil
	case "guide":
		*self = PortalPaymentSystemTypeGuide
		return nil
	}
	if len(s) == 0 {
		*self = PortalPaymentSystemTypeGuide
		return nil
	}
	return errors.New("Unknown PortalPaymentSystemType: " + s)
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

type PortalSignType string

const PortalSignTypeEqual PortalSignType = "equal"
const PortalSignTypeGreater PortalSignType = "greater"
const PortalSignTypeLess PortalSignType = "less"

func (self PortalSignType) GetPtr() *PortalSignType { var v = self; return &v }

func (self PortalSignType) String() string {
	switch self {
	case PortalSignTypeEqual:
		return "equal"
	case PortalSignTypeGreater:
		return "greater"
	case PortalSignTypeLess:
		return "less"
	}
	if len(self) == 0 {
		return "greater"
	}
	panic(errors.New("Invalid value of PortalSignType: " + string(self)))
}

func (self *PortalSignType) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalSignTypeEqual:
		return json.Marshal("equal")
	case PortalSignTypeGreater:
		return json.Marshal("greater")
	case PortalSignTypeLess:
		return json.Marshal("less")
	}
	if len(*self) == 0 {
		return json.Marshal("greater")
	}
	return nil, errors.New("Invalid value of PortalSignType: " + string(*self))
}

func (self *PortalSignType) GetBSON() (interface{}, error) {
	switch *self {
	case PortalSignTypeEqual:
		return "equal", nil
	case PortalSignTypeGreater:
		return "greater", nil
	case PortalSignTypeLess:
		return "less", nil
	}
	if len(*self) == 0 {
		return "greater", nil
	}
	return nil, errors.New("Invalid value of PortalSignType: " + string(*self))
}

func (self *PortalSignType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "equal":
		*self = PortalSignTypeEqual
		return nil
	case "greater":
		*self = PortalSignTypeGreater
		return nil
	case "less":
		*self = PortalSignTypeLess
		return nil
	}
	if len(s) == 0 {
		*self = PortalSignTypeGreater
		return nil
	}
	return errors.New("Unknown PortalSignType: " + s)
}

func (self *PortalSignType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "equal":
		*self = PortalSignTypeEqual
		return nil
	case "greater":
		*self = PortalSignTypeGreater
		return nil
	case "less":
		*self = PortalSignTypeLess
		return nil
	}
	if len(s) == 0 {
		*self = PortalSignTypeGreater
		return nil
	}
	return errors.New("Unknown PortalSignType: " + s)
}

type PortalUserState string

const PortalUserStateAdvertise PortalUserState = "advertise"
const PortalUserStateAuthenticate PortalUserState = "authenticate"
const PortalUserStateAuthorize PortalUserState = "authorize"
const PortalUserStateNew PortalUserState = "new"
const PortalUserStatePass PortalUserState = "pass"

func (self PortalUserState) GetPtr() *PortalUserState { var v = self; return &v }

func (self PortalUserState) String() string {
	switch self {
	case PortalUserStateAdvertise:
		return "advertise"
	case PortalUserStateAuthenticate:
		return "authenticate"
	case PortalUserStateAuthorize:
		return "authorize"
	case PortalUserStateNew:
		return "new"
	case PortalUserStatePass:
		return "pass"
	}
	if len(self) == 0 {
		return "new"
	}
	panic(errors.New("Invalid value of PortalUserState: " + string(self)))
}

func (self *PortalUserState) MarshalJSON() ([]byte, error) {
	switch *self {
	case PortalUserStateAdvertise:
		return json.Marshal("advertise")
	case PortalUserStateAuthenticate:
		return json.Marshal("authenticate")
	case PortalUserStateAuthorize:
		return json.Marshal("authorize")
	case PortalUserStateNew:
		return json.Marshal("new")
	case PortalUserStatePass:
		return json.Marshal("pass")
	}
	if len(*self) == 0 {
		return json.Marshal("new")
	}
	return nil, errors.New("Invalid value of PortalUserState: " + string(*self))
}

func (self *PortalUserState) GetBSON() (interface{}, error) {
	switch *self {
	case PortalUserStateAdvertise:
		return "advertise", nil
	case PortalUserStateAuthenticate:
		return "authenticate", nil
	case PortalUserStateAuthorize:
		return "authorize", nil
	case PortalUserStateNew:
		return "new", nil
	case PortalUserStatePass:
		return "pass", nil
	}
	if len(*self) == 0 {
		return "new", nil
	}
	return nil, errors.New("Invalid value of PortalUserState: " + string(*self))
}

func (self *PortalUserState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "advertise":
		*self = PortalUserStateAdvertise
		return nil
	case "authenticate":
		*self = PortalUserStateAuthenticate
		return nil
	case "authorize":
		*self = PortalUserStateAuthorize
		return nil
	case "new":
		*self = PortalUserStateNew
		return nil
	case "pass":
		*self = PortalUserStatePass
		return nil
	}
	if len(s) == 0 {
		*self = PortalUserStateNew
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
	case "advertise":
		*self = PortalUserStateAdvertise
		return nil
	case "authenticate":
		*self = PortalUserStateAuthenticate
		return nil
	case "authorize":
		*self = PortalUserStateAuthorize
		return nil
	case "new":
		*self = PortalUserStateNew
		return nil
	case "pass":
		*self = PortalUserStatePass
		return nil
	}
	if len(s) == 0 {
		*self = PortalUserStateNew
		return nil
	}
	return errors.New("Unknown PortalUserState: " + s)
}
