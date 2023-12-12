package controller

import (
	"net/http"
	"roomate/delivery/middleware"
	"roomate/model/dto"
	"roomate/usecase"
	"roomate/utils/common"

	"github.com/gin-gonic/gin"
)

type BookingController interface {
	CreateHandler(ctx *gin.Context)
	GetHandler(ctx *gin.Context)
	GetAllHandler(ctx *gin.Context)
	UpdateBookingStatusHandler(ctx *gin.Context)
}

type bookingController struct {
	uc             usecase.BookingUsecase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (b *bookingController) CreateHandler(ctx *gin.Context) {
	var payload dto.CreateBookingParams

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := b.uc.CreateBooking(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusCreated, "booking created", response)
}

func (b *bookingController) GetHandler(ctx *gin.Context) {
	bookingId := ctx.Param("id")

	response, err := b.uc.GetBooking(bookingId)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusNotFound, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "booking found", response)
}

func (b *bookingController) GetAllHandler(ctx *gin.Context) {
	var payload dto.GetAllParams

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := b.uc.GetAllBookings(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "bookings found", response)
}

func (b *bookingController) UpdateBookingStatusHandler(ctx *gin.Context) {
	var payload dto.UpdateBookingStatusParams

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := b.uc.UpdateBookingStatus(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "booking status updated", response)
}

func (b *bookingController) Route() {
	allUser := b.rg.Group("/bookings", b.authMiddleware.RequireToken("admin, ga, employee"))
	allUser.GET("/:id", b.GetHandler)
	allUser.GET("/", b.GetAllHandler)
	b.rg.POST("/bookings/create", b.CreateHandler, b.authMiddleware.RequireToken("employee"))
	b.rg.PUT("/bookings/status/", b.authMiddleware.RequireToken("ga"), b.UpdateBookingStatusHandler)
}

func NewBookingController(uc usecase.BookingUsecase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *bookingController {
	return &bookingController{
		uc:             uc,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}
