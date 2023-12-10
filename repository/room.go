package repository

import (
	"context"
	"roomate/model"
	"roomate/utils/query"
	"time"
)

func (q *Queries) GetRoomById(ctx context.Context, id string) (model.Rooms, error) {
	row := q.db.QueryRowContext(ctx, query.GetRoomById, id)
	var room model.Rooms
	err := row.Scan(
		&room.ID,
		&room.RoomNumber,
		&room.RoomType,
		&room.Capacity,
		&room.Facility,
		&room.Price,
		&room.Status,
		&room.CreatedAt,
		&room.UpdatedAt,
	)

	return room, err
}

func (q *Queries) CreateRoom(ctx context.Context, arg model.Rooms) (model.Rooms, error) {
	row := q.db.QueryRowContext(ctx, query.CreateRoom, arg.ID, arg.RoomNumber, arg.RoomType,
		arg.Capacity, arg.Facility, arg.Price, arg.Status, time.Now())
	var room model.Rooms
	err := row.Scan(
		&room.ID,
		&room.RoomNumber,
		&room.RoomType,
		&room.Capacity,
		&room.Facility,
		&room.Price,
		&room.Status,
		&room.CreatedAt,
		&room.UpdatedAt,
	)

	return room, err
}

func (q *Queries) ListRooms(ctx context.Context) ([]model.Rooms, error) {
	rows, err := q.db.QueryContext(ctx, query.ListRooms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var rooms []model.Rooms
	for rows.Next() {
		var room model.Rooms
		if err := rows.Scan(
			&room.ID,
			&room.RoomNumber,
			&room.RoomType,
			&room.Capacity,
			&room.Facility,
			&room.Price,
			&room.Status,
			&room.CreatedAt,
			&room.UpdatedAt,
		); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (q *Queries) UpdateRoom(ctx context.Context, arg model.Rooms) error {
	_, err := q.db.ExecContext(ctx, query.UpdateRoom,
		arg.ID, arg.RoomNumber, arg.RoomType,
		arg.Capacity, arg.Facility, arg.Price, arg.Status, time.Now(),
	)
	return err
}

func (q *Queries) DeleteRoom(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, query.DeleteRoom, id)
	return err
}
