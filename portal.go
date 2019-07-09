package libwimark

import (
	tc "bitbucket.org/wimarksystems/libwimark/libtc"
)

type GuestControlSettings struct {
	CaptiveRedirect UUID   `json:"captive_redirect" bson:"captive_redirect"`
	MACAuth         []UUID `json:"mac_radius_auth_servers" bson:"mac_radius_auth_servers"`
}

type DnsAddress struct {
	Ip         string `json:"ip" bson:"ip"`
	DomainName string `json:"domain_name" bson:"domain_name"`
}

// struct for manage of cpe http-redirect system
type CaptiveRedirect struct {
	Name         string       `json:"name" bson:"name"`
	RedirectURL  string       `json:"redirect_url" bson:"redirect_url"`
	MACWhiteList []string     `json:"mac_list" bson:"mac_list"`
	URLWhiteList []DnsAddress `json:"url_list" bson:"url_list"`

	DefaultProfile string                  `json:"default_profile" bson:"default_profile"`
	Profiles       map[string]tc.UserClass `json:"user_profiles" bson:"user_profiles"`

	NoMasquerade bool `json:"no_masquerade" bson:"no_masquerade"`
}
