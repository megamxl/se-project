package main

import (
	"github.com/megamxl/se-project/Rental-Server/api"
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("Received request: %s %s", r.Method, r.URL.Path)

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
