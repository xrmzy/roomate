package repository

import (
	"context"
	"roomate/model"
	"roomate/utils/query"
	"time"
)

// type ServiceRepository interface {
// 	GetService(id string) (model.Service, error)
// 	CreateService(payload model.Service) (model.Service, error)
// 	UpdateService(payload model.Service) (model.Service, error)
// 	DeleteService(id string) (model.Service, error)
// 	ListService() ([]model.Service, error)
// }

// type serviceRepository struct {
// 	db *sql.DB
// }

// func (s *serviceRepository) GetService(id string) (model.Service, error) {
// 	var service model.Service
// 	err := s.db.QueryRow(query.GetService, id).Scan(
// 		&service.Id,
// 		&service.Name,
// 		&service.Price,
// 		&service.CreatedAt,
// 		&service.UpdatedAt,
// 	)
// 	if err != nil {
// 		return model.Service{}, err
// 	}

// 	return service, nil
// }

// func (s *serviceRepository) CreateService(payload model.Service) (model.Service, error) {
// 	payload.IsDeleted = false
// 	var service model.Service
// 	err := s.db.QueryRow(query.CreateService, payload.Id, payload.Name, payload.Price, time.Now(), time.Now(), payload.IsDeleted).Scan(
// 		&service.Id,
// 		&service.Name,
// 		&service.Price,
// 		&service.CreatedAt,
// 		&service.UpdatedAt,
// 	)

// 	if err != nil {
// 		return model.Service{}, err
// 	}

// 	return service, nil
// }

// func (s *serviceRepository) UpdateService(payload model.Service) (model.Service, error) {
// 	var service model.Service
// 	err := s.db.QueryRow(query.UpdateService,
// 		payload.Id, payload.Name, payload.Price, time.Now()).
// 		Scan(
// 			&service.Id,
// 			&service.Name,
// 			&service.Price,
// 			&service.CreatedAt,
// 			&service.UpdatedAt,
// 		)

// 	if err != nil {
// 		return model.Service{}, err
// 	}

// 	return service, nil
// }

// func (s *serviceRepository) DeleteService(id string) (model.Service, error) {
// 	var service model.Service
// 	err := s.db.QueryRow(query.DeleteService, id).
// 		Scan(
// 			&service.Id,
// 			&service.Name,
// 			&service.Price,
// 			&service.CreatedAt,
// 			&service.UpdatedAt,
// 		)

// 	if err != nil {
// 		return model.Service{}, err
// 	}

// 	return service, nil
// }

// func (s *serviceRepository) ListService() ([]model.Service, error) {
// 	rows, err := s.db.Query(query.ListService)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var services []model.Service

// 	for rows.Next() {
// 		var service model.Service
// 		if err := rows.Scan(
// 			&service.Id,
// 			&service.Name,
// 			&service.Price,
// 			&service.CreatedAt,
// 			&service.UpdatedAt,
// 		); err != nil {
// 			return nil, err
// 		}
// 		services = append(services, service)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return services, nil
// }

// func NewServiceRepository(db *sql.DB) ServiceRepository {
// 	return &serviceRepository{db: db}
// }

func (s *Queries) GetService(ctx context.Context, id string) (model.Service, error) {
	var service model.Service
	err := s.db.QueryRowContext(ctx, query.GetService, id).Scan(
		&service.Id,
		&service.Name,
		&service.Price,
		&service.CreatedAt,
		&service.UpdatedAt,
	)
	if err != nil {
		return model.Service{}, err
	}

	return service, nil
}

func (s *Queries) CreateService(ctx context.Context, payload model.Service) (model.Service, error) {
	payload.IsDeleted = false
	var service model.Service
	err := s.db.QueryRowContext(ctx, query.CreateService, payload.Id, payload.Name, payload.Price, time.Now(), time.Now(), payload.IsDeleted).Scan(
		&service.Id,
		&service.Name,
		&service.Price,
		&service.CreatedAt,
		&service.UpdatedAt,
	)

	if err != nil {
		return model.Service{}, err
	}

	return service, nil
}

func (s *Queries) UpdateService(ctx context.Context, payload model.Service) error {
	_, err := s.db.ExecContext(ctx, query.UpdateService,
		payload.Name, payload.Price, time.Now(), payload.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Queries) DeleteService(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, query.DeleteService, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Queries) ListService(ctx context.Context) ([]model.Service, error) {
	rows, err := s.db.QueryContext(ctx, query.ListService)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []model.Service

	for rows.Next() {
		var service model.Service
		if err := rows.Scan(
			&service.Id,
			&service.Name,
			&service.Price,
			&service.CreatedAt,
			&service.UpdatedAt,
		); err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return services, nil
}
