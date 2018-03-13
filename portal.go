package libwimark

type GuestControlSettings struct {
	CaptiveRedirect UUID `json:"captive_redirect" bson:"captive_redirect"`
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
}
