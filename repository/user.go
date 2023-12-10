package repository

import (
	"context"
	"roomate/model"
	"roomate/utils/query"
	"time"
)

func (q *Queries) GetUserById(ctx context.Context, id string) (model.User, error) {
	row := q.db.QueryRowContext(ctx, query.GetUserById, id)
	var user model.User
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

func (q *Queries) CreateUser(ctx context.Context, arg model.User) (model.User, error) {
	row := q.db.QueryRowContext(ctx, query.CreateUser, arg.Name, arg.Email, arg.Password, arg.RoleID, time.Now())
	var user model.User
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

func (q *Queries) ListUsers(ctx context.Context) ([]model.User, error) {
	rows, err := q.db.QueryContext(ctx, query.ListUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
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

func (q *Queries) UpdateUser(ctx context.Context, arg model.User) error {
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
