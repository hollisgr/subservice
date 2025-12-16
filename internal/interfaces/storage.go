package interfaces

import (
	"context"
	"main/internal/model"
)

type Storage interface {
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, sub model.Subscription) error
	LoadList(ctx context.Context, limit int, offset int) ([]model.Subscription, error)
	Load(ctx context.Context, id int) (model.Subscription, error)
	Create(ctx context.Context, sub model.Subscription) (int, error)
}
