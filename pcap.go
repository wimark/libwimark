package libwimark

type EthernetHeader struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Type        int    `json:"type"`
}

type IPHeader struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Protocol    int    `json:"protocol"`
}

type UDPHeader struct {
	SourcePort      int `json:"src_port"`
	DestinationPort int `json:"dst_port"`
}

type DHCPHeader struct {
	ClientAddr string `json:"client_addr"`
	YourAddr   string `json:"your_addr"`
	ServerAddr string `json:"server_addr"`
	RelayAddr  string `json:"relay_addr"`
	HWAddr     string `json:"hw_addr"`
}

type DHCPOptions struct {
	Type       string   `json:"type"`
	Subnet     string   `json:"subnet"`
	Routers    []string `json:"routers"`
	DNS        []string `json:"dns"`
	HostName   string   `json:"host_name"`
	DomainName string   `json:"domain_name"`
}

type DHCPPackage struct {
	Header      DHCPHeader  `json:"header"`
	Options     DHCPOptions `json:"options"`
	Fingerprint []int       `json:"fingerprint"`
}

type ClientAddr struct {
	Timestamp int64           `json:"timestamp"`
	Ethernet  *EthernetHeader `json:"ethernet"`
	IP        *IPHeader       `json:"ip"`
	UDP       *UDPHeader      `json:"udp"`
	DHCP      *DHCPPackage    `json:"dhcp"`
	Error     string          `json:"error"`
}
