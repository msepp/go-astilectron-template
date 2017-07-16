// Package message implements message types for IPC between GUI and backend.
//
// All int conversions assume original datatype is float64 because that's the
// type used when receiving data from Electron (JavaScript Number is a float).
package message

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// Message describes an message sent by GUI to our Go process.
type Message struct {
	ID   string      `json:"id,omitempty"`   // message ID, set by sender and replayed by responder.
	Key  Key         `json:"key,omitempty"`  // message target, what is requested.
	Type Type        `json:"type"`           // type tells what kind of a message this is.
	Data interface{} `json:"data,omitempty"` // target data, will be parsed based on Target.
}

func (m Message) String() string {
	switch m.Type {
	case Error:
		return fmt.Sprintf("[%s:%s '%s' for ID %s]", m.Type, m.Key, m.Data, m.ID)

	case Event:
		return fmt.Sprintf("[%s:%s]", m.Type, m.Key)

	default:
		return fmt.Sprintf("[%s:%s for ID %s]", m.Type, m.Key, m.ID)
	}
}

// New returns an Message with given values
func New(id string, t Type, key Key, data interface{}) *Message {
	return &Message{
		ID:   id,
		Key:  key,
		Type: t,
		Data: data,
	}
}

// NewEvent returns an event message with given data.
func NewEvent(event Key, data interface{}) *Message {
	return New("", Event, event, data)
}

// NewResponse returns an response message with given data.
func NewResponse(id string, key Key, data interface{}) *Message {
	return New(id, Response, key, data)
}

// NewError returns an error message.
func NewError(id string, err string) *Message {
	return New(id, Error, ErrorOperationFailed, err)
}

// NewAlert returns an alert message with given data
func NewAlert(data interface{}) *Message {
	return New("", Alert, "", data)
}

// DataString returns the data as string if possible and if the value was a string.
func (m *Message) DataString() (v string, ok bool) {
	v, ok = m.Data.(string)
	return
}

// DataBool returns the data as bool if possible and if the value was a bool.
func (m *Message) DataBool() (v bool, ok bool) {
	v, ok = m.Data.(bool)
	return
}

// DataInt returns the data as int if possible and truth value if data conversion
// was possible.
func (m *Message) DataInt() (int, bool) {
	f, ok := m.DataFloat64()
	if !ok {
		return 0, ok
	}

	return int(f), true
}

// DataFloat64 returns the data as float64 if possible and truth value if data
// conversion worked
func (m *Message) DataFloat64() (v float64, ok bool) {
	v, ok = m.Data.(float64)
	return
}

// DataMap returns the data as DataMap if possible and truth value if data
// conversion worked
func (m *Message) DataMap() (DataMap, bool) {
	mi, ok := m.Data.(map[string]interface{})
	if !ok {
		return nil, false
	}

	return DataMap(mi), true
}

// Into attempts to coerce the message data into given interface value.
// Uses "json" tags for hinting.
func (m *Message) Into(tgt interface{}) error {
	d, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: tgt,
	})
	return d.Decode(m.Data)
}
