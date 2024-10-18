package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/ilhaamms/user-management-api/helper"
	"github.com/ilhaamms/user-management-api/models/entity"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		if tokenString == "" || len(tokenString) < 7 {
			helper.ResponseJsonError(w, http.StatusUnauthorized, "authorization is required")
			return
		}

		tokenString = tokenString[7:]

		claims := &entity.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return entity.JwtKey, nil
		})

		if err != nil || !token.Valid {
			helper.ResponseJsonError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		next.ServeHTTP(w, r)
	})
}
