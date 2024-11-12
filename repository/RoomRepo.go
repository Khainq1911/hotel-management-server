package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"fmt"
	"time"
)

type RoomRepo interface {
	AddRoomRepo(ctx context.Context, room model.AddRoom) error
	UpdateRoomRepo(ctx context.Context, roomId string, room model.UpdateRoom) error
	DeleteRoomRepo(ctx context.Context, roomId string, Room model.DeleteRoom) error
	ViewListRoomRepo(ctx context.Context) ([]model.Room, error)
	ViewDetailRoomRepo(ctx context.Context, roomId string) ([]model.Room, error)
}

type RoomSql struct {
	Sql *database.Sql
}

func NewRoomRepo(sql *database.Sql) RoomRepo {
	return &RoomSql{
		Sql: sql,
	}
}

// add type room
func (db *RoomSql) AddRoomRepo(ctx context.Context, room model.AddRoom) error {
	query := `insert into room (
		room_name,
		type_id, 
		floor,
		booking_status,
		price_per_hour,
		price_per_day,
		cleaning_status,
		check_in_time,
		check_out_time,
		current_guest,
		note, 
		created_at,
		created_by, 
		updated_at,
		updated_by) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`

	current := time.Now()

	if _, err := db.Sql.Db.Exec(query,
		room.RoomName,
		room.TypeID,
		room.Floor,
		room.BookingStatus,
		room.PricePerHour,
		room.PricePerDay,
		room.CleaningStatus,
		room.CheckInTime,
		room.CheckOutTime,
		room.CurrentGuest,
		room.Note,
		current,
		room.CreatedBy,
		current,
		room.UpdatedBy); err != nil {
		return err
	}

	return nil
}

// view list room
func (db *RoomSql) ViewListRoomRepo(ctx context.Context) ([]model.Room, error) {
	data := []model.Room{}

	query := `SELECT 
    room.room_id,
    room.room_name,
    room.floor,
    room.booking_status,
    room.price_per_day,
	room.price_per_hour,
    room.discount,
    room.cleaning_status,
    room.check_in_time,
    room.check_out_time,
    room.current_guest,
    room.note,
    room.created_at,
    room.updated_at,
    typeroom.type_id,
    typeroom.type_name,
    typeroom.description,
    typeroom.max_occupancy,
    typeroom.room_size
FROM 
    room
JOIN 
    typeroom ON room.type_id = typeroom.type_id `

	if err := db.Sql.Db.Select(&data, query); err != nil {
		return []model.Room{}, err
	}

	return data, nil
}

// view detail room
func (db *RoomSql) ViewDetailRoomRepo(ctx context.Context, roomId string) ([]model.Room, error) {
	data := []model.Room{}
	query := `SELECT 
    room.room_id,
    room.room_name,
    room.floor,
    room.booking_status,
    room.price_per_day,
	room.price_per_hour,
    room.discount,
    room.cleaning_status,
    room.check_in_time,
    room.check_out_time,
    room.current_guest,
    room.note,
    room.created_at,
    room.updated_at,
    typeroom.type_id,
    typeroom.type_name,
    typeroom.description,
    typeroom.max_occupancy,
    typeroom.room_size
FROM 
    room
JOIN 
    typeroom ON room.type_id = typeroom.type_id where room_id = $1`

	if err := db.Sql.Db.Select(&data, query, roomId); err != nil {
		return []model.Room{}, err
	}

	return data, nil
}

// update room
func (db *RoomSql) UpdateRoomRepo(ctx context.Context, roomId string, room model.UpdateRoom) error {
	query := `UPDATE room SET
		room_name = $1,
		type_id = $2, 
		price_per_day = $3,
		price_per_hour = $4,
		updated_at = $5,
		updated_by = $6,
		discount = $7
	WHERE room_id = $8`

	current := time.Now()

	if _, err := db.Sql.Db.Exec(query,
		room.RoomName,
		room.TypeID,
		room.PricePerDay,
		room.PricePerHour,
		current,
		room.UpdatedBy,
		room.Discount,
		roomId); err != nil {
		return err
	}

	return nil
}

// delete room
func (db *RoomSql) DeleteRoomRepo(ctx context.Context, roomId string, Room model.DeleteRoom) error {
	query := `update room
	set deletetime = $1,
		deleteby = $2
		is_deleted = $3
	where room_id = $4`

	current := time.Now()
	isDeleted := false

	result, err := db.Sql.Db.Exec(query, current, Room.DeleteBy, isDeleted, roomId)
	if err != nil {
		return err
	}

	rowwAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowwAffected == 0 {
		return fmt.Errorf("column is not exist")
	}

	return nil

}
