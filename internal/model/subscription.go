package model

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	Id          int       `json:"id" db:"id"`
	ServiceName string    `json:"service_name" db:"service_name"`
	Price       uint      `json:"price" db:"price"`
	UserId      uuid.UUID `json:"user_id" db:"user_id"`
	StartDate   time.Time `json:"start_date" db:"start_date"`
	EndDate     time.Time `json:"end_date" db:"end_date"`
}
