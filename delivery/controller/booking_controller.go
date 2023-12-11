package controller

import (
	"net/http"
	"roomate/model/dto"
	"roomate/usecase"
	"roomate/utils/common"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	uc usecase.BookingUsecase
	rg *gin.RouterGroup
}

func (b *BookingController) CreateHandler(ctx *gin.Context) {
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

func (b *BookingController) GetHandler(ctx *gin.Context) {
	bookingId := ctx.Param("id")

	response, err := b.uc.GetBooking(bookingId)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusNotFound, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "booking found", response)
}

func (b *BookingController) GetAllHandler(ctx *gin.Context) {
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

func (b *BookingController) Route() {
	b.rg.GET("/bookings/:id", b.GetHandler)
	b.rg.GET("/bookings", b.GetAllHandler)
	b.rg.POST("/bookings", b.CreateHandler)
}

func NewBookingController(uc usecase.BookingUsecase, rg *gin.RouterGroup) *BookingController {
	return &BookingController{
		uc: uc,
		rg: rg,
	}
}
