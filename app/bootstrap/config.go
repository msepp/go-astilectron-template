package bootstrap

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

// Variables set during build
var (
	// electronVersion is a string that tells the embedded electron version.
	electronVersion string

	// astliectronVersion is a string that tells the embedded astilectron version.
	astilectronVersion string

	// devTools enbables or disables developer tool support. Set to "true" to enable
	// developer tools.
	devTools string

	// buildVersion is the applications version
	buildVersion string

	// name is the name of the application.
	name string

	// prefix gives a filesystem prefix for generated content (dirs etc).
	prefix string

	// resourcesDir is a path prefix for assets
	resourcesDir string
)

// Application icon embedded paths
const (
	EmbeddedIconPng  = "icons/app.png"  // linux
	EmbeddedIconIcns = "icons/app.icns" // mac
	EmbeddedIconIco  = "icons/app.ico"  // windows
)

// EmbeddedUIPackageName is the embedded ui package path
const EmbeddedUIPackageName = "ui/ui.asar"

// Window size constants. These are both the minimum and default values for
// application window size.
const (
	WindowWidth  = 1024
	WindowHeight = 768
)

// Dimensions describes size in width and height in pixels.
type Dimensions struct {
	Height int
	Width  int
}

// Name returns the application name
func Name() string {
	return name
}

// Build returns current build version
func Build() string {
	return buildVersion
}

// ElectronVersion returns bundled electron version
func ElectronVersion() string {
	return electronVersion
}

// AstilectronVersion returns bundled astilectron version
func AstilectronVersion() string {
	return astilectronVersion
}

// DevTools returns if devtools should be enabled
func DevTools() bool {
	return strings.EqualFold(devTools, "true")
}

// EmbeddedResources returns a list of resources that need to be unpacked
func EmbeddedResources() []string {
	return []string{EmbeddedUIPath(), EmbeddedIconPath()}
}

// EmbeddedElectronPath returns the embedded Electron package path
func EmbeddedElectronPath() string {
	return fmt.Sprintf("%s/runtime/electron-%s-%s-%s.zip", resourcesDir, electronVersion, runtime.GOOS, runtime.GOARCH)
}

// EmbeddedAstilectronPath returns the embedded Astilectron package path
func EmbeddedAstilectronPath() string {
	return fmt.Sprintf("%s/runtime/astilectron-%s.zip", resourcesDir, astilectronVersion)
}

// EmbeddedIconPath returns the embedded path of application icon for current
// OS.
func EmbeddedIconPath() string {
	switch runtime.GOOS {
	case "windows":
		return path.Join(resourcesDir, EmbeddedIconIco)

	case "darwin":
		return path.Join(resourcesDir, EmbeddedIconIcns)

	default:
		return path.Join(resourcesDir, EmbeddedIconPng)
	}
}

// EmbeddedUIPath returns path to the UI package
func EmbeddedUIPath() string {
	return path.Join(resourcesDir, EmbeddedUIPackageName)
}

// EmbeddedUIMountPoint return the path to give for Electron to boot up the GUI
func EmbeddedUIMountPoint() string {
	return path.Join(EmbeddedUIPath(), "index.html")
}

// WindowSize returns the effective size to be used for new application window
// to get correct effective size (takes window chrome into account).
func WindowSize() Dimensions {
	switch runtime.GOOS {
	case "windows": // In windows window height is content + chrome, thus offset.
		return Dimensions{Height: WindowHeight + 38, Width: WindowWidth + 26}
	default:
		return Dimensions{Height: WindowHeight, Width: WindowWidth}
	}
}

// TmpDataDir returns an directory path for placing temporary assets
func TmpDataDir() (string, error) {
	return ioutil.TempDir("", prefix)
}

// PersistentDataDir returns an directory path for storing persistent data
func PersistentDataDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	tgt := filepath.Join(home, "."+prefix)
	if err = os.Mkdir(tgt, 0700); err != nil && os.IsExist(err) == false {
		return "", err
	}

	return tgt, nil
}
