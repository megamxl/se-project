package main

import (
	"fmt"
	"github.com/megamxl/se-project/Rental-Server/api"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		slog.Debug(fmt.Sprintf("Received request: %s %s from %s ", r.Method, r.URL.Path, r.RemoteAddr))

		// Iterate over headers and log them
		for name, values := range r.Header {
			for _, value := range values {
				log.Printf("Header: %s = %s", name, value)
			}
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	server := api.NewServer()

	r := http.NewServeMux()

	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(server, r)

	hWithMiddleware := LoggingMiddleware(h)

	s := &http.Server{
		Handler: hWithMiddleware,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
