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
	GetByEmail(email string) (entity.User, error)
	Create(user entity.User) (entity.User, error)
	Update(id string, user entity.User) (entity.User, error)
	UpdatePassword(id, password string) (entity.User, error)
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
			&user.UpdatedAt)

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
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.RoleId,
			&user.RoleName,
			&user.CreatedAt,
			&user.UpdatedAt)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *userRepository) GetByEmail(email string) (entity.User, error) {
	var user entity.User
	err := u.db.QueryRow(query.GetByEmail, email).
		Scan(
			&user.Id,
			&user.RoleName,
			&user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) Create(user entity.User) (entity.User, error) {
	err := u.db.QueryRow(query.CreateUser,
		user.Name,
		user.Email,
		user.Password,
		user.RoleId,
		user.RoleName,
		time.Now().Truncate(time.Second),
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.RoleId,
		&user.RoleName,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) Update(id string, user entity.User) (entity.User, error) {
	err := u.db.QueryRow(query.UpdateUser,
		id,
		user.Name,
		user.Email,
		user.RoleId,
		user.RoleName,
		time.Now().Truncate(time.Second),
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.RoleId,
		&user.RoleName,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) UpdatePassword(id, password string) (entity.User, error) {
	var user entity.User
	err := u.db.QueryRow(query.UpdatePassword, id, password).
		Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.RoleId,
			&user.RoleName,
			&user.CreatedAt,
			&user.UpdatedAt)

	if err != nil {
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
