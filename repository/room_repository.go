package repository

import (
	"context"
	"roomate/model/entity"
	"roomate/utils/query"
	"time"
)

func (q *Queries) GetRoomById(ctx context.Context, id string) (entity.Room, error) {
	row := q.db.QueryRowContext(ctx, query.GetRoomById, id)
	var room entity.Room
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

func (q *Queries) CreateRoom(ctx context.Context, arg entity.Room) (entity.Room, error) {
	row := q.db.QueryRowContext(ctx, query.CreateRoom, arg.ID, arg.RoomNumber, arg.RoomType,
		arg.Capacity, arg.Facility, arg.Price, arg.Status, time.Now())
	var room entity.Room
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

func (q *Queries) ListRooms(ctx context.Context) ([]entity.Room, error) {
	rows, err := q.db.QueryContext(ctx, query.ListRooms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Room []entity.Room
	for rows.Next() {
		var room entity.Room
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
		Room = append(Room, room)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return Room, nil
}

func (q *Queries) UpdateRoom(ctx context.Context, arg entity.Room) error {
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
