package controller

import (
	"net/http"
	"roomate/model/dto"
	"roomate/model/entity"
	"roomate/usecase"
	"roomate/utils/common"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	CreateHandler(ctx *gin.Context)
	GetHandler(ctx *gin.Context)
	GetAllHandler(ctx *gin.Context)
	UpdateHandler(ctx *gin.Context)
	DeleteHandler(ctx *gin.Context)
	Route()
}

type customerController struct {
	uc usecase.CustomerUseCase
	rg *gin.RouterGroup
}

func (c *customerController) CreateHandler(ctx *gin.Context) {
	var payload entity.Customer

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := c.uc.CreateCustomer(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusCreated, "customer created", response)
}

func (c *customerController) GetHandler(ctx *gin.Context) {
	userId := ctx.Param("id")

	response, err := c.uc.GetCustomer(userId)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusNotFound, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "customer found", response)
}

func (c *customerController) GetAllHandler(ctx *gin.Context) {
	var payload dto.GetAllParams

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := c.uc.GetAllCustomers(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendPagedResponse(ctx, http.StatusOK, "customers found", response, gin.H{"Start": payload.Offset, "End": payload.Offset + payload.Limit})
}

func (c *customerController) UpdateHandler(ctx *gin.Context) {
	userId := ctx.Param("id")
	var payload entity.Customer

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := c.uc.UpdateCustomer(userId, payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "customer updated", response)
}

func (c *customerController) DeleteHandler(ctx *gin.Context) {
	userId := ctx.Param("id")

	err := c.uc.DeleteCustomer(userId)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "customer deleted", nil)
}

func (c *customerController) Route() {
	c.rg.GET("/customers/:id", c.GetHandler)
	c.rg.GET("/customers", c.GetAllHandler)
	c.rg.POST("/customers", c.CreateHandler)
	c.rg.PUT("/customers/:id", c.UpdateHandler)
	c.rg.DELETE("/customers/:id", c.DeleteHandler)
}

func NewCustomerController(uc usecase.CustomerUseCase, rg *gin.RouterGroup) *customerController {
	return &customerController{
		uc: uc,
		rg: rg,
	}
}
