package message

// Backend status strings
const (
	BackendIdle = "idle"
)

// BackendStatus indicates status of backend
type BackendStatus struct {
	// Status is the backend status
	Status string `json:"status"`
}

// NewBackendStatus generates an new backend status event message
func NewBackendStatus(status string) *Message {
	return New("", Event, EventBackendStatusChanged, &BackendStatus{
		Status: status,
	})
}
