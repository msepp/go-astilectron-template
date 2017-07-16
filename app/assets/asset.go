// Package assets defines an interface for accessing required files from arbitrary
// filesystem or media
package assets

import (
	"bytes"
	"io"
)

// GetAssetFunc defines an interface function for retrieving data of a named file
type GetAssetFunc func(path string) (io.ReadCloser, error)

// Asset is an embedded asset that implements io.ReadCloser
type Asset struct {
	*bytes.Reader
}

// Close method for completing io.Closer
func (a *Asset) Close() error { return nil }

// New returns an asset from given bytes
func New(b []byte) *Asset {
	return &Asset{Reader: bytes.NewReader(b)}
}
