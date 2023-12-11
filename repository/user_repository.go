package repository

import (
	"database/sql"
	"roomate/model/entity"
	query "roomate/utils/common"
	"time"
)

type UserRepository interface {
	Get(id string) (entity.User, error)
	GetAll(limit, offset int) ([]entity.User, error)
	Create(user entity.User) (entity.User, error)
	Update(id string, user entity.User) (entity.User, error)
	Delete(id string) error
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Get(id string) (entity.User, error) {
	var user entity.User
	err := u.db.QueryRow(query.GetUser, id).
		Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.RoleId,
			&user.RoleName,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.IsDeleted)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) GetAll(limit, offset int) ([]entity.User, error) {
	var users []entity.User
	rows, err := u.db.Query(query.GetAllUsers, limit, offset)

	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user entity.User
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.RoleId,
			&user.RoleName,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.IsDeleted)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *userRepository) Create(user entity.User) (entity.User, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return user, err
	}

	err = tx.QueryRow(query.CreateUser,
		user.Name,
		user.Email,
		user.Password,
		user.RoleId,
		time.Now(),
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.RoleId,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsDeleted)

	if err != nil {
		tx.Rollback()
		return user, err
	}

	err = tx.QueryRow(query.GetRoleName, user.RoleId).Scan(&user.RoleName)
	if err != nil {
		tx.Rollback()
		return user, err
	}

	_, err = tx.Exec(query.UpdateRoleName, user.Id, user.RoleName)
	if err != nil {
		tx.Rollback()
		return user, err
	}

	if err := tx.Commit(); err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) Update(id string, user entity.User) (entity.User, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return user, err
	}

	err = tx.QueryRow(query.UpdateUser,
		id,
		user.Name,
		user.Email,
		user.Password,
		user.RoleId,
		time.Now(),
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.RoleId,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsDeleted)

	if err != nil {
		return user, tx.Rollback()
	}

	err = tx.QueryRow(query.GetRoleName, user.RoleId).Scan(&user.RoleName)
	if err != nil {
		return user, tx.Rollback()
	}

	_, err = tx.Exec(query.UpdateRoleName, user.Id, user.RoleName)
	if err != nil {
		return user, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) Delete(id string) error {
	_, err := u.db.Exec(query.DeleteUser, id)

	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
