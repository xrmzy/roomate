package manager

import "roomate/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	RoleRepo() repository.RoleRepository
	CustomerRepo() repository.CustomerRepository
	RoomRepo() repository.RoomRepository
	ServiceRepo() repository.ServiceRepository
	BookingRepo() repository.BookingRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func (r *repoManager) RoleRepo() repository.RoleRepository {
	return repository.NewRoleRepository(r.infra.Conn())
}

func (r *repoManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.Conn())
}

func (r *repoManager) RoomRepo() repository.RoomRepository {
	return repository.NewRoomRepository(r.infra.Conn())
}

func (r *repoManager) ServiceRepo() repository.ServiceRepository {
	return repository.NewServiceRepository(r.infra.Conn())
}

func (r *repoManager) BookingRepo() repository.BookingRepository {
	return repository.NewBookingRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
