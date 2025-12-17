package subscriptions

import (
	"context"
	"errors"
	"main/internal/dto"
	"main/internal/interfaces"
	"main/internal/mappers"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	ErrEndIsLess = errors.New("end date is less than start date")
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
	logrus.Info("sub service: create")

	// если дату окончания не дали, будем считать что дата окончания конец столетия

	if data.EndDate == "" {
		data.EndDate = "12-2099"
	}

	newSub := mappers.CreateWebToModel(data)

	id, err := s.storage.Create(ctx, newSub)
	if err != nil {
		logrus.Error(err)
		return id, err
	}
	logrus.Info("sub service: create success")
	return id, nil
}

func (s *sub) Load(ctx context.Context, id int) (dto.LoadSubResponce, error) {
	logrus.Info("sub service: load")
	res := dto.LoadSubResponce{}
	data, err := s.storage.Load(ctx, id)
	if err != nil {
		logrus.Error(err)
		return res, err
	}
	res = mappers.ModelToLoadWeb(data)
	logrus.Info("sub service: load success")
	return res, nil
}

func (s *sub) LoadList(ctx context.Context, limit int, offset int) ([]dto.LoadSubResponce, error) {
	logrus.Info("sub service: load list")
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
	logrus.Info("sub service: load list success")
	return res, nil
}

func (s *sub) Update(ctx context.Context, data dto.UpdateSubRequest) error {
	logrus.Info("sub service: update")
	sub := mappers.UpdateWebToModel(data)
	err := s.storage.Update(ctx, sub)
	if err != nil {
		logrus.Error(err)
		return err
	}
	logrus.Info("sub service: update success")
	return nil
}

func (s *sub) Delete(ctx context.Context, id int) error {
	logrus.Info("sub service: delete")
	err := s.storage.Delete(ctx, id)
	if err != nil {
		logrus.Error(err)
		return err
	}
	logrus.Info("sub service: delete success")
	return nil
}

func (s *sub) Cost(ctx context.Context, data dto.CostRequest) (dto.CostResponce, error) {
	result := dto.CostResponce{}

	start := mappers.ConvertStringToDate(data.StartDate)
	end := mappers.ConvertStringToDate(data.EndDate)

	// если дата окончания раньше старта, то возвращаем ошибку
	if end.Before(start) {
		logrus.Error(ErrEndIsLess)
		return result, ErrEndIsLess
	}

	dbData := mappers.CostRequestToCostDB(data)

	sub, err := s.storage.Cost(ctx, dbData)
	if err != nil {
		logrus.Error(err)
		return result, err
	}

	// определить начальную дату
	if start.Before(sub.StartDate) {
		start = sub.StartDate
	}

	// определить конечную дату
	if sub.EndDate.Before(end) {
		end = sub.EndDate
	}

	// +1 потому что учитываем мес включительно
	monthsCount := monthDiff(start, end) + 1

	cost := monthsCount * int(sub.Price)

	result.Cost = cost
	result.MonthsCount = monthsCount
	result.ServiceName = sub.ServiceName
	result.UserId = sub.UserId

	return result, nil
}

func monthDiff(start, end time.Time) int {
	startYear, startMonth, _ := start.Date()
	endYear, endMonth, _ := end.Date()

	totalStart := startYear*12 + int(startMonth)
	totalEnd := endYear*12 + int(endMonth)

	return totalEnd - totalStart
}
