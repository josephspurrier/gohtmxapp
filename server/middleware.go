package server

import (
	"context"
	"net/http"

	"github.com/josephspurrier/gohtmxapp/pkg"
)

// chain middleware - ensure the first one gets executed first.
func chain(handler http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	return handler
}

// corsMiddleware applies CORS to the API endpoints.
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}

// templateVariableMiddleware adds variables for the templates.
func (s *Server) templateVariableMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ensure templates know when to hash assets or not.
		ctx := context.WithValue(r.Context(), pkg.ContextHashed, s.hashedAssets)
		// Store the page URL.
		ctx = context.WithValue(ctx, pkg.ContextPageURL, r.URL.Path)
		// Store the submenu name.
		ctx = context.WithValue(ctx, pkg.ContextSubmenuName, r.URL.Query().Get(pkg.SubmenuQueryParameter))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
