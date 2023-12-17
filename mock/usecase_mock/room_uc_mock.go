package usecasemock

import (
	"roomate/model/dto"
	"roomate/model/entity"

	"github.com/stretchr/testify/mock"
)

type RoomUseCaseMock struct {
	mock.Mock
}

func (r *RoomUseCaseMock) GetRoom(id string) (entity.Room, error) {
	args := r.Called(id)
	return args.Get(0).(entity.Room), args.Error(1)
}

func (r *RoomUseCaseMock) GetAllRooms(payload dto.GetAllParams) ([]entity.Room, error) {
	args := r.Called(payload)
	return args.Get(0).([]entity.Room), args.Error(1)
}

func (r *RoomUseCaseMock) CreateRoom(room entity.Room) (entity.Room, error) {
	args := r.Called(room)
	return args.Get(0).(entity.Room), args.Error(1)
}

func (r *RoomUseCaseMock) UpdateRoom(id string, room entity.Room) (entity.Room, error) {
	args := r.Called(id, room)
	return args.Get(0).(entity.Room), args.Error(1)
}

func (r *RoomUseCaseMock) UpdateStatus(id string) error {
	args := r.Called(id)
	return args.Error(0)
}

func (r *RoomUseCaseMock) DeleteRoom(id string) error {
	args := r.Called(id)
	return args.Error(0)
}
