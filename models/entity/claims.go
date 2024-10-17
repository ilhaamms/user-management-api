package entity

import "github.com/golang-jwt/jwt"

var jwtKey = []byte("user-management-api-secret")

type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}
