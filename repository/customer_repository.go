package repository

import (
	"database/sql"
	"roomate/model/entity"
	query "roomate/utils/common"
	"time"
)

type CustomerRepository interface {
	Get(id string) (entity.Customer, error)
	GetAll(limit, offset int) ([]entity.Customer, error)
	Create(customer entity.Customer) (entity.Customer, error)
	Update(id string, customer entity.Customer) (entity.Customer, error)
	Delete(id string) error
}

type customerRepository struct {
	db *sql.DB
}

func (r *customerRepository) Get(id string) (entity.Customer, error) {
	var customer entity.Customer
	err := r.db.QueryRow(query.GetCustomer, id).
		Scan(
			&customer.Id,
			&customer.Name,
			&customer.Email,
			&customer.Address,
			&customer.PhoneNumber,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *customerRepository) GetAll(limit, offset int) ([]entity.Customer, error) {
	var customers []entity.Customer

	rows, err := r.db.Query(query.GetAllCustomers, limit, offset)
	if err != nil {
		return customers, err
	}

	for rows.Next() {
		var customer entity.Customer
		err := rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.Email,
			&customer.Address,
			&customer.PhoneNumber,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		)

		if err != nil {
			return customers, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (r *customerRepository) Create(customer entity.Customer) (entity.Customer, error) {
	err := r.db.QueryRow(query.CreateCustomer,
		customer.Name,
		customer.Email,
		customer.Address,
		customer.PhoneNumber,
		time.Now().Truncate(time.Second),
	).Scan(
		&customer.Id,
		&customer.Name,
		&customer.Email,
		&customer.Address,
		&customer.PhoneNumber,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *customerRepository) Update(id string, customer entity.Customer) (entity.Customer, error) {
	err := r.db.QueryRow(query.UpdateCustomer,
		id,
		customer.Name,
		customer.Email,
		customer.Address,
		customer.PhoneNumber,
		time.Now().Truncate(time.Second),
	).Scan(
		&customer.Id,
		&customer.Name,
		&customer.Email,
		&customer.Address,
		&customer.PhoneNumber,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *customerRepository) Delete(id string) error {
	_, err := r.db.Exec(query.DeleteCustomer, id)
	if err != nil {
		return err
	}

	return nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}
