package libwimark

type GuestControlSettings struct {
	CaptiveRedirect UUID `json:"captive_redirect" bson:"captive_redirect"`
}

// struct for manage of cpe http-redirect system
type CaptiveRedirect struct {
	Name        string   `json:"name" bson:"name"`
	RedirectURL string   `json:"redirect_url" bson:"redirect_url"`
	MACs        []string `json:"mac_list" bson:"mac_list"`
	URLs        []string `json:"url_list" bson:"url_list"`
	SwitchURL   string   `json:"switch_url" bson:"switch_url"`
}

type PortalAuthRADIUS struct {
	RadiusID string
}

type PortalAuthSMS struct {
	Gateway  string
	Token    string
	Username string
	Password string
}

type PortalAuthOAuth2 struct {
	ClientID    string
	Scope       string
	CallBackURI string
	RedirectURI string
	RESTURI     string
}

type PortalAuthExternal struct {
	URL string

	NeededFields []string
}

// struct for portal description
type Portal struct {
	Name        string             `json:"name" bson:"name"`
	Desc        string             `json:"desc" bson:"desc"`
	Auth        []PortalAuthObject `json:"auth" bson:"auth"`
	RedirectURL string             `json:"redirect_url" bson:"redirect_url"`

	PortalHostname string `json:"hostname" bson:"hostname"`
	PortalHTTPS    bool   `json:"https" bson:"https"`
	PortlaSecure   bool   `json:"secure" bson:"secure"`
	PortalURL      string `json:"url" bson:"url"`
}
