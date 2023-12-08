package repository

import (
	"context"
	"database/sql"
	entity "roomate/model/entitiy"
	"time"
)

const (
	createBookDetail = `INSERT INTO booking_details ( booking_id, room_id, services_id, updated_at ) VALUES ( $1, $2, $3, $4) `

	createBooking = `INSERT INTO bookings (check_in, check_out, user_id, customer_name, status, information, updated_at) VALUES $1, $2, $3, $4, $5, $6, $7`
)

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
		row := qtx.db.QueryRowContext(ctx, createBookDetail, v.BookingID, v.RoomID, v.ServiceID.ID, v.SubTotal, time.Now())
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

	const (
		createBookDetail = `INSERT INTO booking_details ( booking_id, room_id, services_id, updated_at ) VALUES ( $1, $2, $3, $4) `

		createBooking = `INSERT INTO bookings (check_in, check_out, user_id, customer_name, status, information, total_price, updated_at) VALUES $1, $2, $3, $4, $5, $6, $7`
	)

	row := qtx.db.QueryRowContext(ctx, createBooking, checkIn, checkOut, payload.UserID, payload.CustomerName, payload.Status, payload.Information, payload.TotalPrice, bookingDetails, time.Now())
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
