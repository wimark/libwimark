package libwimark

type L2Chain struct {
	Name      string            `json:"name" bson:"name"`
	Policy    FirewallPolicy    `json:"policy" bson:"policy"`
	Rules     []L2Rule          `json:"rules" bson:"rules"`
	Direction FirewallDirection `json:"direction" bson:"direction"`
}

type L2Rule struct {
	Protocol L3Protocol   `json:"protocol" bson:"protocol"`
	SrcMAC   []string     `json:"source" bson:"source"`
	DstMAC   []string     `json:"destination" bson:"destination"`
	SrcIP    []string     `json:"ip-source" bson:"ip-source"`
	DstIP    []string     `json:"ip-destination" bson:"ip-destination"`
	IPProto  L4Protocol   `json:"ip-protocol" bson:"ip-protocol"`
	SrcPort  []string     `json:"ip-source-port" bson:"ip-source-port"`
	DstPort  []string     `json:"ip-destination-port" bson:"ip-destination-port"`
	Jump     FirewallJump `json:"jump" bson:"jump"`
}
