// Package web contains the web assets, scripts, and templates.
package web

import (
	"embed"

	"github.com/benbjohnson/hashfs"
)

// Files contains the embedded assets.
//
//go:embed "assets"
var Files embed.FS

// HashedFiles is the hashed assets.
var HashedFiles = hashfs.NewFS(Files)
