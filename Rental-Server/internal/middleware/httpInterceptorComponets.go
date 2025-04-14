package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func CheckIfRouteAndMethodMatch(r *http.Request, route string, method string) bool {

	if r.URL.Path == route && r.Method == method {
		return true
	}
	return false
}

func SetRequestContext(r *http.Request, userID string, roles string) *http.Request {
	ctx := context.WithValue(r.Context(), ContextKeyUserID, userID)
	ctx = context.WithValue(ctx, ContextKeyRoles, roles)

	// Update the request context with the new context
	r = r.WithContext(ctx)
	return r
}

func CheckIfTokenExistsAndIsValid(r *http.Request) (jwt.MapClaims, error) {
	token, err := ExtractToken(r)
	if err != nil {
		return nil, err

	}

	claims, err := ValidateAndReturnClaimsFromJWT(token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func MonoMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		r.URL.Path = strings.TrimRight(r.URL.Path, "/")

		routesWithoutAuth :=
			CheckIfRouteAndMethodMatch(r, "/login", "POST") ||
				CheckIfRouteAndMethodMatch(r, "/users", "POST")

		if routesWithoutAuth {
			next.ServeHTTP(w, r)
			return
		}

		if r.Method == "GET" && strings.Contains(r.URL.Path, "/bookings/rpc/in_range") {
			next.ServeHTTP(w, r)
			return
		}

		claims, err := CheckIfTokenExistsAndIsValid(r)
		if err != nil {
			http.Error(w, "Token Invalid or Missing", http.StatusUnauthorized)
			return
		}

		userID, ok1 := claims["sub"].(string)
		roles, ok := claims["roles"].(string)

		if !ok || !ok1 {
			http.Error(w, "Unauthorized: missing subject", http.StatusUnauthorized)
			return
		}

		r = SetRequestContext(r, userID, roles)

		next.ServeHTTP(w, r)
	})
}

func UserServiceMiddleware(next http.Handler) http.Handler {
	return pathBlocker(next, []string{"cars", "booking"})
}

func CarsServiceMiddleware(next http.Handler) http.Handler {
	return pathBlocker(next, []string{"login", "booking", "users"})
}

func BookingsServiceMiddleware(next http.Handler) http.Handler {
	return pathBlocker(next, []string{"login", "cars", "users"})
}

func pathBlocker(next http.Handler, paths []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r.URL.Path = strings.TrimRight(r.URL.Path, "/")

		for _, path := range paths {
			if strings.Contains(r.URL.Path, path) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}

		next.ServeHTTP(w, r)

	})
}
