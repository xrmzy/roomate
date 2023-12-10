package query

const (
	GetRole     = `SELECT id, role_name, created_at, updated_at FROM roles WHERE id = $1 AND is_deleted = false`
	CreateRole  = `INSERT INTO roles (role_name, created_at, updated_at) VALUES ($1,$2,$3) RETURNING id, role_name, created_at, updated_at`
	UpdatedRole = `UPDATE roles set role_name = $2, updated_at = $3 WHERE id = $1 RETURNING id, role_name, created_at, updated_at`
	DeleteRole  = `UPDATE roles SET is_deleted = TRUE WHERE id = $1`
	ListRole    = `SELECT id, role_name, created_at, updated_at FROM roles WHERE is_deleted = false`
)
