package query

const (
	CreateUser  = `INSERT INTO users (name, email, password, role_id, updated_at) VALUE ($1, $2, $3, $4, $5) RETURNING id, name, email, password, role_id, role_name, created_at, updated_at;`
	GetUserById = `SELECT id, name, email, role_id, role_name, created_at, updated_at FROM users WHERE id = $1 AND is_deleted = false;`
	ListUsers   = `SELECT id, name, email, role_id, role_name, created_at, updated_at FROM users WHERE is_deleted = false;`
	UpdateUser  = `UPDATE users SET name = $2, email = $3, password = $4, role_id = $5, updated_at = $6 WHERE is_deleted = false AND id = $1; `
	DeleteUser  = `UPDATE users SET is_deleted = true WHERE id = $1`
)
