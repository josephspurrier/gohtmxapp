package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/benbjohnson/hashfs"
	"github.com/josephspurrier/gohtmxapp/web"
)

// RegisterRoutes adds routes to the mux.
func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	if s.hashedAssets {
		// Hashed assets.
		mux.Handle("/assets/", hashfs.FileServer(web.HashedFiles))
	} else {
		// Unhashed assets.
		mux.Handle("/assets/", http.FileServer(http.FS(web.Files)))
	}

	// Register pages.
	mux.Handle("/dashboard", templ.Handler(web.DashboardPage()))
	mux.Handle("/users", templ.Handler(web.HelloForm()))
	mux.Handle("/products", templ.Handler(web.PagePlaceholder("Products")))
	mux.Handle("/billing", templ.Handler(web.PagePlaceholder("Billing")))
	mux.Handle("/invoice", templ.Handler(web.PagePlaceholder("Invoice")))
	mux.Handle("/kanban", templ.Handler(web.PagePlaceholder("Kanban")))
	mux.Handle("/inbox", templ.Handler(web.PagePlaceholder("Inbox")))
	mux.Handle("/preferences", templ.Handler(web.PagePlaceholder("Preferences")))
	mux.Handle("/settings", templ.Handler(web.SettingsPage()))

	// Register content.
	mux.HandleFunc("/hello", web.HelloHandler)

	// Determines if hot reload endpoint should be enabled.
	hotReload(mux)

	// Wrap the mux with middleware. The top gets executed first.
	return chain(
		mux,
		s.corsMiddleware,
		s.templateVariableMiddleware,
	)
}
