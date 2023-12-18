package repository

import (
	"database/sql"
	"roomate/model/entity"
	query "roomate/utils/common"
	"time"
)

type ServiceRepository interface {
	Get(id string) (entity.Service, error)
	GetAll(limit, offset int) ([]entity.Service, error)
	Create(service entity.Service) (entity.Service, error)
	Update(id string, service entity.Service) (entity.Service, error)
	Delete(id string) error
}

type serviceRepository struct {
	db *sql.DB
}

func (r *serviceRepository) Get(id string) (entity.Service, error) {
	var service entity.Service
	err := r.db.QueryRow(query.GetService, id).
		Scan(
			&service.Id,
			&service.Name,
			&service.Price,
			&service.CreatedAt,
			&service.UpdatedAt,
		)

	if err != nil {
		return service, err
	}

	return service, nil
}

func (r *serviceRepository) GetAll(limit, offset int) ([]entity.Service, error) {
	var services []entity.Service

	rows, err := r.db.Query(query.GetAllServices, limit, offset)
	if err != nil {
		return services, err
	}

	for rows.Next() {
		var service entity.Service
		err := rows.Scan(
			&service.Id,
			&service.Name,
			&service.Price,
			&service.CreatedAt,
			&service.UpdatedAt,
		)

		if err != nil {
			return services, err
		}

		services = append(services, service)
	}

	return services, nil
}

func (r *serviceRepository) Create(service entity.Service) (entity.Service, error) {
	err := r.db.QueryRow(query.CreateService,
		service.Id,
		service.Name,
		service.Price,
		time.Now().Truncate(time.Second),
	).Scan(
		&service.Id,
		&service.Name,
		&service.Price,
		&service.CreatedAt,
		&service.UpdatedAt,
	)

	if err != nil {
		return service, err
	}

	return service, nil
}

func (r *serviceRepository) Update(id string, service entity.Service) (entity.Service, error) {
	err := r.db.QueryRow(query.UpdateService,
		id,
		service.Name,
		service.Price,
		time.Now().Truncate(time.Second),
	).Scan(
		&service.Id,
		&service.Name,
		&service.Price,
		&service.CreatedAt,
		&service.UpdatedAt,
	)

	if err != nil {
		return service, err
	}

	return service, nil
}

func (r *serviceRepository) Delete(id string) error {
	_, err := r.db.Exec(query.DeleteService, id)
	if err != nil {
		return err
	}

	return nil
}

func NewServiceRepository(db *sql.DB) ServiceRepository {
	return &serviceRepository{
		db: db,
	}
}
