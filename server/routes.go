package server

import (
	"net/http"

	"github.com/benbjohnson/hashfs"
	"github.com/josephspurrier/gohtmxapp/web"

	"github.com/a-h/templ"
)

// RegisterRoutes adds routes to the mux.
func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Register routes
	// mux.HandleFunc("/", s.HelloWorldHandler)

	if s.hashedAssets {
		// Hashed assets.
		mux.Handle("/assets/", hashfs.FileServer(web.HashedFiles))
	} else {
		// Unhashed assets.
		mux.Handle("/assets/", http.FileServer(http.FS(web.Files)))
	}

	mux.Handle("/web", templ.Handler(web.HelloForm()))
	mux.HandleFunc("/hello", web.HelloHandler)

	mux.Handle("/dashboard", templ.Handler(web.DashboardPage()))
	mux.Handle("/settings", templ.Handler(web.SettingsPage()))

	// Determines if hot reload endpoint should be enabled.
	hotReload(mux)

	// Wrap the mux with middleware. The top gets executed first.
	return chain(
		mux,
		s.corsMiddleware,
		s.templateVariableMiddleware,
	)
}
