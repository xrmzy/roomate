package manager

import "roomate/usecase"

type UseCaseManager interface {
	UserUsecase() usecase.UserUseCase
	RoleUsecase() usecase.RoleUseCase
	CustomerUseCase() usecase.CustomerUseCase
	RoomUseCase() usecase.RoomUseCase
	ServiceUseCase() usecase.ServiceUseCase
	BookingUseCase() usecase.BookingUsecase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) UserUsecase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo(), u.RoleUsecase())
}

func (u *useCaseManager) RoleUsecase() usecase.RoleUseCase {
	return usecase.NewRoleUseCase(u.repo.RoleRepo())
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repo.CustomerRepo())
}

func (u *useCaseManager) RoomUseCase() usecase.RoomUseCase {
	return usecase.NewRoomUseCase(u.repo.RoomRepo())
}

func (u *useCaseManager) ServiceUseCase() usecase.ServiceUseCase {
	return usecase.NewServiceUseCase(u.repo.ServiceRepo())
}

func (u *useCaseManager) BookingUseCase() usecase.BookingUsecase {
	return usecase.NewBookingUseCase(u.repo.BookingRepo(), u.RoomUseCase(), u.ServiceUseCase())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
