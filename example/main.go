// +build linux windows
// +build 386 amd64

package main

import (
	"log"

	"github.com/msepp/go-astilectron-template/app/bootstrap"
)

// Various handles that are used globally
var gApp *bootstrap.App

func main() {

	// Init new application
	gApp = bootstrap.New(Asset, RestoreAsset, HandleGUIMessage)

	// Bootstrap to get things going
	if err := gApp.Bootstrap(); err != nil {
		log.Fatalln(err)
	}

	// Wait for app to exit
	gApp.Wait()
}
