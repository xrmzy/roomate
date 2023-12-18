package controller

import (
	"net/http"
	"roomate/model/dto"
	"roomate/model/entity"
	"roomate/usecase"
	"roomate/utils/common"

	"github.com/gin-gonic/gin"
)

type RoomController interface {
	CreateHandler(ctx *gin.Context)
	GetHandler(ctx *gin.Context)
	GetAllHandler(ctx *gin.Context)
	UpdateHandler(ctx *gin.Context)
	DeleteHandler(ctx *gin.Context)
	Route()
}

type roomController struct {
	uc usecase.RoomUseCase
	rg *gin.RouterGroup
}

func (r *roomController) CreateHandler(ctx *gin.Context) {
	var payload entity.Room

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := r.uc.CreateRoom(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusCreated, "room created", response)
}

func (r *roomController) GetHandler(ctx *gin.Context) {
	roomId := ctx.Param("id")

	response, err := r.uc.GetRoom(roomId)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusNotFound, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "room found", response)
}

func (r *roomController) GetAllHandler(ctx *gin.Context) {
	var payload dto.GetAllParams

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := r.uc.GetAllRooms(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendPagedResponse(ctx, http.StatusOK, "rooms found", response, gin.H{"Start": payload.Offset, "End": payload.Offset + payload.Limit})
}

func (r *roomController) UpdateHandler(ctx *gin.Context) {
	roomId := ctx.Param("id")
	var payload entity.Room

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := r.uc.UpdateRoom(roomId, payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "room updated", response)
}

func (r *roomController) DeleteHandler(ctx *gin.Context) {
	roomId := ctx.Param("id")

	err := r.uc.DeleteRoom(roomId)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "room deleted", nil)
}

func (r *roomController) Route() {
	r.rg.GET("/rooms/:id", r.GetHandler)
	r.rg.GET("/rooms", r.GetAllHandler)
	r.rg.POST("/rooms", r.CreateHandler)
	r.rg.PUT("/rooms/:id", r.UpdateHandler)
	r.rg.DELETE("/rooms/:id", r.DeleteHandler)
}

func NewRoomController(uc usecase.RoomUseCase, rg *gin.RouterGroup) *roomController {
	return &roomController{
		uc: uc,
		rg: rg,
	}
}
