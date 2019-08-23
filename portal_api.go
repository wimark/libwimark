package libwimark

import (
	"time"
)

const (
	COLL_PORTAL_COLLECTION      = "portal_client_sessions"
	COLL_PORTAL_AUTHENTICATIONS = "portal_client_authentications"
	COLL_PORTAL_PROFILES        = "portal_profiles"
	COLL_AD_PROFILES            = "ad_profiles"
	COLL_PAGE_PROFILES          = "page_profiles"
	COLL_PORTAL_UACCOUNTS       = "portal_user_accounts"
)

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

	// Type of choosen type
	Type string `json:"type"`

	// for internal using
	Timeout int64 `json:"-" validate:"-"`
}

type PortalResponseObject struct {
	Status PortalResponseStatus `json:"status,omitempty"`
	Code   int                  `json:"code"`

	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`

	State    PortalUserState `json:"state"`
	Substate string          `json:"substate,omitempty"`
	Types    map[string]bool `json:"types,omitempty"`
}

type HTTPResponseObject struct {
	Status string `json:"status,omitempty"`
	Code   int    `json:"code"`

	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`

	State string `json:"state"`
}

// struct for store redirect session on platform
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

	Create   time.Time `json:"create" bson:"create"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
	StartAt  int64     `json:"start_at" bson:"create_at"`
	StopAt   int64     `json:"stop_at" bson:"stop_at"`

	SessionConfig PortalSessionConfig `json:"session_config" bson:"session_config"`
	Auth          []PortalAuthObject  `json:"auth" bson:"auth"`

	// new states
	State          PortalUserState            `json:"state" bson:"state"`
	Authentication PortalClientAuthentication `json:"authentication" bson:"authontication"`
	Authorization  PortalAuthorizationData    `json:"authorization" bson:"authorization"`
	Advertisement  interface{}                `json:"advertisement" bson:"advertisement"`

	// will be deprecated
	Status   string `json:"status" bson:"status"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// struct for store portal client
type PortalClientAuthentication struct {
	Id string `json:"id" bson:"_id"`

	// identification
	MAC  string `json:"mac" bson:"mac"`
	WLAN string `json:"wlan_id" bson:"wlan_id"`

	State PortalAuthenticationState `json:"state" bson:"state"`

	Identity string `json:"identity" bson:"identity"`
	Password string `json:"password" bson:"password"`

	CreateAt time.Time `json:"create_at" bson:"create_at"`
	StartAt  int64     `json:"start_at" bson:"start_at"`
	StopAt   int64     `json:"stop_at" bson:"stop_at"`
	Remember int64     `json:"config" bson:"config"`

	// current cpe -- from auth
	// CPE string `json:"cpe" bson:"cpe"`
	// Ip string `json:"ip"`
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
	// // store in DB -- 24 hours as example
	// StoreTimeout    int64 `json:"store_timeout" bson:"store_timeout"`

	// one-time online timeout -- 30 min as example
	Timeout int64 `json:"timeout" bson:"timeout"`

	// timeout to resend
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
	CC  []string `json:"cc" bson:"cc"`
	NDC []string `json:"ndc" bson:"ndc"`

	CCNDC []string `json:"-" bson:"ccndc"`
}

func (p *PortalMSISDNConfig) Map() {
	p.CCNDC = make([]string, 0, len(p.CC)*len(p.NDC))
	for _, c := range p.CC {
		for _, n := range p.NDC {
			p.CCNDC = append(p.CCNDC, c+n)
		}
	}
}

// PortalAuthenticationConfig config of possible authentications
type PortalAuthenticationConfig struct {
	Type        PortalAuthenticationType `json:"type" bson:"type"`
	Name        string                   `json:"name" bson:"name"`
	PassGateway string                   `json:"pass_gateway" bson:"pass_gateway"`

	Remember int64 `json:"config" bson:"config"`
}

type PortalAuthenticationData struct {
	Type  PortalAuthenticationType  `json:"type" bson:"type"`
	State PortalAuthenticationState `json:"state" bson:"state"`

	Identity string `json:"identity" bson:"identity'`
	Password string `json:"password" bson:"password"`

	Remember int64 `json:"config" bson:"config"`
}

type PortalAuthorizationData struct {
	Type  PortalAuthorizationType  `json:"type" bson:"type"`
	State PortalAuthorizationState `json:"state" bson:"state"`

	Identity string `json:"identity" bson:"identity'`
	Password string `json:"password" bson:"password"`

	StartAt int64 `json:"start_at" bson:"start_at"`
}

type PortalAuthorizationConfig struct {
	Type        PortalAuthorizationType `json:"type" bson:"type"`
	Name        string                  `json:"name" bson:"name"`
	PassGateway string                  `json:"pass_gateway,omitempty" bson:"pass_gateway"`
	Vouchers    []string                `json:"vouchers,omitempty" bson:"vouchers"`

	Config PortalSessionConfig `json:"config" bson:"config"`
}

// PortalProfile to link portal and it's config
type PortalProfile struct {
	Id string `json:"id" bson:"_id"`

	// condition to check
	Condition PortalCondition `json:"condition" bson:"condition"`

	// flow
	States []PortalUserState `json:"states" bson:"states"`

	// authentication types
	Authentications map[string]PortalAuthenticationConfig `json:"authentications" bson:"authentications"`

	// authorization types
	Authorizations map[string]PortalAuthorizationConfig `json:"authorizations" bson:"authorizations"`

	// advertisement types

	// session configuration (timeout and block timeout)
	SessionConfig PortalSessionConfig `json:"session_config" bson:"session_config"`

	// maximum number of online sessions
	MaxNum int `json:"session_max_num" bson:"session_max_num"`

	// true for whitelist, false for blacklist
	AccessList map[string]bool `json:"access_list" bson:"access_list"`

	// available MSISDN prefixes --
	MSISDNPrefixes []string `json:"msisdn_prefixes" bson:"msisdn_prefixes"`

	//
	MSISDNConfig PortalMSISDNConfig `json:"msisdn_config" bson:"msisdn_config'`
}

func (p PortalProfile) NextState(state PortalUserState) PortalUserState {
	for i, v := range p.States {
		if state == v {
			if i < len(p.States) {
				return p.States[i+1]
			}
		}
	}
	return p.States[len(p.States)-1]
}

// PortalPageProfile provide page information
type PortalPageProfile struct {
	Id string `json:"id" bson:"_id"`

	IdPath string `json:"path_id" bson:"path_id"`

	// URL generated id
	IdURL string `json:"url_id" bson:"url_id"`

	// condition to check
	// Condition PortalCondition `json:"condition" bson:"condition"`

	// title of webpage
	Title string `json:"title" bson:"title"`

	// interface features
	Interface struct {
		Logo        string `json:"logo" bson:"logo"`
		Background  string `json:"background" bson:"background"`
		ButtonColor string `json:"button_color" bson:"button_color"`
		LogoFooter  string `json:"logo_footer" bson:"logo_footer"`
	} `json:"interface" bson:"interface"`
}

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
