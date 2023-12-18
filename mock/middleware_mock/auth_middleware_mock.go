package middlewaremock

import "github.com/gin-gonic/gin"

type AuthMiddlewareMock struct {
	// sesuaikan struktur sesuai kebutuhan, termasuk jika diperlukan implementasi antarmuka AuthorizeJWT
}

func (a *AuthMiddlewareMock) RequireToken(roles ...string) gin.HandlerFunc {
	return func(context *gin.Context) {
		// yg penting sebuah handlerfunc
	}
}
