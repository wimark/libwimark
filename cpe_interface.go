package libwimark

import (
	"encoding/json"
)

type CPEAgentError struct {
	Type CPEAgentStatusType `json:"type"`
	Data string             `json:"data"`
}

type CPEAgentResponse struct {
	Result json.RawMessage `json:"result"`
	Status CPEAgentError   `json:"error"`
}

type ConfigResponceCpeMap map[UUID]struct {
	Status ConfigurationStatus `json:"status"`
	Errors []string            `json:"errors,omitempty"`
}

type ConfigResponce struct {
	Status ConfigurationStatus  `json:"status"`
	Errors []ModelError         `json:"errors,omitempty"`
	Data   ConfigResponceCpeMap `json:"data,omitempty"`
}
