package query

const (
	CreateBookingDetail = `INSERT INTO booking_details ( booking_id, room_id, services_id, updated_at ) VALUES ( $1, $2, $3, $4) `

	CreateBooking = `
	INSERT INTO bookings (check_in, check_out, user_id, customer_name, status, information, total_price, updated_at) VALUES $1, $2, $3, $4, $5, $6, $7`

	UpdateBooking = `
	UPDATE bookings
	SET night = $2, check_in = $3, check_out = $4, customer_name = $5, status = $6, information = $7, total_price = $8, updated_at = $9
	WHERE id = $1
	`

	UpdateBookingDetail = `
	UPDATE booking_details
	SET room_id = $2, services_id = $3, updated_at = $4
	WHERE id = $1
	`

	GetBookingByID = `SELECT
    b.id AS booking_id,
    b.night,
    b.check_in,
    b.check_out,
    b.user_id,
    b.customer_id,
    b.customer_name,
    b.status,
    b.information,
    b.total_price,
    b.created_at AS booking_created_at,
    b.updated_at AS booking_updated_at,
    b.is_deleted AS booking_is_deleted,
    bd.id AS booking_detail_id,
    bd.room_id,
    bd.services_id,
    bd.updated_at AS detail_updated_at
FROM
    bookings b
JOIN
    booking_details bd ON 
	b.id = bd.booking_id
WHERE
    b.id = $1
`
	GetAllBooking = `
	SELECT
        id,
        night,
        check_in,
        check_out,
        user_id,
        customer_id,
        customer_name,
        status,
        information,
        total_price,
        created_at,
        updated_at,
		is_deleted
    FROM
        bookings
    WHERE
        is_deleted = false
`

	DeleteBooking = `
	UPDATE bookings AS b
	SET is deleted = true
	FROM booking_details AS bd
	WHERE b.id = bd.booking_id
	AND b.id = $1
	`
)
