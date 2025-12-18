package mappers

import (
	"fmt"
	"main/internal/dto"
	"main/internal/model"
	"strings"
	"time"
)

func CreateWebToModel(data dto.CreateSubRequest) model.Subscription {
	res := model.Subscription{
		ServiceName: data.ServiceName,
		Price:       data.Price,
		UserId:      data.UserId,
		StartDate:   ConvertStringToDate(data.StartDate),
	}
	if data.EndDate != "" {
		res.EndDate = ConvertStringToDate(data.EndDate)
	}
	return res
}

func ModelToLoadWeb(data model.Subscription) dto.LoadSubResponce {
	res := dto.LoadSubResponce{
		Id:          data.Id,
		ServiceName: data.ServiceName,
		Price:       data.Price,
		UserId:      data.UserId,
		StartDate:   ConvertDateToString(data.StartDate),
	}
	if data.EndDate != time.Unix(0, 0) {
		res.EndDate = ConvertDateToString(data.EndDate)
	}
	return res
}

func UpdateWebToModel(data dto.UpdateSubRequest) model.Subscription {
	res := model.Subscription{
		Id:          data.Id,
		ServiceName: data.ServiceName,
		Price:       data.Price,
		UserId:      data.UserId,
		StartDate:   ConvertStringToDate(data.StartDate),
	}
	if data.EndDate != "" {
		res.EndDate = ConvertStringToDate(data.EndDate)
	}
	return res
}

func CostRequestToCostDB(data dto.CostRequest) dto.CostRequestToDB {
	return dto.CostRequestToDB{
		ServiceName: data.ServiceName,
		UserId:      data.UserId,
		StartDate:   ConvertStringToDate(data.StartDate),
		EndDate:     ConvertStringToDate(data.EndDate),
	}
}

func ConvertStringToDate(str string) (date time.Time) {
	split := strings.Split(str, "-")
	dateStr := fmt.Sprintf("%s-%s-01", split[1], split[0])
	date, _ = time.Parse("2006-01-02", dateStr)
	return date
}

func ConvertDateToString(date time.Time) (str string) {
	if date.IsZero() {
		return ""
	}
	year, month, _ := date.Date()
	return fmt.Sprintf("%02d-%d", month, year)
}
