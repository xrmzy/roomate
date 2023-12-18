package middleware

import (
	"net/http"
	"roomate/utils/common"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	RequireToken(roles ...string) gin.HandlerFunc
}

type authMiddleware struct {
	jwtService common.JwtToken
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func (a *authMiddleware) RequireToken(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := authHeader{}
		err := ctx.ShouldBindHeader(&authHeader)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message: ": "Unauthorized"})
			return
		}

		tokenString := strings.Replace(authHeader.AuthorizationHeader, "Bearer ", "", -1)
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message: ": "Unauthorized"})
			return
		}

		claims, err := a.jwtService.VerifyToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message: ": "Unauthorized"})
			return
		}

		ctx.Set("user", claims["userId"])

		var validRole bool
		for _, role := range roles {
			if role == claims["role"] {
				validRole = true
				break
			}
		}

		if !validRole {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message: ": "This role is not allowed to access this resource"})
			return
		}

		ctx.Next()
	}
}

func NewAuthMiddleware(jwtService common.JwtToken) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}
