package usecase

import (
	"context"
	"roomate/model/entity"
	"roomate/repository"
)

type UserUC interface {
	FindUserById(ctx context.Context, id string) (entity.User, error)
	RegisterUser(ctx context.Context, arg entity.User) (entity.User, error)
	EditUser(ctx context.Context, arg entity.User) error
	FetchingAllUsers(ctx context.Context) ([]entity.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type CustomerUC interface{}

type RoomUC interface {
	FindRoomById(ctx context.Context, id string) (entity.Room, error)
	RegisterRoom(ctx context.Context, arg entity.Room) (entity.Room, error)
	EditRoom(ctx context.Context, arg entity.Room) error
	FetchingAllRooms(ctx context.Context) ([]entity.Room, error)
	DeleteRoom(ctx context.Context, id string) error
}

type RoleUC interface{}

type ServiceUC interface{}

type BookingUC interface{}

type Usecase struct {
	repo *repository.Queries
}

type UC struct {
	UserUC     *Usecase
	CustomerUC *Usecase
	RoomUC     *Usecase
	RoleUC     *Usecase
	ServiceUC  *Usecase
	BookingUC  *Usecase
}

func NewUsecase(repo *repository.Queries) *UC {
	userUC := &Usecase{repo: repo}
	customerUC := &Usecase{repo: repo}
	roomUC := &Usecase{repo: repo}
	roleUC := &Usecase{repo: repo}
	serviceUC := &Usecase{repo: repo}
	bookingUC := &Usecase{repo: repo}
	return &UC{
		UserUC:     userUC,
		CustomerUC: customerUC,
		RoomUC:     roomUC,
		RoleUC:     roleUC,
		ServiceUC:  serviceUC,
		BookingUC:  bookingUC,
	}
}
