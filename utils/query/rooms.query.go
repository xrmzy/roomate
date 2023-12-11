package query

const (
	CreateRoom  = `INSERT INTO rooms (id, room_number, room_type, capacity, facility, price, status, updated_at) VALUE ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, room_number, room_type, capacity, facility, price, status, created_at, updated_at;`
	GetRoomById = `SELECT id, room_number, room_type, capacity, facility, price, status, created_at, updated_at FROM rooms WHERE id = $1 AND is_deleted = false;`
	ListRooms   = `SELECT id, room_number, room_type, capacity, facility, price, status, created_at, updated_at FROM users WHERE is_deleted = false;`
	UpdateRoom  = `UPDATE users SET room_number = $2, room_type = $3, capacity = $4, facility = $5, price = $6, status = $7, updated_at = $8 WHERE is_deleted = false AND id = $1; `
	DeleteRoom  = `UPDATE rooms SET is_deleted = true WHERE id = $1`
)
