package usecase

import (
	"roomate/model/dto"
	"roomate/model/entity"
	"roomate/repository"
	"roomate/utils/common"
)

type RoomUseCase interface {
	GetAllRooms(payload dto.GetAllParams) ([]entity.Room, error)
	GetRoom(id string) (entity.Room, error)
	CreateRoom(room entity.Room) (entity.Room, error)
	UpdateRoom(id string, room entity.Room) (entity.Room, error)
	DeleteRoom(id string) error
}

type roomUseCase struct {
	roomRepo repository.RoomRepository
}

var GenerateRoomID = common.GenerateID("R")

func (u *roomUseCase) GetAllRooms(payload dto.GetAllParams) ([]entity.Room, error) {
	rooms, err := u.roomRepo.GetAll(payload.Limit, payload.Offset)
	if err != nil {
		return rooms, err
	}

	return rooms, nil
}

func (u *roomUseCase) GetRoom(id string) (entity.Room, error) {
	room, err := u.roomRepo.Get(id)
	if err != nil {
		return room, err
	}

	return room, nil
}

func (u *roomUseCase) CreateRoom(room entity.Room) (entity.Room, error) {
	room.Id = GenerateRoomID
	room, err := u.roomRepo.Create(room)
	if err != nil {
		return room, err
	}

	return room, nil
}

func (u *roomUseCase) UpdateRoom(id string, room entity.Room) (entity.Room, error) {
	room, err := u.roomRepo.Update(id, room)
	if err != nil {
		return room, err
	}

	return room, nil
}

func (u *roomUseCase) DeleteRoom(id string) error {
	err := u.roomRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewRoomUseCase(roomRepo repository.RoomRepository) RoomUseCase {
	return &roomUseCase{
		roomRepo: roomRepo,
	}
}
