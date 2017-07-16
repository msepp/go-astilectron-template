package bootstrap

import (
	"io"
	"path"

	"github.com/msepp/go-astilectron-template/app/assets"
)

// AssetRestoreFn defines an accessor function type for restoring embedded assets to
// file system.
type AssetRestoreFn func(dir, name string) error

// UnpackEmbeddedAssets unloads certain required files to given directory from the
// inmemory copy.
func UnpackEmbeddedAssets(toDir string, assetFn AssetRestoreFn) error {
	for _, asset := range EmbeddedResources() {
		if err := assetFn(toDir, asset); err != nil {
			return err
		}
	}

	return nil
}

// AssetReader returns an function for reading in-memory firmware asset as a io.Reader
func (a *App) AssetReader() assets.GetAssetFunc {
	return func(fname string) (io.ReadCloser, error) {
		b, err := a.assetData(path.Join(resourcesDir, fname))
		if err != nil {
			return nil, err
		}

		return assets.New(b), nil
	}
}
