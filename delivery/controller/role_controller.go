package controller

import (
	"net/http"
	"roomate/model/dto"
	"roomate/model/entity"
	"roomate/usecase"
	"roomate/utils/common"

	"github.com/gin-gonic/gin"
)

type RoleController interface {
	CreateHandler(ctx *gin.Context)
	GetHandler(ctx *gin.Context)
	GetAllHandler(ctx *gin.Context)
	UpdateHandler(ctx *gin.Context)
	DeleteHandler(ctx *gin.Context)
	Route()
}

type roleController struct {
	uc usecase.RoleUseCase
	rg *gin.RouterGroup
}

func (r *roleController) CreateHandler(ctx *gin.Context) {
	var payload entity.Role

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := r.uc.CreateRole(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusCreated, "role created", response)
}

func (r *roleController) GetHandler(ctx *gin.Context) {
	roleId := ctx.Param("id")

	response, err := r.uc.GetRole(roleId)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusNotFound, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "role found", response)
}

func (r *roleController) GetAllHandler(ctx *gin.Context) {
	var payload dto.GetAllParams

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := r.uc.GetAllRoles(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendPagedResponse(ctx, http.StatusOK, "roles found", response, gin.H{"Start": payload.Offset, "End": payload.Offset + payload.Limit})
}

func (r *roleController) UpdateHandler(ctx *gin.Context) {
	roleId := ctx.Param("id")
	var payload entity.Role

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := r.uc.UpdateRole(roleId, payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "role updated", response)
}

func (r *roleController) DeleteHandler(ctx *gin.Context) {
	roleId := ctx.Param("id")

	err := r.uc.DeleteRole(roleId)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "role deleted", nil)
}

func (r *roleController) Route() {
	r.rg.GET("/roles/:id", r.GetHandler)
	r.rg.GET("/roles", r.GetAllHandler)
	r.rg.POST("/roles", r.CreateHandler)
	r.rg.PUT("/roles/:id", r.UpdateHandler)
	r.rg.DELETE("/roles/:id", r.DeleteHandler)
}

func NewRoleController(uc usecase.RoleUseCase, rg *gin.RouterGroup) *roleController {
	return &roleController{
		uc: uc,
		rg: rg,
	}
}
