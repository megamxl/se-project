package config

import (
	"fmt"
	"github.com/megamxl/se-project/Rental-Server/api"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func ListenAndServeServer(hWithMiddleware http.Handler, address string) {
	s := &http.Server{
		Handler: hWithMiddleware,
		Addr:    address,
	}

	slog.Info("Server started on " + s.Addr)

	// And we serve HTTP until the world ends.
	err1 := s.ListenAndServe()
	if err1 != nil {
		log.Fatal("ListenAndServe failed : ", err1)
	}
}

func BasicServerSetup() http.Handler {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	server := api.NewServer(GenerateDSNStringFromEnvVariables())

	r := http.NewServeMux()

	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(server, r)
	return h
}

func GenerateDSNStringFromEnvVariables() string {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPW := os.Getenv("DB_PASSWORD")
	dbSSL := os.Getenv("DB_SSL-MODE")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUsername, dbPW, dbName, dbPort, dbSSL)
}
