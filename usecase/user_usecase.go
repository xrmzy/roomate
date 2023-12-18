package usecase

import (
	"errors"
	"fmt"
	"roomate/model/dto"
	"roomate/model/entity"
	"roomate/repository"
	"roomate/utils/common"
)

type UserUseCase interface {
	GetAllUsers(payload dto.GetAllParams) ([]entity.User, error)
	GetUser(id string) (entity.User, error)
	GetByEmailPassword(email, password string) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(id string, user entity.User) (entity.User, error)
	UpdatePassword(id, password string) (entity.User, error)
	DeleteUser(id string) error
}

type userUseCase struct {
	userRepo repository.UserRepository
	roleUc   RoleUseCase
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
		return user, fmt.Errorf("user with ID %s not found", id)
	}

	return user, nil
}

func (u *userUseCase) GetByEmailPassword(email, password string) (entity.User, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return user, err
	}

	// compare password
	err = common.ComparePasswordHash(user.Password, password)
	if err != nil {
		return user, err
	}

	user.Password = ""
	return user, nil
}

func (u *userUseCase) CreateUser(user entity.User) (entity.User, error) {
	// check if user already exist
	userCheck, err := u.userRepo.GetByEmail(user.Email)
	if err != nil {
		return user, nil
	}

	// if user already exist, return error with message "email already exist"
	if userCheck.Id != "" {
		return user, errors.New("email already exist")
	}

	role, err := u.roleUc.GetRole(user.RoleId)
	if err != nil {
		return user, err
	}

	user.RoleName = role.RoleName

	hashedPassword, err := common.GeneratePasswordHash(user.Password)
	if err != nil {
		return user, err
	}

	user.Password = hashedPassword

	user, err = u.userRepo.Create(user)
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

func (u *userUseCase) UpdatePassword(id, password string) (entity.User, error) {
	hashedPassword, err := common.GeneratePasswordHash(password)
	if err != nil {
		return entity.User{}, err
	}

	user, err := u.userRepo.UpdatePassword(id, hashedPassword)
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

func NewUserUseCase(userRepo repository.UserRepository, roleUc RoleUseCase) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
		roleUc:   roleUc,
	}
}
