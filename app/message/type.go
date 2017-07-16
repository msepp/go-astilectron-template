package message

// Type tells the type of the message, if it's a response, error or an event
type Type string

// Message types
const (
	Response = "response"
	Error    = "error"
	Event    = "event"
	Alert    = "alert"
)
