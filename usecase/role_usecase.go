package usecase

import (
	"roomate/model/dto"
	"roomate/model/entity"
	"roomate/repository"
)

type RoleUseCase interface {
	GetAllRoles(payload dto.GetAllParams) ([]entity.Role, error)
	GetRole(id string) (entity.Role, error)
	CreateRole(role entity.Role) (entity.Role, error)
	UpdateRole(id string, role entity.Role) (entity.Role, error)
	DeleteRole(id string) error
}

type roleUseCase struct {
	roleRepo repository.RoleRepository
}

func (u *roleUseCase) GetAllRoles(payload dto.GetAllParams) ([]entity.Role, error) {
	roles, err := u.roleRepo.GetAll(payload.Limit, payload.Offset)
	if err != nil {
		return roles, err
	}

	return roles, nil
}

func (u *roleUseCase) GetRole(id string) (entity.Role, error) {
	role, err := u.roleRepo.Get(id)
	if err != nil {
		return role, err
	}

	return role, nil
}

func (u *roleUseCase) CreateRole(role entity.Role) (entity.Role, error) {
	role, err := u.roleRepo.Create(role)
	if err != nil {
		return role, err
	}

	return role, nil
}

func (u *roleUseCase) UpdateRole(id string, role entity.Role) (entity.Role, error) {
	role, err := u.roleRepo.Update(id, role)
	if err != nil {
		return role, err
	}

	return role, nil
}

func (u *roleUseCase) DeleteRole(id string) error {
	err := u.roleRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewRoleUseCase(roleRepo repository.RoleRepository) RoleUseCase {
	return &roleUseCase{
		roleRepo: roleRepo,
	}
}
