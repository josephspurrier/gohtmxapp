// Package server handles the router configuration.
package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	// Allows loading envrionment variables from a .env file.
	_ "github.com/joho/godotenv/autoload"
)

// Server is a web server.
type Server struct {
	port         int
	hashedAssets bool
}

// NewServer returns an instance of the web server.
func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	// Set the default port if not set.
	if port == 0 {
		port = 8080
	}

	hashedAssets, _ := strconv.ParseBool(os.Getenv("HASH_ASSETS"))

	NewServer := &Server{
		port:         port,
		hashedAssets: hashedAssets,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
