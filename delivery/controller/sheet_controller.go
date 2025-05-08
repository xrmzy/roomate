package controller

import (
	"io"
	"net/http"
	"roomate/delivery/middleware"
	"roomate/model/dto"
	"roomate/usecase"
	"roomate/utils/common"
	"time"

	"github.com/gin-gonic/gin"
)

type SheetController interface {
	dayHandler(ctx *gin.Context)
	monthHandler(ctx *gin.Context)
	yearHandler(ctx *gin.Context)
	Route()
}

type sheetController struct {
	sheetUc        usecase.GSheetUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (s *sheetController) dayHandler(ctx *gin.Context) {
	var payload dto.GetBookingOneDayParams
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	resp, err := s.sheetUc.DailyReport(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	parsedTime, err := time.Parse("2006/01/02", payload.Date)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	formattedDate := parsedTime.Format("2006-01-02")
	newFileName := "DailyReport-" + formattedDate + ".xlsx"

	ctx.Header("Content-Disposition", "attachment; filename="+newFileName)

	// Copy the file from the API response to the client's response
	_, err = io.Copy(ctx.Writer, resp.Body)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

func (s *sheetController) monthHandler(ctx *gin.Context) {
	var payload dto.GetBookingOneMonthParams
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	resp, err := s.sheetUc.MonthlyReport(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	newFileName := "MonthlyReport-" + payload.Month + "-" + payload.Year + ".xlsx"

	ctx.Header("Content-Disposition", "attachment; filename="+newFileName)

	// Copy the file from the API response to the client's response
	_, err = io.Copy(ctx.Writer, resp.Body)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

func (s *sheetController) yearHandler(ctx *gin.Context) {
	var payload dto.GetBookingOneYearParams
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	resp, err := s.sheetUc.YearlyReport(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	newFileName := "YearlyReport-" + payload.Year + ".xlsx"

	ctx.Header("Content-Disposition", "attachment; filename="+newFileName)

	// Copy the file from the API response to the client's response
	_, err = io.Copy(ctx.Writer, resp.Body)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}

func (s *sheetController) Route() {
	onlyAdmin := s.rg.Group("/reports", s.authMiddleware.RequireToken("admin"))
	onlyAdmin.GET("/daily", s.dayHandler)
	onlyAdmin.GET("/monthly", s.monthHandler)
	onlyAdmin.GET("/yearly", s.yearHandler)
}

func NewGSheetController(sheetUc usecase.GSheetUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) SheetController {
	return &sheetController{
		sheetUc:        sheetUc,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}
