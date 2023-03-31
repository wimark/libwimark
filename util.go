package libwimark

import "github.com/google/uuid"

func NewUUID() string {
	return uuid.New().String()
}

func IsWPAEnterprise(secType SecurityType) bool {
	switch secType {
	case SecurityTypeWPAEnterprise,
		SecurityTypeWPA2Enterprise,
		SecurityTypeWPA3Enterprise,
		SecurityTypeWPA23Enterprise:
		return true
	default:
		return false
	}
}
