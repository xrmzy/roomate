package repomock

import (
	"roomate/model/entity"

	"github.com/stretchr/testify/mock"
)

type RoomRepoMock struct {
	mock.Mock
}

func (r *RoomRepoMock) Get(id string) (entity.Room, error) {
	args := r.Called(id)
	return args.Get(0).(entity.Room), args.Error(1)
}

func (r *RoomRepoMock) GetAll(limit, offset int) ([]entity.Room, error) {
	args := r.Called(limit, offset)
	return args.Get(0).([]entity.Room), args.Error(1)
}

func (r *RoomRepoMock) Create(room entity.Room) (entity.Room, error) {
	args := r.Called(room)
	return args.Get(0).(entity.Room), args.Error(1)
}

func (r *RoomRepoMock) Update(id string, room entity.Room) (entity.Room, error) {
	args := r.Called(id, room)
	return args.Get(0).(entity.Room), args.Error(1)
}

func (r *RoomRepoMock) UpdateStatus(id string) error {
	args := r.Called(id)
	return args.Error(0)
}

func (r *RoomRepoMock) Delete(id string) error {
	args := r.Called(id)
	return args.Error(0)
}
