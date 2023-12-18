package controller

import (
	"net/http"
	"roomate/model/dto"
	"roomate/model/entity"
	"roomate/usecase"
	"roomate/utils/common"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateHandler(ctx *gin.Context)
	GetHandler(ctx *gin.Context)
	GetAllHandler(ctx *gin.Context)
	UpdateHandler(ctx *gin.Context)
	UpdatePasswordHandler(ctx *gin.Context)
	DeleteHandler(ctx *gin.Context)
	Route()
}

type userController struct {
	uc usecase.UserUseCase
	rg *gin.RouterGroup
}

func (u *userController) CreateHandler(ctx *gin.Context) {
	var payload entity.User

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := u.uc.CreateUser(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusCreated, "user created", response)
}

func (u *userController) GetHandler(ctx *gin.Context) {
	userId := ctx.Param("id")

	response, err := u.uc.GetUser(userId)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusNotFound, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "user found", response)
}

func (u *userController) GetAllHandler(ctx *gin.Context) {
	var payload dto.GetAllParams

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := u.uc.GetAllUsers(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendPagedResponse(ctx, http.StatusOK, "users found", response, gin.H{"Start": payload.Offset, "End": payload.Offset + payload.Limit})
}

func (u *userController) UpdateHandler(ctx *gin.Context) {
	userId := ctx.Param("id")
	var payload entity.User

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := u.uc.UpdateUser(userId, payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "user updated", response)
}

func (u *userController) UpdatePasswordHandler(ctx *gin.Context) {
	userId := ctx.Param("id")
	var payload entity.User

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := u.uc.UpdatePassword(userId, payload.Password)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "password updated", response)
}

func (u *userController) DeleteHandler(ctx *gin.Context) {
	userId := ctx.Param("id")

	err := u.uc.DeleteUser(userId)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "user deleted", nil)
}

func (u *userController) Route() {
	u.rg.GET("/users/:id", u.GetHandler)
	u.rg.GET("/users", u.GetAllHandler)
	u.rg.POST("/users", u.CreateHandler)
	u.rg.PUT("/users/:id", u.UpdateHandler)
	u.rg.PUT("/users/update-password/:id", u.UpdatePasswordHandler) // update password berguna jika forgot password
	u.rg.DELETE("/users/:id", u.DeleteHandler)
}

func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup) *userController {
	return &userController{
		uc: uc,
		rg: rg,
	}
}
