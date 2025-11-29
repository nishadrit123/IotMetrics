package main

import (
	"context"
	"errors"
	"net/http"

	"iot/internal/env"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

var userContextKey = contextKey("user")

func (app *application) OktaAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		loginURL := env.GetString("FRONTEND_URL", "http://localhost:5173") + "/v1/authentication/login"

		// Read cookie instead of header
		cookie, err := r.Cookie("access_token")
		if err != nil {
			app.unauthorizedErrorResponse(w, r, errors.New("missing access_token cookie"))
			http.Redirect(w, r, loginURL, http.StatusTemporaryRedirect)
			return
		}

		tokenString := cookie.Value

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(env.GetString("JWT_SECRET", "")), nil
		})

		if err != nil || !token.Valid {
			app.unauthorizedErrorResponse(w, r, errors.New("invalid or expired token"))
			http.Redirect(w, r, loginURL, http.StatusTemporaryRedirect)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			app.unauthorizedErrorResponse(w, r, errors.New("invalid claims"))
			http.Redirect(w, r, loginURL, http.StatusTemporaryRedirect)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
