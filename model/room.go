package model

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	RoomID         string     `json:"room_id" db:"room_id"`
	RoomName       string     `json:"room_name" db:"room_name"`
	Floor          int        `json:"floor" db:"floor"`
	TempPrice      *float32   `json:"price_temporary" db:"price_temporary"`
	BookingStatus  bool       `json:"booking_status" db:"booking_status"`
	CleaningStatus string     `json:"cleaning_status" db:"cleaning_status"`
	CreatedAt      *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at"`
	SelectTypeRoom
	View
	Price
}

type ListRoom struct {
	RoomID         string     `json:"room_id" db:"room_id"`
	RoomName       string     `json:"room_name" db:"room_name"`
	Floor          int        `json:"floor" db:"floor"`
	TempPrice      *float32   `json:"price_temporary" db:"price_temporary"`
	BookingStatus  bool       `json:"booking_status" db:"booking_status"`
	CleaningStatus string     `json:"cleaning_status" db:"cleaning_status"`
	CreatedAt      *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at"`
	TypeRoom       SelectTypeRoom
	View           View
	Price          Price
}

type AddRoom struct {
	RoomName       string    `json:"room_name" db:"room_name"`
	PricingId      uuid.UUID `json:"pricing_id" db:"pricing_id"`
	Floor          int       `json:"floor" db:"floor"`
	BookingStatus  bool      `json:"booking_status" db:"booking_status"`
	PriceTemporary *float32  `json:"price_temporary" db:"price_temporary"`
	CleaningStatus string    `json:"cleaning_status" db:"cleaning_status"`
}

type UpdateRoom struct {
	BookingStatus  bool     `json:"booking_status" db:"booking_status"`
	CleaningStatus string   `json:"cleaning_status" db:"cleaning_status"`
	TempPrice      *float32 `json:"price_temporary" db:"price_temporary"`
	UpdatedBy string `json:"updated_by,omitempty" db:"updated_by"`
}
type DeleteRoom struct {
	DeletedBy uuid.UUID `json:"deleted_by" db:"deleted_by"`
}
