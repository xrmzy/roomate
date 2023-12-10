package repository

import (
	"context"
	"database/sql"
	entity "roomate/model/entitiy"
	"roomate/utils/query"
	"time"

	"github.com/jmoiron/sqlx"
)

// const (
// 	createBookDetail = `INSERT INTO booking_details ( booking_id, room_id, services_id, updated_at ) VALUES ( $1, $2, $3, $4) `

// 	createBooking = `
// 	INSERT INTO bookings (check_in, check_out, user_id, customer_name, status, information, total_price, updated_at) VALUES $1, $2, $3, $4, $5, $6, $7`

// 	updateBooking = `
// 	UPDATE bookings
// 	SET night = $2, check_in = $3, check_out = $4, customer_name = $5, status = $6, information = $7, total_price = $8, updated_at = $9
// 	WHERE id = $1
// 	`

// 	updateBookingDetail = `
// 	UPDATE booking_details
// 	SET room_id = $2, services_id = $3, updated_at = $4
// 	WHERE id = $1
// 	`

// 	getBookingByID = `SELECT
//     b.id AS booking_id,
//     b.night,
//     b.check_in,
//     b.check_out,
//     b.user_id,
//     b.customer_id,
//     b.customer_name,
//     b.status,
//     b.information,
//     b.total_price,
//     b.created_at AS booking_created_at,
//     b.updated_at AS booking_updated_at,
//     b.is_deleted AS booking_is_deleted,
//     bd.id AS booking_detail_id,
//     bd.room_id,
//     bd.services_id,
//     bd.updated_at AS detail_updated_at
// FROM
//     bookings b
// JOIN
//     booking_details bd ON
// 	b.id = bd.booking_id
// WHERE
//     b.id = $1
// `
// 	getAllBooking = `
// 	SELECT
//         id,
//         night,
//         check_in,
//         check_out,
//         user_id,
//         customer_id,
//         customer_name,
//         status,
//         information,
//         total_price,
//         created_at,
//         updated_at,
// 		is_deleted
//     FROM
//         bookings
//     WHERE
//         is_deleted = false
// `

// 	deleteBooking = `
// 	UPDATE bookings AS b
// 	SET is deleted = true
// 	FROM booking_details AS bd
// 	WHERE b.id = bd.booking_id
// 	AND b.id = $1
// 	`
// )

func (q *Queries) CreateBooking(ctx context.Context, payload entity.Booking) (entity.Booking, error) {
	var db *sql.DB
	tx, err := db.BeginTx(ctx, nil)
	qtx := q.WithTx(tx)
	if err != nil {
		return entity.Booking{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var booking entity.Booking

	checkIn, err := time.Parse(time.RFC3339, payload.CheckIn)
	if err != nil {
		return entity.Booking{}, err
	}

	checkOut, err := time.Parse(time.RFC3339, payload.CheckOut)
	if err != nil {
		return entity.Booking{}, err
	}

	var bookingDetails []entity.BookingDetail
	for _, v := range payload.BookingDetails {
		var bookingDetail entity.BookingDetail
		row := qtx.db.QueryRowContext(ctx, query.CreateBookingDetail, v.BookingID, v.RoomID, v.ServiceID.ID, v.SubTotal, time.Now())
		err = row.Scan(
			&bookingDetail.Id,
			&bookingDetail.BookingID,
			&bookingDetail.RoomID,
			&bookingDetail.ServiceID.ID,
			&bookingDetail.SubTotal,
			&bookingDetail.CreatedAt,
			&bookingDetail.UpadtedAt,
			&bookingDetail.IsDeleted,
		)
		bookingDetail.ServiceID = v.ServiceID
		bookingDetails = append(bookingDetails, bookingDetail)
		if err != nil {
			return entity.Booking{}, err
		}
	}

	row := qtx.db.QueryRowContext(ctx, query.CreateBooking, checkIn, checkOut, payload.UserID, payload.CustomerName, payload.Status, payload.Information, payload.TotalPrice, bookingDetails, time.Now())
	err = row.Scan(
		&booking.ID,
		&booking.Night,
		&booking.CheckIn,
		&booking.CheckOut,
		&booking.UserID,
		&booking.CustomerID,
		&booking.CustomerName,
		&booking.Status,
		&booking.BookingDetails,
		&booking.Information,
		&booking.TotalPrice,
		&booking.CreatedAt,
		&booking.UpadatedAt,
		&booking.IsDeleted,
	)

	if err := tx.Commit(); err != nil {
		return entity.Booking{}, err
	}
	return booking, nil
}

func (q *Queries) GetByID(ctx context.Context, bookingID string) (entity.Booking, error) {
	var booking entity.Booking
	rows, err := q.db.QueryContext(ctx, query.GetBookingByID, bookingID)
	if err != nil {
		return entity.Booking{}, err
	}
	defer rows.Close()

	bookingDetailMap := make(map[string]entity.BookingDetail)

	for rows.Next() {
		var bookingID string
		var bookingDetail entity.BookingDetail

		err := rows.Scan(
			&bookingID,
			&booking.Night,
			&booking.CheckIn,
			&booking.CheckOut,
			&booking.UserID,
			&booking.CustomerID,
			&booking.CustomerName,
			&booking.Status,
			&booking.BookingDetails,
			&booking.Information,
			&booking.TotalPrice,
			&booking.CreatedAt,
			&booking.UpadatedAt,
			&booking.IsDeleted,
			&bookingDetail.Id,
			&bookingDetail.RoomID,
			&bookingDetail.ServiceID,
			&bookingDetail.UpadtedAt,
		)
		if err != nil {
			return entity.Booking{}, err
		}
	}
	if len(bookingDetailMap) > 0 {
		for _, detail := range bookingDetailMap {
			booking.BookingDetails = append(booking.BookingDetails, detail)
		}
	}

	return booking, nil
}

func (q *Queries) GetAllBooking(ctx context.Context) ([]entity.Booking, error) {
	var bookings []entity.Booking

	rows, err := q.db.QueryContext(ctx, query.GetAllBooking)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var booking entity.Booking

		err := rows.Scan(
			&booking.ID,
			&booking.Night,
			&booking.CheckIn,
			&booking.CheckOut,
			&booking.UserID,
			&booking.CustomerID,
			&booking.CustomerName,
			&booking.Status,
			&booking.BookingDetails,
			&booking.Information,
			&booking.TotalPrice,
			&booking.CreatedAt,
			&booking.UpadatedAt,
			&booking.IsDeleted,
		)
		if err != nil {
			return []entity.Booking{}, err
		}
		bookings = append(bookings, booking)
	}
	return bookings, nil
}

func (q *Queries) DeleteBooking(ctx context.Context, db *sqlx.DB, bookingID string) error {
	_, err := q.db.ExecContext(ctx, query.DeleteBooking, bookingID)
	if err != nil {
		return err
	}
	return nil
}

func (q *Queries) UpdateBooking(ctx context.Context, booking entity.Booking) (entity.Booking, error) {
	var db *sql.DB

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return entity.Booking{}, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Update bookings table
	_, err = tx.ExecContext(ctx, query.UpdateBooking, booking.Night, booking.CheckIn, booking.CheckOut, booking.CustomerName, booking.Status, booking.Information, booking.TotalPrice, time.Now(), booking.ID)
	if err != nil {
		return entity.Booking{}, err
	}

	// Update booking_details table
	for _, detail := range booking.BookingDetails {
		_, err = tx.ExecContext(ctx, query.UpdateBookingDetail, detail.RoomID, detail.ServiceID.ID, time.Now(), detail.Id)
		if err != nil {
			return entity.Booking{}, err
		}
	}

	return booking, nil
}
