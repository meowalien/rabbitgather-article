package handler

type StandardResponseTypes string

const (
	ErrorResponse StandardResponseTypes = "err"
)

type StandardJsonResponse struct {
	Type StandardResponseTypes `json:"type"`
	Msg  string                `json:"msg,omitempty"`
	Obj  interface{}           `json:"obj,omitempty"`
}
