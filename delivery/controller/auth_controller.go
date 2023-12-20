package controller

import (
	"net/http"
	"roomate/model/dto"
	"roomate/usecase"
	"roomate/utils/common"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	loginHandler(ctx *gin.Context)
	refreshTokenHandler(ctx *gin.Context)
	Route()
}

type authController struct {
	authUc     usecase.AuthUseCase
	rg         *gin.RouterGroup
	jwtService common.JwtToken
}

func (a *authController) loginHandler(ctx *gin.Context) {
	var payload dto.AuthRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := a.authUc.Login(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "successfully logged in", response)
}

func (a *authController) refreshTokenHandler(ctx *gin.Context) {
	tokenString := strings.Replace(ctx.GetHeader("Authorization"), "Bearer ", "", -1)
	newToken, err := a.jwtService.RefreshToken(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "successfully generate new token", newToken)
}

func (a *authController) Route() {
	ag := a.rg.Group("/auth")
	ag.GET("/login", a.loginHandler)
	ag.GET("/refresh-token", a.refreshTokenHandler)
}

func NewAuthController(authUc usecase.AuthUseCase, rg *gin.RouterGroup, jwtService common.JwtToken) AuthController {
	return &authController{
		authUc:     authUc,
		rg:         rg,
		jwtService: jwtService,
	}
}
