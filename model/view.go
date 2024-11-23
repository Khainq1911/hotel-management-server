package model

import "github.com/google/uuid"

type View struct {
	ViewId          uuid.UUID `json:"view_id" db:"view_id"`
	ViewName        string    `json:"view_name" db:"view_name"`
	ViewDescription string    `json:"view_description" db:"view_description"`
}
