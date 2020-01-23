package libwimark

import (
	"sort"
	"strings"
	"time"
)

// PortalAuthObject is struct for store every auth attempt
// (Deprecated)
type PortalAuthObject struct {
	Timestamp int64 `json:"timestamp" bson:"timestamp"`

	CPE       string `json:"cpe_id" bson:"cpe_id"`
	Ip        string `json:"client_ip" bson:"client_ip"`
	Useragent string `json:"useragent" bson:"useragent"`

	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// PortalClientSession is struct for store portal client
type PortalClientSession struct {
	Id string `json:"id" bson:"_id"`

	// identity
	MAC  string `json:"mac" bson:"mac"`
	WLAN string `json:"wlan_id" bson:"wlan_id"`

	// just info data
	CPE     string `json:"cpe" bson:"cpe"`
	IP      string `json:"ip" bson:"ip"`
	Profile string `json:"profile" bson:"profile"`

	// creation time (index needs)
	Create time.Time `json:"create" bson:"create"`

	// from create to end of internet + block
	CreateAt int64 `json:"create_at" bson:"create_at"`
	ExpireAt int64 `json:"expire_at" bson:"expire_at"`

	// real internet part
	StartAt  int64 `json:"start_at" bson:"start_at"`
	StopAt   int64 `json:"stop_at" bson:"stop_at"`
	Duration int64 `json:"duration" bson:"duration"`

	SessionConfig PortalSessionConfig `json:"session_config" bson:"session_config"`
	Auth          []PortalAuthObject  `json:"auth" bson:"auth"`

	State       PortalUserState `json:"state" bson:"state"`
	AuthenState PortalAuthenticationState
	AuthState   PortalAuthorizationState

	Authentication PortalAuthenticationData `json:"authentication" bson:"authentication"`
	Authorization  PortalAuthorizationData  `json:"authorization" bson:"authorization"`

	// will be deprecated
	Status   string `json:"status" bson:"status"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// PortalClientAuthentication struct for store portal client
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
}

// PortalCondition struct for apply profile condition
type PortalCondition struct {
	// MAC   map[string]bool `json:"mac" bson:"mac"`
	WLAN  []string `json:"wlan" bson:"wlan"`
	CPE   []string `json:"cpe" bson:"cpe"`
	NasID []string `json:"nas_id" bson:"nas_id"`
}

// Empty func to check struct for empty
func (pc *PortalCondition) Empty() bool {
	return len(pc.WLAN) == 0 && len(pc.CPE) == 0 && len(pc.NasID) == 0
}

// PortalSessionConfig struct for flexible session config
type PortalSessionConfig struct {

	// session timeout -- 30 min as example
	Timeout int64 `json:"timeout" bson:"timeout"`

	// timeout to remember
	AuthTimeout int64 `json:"auth_timeout" bson:"auth_timeout"`

	// block after using timeout for (example: 3600 is for seconds online)
	BlockAfterTimeout int64 `json:"block_after" bson:"block_after"`

	// expiration of block after (example: 7200 is for seconds to nullify block_after)
	BlockExpireTimeout int64 `json:"block_expire" bson:"block_expire"`

	// traffic limit
	DownloadLimit int `json:"download_limit" bson:"download_limit"`

	// max number
	MaxSessions int `json:"max_sessions" bson:"max_sessions"`

	// Deprecated field for block timeout
	BlockTimeout int64 `json:"block_timeout" bson:"block_timeout"`
}

// PortalMSISDNConfig config of possible CC and NDC prefixes in MSISDN
type PortalMSISDNConfig struct {
	// example: {"992": ["90","92"...]}
	Prefix map[string][]string `json:"prefix" bson:"prefix"`
	PrMap  []string            `json:"-" bson:"map"`
}

// Map (PortalMSISDNConfig) func to Map DEF / CC codes
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

type PortalAuthenticationData struct {
	Type PortalAuthenticationType `json:"type" bson:"type"`

	Identity string `json:"identity" bson:"identity"`
	Password string `json:"password" bson:"password"`

	Remember      int64 `json:"remember" bson:"remember"`
	PushAgreement bool  `json:"push_agree" bson:"push_agree"`

	// additional data if needed (for example full ESIA string)
	Data string `json:"data" bson:"data"`
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

	Validate struct {
		Regex    string `json:"regex" bson:"regex"`
		ErrorMsg string `json:"error" bson:"error"`
	} `json:"validate" bson:"validate"`
}

type PortalAuthorizationData struct {
	ConfigId string `json:"config_id" bson:"config_id"`

	Type  PortalAuthorizationType  `json:"type" bson:"type"`
	State PortalAuthorizationState `json:"state" bson:"state"`

	Config PortalSessionConfig `json:"config" bson:"config"`

	Ads        []PortalAd `json:"ads" bson:"ads"`
	AdsToWatch int        `json:"ads_to_watch" bson:"ads_to_watch"`

	RedirectURL string `json:"redirect_url" bson:"redirect_url"`
}

type PortalAuthorizationConfig struct {
	Id string `json:"id" bson:"_id"`

	Type        PortalAuthorizationType `json:"type" bson:"type"`
	Name        string                  `json:"name" bson:"name"`
	Description string                  `json:"description" bson:"description"`

	// pass gateway needed if external authorization
	PassGateway string `json:"-" bson:"pass_gateway"`

	// advertisements for user
	Ads        []PortalAd `json:"ads" bson:"ads"`
	AdsToWatch int        `json:"ads_to_watch" bson:"ads_to_watch"`

	// session configuration
	Config PortalSessionConfig `json:"config" bson:"config"`

	// landing page redirect URL
	RedirectURL string `json:"redirect_url" bson:"redirect_url"`
}

func (p *PortalAuthorizationConfig) SortAd() {
	sort.Slice(p.Ads, func(i, j int) bool {
		return p.Ads[i].Priority > p.Ads[j].Priority
	})
}

// PortalProfile to link portal and it's config
type PortalProfile struct {
	Id string `json:"id" bson:"_id"`

	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	// condition to check
	Condition PortalCondition `json:"condition" bson:"condition"`

	// authentication types
	Authentications []PortalAuthenticationConfig `json:"authentications" bson:"authentications"`

	// available MSISDN prefixes --
	MSISDNConfig PortalMSISDNConfig `json:"msisdn_config" bson:"msisdn_config"`

	// authorization types
	Authorizations []PortalAuthorizationConfig `json:"authorizations" bson:"authorizations"`

	// found for whitelist
	AccessList map[string]bool `json:"access_list" bson:"access_list"`
	BlackList  map[string]bool `json:"black_list" bson:"black_list"`

	// default session configuration (timeout and block timeout)
	SessionConfig PortalSessionConfig `json:"session_config" bson:"session_config"`

	// UTC diff (plus or minus from UTC time)
	UTCDiff int `json:"utc_diff"`
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

func (p *PortalProfile) SortAd() {
	for i := range p.Authorizations {
		p.Authorizations[i].SortAd()
	}
}

type pageColor struct {
	Main       string `json:"main" bson:"main"`
	Light      string `json:"light" bson:"light"`
	Dark       string `json:"dark" bson:"dark"`
	Error      string `json:"error" bson:"error"`
	Background string `json:"background" bson:"background"`
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

	// footer text
	Footer string `json:"footer" bson:"footer"`

	// interface features
	Interface struct {
		// Favicon    string `json:"favicon" bson:"-"`
		Logo       string `json:"logo" bson:"-"`
		LogoFooter string `json:"logo_footer" bson:"-"`
		Background string `json:"background" bson:"-"`

		// FaviconURL    string `json:"favicon_url" bson:"favicon_url"`
		LogoURL       string `json:"logo_url" bson:"logo_url"`
		LogoFooterURL string `json:"logo_footer_url" bson:"logo_footer_url"`
		BackgroundURL string `json:"background_url" bson:"background_url"`
		ButtonColor   string `json:"button_color" bson:"button_color"`

		Color pageColor `json:"color" bson:"color"`
	} `json:"interface" bson:"interface"`

	// locales
	Locales map[string]string `json:"locales" bson:"locales"`

	// service aggrement per locale
	Agreements map[string]string `json:"agreements" bson:"agreements"`

	// service aggrement plain text
	Agreement string `json:"agreement" bson:"agreement"`
}

type PortalAdData struct {
	Type PortalAdvertisementType `json:"type" bson:"type"`

	// URL of iframe or URL to file on portal
	URL  string `json:"url" bson:"url"`
	File string `json:"file,omitempty" bson:"-"`

	// question and poll variants if type == poll
	Question     string   `json:"question" bson:"question"`
	QuestionDesc string   `json:"question_desc" bson:"question_desc"`
	PollVariants []string `json:"poll_variants" bson:"poll_variants"`
	SelfVariant  bool     `json:"self_variant" bson:"self_variant"`

	Skip         bool  `json:"skip" bson:"skip"`
	Duration     int64 `json:"duration" bson:"duration"`
	SkipDuration int64 `json:"skip_after" bson:"skip_after"`

	// for redirect to client's site
	RedirectURL string `json:"redirect_url" bson:"redirect_url"`
	SkipURL     string `json:"skip_url" bson:"skip_url"`

	// title text
	Text string `json:"text" bson:"text"`

	// color options
	Color pageColor `json:"color" bson:"color"`
}

// PortalAd object for ihot
type PortalAd struct {
	Id string `json:"id" bson:"_id"`

	// common data
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	// more specific data to link with profiles and authorization
	Profile       string `json:"profile" bson:"profile"`
	Authorization string `json:"authorization" bson:"authorization"`

	// number beetween 0 and 100 - more is higher
	Priority int `json:"priority" bson:"priority"`

	// schedule of ads to work start-stop and number of views to show
	Schedule struct {
		Start     int64 `json:"start" bson:"start"`
		Stop      int64 `json:"stop" bson:"stop"`
		Views     int   `json:"views" bson:"views"`
		TimeOfDay struct {
			Morning   bool `json:"morning" bson:"morning"`
			Afternoon bool `json:"afternoon" bson:"afternoon"`
			Evening   bool `json:"evening" bson:"evening"`
			Night     bool `json:"night" bson:"night"`
		} `json:"time_of_day" bson:"time_of_day"`
	} `json:"schedule" bson:"schedule"`

	Data PortalAdData `json:"data" bson:"data"`
}

type PortalAdStatRequest struct {
	Id            string `json:"id"`
	Profile       string `json:"profile"`
	Authorization string `json:"authorization"`

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

	// more specific data to link with profiles and authorization
	Profile       string `json:"profile" bson:"profile"`
	Authorization string `json:"authorization" bson:"authorization"`

	Data PortalAdData `json:"data" bson:"data"`

	Counter        int            `json:"counter" bson:"counter"`
	CounterSkipped int            `json:"counter_skipped" bson:"counter_skipped"`
	Duration       int64          `json:"duration" bson:"duration"`
	PollCounter    map[string]int `json:"poll_counter" bson:"poll_counter"`
}
