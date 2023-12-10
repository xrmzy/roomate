package query

const (
	CreateCustomer = `INSERT INTO customers 
	(name, email, address, phone_number, updated_at,) 
	VALUES $1, $2, $3, $4, $5`

	UpdateCustomer = `UPDATE customers 
	SET 
		name = $2,
		email = $3,
		address = $4,
		phone_number = $5,
		updated_at = $6
	WHERE id = $1
	RETURNING id, name, email, address, phone_number, created_at, updated_at
	`

	GetAllCustomer = `SELECT 
		(name, email, address, phone_number, created_at, updated_at) 
	FROM 
		customers
	WHERE is_deleted = false`

	GetCustomerByID = `SELECT 
	(name, email, address, phone_number, created_at, updated_at) 
FROM 
	customers
WHERE id = $1 AND is_deleted = false`

	DeleteCustomer = `UPDATE customers
	SET
		is_deleted = true
	WHERE id = $1 
	`
)
