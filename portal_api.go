package libwimark

import (
	"strings"
	"time"
)

const (

	// user clients
	COLL_PORTAL_COLLECTION = "portal_client_sessions"

	// user authentications
	COLL_PORTAL_AUTHENTICATIONS = "portal_client_authentications"

	// webpage of portal settings
	COLL_PAGE_PROFILES = "portal_pages"
	// portal profile settings
	COLL_PORTAL_PROFILES = "portal_profiles"

	// available authentications
	COLL_PORTAL_AUTHEN_CFG = "portal_authentications"
	// available authorizations
	COLL_PORTAL_AUTH_CFG = "portal_authorizations"

	// portal advertisements
	COLL_PORTAL_ADS       = "portal_ads"
	COLL_PORTAL_ADS_STATS = "portal_ads_stats"

	// futher accounts
	COLL_PORTAL_USER_ACCOUNTS = "portal_user_accounts"

	// futher vouchers
	COLL_PORTAL_VOUCHERS = "portal_vouchers"
)

// Platform part
// --------

// RedirectRequestObject struct for request payload from external captive portal
type RedirectRequestObject struct {
	// Basic and Needed client data
	MAC  string `json:"mac" bson:"mac" form:"mac" query:"mac" validate:"required,mac"`
	CPE  string `json:"cpe_id" bson:"cpe_id" form:"cpe_id" query:"cpe_id" validate:"uuid"`
	WLAN string `json:"wlan_id" bson:"wlan_id" form:"wlan_id" query:"wlan_id" validate:"uuid"`

	// Username will be use in RADIUS Acct
	Username string `json:"username" bson:"username" form:"username" query:"username" validate:"-"`
	Password string `json:"password" bson:"password" form:"password" query:"password" validate:"-"`

	// passing NAS-ID to use in next RADIUS Acct
	NasId string `json:"nas-id,omitempty" bson:"nas-id" form:"nas-id" query:"nas-id" validate:"-"`

	// two way to set session-timeout
	Timeout  int64 `json:"session-timeout,omitempty" bson:"session-timeout" form:"session-timeout" query:"session-timeout" validate:"-"`
	Timeout2 int64 `json:"wimark-session-timeout,omitempty" bson:"wimark-session-timeout" validate:"-"`

	// ACL group -- not using now
	Group string `json:"wimark-client-group,omitempty" bson:"wimark-client-group" validate:"-"`

	// redirect every connect
	AlwaysRedirect bool `json:"always_redirect"`

	// Auth additional data
	UserAgent string `json:"useragent"`
}

// RedirectClientSession struct for store redirect session on platform
type RedirectClientSession struct {
	ID string `json:"id" bson:"_id"`

	Session   string `json:"session_id" bson:"session_id"`
	MAC       string `json:"mac" bson:"mac"`
	WLAN      string `json:"wlan_id" bson:"wlan_id"`
	CPE       string `json:"cpe_id" bson:"cpe_id"`
	Redirect  string `json:"redirect_id" bson:"redirect_id"`
	Radio     string `json:"radio_id" bson:"radio_id"`
	UserAgent string `json:"useragent" bson:"useragent"`

	AcctStart int64 `json:"acct_start" bson:"acct_start"`
	AuthStart int64 `json:"auth_start" bson:"auth_start"`
	AuthStop  int64 `json:"auth_stop" bson:"auth_stop"`
	Timeout   int64 `json:"timeout" bson:"timeout"`

	// data from acct start
	UserName         string `json:"User-Name"`
	CallingStationId string `json:"Calling-Station-Id"`
	CalledStationId  string `json:"Called-Station-Id"`
	NasIdentifier    string `json:"NAS-Identifier"`
	NasIPAddress     string `json:"NAS-IP-Address"`
	FramedIPAddress  string `json:"Framed-IP-Address"`
	UserQosGroup     string `json:"Qos-Group"`

	// data for interim update and acct stop
	AcctSessionTime     int `json:"Acct-Session-Time"`
	AcctInputGigawords  int `json:"Acct-Input-Gigawords"`
	AcctOutputGigawords int `json:"Acct-Output-Gigawords"`
	AcctInputOctets     int `json:"Acct-Input-Octets"`
	AcctOutputOctets    int `json:"Acct-Output-Octets"`
	AcctInputPackets    int `json:"Acct-Input-Packets"`
	AcctOutputPackets   int `json:"Acct-Output-Packets"`
}

// struct for update acct of redirect session on platform
type RedirectClientSessionAcct struct {
	AcctSessionTime     int `json:"Acct-Session-Time"`
	AcctInputGigawords  int `json:"Acct-Input-Gigawords"`
	AcctOutputGigawords int `json:"Acct-Output-Gigawords"`
	AcctInputOctets     int `json:"Acct-Input-Octets"`
	AcctOutputOctets    int `json:"Acct-Output-Octets"`
	AcctInputPackets    int `json:"Acct-Input-Packets"`
	AcctOutputPackets   int `json:"Acct-Output-Packets"`
}

type HTTPResponseObject struct {
	Status string `json:"status,omitempty"`
	Code   int    `json:"code"`

	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`

	State string `json:"state"`
}

// Portal part
// --------

// PortalRequestObject struct for request payload from web
type PortalRequestObject struct {
	// Needed client data
	MAC  string `json:"mac" bson:"mac" form:"mac" query:"mac" validate:"required,mac"`
	CPE  string `json:"cpe_id" bson:"cpe_id" form:"cpe_id" query:"cpe_id" validate:"required,uuid"`
	WLAN string `json:"wlan_id" bson:"wlan_id" form:"wlan_id" query:"wlan_id" validate:"required,uuid"`
	Ip   string `json:"client_ip" bson:"client_ip" form:"client_ip" query:"client_ip"`

	// client credentials
	Username string `json:"username,omitempty" bson:"username" form:"username" query:"username" validate:"-"`
	Password string `json:"password,omitempty" bson:"password" form:"password" query:"password" validate:"-"`

	// browser specific data
	UserAgent string `json:"useragent"  bson:"useragent" form:"useragent" query:"useragent" validate:"-"`

	// Address of platform CoA manager
	SwitchURL string `json:"switch_url" validate:"-"`

	// Remember period for user accounts
	Remember int64 `json:"remember"`

	// push aggrement
	PushAgreement bool `json:"push_agree"`

	// Type of choosen type
	Type string `json:"type"`

	Ad PortalAdStatRequest `json:"ad"`

	// for internal using
	Timeout int64 `json:"-" validate:"-"`
}

type PortalResponseObject struct {
	Status PortalResponseStatus `json:"status,omitempty"`
	Code   int                  `json:"code"`

	Description string `json:"description,omitempty"`
	// Data        interface{} `json:"data,omitempty"`

	State    PortalUserState `json:"state,omitempty"`
	Substate string          `json:"substate,omitempty"`
	Data     interface{}     `json:"data,omitempty"`
}

// struct for store every auth attempt
type PortalAuthObject struct {
	Timestamp int64 `json:"timestamp" bson:"timestamp"`

	CPE       string `json:"cpe_id" bson:"cpe_id"`
	Ip        string `json:"client_ip" bson:"client_ip"`
	Useragent string `json:"useragent" bson:"useragent"`

	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// struct for store portal client
type PortalClientSession struct {
	Id string `json:"id" bson:"_id"`

	// identity
	MAC  string `json:"mac" bson:"mac"`
	WLAN string `json:"wlan_id" bson:"wlan_id"`

	// just info data
	CPE string `json:"cpe" bson:"cpe"`
	IP  string `json:"ip" bson:"ip"`

	Create time.Time `json:"create" bson:"create"`

	// from create to end of internet + block
	CreateAt int64 `json:"create_at" bson:"create_at"`
	ExpireAt int64 `json:"expire_at" bson:"expire_at"`

	// real internet part
	StartAt int64 `json:"start_at" bson:"start_at"`
	StopAt  int64 `json:"stop_at" bson:"stop_at"`
	// BlockAt int64 `json:"block_at" bson:"block_at"`

	SessionConfig PortalSessionConfig `json:"session_config" bson:"session_config"`
	Auth          []PortalAuthObject  `json:"auth" bson:"auth"`

	State       PortalUserState `json:"state" bson:"state"`
	AuthenState PortalAuthenticationState
	AuthState   PortalAuthorizationState

	Authentication PortalAuthenticationData `json:"authentication" bson:"authentication"`
	Authorization  PortalAuthorizationData  `json:"authorization" bson:"authorization"`
	// Advertisement  PortalAdvertisementData  `json:"advertisement" bson:"advertisement"`

	// will be deprecated
	Status   string `json:"status" bson:"status"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// struct for store portal client
type PortalClientAuthentication struct {
	Id string `json:"id" bson:"_id"`

	// client identification
	MAC  string `json:"mac" bson:"mac"`
	WLAN string `json:"wlan_id" bson:"wlan_id"`

	// data with identity and authentication info
	Data PortalAuthenticationData `json:"data" bson:"data"`

	Create time.Time `json:"create" bson:"create"`

	// valid time
	CreateAt int64 `json:"create_at" bson:"create_at"`
	ExpireAt int64 `json:"expire_at" bson:"expire_at"`

	// // time to remember (from create to expire)
	// Remember int64 `json:"remember" bson:"remember"`

	// // client agree to take pushes
	// PushAgreement bool `json:"push_agree" bson:"push_agree"`
}

// portal condition
type PortalCondition struct {
	// MAC   map[string]bool `json:"mac" bson:"mac"`
	WLAN  []string `json:"wlan" bson:"wlan"`
	CPE   []string `json:"cpe" bson:"cpe"`
	NasId []string `json:"nas_id" bson:"nas_id"`
}

// func to check struct for empty
func (pc *PortalCondition) Empty() bool {
	return len(pc.WLAN) == 0 && len(pc.CPE) == 0 && len(pc.NasId) == 0
}

// struct for flexible session config
type PortalSessionConfig struct {

	// session timeout -- 30 min as example
	Timeout int64 `json:"timeout" bson:"timeout"`

	// timeout to remember
	AuthTimeout int64 `json:"auth_timeout" bson:"auth_timeout"`

	// block after using timeout for
	BlockTimeout int64 `json:"block_timeout" bson:"block_timeout"`

	// traffic limit
	DownloadLimit int `json:"download_limit" bson:"download_limit"`

	// max number
	MaxSessions int `json:"max_sessions" bson:"max_sessions"`
}

// PortalMSISDNConfig config of possible CC and NDC prefixes in MSISDN
type PortalMSISDNConfig struct {
	// example: {"992": ["90","92"...]}
	Prefix map[string][]string `json:"prefix" bson:"prefix"`
	PrMap  []string            `json:"-" bson:"map"`
}

func (p *PortalMSISDNConfig) Map() {
	p.PrMap = []string{}
	for k, v := range p.Prefix {
		if len(v) == 0 {
			p.PrMap = append(p.PrMap, k)
		} else {
			for _, vv := range v {
				p.PrMap = append(p.PrMap, k+vv)
			}
		}
	}
}

func (p PortalMSISDNConfig) Check(phone string) bool {
	if len(p.PrMap) > 0 && len(phone) > 0 {
		for _, v := range p.PrMap {
			if strings.HasPrefix(phone, v) {
				return true
			}
		}
		return false
	}
	return true
}

type PortalVoucher struct {
	Config PortalSessionConfig
}

type PortalAuthenticationData struct {
	// ConfigId string `json:"config_id" bson:"config_id"`

	Type PortalAuthenticationType `json:"type" bson:"type"`

	Identity string `json:"identity" bson:"identity"`
	Password string `json:"password" bson:"password"`

	Remember      int64 `json:"remember" bson:"remember"`
	PushAgreement bool  `json:"push_agree" bson:"push_agree"`
}

// PortalAuthenticationConfig config of possible authentications
type PortalAuthenticationConfig struct {
	Id string `json:"id" bson:"_id"`

	Type PortalAuthenticationType `json:"type" bson:"type"`

	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	Remember  int64 `json:"remember" bson:"remember"`
	OPTLength int   `json:"otp_length" bson:"otp_length"`

	PassGateway string `json:"-" bson:"pass_gateway"`
}

type PortalAuthorizationData struct {
	ConfigId string `json:"config_id" bson:"config_id"`

	Type  PortalAuthorizationType  `json:"type" bson:"type"`
	State PortalAuthorizationState `json:"state" bson:"state"`

	Config      PortalSessionConfig `json:"config" bson:"config"`
	Ads         []PortalAd          `json:"ads" bson:"ads"`
	AdsToWatch  int                 `json:"ads_to_watch" bson:"ads_to_watch"`
	RedirectURL string              `json:"redirect_url" bson:"redirect_url"`
}

type PortalAuthorizationConfig struct {
	Id string `json:"id" bson:"_id"`

	Type        PortalAuthorizationType `json:"type" bson:"type"`
	Name        string                  `json:"name" bson:"name"`
	Description string                  `json:"description" bson:"description"`

	// pass gateway needed if external authorization
	PassGateway string `json:"-" bson:"pass_gateway"`

	// Vouchers    []string `json:"vouchers,omitempty" bson:"vouchers"`

	// advertisements
	Ads []PortalAd `json:"ads" bson:"ads"`

	// session configuration
	Config PortalSessionConfig `json:"config" bson:"config"`

	// landing page redirect URL
	RedirectURL string `json:"redirect_url" bson:"redirect_url"`
}

// PortalProfile to link portal and it's config
type PortalProfile struct {
	Id string `json:"id" bson:"_id"`

	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	// condition to check
	Condition PortalCondition `json:"condition" bson:"condition"`

	// flow of client portal
	// States []PortalUserState `json:"states" bson:"states"`

	// authentication types
	Authentications []PortalAuthenticationConfig `json:"authentications" bson:"authentications"`

	// available MSISDN prefixes --
	MSISDNConfig PortalMSISDNConfig `json:"msisdn_config" bson:"msisdn_config'`

	// authorization types
	Authorizations []PortalAuthorizationConfig `json:"authorizations" bson:"authorizations"`

	// found for whitelist
	AccessList map[string]bool `json:"access_list" bson:"access_list"`
	BlackList  map[string]bool `json:"black_list" bson:"black_list"`

	// default session configuration (timeout and block timeout)
	SessionConfig PortalSessionConfig `json:"session_config" bson:"session_config"`
}

func (p *PortalProfile) NextState(state PortalUserState) (PortalUserState, []string) {

	var retState = PortalUserStatePass
	var possible = []string{}
	switch state {
	case PortalUserStateNew:
		if len(p.Authentications) == 0 && p.Authentications[0].Type == PortalAuthenticationTypeNone {
			if len(p.Authorizations) == 0 && p.Authorizations[0].Type == PortalAuthorizationTypeNone {
			} else {
				retState = PortalUserStateAuthorize
				for _, v := range p.Authorizations {
					possible = append(possible, v.Type.String())
				}
			}
		} else {
			retState = PortalUserStateAuthenticate
			for _, v := range p.Authentications {
				possible = append(possible, v.Type.String())
			}
		}
	case PortalUserStateAuthenticate:
		if len(p.Authorizations) == 0 && p.Authorizations[0].Type == PortalAuthorizationTypeNone {

		} else {
			retState = PortalUserStateAuthorize
			for _, v := range p.Authorizations {
				possible = append(possible, v.Type.String())
			}
		}
	}
	return retState, possible
}

// PortalPageProfile provide page information
type PortalPageProfile struct {
	Id string `json:"id" bson:"_id"`

	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	// URL generated id
	IdURL string `json:"url_id" bson:"url_id"`

	// title of webpage
	Title string `json:"title" bson:"title"`

	// interface features
	Interface struct {
		Logo          string `json:"logo" bson:"-"`
		LogoFooter    string `json:"logo_footer" bson:"-"`
		Background    string `json:"background" bson:"-"`
		LogoURL       string `json:"logo_url" bson:"logo_url"`
		LogoFooterURL string `json:"logo_footer_url" bson:"logo_footer_url"`
		BackgroundURL string `json:"background_url" bson:"background_url"`
		ButtonColor   string `json:"button_color" bson:"button_color"`
		Color         struct {
			Main  string `json:"main" bson:"main"`
			Light string `json:"light" bson:"light"`
			Dark  string `json:"dark" bson:"dark"`
			Error string `json:"error" bson:"error"`
		} `json:"color" bson:"color"`
	} `json:"interface" bson:"interface"`

	// locales
	Locales map[string]string `json:"locales" bson:"locales"`

	// service aggrement per locale
	Agreements map[string]string `json:"agreements" bson:"agreements"`
}

type PortalAdData struct {
	Type PortalAdvertisementType `json:"type" bson:"type"`

	URL  string `json:"url" bson:"url"`
	File string `json:"file,omitempty" bson:"-"`

	Question     string   `json:"question" bson:"question"`
	PollVariants []string `json:"poll_variants" bson:"poll_variants"`

	Skip         bool  `json:"skip" bson:"skip"`
	Duration     int64 `json:"duration" bson:"duration"`
	SkipDuration int64 `json:"skip_after" bson:"skip_after"`
}

// PortalAd object for ihot
type PortalAd struct {
	Id string `json:"id" bson:"_id"`

	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	Data PortalAdData `json:"data" bson:"data"`
}

type PortalAdStatRequest struct {
	Id          string `json:"id"`
	Duration    int64  `json:"duration"`
	Skipped     bool   `json:"skipped"`
	PollVariant string `json:"poll_variant"`
}

type PortalAdStatInc struct {
	Counter        int   `json:"counter" bson:"counter"`
	CounterSkipped int   `json:"counter_skipped" bson:"counter_skipped"`
	Duration       int64 `json:"duration" bson:"duration"`
	// PollCounter    int `json:"poll_counter" bson:"poll_counter."`
}

// PortalAdStat
type PortalAdStat struct {
	// same ID as portalad
	Id string `json:"id" bson:"_id"`

	Data PortalAdData `json:"data" bson:"data"`

	Counter        int            `json:"counter" bson:"counter"`
	CounterSkipped int            `json:"counter_skipped" bson:"counter_skipped"`
	Duration       int64          `json:"duration" bson:"duration"`
	PollCounter    map[string]int `json:"poll_counter" bson:"poll_counter"`
}

// type PortalAdConfig struct {
// 	Id string `json:"id" bson:"_id"`

// 	Name        string `json:"name" bson:"name"`
// 	Description string `json:"description" bson:"description"`

// 	Ads []PortalAdData `json:"ads" bson:"ads"`
// }

type PortalClientAccount struct {
	Id string `json:"id" bson:"_id"`

	// user common data
	MAC  string `json:"mac" bson:"mac" form:"mac" query:"mac" validate:"required,mac"`
	WLAN string `json:"wlan_id" bson:"wlan_id" form:"wlan_id" query:"wlan_id" validate:"required,uuid"`

	// identity for PHONE
	Identity string `json:"identity" bson:"identity"`

	CreateAt time.Time `json:"create_at" bson:"create_at"`

	// virtual amount
	Amount int `json:"amount" bson:"amount"`
}
