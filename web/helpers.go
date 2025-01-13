package web

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/josephspurrier/gohtmxapp/pkg"
)

// AssetPath returns the hashed or unhashed asset path.
func assetPath(ctx context.Context, name string) string {
	if pkg.ContextBool(ctx, pkg.ContextHashed) {
		// Ensure to strip the leading slash, then add it back.
		return "/" + HashedFiles.HashName(strings.TrimPrefix(name, "/"))
	}

	return name
}

// hotReload returns true if it should be hot reloaded.
func hotReload() bool {
	hr, _ := strconv.ParseBool(os.Getenv(pkg.EnvHotReload))

	return hr
}

// isActive returns true if the page URL matches.
func isActive(ctx context.Context, urlPath string) string {
	if urlPath == pkg.ContextString(ctx, pkg.ContextPageURL) {
		return "active"
	}

	return ""
}

// navbarSubmenuHidden returns 'hidden' if the submenu is not open.
func navbarSubmenuHidden(ctx context.Context, submenu string) string {
	if submenu != pkg.ContextString(ctx, pkg.ContextSubmenuName) {
		return "hidden"
	}

	return ""
}

// hxValMap returns a JSON string from a map for hx-vals.
func hxValMap(m map[string]string) string {
	bytes, _ := json.Marshal(m)

	return string(bytes)
}
