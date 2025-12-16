package mappers

import (
	"main/internal/dto"
	"main/internal/model"
)

func CreateWebToModel(data dto.CreateSubRequest) model.Subscription {
	return model.Subscription{
		ServiceName: data.ServiceName,
		Price:       data.Price,
		UserId:      data.UserId,
		StartDate:   data.StartDate,
		EndDate:     data.EndDate,
	}
}

func ModelToLoadWeb(data model.Subscription) dto.LoadSubResponce {
	return dto.LoadSubResponce{
		Id:          data.Id,
		ServiceName: data.ServiceName,
		Price:       data.Price,
		UserId:      data.UserId,
		StartDate:   data.StartDate,
		EndDate:     data.EndDate,
	}
}

func UpdateWebToModel(data dto.UpdateSubRequest) model.Subscription {
	return model.Subscription{
		Id:          data.Id,
		ServiceName: data.ServiceName,
		Price:       data.Price,
		UserId:      data.UserId,
		StartDate:   data.StartDate,
		EndDate:     data.EndDate,
	}
}
