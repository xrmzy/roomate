package repository

import (
	"context"
	"roomate/model/entity"
	"roomate/utils/query"
	"time"
)

// const (
// 	createCustomer = `INSERT INTO customers
// 	(name, email, address, phone_number, updated_at,)
// 	VALUES $1, $2, $3, $4, $5`

// 	updateCustomer = `UPDATE customers
// 	SET
// 		name = $2,
// 		email = $3,
// 		address = $4,
// 		phone_number = $5,
// 		updated_at = $6
// 	WHERE id = $1
// 	RETURNING id, name, email, address, phone_number, created_at, updated_at
// 	`

// 	getAllCustomer = `SELECT
// 		(name, email, address, phone_number, created_at, updated_at)
// 	FROM
// 		customers
// 	WHERE is_deleted = false`

// 	getCustomerByID = `SELECT
// 	(name, email, address, phone_number, created_at, updated_at)
// FROM
// 	customers
// WHERE id = $1 AND is_deleted = false`

// 	deleteCustomer = `UPDATE customers
// 	SET
// 		is_deleted = true
// 	WHERE id = $1
// 	`
// )

func (q *Queries) CreateCustomer(ctx context.Context, payload entity.Customer) (entity.Customer, error) {
	payload.IsDeleted = false
	var customer entity.Customer
	err := q.db.QueryRowContext(ctx, query.CreateCustomer, payload.ID, payload.Name, payload.Email, payload.Address, payload.PhoneNumber, time.Now(), time.Now(), payload.IsDeleted).Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.Address,
		&customer.PhoneNumber,
		&customer.CreatedAt,
		&customer.UpdatedAt,
		&customer.CreatedAt,
		&customer.IsDeleted,
	)

	if err != nil {
		return entity.Customer{}, err
	}

	return customer, nil
}

func (q *Queries) UpdateCustomer(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	_, err := q.db.ExecContext(ctx, query.UpdateCustomer, customer.Name, customer.Email, customer.Address, customer.PhoneNumber, time.Now(), customer.ID)

	if err != nil {
		return entity.Customer{}, err
	}
	return customer, nil
}

func (q *Queries) GetAllCustomer(ctx context.Context) ([]entity.Customer, error) {
	var customers []entity.Customer

	rows, err := q.db.QueryContext(ctx, query.GetAllCustomer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer entity.Customer
		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Email,
			&customer.Address,
			&customer.PhoneNumber,
			&customer.CreatedAt,
			&customer.UpdatedAt,
			&customer.CreatedAt,
			&customer.IsDeleted,
		)
		if err != nil {
			return []entity.Customer{}, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (q *Queries) GetCustomerByID(ctx context.Context, customerID string) (entity.Customer, error) {
	var customer entity.Customer
	err := q.db.QueryRowContext(ctx, query.GetCustomerByID, customerID).Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.Address,
		&customer.PhoneNumber,
		&customer.CreatedAt,
		&customer.UpdatedAt,
		&customer.CreatedAt,
		&customer.IsDeleted,
	)
	if err != nil {
		return entity.Customer{}, err
	}
	return customer, nil
}

func (q *Queries) DeleteCustomer(ctx context.Context, customerID string) error {
	_, err := q.db.ExecContext(ctx, query.DeleteCustomer, customerID)
	if err != nil {
		return err
	}
	return nil
}
