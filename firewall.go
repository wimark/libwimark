package libwimark

type FireWallSettings struct {
	L2Chain        UUID `json:"l2_chain" bson:"l2_chain"`
	WanAccessBlock bool `json:"wan_access_block" bson:"wan_access_block"`
}

type L2Chain struct {
	Name      string            `json:"name" bson:"name"`
	Policy    FirewallPolicy    `json:"policy" bson:"policy"`
	Rules     []L2Rule          `json:"rules" bson:"rules"`
	Direction FirewallDirection `json:"direction" bson:"direction"`
}

type L2Rule struct {
	Protocol string   `json:"protocol" bson:"protocol"`
	SrcMAC   []string `json:"source" bson:"source"`
	DstMAC   []string `json:"destination" bson:"destination"`
	SrcIP    []string `json:"ip_source" bson:"ip_source"`
	DstIP    []string `json:"ip_destination" bson:"ip_destination"`
	IPProto  string   `json:"ip_protocol" bson:"ip_protocol"`
	SrcPort  []string `json:"ip_source_port" bson:"ip_source_port"`
	DstPort  []string `json:"ip_destination_port" bson:"ip_destination_port"`
	Jump     string   `json:"jump" bson:"jump"`
}
