package mappers

import (
	"fmt"
	"main/internal/dto"
	"main/internal/model"
	"strings"
	"time"
)

func CreateWebToModel(data dto.CreateSubRequest) model.Subscription {
	return model.Subscription{
		ServiceName: data.ServiceName,
		Price:       data.Price,
		UserId:      data.UserId,
		StartDate:   convertStringToDate(data.StartDate),
		EndDate:     convertStringToDate(data.EndDate),
	}
}

func ModelToLoadWeb(data model.Subscription) dto.LoadSubResponce {
	return dto.LoadSubResponce{
		Id:          data.Id,
		ServiceName: data.ServiceName,
		Price:       data.Price,
		UserId:      data.UserId,
		StartDate:   convertDateToString(data.StartDate),
		EndDate:     convertDateToString(data.EndDate),
	}
}

func UpdateWebToModel(data dto.UpdateSubRequest) model.Subscription {
	return model.Subscription{
		Id:          data.Id,
		ServiceName: data.ServiceName,
		Price:       data.Price,
		UserId:      data.UserId,
		StartDate:   convertStringToDate(data.StartDate),
		EndDate:     convertStringToDate(data.EndDate),
	}
}

func convertStringToDate(str string) (date time.Time) {
	minDate := "01-1999"
	if len(str) == 0 || str <= minDate {
		return date
	}
	split := strings.Split(str, "-")
	dateStr := fmt.Sprintf("%s-%s-01", split[1], split[0])
	date, _ = time.Parse("2006-01-02", dateStr)

	return date
}

func convertDateToString(date time.Time) (str string) {
	if date.IsZero() {
		return ""
	}
	year, month, _ := date.Date()
	return fmt.Sprintf("%02d-%d", month, year)
}
