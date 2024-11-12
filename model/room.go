package model

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	RoomID         string     `json:"room_id" db:"room_id"`
	RoomName       string     `json:"room_name" db:"room_name"`
	Floor          int        `json:"floor" db:"floor"`
	BookingStatus  bool       `json:"booking_status" db:"booking_status"`
	PricePerDay    float32    `json:"price_per_day" db:"price_per_day"`
	PricePerHour   float32    `json:"price_per_hour" db:"price_per_hour"`
	Discount       float64    `json:"discount" db:"discount"`
	CleaningStatus string     `json:"cleaning_status" db:"cleaning_status"`
	CheckInTime    *time.Time `json:"check_in_time" db:"check_in_time"`
	CheckOutTime   *time.Time `json:"check_out_time" db:"check_out_time"`
	CurrentGuest   *int       `json:"current_guest" db:"current_guest"`
	Note           string     `json:"note" db:"note"`
	CreatedAt      *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at"`
	SelectTypeRoom
}

type ListRoom struct {
	RoomID         string     `json:"room_id" db:"room_id"`
	RoomName       string     `json:"room_name" db:"room_name"`
	Floor          int        `json:"floor" db:"floor"`
	BookingStatus  bool       `json:"booking_status" db:"booking_status"`
	PricePerDay    float32    `json:"price_per_day" db:"price_per_day"`
	PricePerHour   float32    `json:"price_per_hour" db:"price_per_hour"`
	Discount       float64    `json:"discount" db:"discount"`
	CleaningStatus string     `json:"cleaning_status" db:"cleaning_status"`
	CheckInTime    *time.Time `json:"check_in_time" db:"check_in_time"`
	CheckOutTime   *time.Time `json:"check_out_time" db:"check_out_time"`
	CurrentGuest   *int       `json:"current_guest" db:"current_guest"`
	Note           string     `json:"note" db:"note"`
	CreatedAt      *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at"`
	TypeRoom       SelectTypeRoom
}

type ViewRoom struct {
	RoomName       string     `json:"room_name" db:"room_name"`
	Floor          int        `json:"floor" db:"floor"`
	BookingStatus  bool       `json:"booking_status" db:"booking_status"`
	PricePerDay    float32    `json:"price_per_day" db:"price_per_day"`
	PricePerHour   float32    `json:"price_per_hour" db:"price_per_hour"`
	Discount       float64    `json:"discount" db:"discount"`
	CleaningStatus string     `json:"cleaning_status" db:"cleaning_status"`
	CheckInTime    *time.Time `json:"check_in_time" db:"check_in_time"`
	CheckOutTime   *time.Time `json:"check_out_time" db:"check_out_time"`
	CurrentGuest   *string    `json:"current_guest" db:"current_guest"`
	Note           string     `json:"note" db:"note"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	SelectTypeRoom
}
type AddRoom struct {
	RoomName       string     `json:"room_name" db:"room_name"`
	TypeID         uuid.UUID  `json:"type_id" db:"type_id"`
	Floor          int        `json:"floor" db:"floor"`
	BookingStatus  bool       `json:"booking_status" db:"booking_status"`
	PricePerDay    float32    `json:"price_per_day" db:"price_per_day"`
	PricePerHour   float32    `json:"price_per_hour" db:"price_per_hour"`
	Discount       float64    `json:"discount" db:"discount"`
	CleaningStatus string     `json:"cleaning_status" db:"cleaning_status"`
	CheckInTime    *time.Time `json:"check_in_time" db:"check_in_time"`
	CheckOutTime   *time.Time `json:"check_out_time" db:"check_out_time"`
	CurrentGuest   *int       `json:"current_guest" db:"current_guest"`
	Note           string     `json:"note" db:"note"`
	CreatedBy      string     `json:"created_by" db:"created_by"`
	UpdatedBy      string     `json:"updated_by" db:"updated_by"`
}

type UpdateRoom struct {
	RoomName     string    `json:"room_name,omitempty" db:"room_name"`
	TypeID       uuid.UUID `json:"type_id,omitempty" db:"type_id"`
	PricePerDay  float32   `json:"price_per_day,omitempty" db:"price_per_day"`
	PricePerHour float32   `json:"price_per_hour,omitempty" db:"price_per_hour"`
	Discount     float32  `json:"discount,omitempty" db:"discount"`
	UpdatedBy    string    `json:"updated_by,omitempty" db:"updated_by"`
}
type DeleteRoom struct {
	DeleteBy string `json:"deleteby" db:"deleteby"`
}
