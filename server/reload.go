package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/josephspurrier/gohtmxapp/pkg"
)

// hotReload enables the hot reload endpoint to support SSE on the front-end.
func hotReload(mux *http.ServeMux) {
	hotReload, _ := strconv.ParseBool(os.Getenv(pkg.EnvHotReload))

	if hotReload {
		// Reload is set to true to force a page reload on application boot.
		reload := true
		maxSeconds := 20
		delaySeconds := 1

		mux.HandleFunc("/sse", func(w http.ResponseWriter, _ *http.Request) {
			flusher, ok := w.(http.Flusher)
			if !ok {
				http.Error(w, "Server does not support flusher!", http.StatusInternalServerError)

				return
			}

			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")

			counter := 0

			for {
				if reload {
					reload = false
					fmt.Fprintf(w, "data: reload\n\n")
				} else {
					fmt.Fprintf(w, "data: skip\n\n")
				}

				flusher.Flush()
				time.Sleep(time.Duration(delaySeconds) * time.Second)

				// Prevent this error by closing the connection before the browser closes it:
				// GET http://localhost:8080/sse net::ERR_INCOMPLETE_CHUNKED_ENCODING 200 (OK)
				counter++
				if counter*delaySeconds >= maxSeconds {
					return
				}
			}
		})
	}
}
