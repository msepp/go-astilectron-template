package bootstrap

import (
	astilectron "github.com/asticode/go-astilectron"
	"github.com/msepp/go-astilectron-template/app/message"
)

// MessageHandlerFn defines an callback for incoming GUI messages
type MessageHandlerFn func(*message.Message) (interface{}, error)

// OnWindowMessage creates a handler for incoming EventNameWindowEventMessage
// events.
func (a *App) onWindowMessage() func(astilectron.Event) bool {
	return func(e astilectron.Event) (deleteListener bool) {
		var err error
		var res interface{}
		var req message.Message
		var resp *message.Message

		if err = e.Message.Unmarshal(&req); err != nil {
			resp = message.NewError(req.ID, err.Error())
			a.msgQueue <- *resp
			return
		}

		switch req.Key {
		// We pass these on for handling elsewhere. For one reason or another,
		// eg. w.Minimize() will never return here, maybe something to do with
		// the context we're in.
		case message.RequestWindowClose, message.RequestWindowMinimize:
			res = "ok"

		default:
			res, err = a.msgHandler(&req)
		}

		// Generate response Message, use Error type if erros were found
		if err != nil {
			resp = message.NewError(req.ID, err.Error())
		} else {
			resp = message.NewResponse(req.ID, req.Key, res)
		}

		// Queue response
		a.msgQueue <- *resp
		return
	}
}
