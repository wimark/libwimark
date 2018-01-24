package libwimark

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
