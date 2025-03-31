package main

import (
	"context"
	"fmt"
	"github.com/megamxl/se-project/Rental-Server/api"
	"github.com/megamxl/se-project/Rental-Server/internal/middleware"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		r.URL.Path = strings.TrimRight(r.URL.Path, "/")

		if r.URL.Path == "/login" && r.Method == "POST" {
			next.ServeHTTP(w, r)
			return
		}

		if r.URL.Path == "/users" && r.Method == "POST" {
			next.ServeHTTP(w, r)
			return
		}

		token, err := middleware.ExtractToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		claims, err := middleware.ValidateAndReturnClaimsFromJWT(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		userID, ok := claims["sub"].(string)
		if !ok {
			http.Error(w, "Unauthorized: missing subject", http.StatusUnauthorized)
			return
		}

		roles, ok := claims["roles"].(string)
		if !ok {
			http.Error(w, "Unauthorized: missing subject", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), middleware.ContextKeyUserID, userID)
		ctx = context.WithValue(ctx, middleware.ContextKeyRoles, roles)

		// Update the request context with the new context
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	server := api.NewServer(generateDSNStringFromEnvVariables())

	r := http.NewServeMux()

	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(server, r)

	hWithMiddleware := Middleware(h)

	s := &http.Server{
		Handler: hWithMiddleware,
		Addr:    os.Getenv("WEB_HOST") + ":" + os.Getenv("WEB_PORT"),
	}

	slog.Info("Server started on " + s.Addr)

	// And we serve HTTP until the world ends.
	err1 := s.ListenAndServe()
	if err1 != nil {
		log.Fatal("ListenAndServe failed : ", err1)
	}
}

func generateDSNStringFromEnvVariables() string {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPW := os.Getenv("DB_PASSWORD")
	dbSSL := os.Getenv("DB_SSL-MODE")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUsername, dbPW, dbName, dbPort, dbSSL)
}
