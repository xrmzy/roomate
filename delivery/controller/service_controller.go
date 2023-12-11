package controller

import (
	"net/http"
	"roomate/model/dto"
	"roomate/model/entity"
	"roomate/usecase"
	"roomate/utils/common"

	"github.com/gin-gonic/gin"
)

type ServiceController struct {
	uc usecase.ServiceUseCase
	rg *gin.RouterGroup
}

func (s *ServiceController) CreateHandler(ctx *gin.Context) {
	var payload entity.Service

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := s.uc.CreateService(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusCreated, "service created", response)
}

func (s *ServiceController) GetHandler(ctx *gin.Context) {
	serviceId := ctx.Param("id")

	response, err := s.uc.GetService(serviceId)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusNotFound, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "service found", response)
}

func (s *ServiceController) GetAllHandler(ctx *gin.Context) {
	var payload dto.GetAllParams

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := s.uc.GetAllServices(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "services found", response)
}

func (s *ServiceController) UpdateHandler(ctx *gin.Context) {
	serviceId := ctx.Param("id")
	var payload entity.Service

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := s.uc.UpdateService(serviceId, payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "service updated", response)
}

// delete handler
func (s *ServiceController) DeleteHandler(ctx *gin.Context) {
	serviceId := ctx.Param("id")

	err := s.uc.DeleteService(serviceId)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "service deleted", nil)
}

func (s *ServiceController) Route() {
	s.rg.GET("/services/:id", s.GetHandler)
	s.rg.GET("/services", s.GetAllHandler)
	s.rg.POST("/services", s.CreateHandler)
	s.rg.PUT("/services/:id", s.UpdateHandler)
	s.rg.DELETE("/services/:id", s.DeleteHandler)
}

func NewServiceController(uc usecase.ServiceUseCase, rg *gin.RouterGroup) *ServiceController {
	return &ServiceController{
		uc: uc,
		rg: rg,
	}
}