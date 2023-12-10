package repository

import (
	"context"
	entity "roomate/model/entity"
	"roomate/utils/query"
	"time"
)

func (s *Queries) GetRole(ctx context.Context, id int) (entity.Role, error) {
	var role entity.Role
	err := s.db.QueryRowContext(ctx, query.GetRole, id).Scan(
		&role.ID,
		&role.RoleName,
		&role.CreatedAt,
		&role.UpdatedAt,
	)
	if err != nil {
		return entity.Role{}, err
	}

	return role, nil
}

func (s *Queries) CreateRole(ctx context.Context, payload entity.Role) (entity.Role, error) {
	payload.IsDeleted = false
	var role entity.Role
	err := s.db.QueryRowContext(ctx, query.CreateRole, payload.RoleName, time.Now(), time.Now()).Scan(
		&role.ID,
		&role.RoleName,
		&role.CreatedAt,
		&role.UpdatedAt,
	)

	if err != nil {
		return entity.Role{}, err
	}

	return role, nil
}

func (s *Queries) UpdateRole(ctx context.Context, payload entity.Role) error {
	_, err := s.db.ExecContext(ctx, query.UpdatedRole, payload.ID,
		payload.RoleName, time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (s *Queries) DeleteRole(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, query.DeleteRole, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Queries) ListRole(ctx context.Context) ([]entity.Role, error) {
	rows, err := s.db.QueryContext(ctx, query.ListRole)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []entity.Role

	for rows.Next() {
		var role entity.Role
		if err := rows.Scan(
			&role.ID,
			&role.RoleName,
			&role.CreatedAt,
			&role.UpdatedAt,
		); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}
