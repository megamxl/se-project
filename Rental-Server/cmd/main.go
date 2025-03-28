package main

import (
	"context"
	"github.com/megamxl/se-project/Rental-Server/api"
	d "github.com/megamxl/se-project/Rental-Server/internal/data/sql"
	"github.com/megamxl/se-project/Rental-Server/internal/middleware"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		if r.URL.Path == "/login" && r.Method == "POST" {
			next.ServeHTTP(w, r)
		}

		if r.URL.Path == "/users" && r.Method == "POST" {
			next.ServeHTTP(w, r)
		}

		token, err := middleware.ExtractBearerToken(r)
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

	d.Db()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	//convService := int.NewSoapService("http://localhost:8080/ws")
	//
	//resp, _ := convService.Convert(req.Request{Amount: 12.0, GivenCurrency: "USD", TargetCurrency: "JPY"})
	//
	//currency, _ := convService.GetAvailableCurrency()
	//for _, s := range currency {
	//	fmt.Println(s)
	//}

	//fmt.Println(resp)

	dsn := "host=localhost user=admin password=admin dbname=main port=5432 sslmode=disable"

	server := api.NewServer(dsn)

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
	err1 := s.ListenAndServe()
	if err1 != nil {
		log.Fatal("ListenAndServe failed : ", err1)
	}
}
