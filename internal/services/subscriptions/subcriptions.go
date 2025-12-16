package subscriptions

import (
	"context"
	"main/internal/dto"
	"main/internal/interfaces"
	"main/internal/mappers"

	"github.com/sirupsen/logrus"
)

type sub struct {
	storage interfaces.Storage
}

func New(s interfaces.Storage) interfaces.Subscriptions {
	return &sub{
		storage: s,
	}
}

func (s *sub) Create(ctx context.Context, data dto.CreateSubRequest) (int, error) {
	newSub := mappers.CreateWebToModel(data)
	id, err := s.storage.Create(ctx, newSub)
	if err != nil {
		logrus.Error(err)
		return id, err
	}
	return id, nil
}

func (s *sub) Load(ctx context.Context, id int) (dto.LoadSubResponce, error) {
	res := dto.LoadSubResponce{}
	data, err := s.storage.Load(ctx, id)
	if err != nil {
		logrus.Error(err)
		return res, err
	}
	res = mappers.ModelToLoadWeb(data)
	return res, nil
}

func (s *sub) LoadList(ctx context.Context, limit int, offset int) ([]dto.LoadSubResponce, error) {
	res := []dto.LoadSubResponce{}
	data, err := s.storage.LoadList(ctx, limit, offset)
	if err != nil {
		logrus.Error(err)
		return res, err
	}
	for _, sub := range data {
		temp := mappers.ModelToLoadWeb(sub)
		res = append(res, temp)
	}
	return res, nil
}

func (s *sub) Update(ctx context.Context, data dto.UpdateSubRequest) error {
	sub := mappers.UpdateWebToModel(data)
	err := s.storage.Update(ctx, sub)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (s *sub) Delete(ctx context.Context, id int) error {
	err := s.storage.Delete(ctx, id)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
