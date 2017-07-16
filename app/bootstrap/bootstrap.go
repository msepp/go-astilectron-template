package bootstrap

import (
	"log"
	"path"

	"github.com/msepp/go-astilectron-template/app/message"
	astilectron "github.com/asticode/go-astilectron"
)

// Bootstrap initializes the app runtime and returns a new application struct.
func (a *App) Bootstrap() error {
	var devTools = DevTools()
	var err error
	var hadCrash bool

	// Get directory for placing assets
	if a.assetDir, err = TmpDataDir(); err != nil {
		return err
	}

	// Unpack assets
	if err = UnpackEmbeddedAssets(a.assetDir, a.assetRestore); err != nil {
		return err
	}

	// Initialize astilectron
	a.Renderer, err = astilectron.New(astilectron.Options{
		AppName:            Name(),
		AppIconDefaultPath: path.Join(a.assetDir, EmbeddedIconPath()),
		AppIconDarwinPath:  path.Join(a.assetDir, EmbeddedIconPath()),
		BaseDirectoryPath:  a.assetDir,
	})
	if err != nil {
		return err
	}

	// Set provisioning to load from binary data.
	a.Renderer.SetProvisioner(
		astilectron.NewDisembedderProvisioner(
			a.assetData,
			EmbeddedAstilectronPath(),
			EmbeddedElectronPath(),
		),
	)

	// Set handling for signals and capture crashes
	a.Renderer.HandleSignals()
	a.Renderer.On(astilectron.EventNameAppCrash, func(e astilectron.Event) (deleteListener bool) {
		log.Printf("[EXIT] GUI has exited")
		log.Printf("[EXIT] event: %s, message: %v", e.Name, e.Message)
		hadCrash = true
		return true
	})

	// Start astilectron
	if err = a.Renderer.Start(); err != nil {
		log.Fatal(err)
	}

	if hadCrash {
		log.Fatal("Crashed during init.")
	}

	// Create main window
	wd := WindowSize()
	if a.Window, err = a.Renderer.NewWindow(path.Join(a.assetDir, EmbeddedUIMountPoint()), &astilectron.WindowOptions{
		Center:    astilectron.PtrBool(true),
		Height:    astilectron.PtrInt(wd.Height),
		Width:     astilectron.PtrInt(wd.Width),
		MinHeight: astilectron.PtrInt(wd.Height),
		MinWidth:  astilectron.PtrInt(wd.Width),
		Title:     astilectron.PtrStr(Name()),
		Icon:      astilectron.PtrStr(path.Join(a.assetDir, EmbeddedIconPath())),
		WebPreferences: &astilectron.WebPreferences{
			DevTools:        &devTools,
			DefaultEncoding: astilectron.PtrStr("utf-8"),
			Webaudio:        astilectron.PtrBool(false),
		},
	}); err != nil {
		log.Fatal(err)
	}

	// Setup queue for message sending
	a.msgQueue = make(chan message.Message, 50)

	// Run routine for handling sending
	go a.messageQueueFlusher()

	// Setup handler for GUI messages
	a.Window.On(astilectron.EventNameWindowEventMessage, a.onWindowMessage())

	// Actually create the window to make it appear.
	a.Window.Create()

	// Open dev tools
	if devTools {
		a.Window.OpenDevTools()
	}

	return nil
}
