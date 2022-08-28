package libwimark

import "fmt"

// GuestControlSettings for WLAN/Wired Guest Control settings
type GuestControlSettings struct {
	CaptiveRedirect UUID   `json:"captive_redirect" bson:"captive_redirect"`
	MACAuth         []UUID `json:"mac_radius_auth_servers" bson:"mac_radius_auth_servers"`
}

// DnsAddress for link IP with domain name
type DnsAddress struct {
	Ip         string `json:"ip" bson:"ip"`
	DomainName string `json:"domain_name" bson:"domain_name"`
}

// CaptiveRedirect struct for manage of Client redirect system
type CaptiveRedirect struct {
	Name string `json:"name" bson:"name"`

	// URL to redirect (Portal Web part)
	RedirectURL string `json:"redirect_url" bson:"redirect_url"`

	MACWhiteList []string `json:"mac_list" bson:"mac_list"`

	URLWhiteList []DnsAddress `json:"url_list" bson:"url_list"`
	PreAuthList  []string     `json:"preauth_list" bson:"preauth_list"`

	NoMasquerade bool `json:"no_masquerade" bson:"no_masquerade"`
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
	IP        string `json:"ip" bson:"ip"`
	CPEIP     string `json:"cpe_ip" bson:"cpe_ip"`

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

	Transferable bool `json:"transferable"`
}

func (session *RedirectClientSession) String() string {
	return fmt.Sprintf("packet (username, mac, session, framed, nas, calling): '%s' '%s' '%s' '%s' '%s' '%s'",
		session.UserName, session.MAC, session.Session, session.FramedIPAddress, session.NasIPAddress, session.CallingStationId)
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
