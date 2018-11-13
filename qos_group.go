package libwimark

const QOS_GROUP_COLLECTION = "qos_groups"

type QosGroup struct {
	Name   string    `json:"name" bson:"name"`
	TxRate int       `json:"tx_rate" bson:"tx_rate"`
	RxRate int       `json:"rx_rate" bson:"rx_rate"`
	ACL    []L2Chain `json:"acl" bson:"acl"`
}
