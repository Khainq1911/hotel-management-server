package model

import "github.com/google/uuid"

type SelectTypeRoom struct {
	TypeID          uuid.UUID `json:"type_id" db:"type_id"`
	TypeName        string    `json:"type_name" db:"type_name"`
	TypeDescription string    `json:"type_description" db:"description"`
	MaxOccupancy    int       `json:"max_occupancy" db:"max_occupancy"`
	RoomSize        float32   `json:"room_size" db:"room_size"`
}

type RoomPricing struct {
	PricingId    uuid.UUID `json:"pricing_id" db:"pricing_id"`
	PricePerDay  float32   `json:"price_per_day" db:"price_per_day"`
	PricePerHour float32   `json:"price_per_hour" db:"price_per_hour"`
	Discount     float32   `json:"discount" db:"discount"`
	SelectTypeRoom
	View
}

type ListRoomPricing struct {
	PricingId    uuid.UUID `json:"pricing_id" db:"pricing_id"`
	PricePerDay  float32   `json:"price_per_day" db:"price_per_day"`
	PricePerHour float32   `json:"price_per_hour" db:"price_per_hour"`
	Discount     float32   `json:"discount" db:"discount"`
	TypeRoom     SelectTypeRoom
	View         View
}
type UpdatePrice struct {
	PricePerDay  float32 `json:"price_per_day" db:"price_per_day"`
	PricePerHour float32 `json:"price_per_hour" db:"price_per_hour"`
	Discount     float32 `json:"discount" db:"discount"`
}
