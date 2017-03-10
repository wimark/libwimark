package libwimark

type ErrorCode int

const (
	ErrorUnknown = ErrorCode(iota)
	ErrorJson
	ErrorDB
	ErrorNotFound
	ErrorInvalidRequestModel
	ErrorInvalidUUID
)

type ModelError struct {
	Code        ErrorCode `json:"code"`
	Description string    `json:"description"`
	Model       *string   `json:"model"`
	Id          *UUID     `json:"uuid"`
}
