package libwimark

type ModelError struct {
	Module         Module          `json:"module"`
	ModuleId       UUID            `json:"module_id"`
	Object         string          `json:"object,omitempty"`
	ObjectId       UUID            `json:"object_id,omitempty"`
	Type           WimarkErrorCode `json:"type"`
	Description    string          `json:"description"`
	Recommendation string          `json:"recommendation"`
	Data           interface{}     `json:"data,omitempty"`
}
