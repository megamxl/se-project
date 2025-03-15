package main

import (
	"github.com/megamxl/se-project/Rental-Server/api"
	"log"
	"net/http"
)

func main() {

	server := api.NewServer()

	r := http.NewServeMux()

	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
