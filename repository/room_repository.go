package repository

import (
	"database/sql"
	"roomate/model/entity"
	query "roomate/utils/common"
	"time"
)

type RoomRepository interface {
	Get(id string) (entity.Room, error)
	GetAll(limit, offset int) ([]entity.Room, error)
	Create(room entity.Room) (entity.Room, error)
	Update(id string, room entity.Room) (entity.Room, error)
	Delete(id string) error
}

type roomRepository struct {
	db *sql.DB
}

func (r *roomRepository) Get(id string) (entity.Room, error) {
	var room entity.Room
	err := r.db.QueryRow(query.GetRoom, id).
		Scan(
			&room.Id,
			&room.RoomNumber,
			&room.RoomType,
			&room.Capacity,
			&room.Facility,
			&room.Price,
			&room.Status,
			&room.CreatedAt,
			&room.UpdatedAt,
			&room.IsDeleted)

	if err != nil {
		return room, err
	}

	return room, nil
}

func (r *roomRepository) GetAll(limit, offset int) ([]entity.Room, error) {
	var rooms []entity.Room

	rows, err := r.db.Query(query.GetAllRooms, limit, offset)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room entity.Room
		err := rows.Scan(
			&room.Id,
			&room.RoomNumber,
			&room.RoomType,
			&room.Capacity,
			&room.Facility,
			&room.Price,
			&room.Status,
			&room.CreatedAt,
			&room.UpdatedAt,
			&room.IsDeleted)

		if err != nil {
			return rooms, err
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r *roomRepository) Create(room entity.Room) (entity.Room, error) {
	room.Id = query.GenerateRoomID("R") // Membuat ID dengan prefiks "R"

	err := r.db.QueryRow(query.CreateRoom,
		room.Id,
		room.RoomNumber,
		room.RoomType,
		room.Capacity,
		room.Facility,
		room.Price,
		"available",
		time.Now(),
	).Scan(
		&room.Id,
		&room.RoomNumber,
		&room.RoomType,
		&room.Capacity,
		&room.Facility,
		&room.Price,
		&room.Status,
		&room.CreatedAt,
		&room.UpdatedAt,
		&room.IsDeleted)

	if err != nil {
		return room, err
	}

	return room, nil
}

func (r *roomRepository) Update(id string, room entity.Room) (entity.Room, error) {
	err := r.db.QueryRow(query.UpdateRoom,
		id,
		room.RoomNumber,
		room.RoomType,
		room.Capacity,
		room.Facility,
		room.Price,
		room.Status,
		time.Now(),
	).Scan(
		&room.Id,
		&room.RoomNumber,
		&room.RoomType,
		&room.Capacity,
		&room.Facility,
		&room.Price,
		&room.Status,
		&room.CreatedAt,
		&room.UpdatedAt,
		&room.IsDeleted)

	if err != nil {
		return room, err
	}

	return room, nil
}

func (r *roomRepository) Delete(id string) error {
	_, err := r.db.Exec(query.DeleteRoom, id)
	if err != nil {
		return err
	}

	return nil
}

func NewRoomRepository(db *sql.DB) RoomRepository {
	return &roomRepository{
		db: db,
	}
}
