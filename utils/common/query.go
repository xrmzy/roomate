package common

const (
	// User
	CreateUser = `INSERT INTO users (name, email, password, role_id, role_name, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, email, role_id, role_name, created_at, updated_at, is_deleted`
	// update user without updating password
	UpdateUser  = `UPDATE users SET name = $2, email = $3, role_id = $4, role_name = $5, updated_at = $6 WHERE id = $1 RETURNING id, name, email, role_id, role_name, created_at, updated_at, is_deleted`
	DeleteUser  = `UPDATE users SET is_deleted = true WHERE id = $1`
	GetUser     = `SELECT id, name, email, role_id, role_name, created_at, updated_at, is_deleted FROM users WHERE id = $1`
	GetAllUsers = `SELECT id, name, email, role_id, role_name, created_at, updated_at, is_deleted FROM users WHERE is_deleted = false ORDER BY id LIMIT $1 OFFSET $2`

	// Role
	CreateRole  = `INSERT INTO roles (role_name, updated_at) VALUES ($1, $2) RETURNING id, role_name, created_at, updated_at, is_deleted`
	UpdateRole  = `UPDATE roles SET role_name = $2, updated_at = $3 WHERE id = $1 RETURNING id, role_name, created_at, updated_at, is_deleted`
	DeleteRole  = `UPDATE roles SET is_deleted = true WHERE id = $1`
	GetRole     = `SELECT id, role_name, created_at, updated_at, is_deleted FROM roles WHERE id = $1`
	GetAllRoles = `SELECT id, role_name, created_at, updated_at, is_deleted FROM roles WHERE is_deleted = false ORDER BY id LIMIT $1 OFFSET $2`

	// Customer
	CreateCustomer  = `INSERT INTO customers (name, email, address, phone_number, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, email, address, phone_number, created_at, updated_at, is_deleted`
	UpdateCustomer  = `UPDATE customers SET name = $2, email = $3, address = $4, phone_number = $5, updated_at = $6 WHERE id = $1 RETURNING id, name, email, address, phone_number, created_at, updated_at, is_deleted`
	DeleteCustomer  = `UPDATE customers SET is_deleted = true WHERE id = $1`
	GetCustomer     = `SELECT id, name, email, address, phone_number, created_at, updated_at, is_deleted FROM customers WHERE id = $1`
	GetAllCustomers = `SELECT id, name, email, address, phone_number, created_at, updated_at, is_deleted FROM customers WHERE is_deleted = false ORDER BY id LIMIT $1 OFFSET $2`

	// Room
	CreateRoom  = `INSERT INTO rooms (id, room_number, room_type, capacity, facility, price, status, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, room_number, room_type, capacity, facility, price, status, created_at, updated_at, is_deleted`
	UpdateRoom  = `UPDATE rooms SET room_number = $2, room_type = $3, capacity = $4, facility = $5, price = $6, status = $7, updated_at = $8 WHERE id = $1 RETURNING id, room_number, room_type, capacity, facility, price, status, created_at, updated_at, is_deleted`
	DeleteRoom  = `UPDATE rooms SET is_deleted = true WHERE id = $1`
	GetRoom     = `SELECT id, room_number, room_type, capacity, facility, price, status, created_at, updated_at, is_deleted FROM rooms WHERE id = $1`
	GetAllRooms = `SELECT id, room_number, room_type, capacity, facility, price, status, created_at, updated_at, is_deleted FROM rooms WHERE is_deleted = false ORDER BY id LIMIT $1 OFFSET $2`

	// Service
	CreateService  = `INSERT INTO services (id, name, price, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, name, price, created_at, updated_at, is_deleted`
	UpdateService  = `UPDATE services SET name = $2, price = $3, updated_at = $4 WHERE id = $1 RETURNING id, name, price, created_at, updated_at, is_deleted`
	DeleteService  = `UPDATE services SET is_deleted = true WHERE id = $1`
	GetService     = `SELECT id, name, price, created_at, updated_at, is_deleted FROM services WHERE id = $1`
	GetAllServices = `SELECT id, name, price, created_at, updated_at, is_deleted FROM services WHERE is_deleted = false ORDER BY id LIMIT $1 OFFSET $2`

	// Booking
	CreateBooking  = `INSERT INTO bookings (night, check_in, check_out, user_id, customer_id, total_price, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, night, check_in, check_out, user_id, customer_id, is_agree, information, total_price, created_at, updated_at, is_deleted`
	UpdateBooking  = `UPDATE bookings SET is_agree = $2, information = $3, total_price = $4, updated_at = $5 WHERE id = $1 RETURNING id, night, check_in, check_out, user_id, customer_id, is_agree, information, total_price, created_at, updated_at, is_deleted`
	DeleteBooking  = `UPDATE bookings SET is_deleted = true WHERE id = $1`
	GetBooking     = `SELECT id, night, check_in, check_out, user_id, customer_id, is_agree, information, total_price, created_at, updated_at, is_deleted FROM bookings WHERE id = $1`
	GetAllBookings = `SELECT id, night, check_in, check_out, user_id, customer_id, is_agree, information, total_price, created_at, updated_at, is_deleted FROM bookings WHERE is_deleted = false ORDER BY id LIMIT $1 OFFSET $2`

	// Booking Detail
	CreateBookingDetail  = `INSERT INTO booking_details (booking_id, room_id, sub_total, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, booking_id, room_id, sub_total, created_at, updated_at, is_deleted`
	UpdateBookingDetail  = `UPDATE booking_details SET sub_total = $2, updated_at = $3 WHERE id = $1 RETURNING id, booking_id, room_id, sub_total, created_at, updated_at, is_deleted`
	DeleteBookingDetail  = `UPDATE booking_details SET is_deleted = true WHERE id = $1`
	GetAllBookingDetails = `SELECT id, booking_id, room_id, sub_total, created_at, updated_at, is_deleted FROM booking_details WHERE booking_id = $1`

	// Booking Detail Service
	CreateBookingDetailService  = `INSERT INTO booking_detail_services (booking_detail_id, service_id, service_name, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, booking_detail_id, service_id, service_name, created_at, updated_at, is_deleted`
	DeleteBookingDetailService  = `UPDATE booking_detail_services SET is_deleted = true WHERE id = $1`
	GetAllBookingDetailServices = `SELECT id, booking_detail_id, service_id, created_at, updated_at, is_deleted FROM booking_detail_services WHERE booking_detail_id = $1`
)

// custom queries
const (
	GetRoleName         = `SELECT role_name FROM roles WHERE id = $1`
	GetByEmail          = `SELECT id, role_name, password FROM users WHERE email = $1`
	UpdatePassword      = `UPDATE users SET password = $2 WHERE id = $1 RETURNING id, name, email, role_id, role_name, created_at, updated_at, is_deleted`
	UpdateBookingStatus = `UPDATE bookings SET is_agree = $2, information = $3 WHERE id = $1 RETURNING id, night, check_in, check_out, user_id, customer_id, is_agree, information, total_price, created_at, updated_at, is_deleted`
	GetBookingOneDay    = `SELECT id, check_in, check_out, user_id, customer_id, is_agree, information, total_price FROM bookings WHERE check_in = $1`
	GetBookingOneMonth  = `SELECT id, check_in, check_out, user_id, customer_id, is_agree, information, total_price FROM bookings WHERE EXTRACT(MONTH FROM "check_in") = $1 AND EXTRACT(YEAR FROM "check_in") = $2`
	GetBookingOneYear   = `SELECT id, check_in, check_out, user_id, customer_id, is_agree, information, total_price FROM bookings WHERE EXTRACT(YEAR FROM "check_in") = $1`
)
