package main

import (
	"fmt"
	"github.com/megamxl/se-project/Rental-Server/api"
	req "github.com/megamxl/se-project/Rental-Server/internal/communication/converter"
	int "github.com/megamxl/se-project/Rental-Server/internal/communication/converter/soap"
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

	convService := int.NewSoapService("http://localhost:8080/ws")

	resp, _ := convService.Convert(req.Request{Amount: 12.0, GivenCurrency: "USD", TargetCurrency: "JPY"})

	currency, _ := convService.GetAvailableCurrency()
	for _, s := range currency {
		fmt.Println(s)
	}

	fmt.Println(resp)

	server := api.NewServer()

	r := http.NewServeMux()

	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(server, r)

	hWithMiddleware := LoggingMiddleware(h)

	s := &http.Server{
		Handler: hWithMiddleware,
		Addr:    "0.0.0.0:8098",
	}

	slog.Info("Server started on " + s.Addr)

	// And we serve HTTP until the world ends.
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe failed : ", err)
	}
}
