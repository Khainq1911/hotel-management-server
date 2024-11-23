package model

import "github.com/google/uuid"

type Price struct {
	PricingId    uuid.UUID `json:"pricing_id" db:"pricing_id"`
	PricePerDay  float32   `json:"price_per_day" db:"price_per_day"`
	PricePerHour float32   `json:"price_per_hour" db:"price_per_hour"`
	Discount     float64   `json:"discount" db:"discount"`
}
