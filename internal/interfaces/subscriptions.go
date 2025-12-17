package interfaces

import (
	"context"
	"main/internal/dto"
)

type Subscriptions interface {
	Create(ctx context.Context, data dto.CreateSubRequest) (int, error)
	Load(ctx context.Context, id int) (dto.LoadSubResponce, error)
	LoadList(ctx context.Context, limit int, offset int) ([]dto.LoadSubResponce, error)
	Update(ctx context.Context, data dto.UpdateSubRequest) error
	Delete(ctx context.Context, id int) error
	Cost(ctx context.Context, data dto.CostRequest) (dto.CostResponce, error)
}
