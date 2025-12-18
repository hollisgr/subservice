package subscriptions

import (
	"context"
	"errors"
	"fmt"
	"main/internal/dto"
	"main/internal/interfaces"
	"main/internal/mappers"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	ErrEndIsLess      = errors.New("end date is less than start date")
	ErrIncorrectDate  = errors.New("start or end date is incorrect")
	ErrIncorrectValue = errors.New("incorrect value")
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
	id := 0

	logrus.Info("sub service: create")

	ok := checkDateStr(data.StartDate)

	if !ok {
		logrus.Error(ErrIncorrectDate)
		return id, ErrIncorrectDate
	}

	// если дату окончания не дали, будем считать что дата окончания конец столетия

	if data.EndDate == "" {
		data.EndDate = "12-2099"
	}

	ok = checkDateStr(data.EndDate)

	if !ok {
		logrus.Error(ErrIncorrectDate)
		return id, ErrIncorrectDate
	}

	newSub := mappers.CreateWebToModel(data)

	if newSub.EndDate.Before(newSub.StartDate) {
		logrus.Error(ErrEndIsLess)
		return id, ErrEndIsLess
	}

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

	ok := checkDateStr(data.StartDate)

	if !ok {
		logrus.Error(ErrIncorrectDate)
		return ErrIncorrectDate
	}

	// если дату окончания не дали, будем считать что дата окончания конец столетия

	if data.EndDate == "" {
		data.EndDate = "12-2099"
	}

	ok = checkDateStr(data.EndDate)

	if !ok {
		logrus.Error(ErrIncorrectDate)
		return ErrIncorrectDate
	}

	sub := mappers.UpdateWebToModel(data)

	if sub.EndDate.Before(sub.StartDate) {
		logrus.Error(ErrEndIsLess)
		return ErrEndIsLess
	}

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

	ok := checkDateStr(data.StartDate)

	if !ok {
		logrus.Error(ErrIncorrectDate)
		return result, ErrIncorrectDate
	}

	ok = checkDateStr(data.EndDate)

	if !ok {
		logrus.Error(ErrIncorrectDate)
		return result, ErrIncorrectDate
	}

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

func checkDateStr(date string) bool {
	var month, year int
	n, err := fmt.Sscanf(date, "%d-%d", &month, &year)
	if err != nil || n != 2 {
		return false
	}

	// Захардкодил проверку на месяц и год, минимальный допустимый год 1999, максимальный 2099

	if month > 12 || month <= 0 {
		return false
	}

	if year < 1999 || year > 2099 {
		return false
	}

	return true
}
