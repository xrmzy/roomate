package modelutil

import "github.com/golang-jwt/jwt/v5"

type JwtTokenClaims struct {
	jwt.RegisteredClaims
	// list yang dibuat di payload
	UserId string `json:"userId"`
	Role   string `json:"role"`
}
