package usecase

import (
	"roomate/model/dto"
	"roomate/model/entity"
	"roomate/repository"
)

type UserUseCase interface {
	GetAllUsers(payload dto.GetAllParams) ([]entity.User, error)
	GetUser(id string) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(id string, user entity.User) (entity.User, error)
	DeleteUser(id string) error
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func (u *userUseCase) GetAllUsers(payload dto.GetAllParams) ([]entity.User, error) {
	users, err := u.userRepo.GetAll(payload.Limit, payload.Offset)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (u *userUseCase) GetUser(id string) (entity.User, error) {
	user, err := u.userRepo.Get(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userUseCase) CreateUser(user entity.User) (entity.User, error) {
	user, err := u.userRepo.Create(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userUseCase) UpdateUser(id string, user entity.User) (entity.User, error) {
	user, err := u.userRepo.Update(id, user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userUseCase) DeleteUser(id string) error {
	err := u.userRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}
