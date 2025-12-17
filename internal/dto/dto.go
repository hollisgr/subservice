package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateSubRequest struct {
	ServiceName string    `json:"service_name" example:"Yandex Plus"`
	Price       uint      `json:"price" example:"400"`
	UserId      uuid.UUID `json:"user_id" example:"dceb1963-e152-47ff-a562-81a360627309"`
	StartDate   string    `json:"start_date" example:"01-2025"`
	EndDate     string    `json:"end_date,omitempty" example:"02-2025"`
}

type LoadListRequest struct {
	Limit  uint `json:"limit" example:"10"`
	Offset uint `json:"offset" example:"1"`
}

type CreateSubResponce struct {
	Success        bool `json:"success" example:"true"`
	SubscriptionId int  `json:"subscription_id" example:"123"`
}

type UpdateSubResponce struct {
	Success bool `json:"success" example:"true"`
}

type DeleteSubResponce struct {
	Success bool `json:"success" example:"true"`
}

type LoadSubResponce struct {
	Id          int       `json:"id" example:"1"`
	ServiceName string    `json:"service_name" example:"Yandex Plus"`
	Price       uint      `json:"price" example:"400"`
	UserId      uuid.UUID `json:"user_id" example:"dceb1963-e152-47ff-a562-81a360627309"`
	StartDate   string    `json:"start_date" example:"01-2025"`
	EndDate     string    `json:"end_date,omitempty" example:"02-2025"`
}

type UpdateSubRequest struct {
	Id          int       `json:"id" example:"1"`
	ServiceName string    `json:"service_name" example:"Yandex Minus"`
	Price       uint      `json:"price" example:"399"`
	UserId      uuid.UUID `json:"user_id" example:"dceb1963-e152-47ff-a562-81a360627309"`
	StartDate   string    `json:"start_date" example:"05-2025"`
	EndDate     string    `json:"end_date,omitempty" example:"07-2025"`
}

type CostRequest struct {
	ServiceName string    `json:"service_name" example:"Yandex Plus"`
	UserId      uuid.UUID `json:"user_id" example:"dceb1963-e152-47ff-a562-81a360627309"`
	StartDate   string    `json:"start_date" example:"01-2025"`
	EndDate     string    `json:"end_date" example:"02-2025"`
}

type CostRequestToDB struct {
	ServiceName string
	UserId      uuid.UUID
	StartDate   time.Time
	EndDate     time.Time
}

type CostResponce struct {
	ServiceName string    `json:"service_name" example:"Yandex Plus"`
	UserId      uuid.UUID `json:"user_id" example:"dceb1963-e152-47ff-a562-81a360627309"`
	Cost        int       `json:"cost" example:"900"`
	MonthsCount int       `json:"months_count" example:"3"`
}
