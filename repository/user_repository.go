package repository

// import (
// 	"database/sql"
// 	"roomate/model"
// )

// type UserRepository interface {
// 	Get(id string) (model.User, error)
// }

// type userRepository struct {
// 	db *sql.DB
// }

// const (

// )

// func (u *userRepository) Get(id string) (model.User, error) {
// 	var user model.User
// 	err := u.db.QueryRow(`SELECT id, name, email, password, role_id, role_name, created_at, updated_at WHERE id = $1 AND is_deleted = false`, id).Scan(
// 		&user.ID,
// 		&user.Name,
// 		&user.Email,
// 		&user.Password,
// 		&user.RoleID,
// 		&user.RoleName,
// 		&user.CreatedAt,
// 		&user.UpdatedAt,
// 	)

// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	return user, nil
// }

// func NewUserRepository(db *sql.DB) UserRepository {
// 	return &userRepository{db: db}
// }


