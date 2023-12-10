package query

const (
	GetService    = `SELECT id, name, price, created_at, updated_at FROM services WHERE id = $1 AND is_deleted = false`
	CreateService = `INSERT INTO services (id, name, price, created_at, updated_at) VALUES ($1,$2,$3,$4,$5) RETURNING id, name, price, created_at, updated_at`
	UpdateService = `UPDATE services SET name = $2, price = $3, updated_at = $4 WHERE id = $1 RETURNING id, name, price, created_at, updated_at`
	DeleteService = `UPDATE services SET is_deleted = TRUE WHERE id = $1`
	ListService   = `SELECT id, name, price, created_at, updated_at FROM services WHERE is_deleted = false`
)
