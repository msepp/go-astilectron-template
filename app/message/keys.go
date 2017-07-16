package message

// Key is a message key, which tells what the message purpose is.
type Key string

// Error keys
const (
	ErrorOperationFailed = Key("failure")
)

// Request keys
const (
	RequestWindowClose    = Key("window.close")
	RequestWindowMinimize = Key("window.minimize")
	RequestAppVersions    = Key("get.versions")
)

// Event keys
const (
	EventBackendStatusChanged = Key("backend.status")
)
