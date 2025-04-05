package middleware

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/megamxl/se-project/Rental-Server/internal/data"
	"net/http"
	"os"
	"strings"
	"time"
)

func CreateJWForUser(user data.RentalUser) (string, error) {

	claims := jwt.MapClaims{
		"sub":   user.Id, // subject (e.g., user ID)
		"name":  user.Name,
		"roles": "customer",
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // Expires in 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(getSecretFromEnv())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateAndReturnClaimsFromJWT(tokenString string) (jwt.MapClaims, error) {

	secretKey := getSecretFromEnv()

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	}, jwt.WithLeeway(2*time.Second)) // Optional: allow 2s clock skew

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, err
	} else {
		return nil, fmt.Errorf("can't get roles claim")
	}
}

func getSecretFromEnv() []byte {
	secretKey := []byte(os.Getenv("jwtSecret"))
	return secretKey
}

var ErrMissingAuthHeader = errors.New("missing Authorization header")
var ErrInvalidAuthHeader = errors.New("invalid Authorization header format")

func ExtractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {

		cookie, err := r.Cookie("jwt")
		if err != nil {
			return "", ErrMissingAuthHeader
		}

		return cookie.Value, nil
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}

type contextKey string

const (
	ContextKeyUserID contextKey = "userID"
	ContextKeyRoles  contextKey = "roles"
)
