package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"time"
)

type TypeRoomRepo interface {
	ListPriceRepo(ctx context.Context) ([]model.RoomPricing, error)
	UpdatePriceRepo(ctx context.Context, data model.UpdatePrice, pricing_id string, employee_id string) error
}

type Sql struct {
	Sql *database.Sql
}

func NewTypeRoomRepo(sql *database.Sql) TypeRoomRepo {
	return &Sql{
		Sql: sql,
	}
}

func (db *Sql) ListPriceRepo(ctx context.Context) ([]model.RoomPricing, error) {
	RoomPricing := []model.RoomPricing{}
	query := `select rp.pricing_id, 
				rp.price_per_hour, 
				rp.price_per_day, 
				rp.discount, 
				tr.type_id, 
				tr.type_name, 
				tr.description, 
				tr.max_occupancy, 
				tr.room_size,
				v.view_id,
				v.view_name,
				v.view_description from roompricing as rp 
				join typeroom as tr on rp.type_id = tr.type_id
				join view as v on rp.view_id = v.view_id`
	if err := db.Sql.Db.Select(&RoomPricing, query); err != nil {
		return []model.RoomPricing{}, err
	}
	return RoomPricing, nil
}

func (db *Sql) UpdatePriceRepo(ctx context.Context, data model.UpdatePrice, pricing_id string, employee_id string) error {
	query := `UPDATE roompricing
SET 
    price_per_hour = $1,
    price_per_day = $2,
    discount = $3,
    updated_at = $4,
    updated_by = $5
WHERE 
    pricing_id = $6;`
	current := time.Now()
	if _, err := db.Sql.Db.Exec(query, data.PricePerHour, data.PricePerDay, data.Discount, current, employee_id, pricing_id); err != nil {
		return err
	}
	return nil
}
