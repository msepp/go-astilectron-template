package bootstrap

import (
	"os"
	"sync"

	astilectron "github.com/asticode/go-astilectron"
	"github.com/msepp/go-astilectron-template/app/message"
)

// App defines an GUI application and the required component handles
type App struct {
	Window       *astilectron.Window      // Main application window
	Renderer     *astilectron.Astilectron // Renderer process handle
	assetDir     string                   // Path to unpacked assets
	assetData    astilectron.Disembedder  // Function for reading embedded data
	assetRestore AssetRestoreFn           // function for unpacking embedded data
	msgQueue     chan message.Message     // Queue for outbound GUI messages
	msgHandler   MessageHandlerFn         // Handle function for incoming GUI messages
	opRunning    bool                     // True if performing something that blocks other similar requests
	lock         sync.RWMutex
}

// New returns a new initialized application ready for boostrapping.
func New(dataFn astilectron.Disembedder, restoreFn AssetRestoreFn, msgFn MessageHandlerFn) *App {
	return &App{
		assetData:    dataFn,
		assetRestore: restoreFn,
		msgHandler:   msgFn,
	}
}

// Wait waits until process has finished.
func (a *App) Wait() {
	if a.Renderer == nil {
		return
	}

	a.Renderer.Wait()
	a.Renderer.Close()

	if UseTemp() && a.assetDir != "" {
		os.RemoveAll(a.assetDir)
	}
}

// Busy tells if application state is marked busy.
// This call is synchronized via mutex and is thread safe
func (a *App) Busy() bool {
	a.lock.RLock()
	defer a.lock.RUnlock()
	busy := a.busy()
	return busy
}

func (a *App) busy() bool {
	return a.opRunning
}

// SetOpRunning activates or disables run status. If return value is different
// than input value, then operation was refused
func (a *App) SetOpRunning(active bool) bool {
	a.lock.Lock()
	defer a.lock.Unlock()

	// Not allowed to activate if something is active
	if active && a.busy() {
		return false
	}

	a.opRunning = active
	return active
}
