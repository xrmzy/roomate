package repository

import (
	"database/sql"
	"roomate/model/entity"
	query "roomate/utils/common"
	"time"
)

type RoleRepository interface {
	Get(id string) (entity.Role, error)
	GetAll(limit, offset int) ([]entity.Role, error)
	Create(role entity.Role) (entity.Role, error)
	Update(id string, role entity.Role) (entity.Role, error)
	Delete(id string) error
}

type roleRepository struct {
	db *sql.DB
}

func (r *roleRepository) Get(id string) (entity.Role, error) {
	var role entity.Role
	err := r.db.QueryRow(query.GetRole, id).
		Scan(
			&role.Id,
			&role.RoleName,
			&role.CreatedAt,
			&role.UpdatedAt,
			&role.IsDeleted)

	if err != nil {
		return role, err
	}

	return role, nil
}

func (r *roleRepository) GetAll(limit, offset int) ([]entity.Role, error) {
	var roles []entity.Role

	rows, err := r.db.Query(query.GetAllRoles, limit, offset)
	if err != nil {
		return roles, err
	}

	for rows.Next() {
		var role entity.Role
		err := rows.Scan(
			&role.Id,
			&role.RoleName,
			&role.CreatedAt,
			&role.UpdatedAt,
			&role.IsDeleted)

		if err != nil {
			return roles, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}

func (r *roleRepository) Create(role entity.Role) (entity.Role, error) {
	err := r.db.QueryRow(query.CreateRole,
		role.RoleName,
		time.Now(),
	).Scan(
		&role.Id,
		&role.RoleName,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.IsDeleted)

	if err != nil {
		return role, err
	}

	return role, nil
}

func (r *roleRepository) Update(id string, role entity.Role) (entity.Role, error) {
	err := r.db.QueryRow(query.UpdateRole,
		id,
		role.RoleName,
		time.Now(),
	).Scan(
		&role.Id,
		&role.RoleName,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.IsDeleted)

	if err != nil {
		return role, err
	}

	return role, nil
}

func (r *roleRepository) Delete(id string) error {
	_, err := r.db.Exec(query.DeleteRole, id)
	if err != nil {
		return err
	}

	return nil
}

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}
