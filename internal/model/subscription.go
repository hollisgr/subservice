package model

import "github.com/google/uuid"

type Subscription struct {
	Id          int       `json:"id" db:"id"`
	ServiceName string    `json:"service_name" db:"service_name"`
	Price       int       `json:"price" db:"price"`
	UserId      uuid.UUID `json:"user_id" db:"user_id"`
	StartDate   string    `json:"start_date" db:"start_date"`
	EndDate     string    `json:"end_date" db:"end_date"`
}
