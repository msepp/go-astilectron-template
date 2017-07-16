package main

import (
	"fmt"

	"github.com/msepp/go-astilectron-template/app/bootstrap"
	"github.com/msepp/go-astilectron-template/app/message"
)

// HandleGUIMessage is called when we receive messages from the user interface.
func HandleGUIMessage(msg *message.Message) (interface{}, error) {
	switch msg.Key {
	case message.RequestAppVersions:
		return HandleGetAppVersions(msg)

	default:
		return nil, fmt.Errorf("Unrecognized GUI message key: %s", msg.Key)
	}
}

// HandleGetAppVersions returns versions used in the updater
func HandleGetAppVersions(msg *message.Message) (interface{}, error) {
	return map[string]string{
		"app":         bootstrap.Version(),
		"build":       bootstrap.Build(),
		"electron":    bootstrap.ElectronVersion(),
		"astilectron": bootstrap.AstilectronVersion(),
	}, nil
}
