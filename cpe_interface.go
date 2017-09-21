package libwimark

import (
	"encoding/json"
)

// CPE responce for scripts request
type CPEAgentResponse struct {
	Result json.RawMessage `json:"result"`
	Status struct {
		Type CPEAgentStatusType `json:"type"`
		Data string             `json:"data"`
	} `json:"error"`
}

// CPE STATUS request
type CPEResult struct {
	Type      SystemEventType  `json:"type"`
	Id        UUID             `json:"subject_id"`
	Level     SystemEventLevel `json:"level"`
	Timestamp int64            `json:"timestamp"`
	Data      json.RawMessage  `json:"data"`
}
