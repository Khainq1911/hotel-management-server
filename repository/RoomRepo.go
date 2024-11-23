package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"fmt"
	"time"
)

type RoomRepo interface {
	AddRoomRepo(ctx context.Context, room model.AddRoom, employee_id string) error
	UpdateRoomRepo(ctx context.Context, roomId string, room model.UpdateRoom, employeeID string) error
	DeleteRoomRepo(ctx context.Context, roomId string, employeeID string) error
	ViewListRoomRepo(ctx context.Context) ([]model.Room, error)
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
func (db *RoomSql) AddRoomRepo(ctx context.Context, room model.AddRoom, employee_id string) error {
	query := `INSERT INTO room (
		room_name,
		floor,
		pricing_id,
		booking_status,
		cleaning_status,
		price_temporary,
		created_at,
		created_by,
		updated_at,
		updated_by
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	current := time.Now()

	_, err := db.Sql.Db.Exec(query,
		room.RoomName,
		room.Floor,
		room.PricingId,
		room.BookingStatus,
		room.CleaningStatus,
		room.PriceTemporary,
		current,
		employee_id,
		current,
		employee_id,
	)

	return err
}

// view list room
func (db *RoomSql) ViewListRoomRepo(ctx context.Context) ([]model.Room, error) {
	data := []model.Room{}

	query := `SELECT 
    room.room_id,
    room.room_name,
    room.floor,
    room.booking_status,
	room.cleaning_status,
	room.price_temporary,
    room.created_at,
    room.updated_at,
    typeroom.type_id,
    typeroom.type_name,
    typeroom.description,
    typeroom.max_occupancy,
    typeroom.room_size,
	roompricing.pricing_id,
    roompricing.price_per_day,
    roompricing.price_per_hour,
    roompricing.discount,
	view.view_id,
	view.view_name,
	view.view_description
FROM 
    room
JOIN 
	roompricing ON room.pricing_id = roompricing.pricing_id
JOIN 
    typeroom ON roompricing.type_id = typeroom.type_id
JOIN 
    view ON view.view_id = roompricing.view_id
WHERE 
    room.is_deleted = false;

`
	if err := db.Sql.Db.Select(&data, query); err != nil {
		return nil, err
	}

	return data, nil
}

// update room
func (db *RoomSql) UpdateRoomRepo(ctx context.Context, roomId string, room model.UpdateRoom, employeeID string) error {
	query := `UPDATE room SET
		price_temporary = $1,
		booking_status = $2,
		cleaning_status = $3,
		updated_at = $4,
		updated_by = $5
	WHERE room_id = $6`

	current := time.Now()

	if _, err := db.Sql.Db.Exec(query,
		room.TempPrice,
		room.BookingStatus,
		room.CleaningStatus,
		current,
		employeeID,
		roomId); err != nil {
		return err
	}

	return nil
}

// delete room
func (db *RoomSql) DeleteRoomRepo(ctx context.Context, roomId string, employeeID string) error {
	query := `update room

	set deleted_at = $1,
		deleted_by = $2,
		is_deleted = $3
	where room_id = $4`

	current := time.Now()
	isDeleted := true

	result, err := db.Sql.Db.Exec(query, current, employeeID, isDeleted, roomId)
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
