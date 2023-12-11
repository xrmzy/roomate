package repository

import (
	"context"

	"roomate/model/entity"
	"roomate/utils/query"
	"time"
)

func (q *Queries) GetUserById(ctx context.Context, id string) (entity.User, error) {
	row := q.db.QueryRowContext(ctx, query.GetUserById, id)
	var user entity.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.RoleID,
		&user.RoleName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}

func (q *Queries) CreateUser(ctx context.Context, arg entity.User) (entity.User, error) {
	row := q.db.QueryRowContext(ctx, query.CreateUser, arg.Name, arg.Email, arg.Password, arg.RoleID, time.Now())
	var user entity.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.RoleID,
		&user.RoleName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}

func (q *Queries) ListUsers(ctx context.Context) ([]entity.User, error) {
	rows, err := q.db.QueryContext(ctx, query.ListUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.RoleID,
			&user.RoleName,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (q *Queries) UpdateUser(ctx context.Context, arg entity.User) error {
	_, err := q.db.ExecContext(ctx, query.UpdateUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.RoleID,
		time.Now(),
	)
	return err
}

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, query.DeleteUser, id)
	return err
}
