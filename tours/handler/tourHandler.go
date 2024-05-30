package handler

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"tours.xws.com/model"
	"tours.xws.com/proto/tours"
	"tours.xws.com/service"
)

type TourHandler struct {
	tours.UnimplementedToursServiceServer
	TourService *service.TourService
}

func TourToRpc(tour *model.Tour) *tours.TourResponse {
	return &tours.TourResponse{
		Id:            tour.Id,
		Name:          tour.Name,
		Description:   tour.Description,
		Difficult:     tour.Difficult,
		Price:         float32(tour.Price),
		Status:        tours.TourStatus(tour.Status),
		AuthorId:      tour.AuthorId,
		Length:        float32(tour.Length),
		PublishTime:   timestamppb.New(tour.PublishTime),
		ArchiveTime:   timestamppb.New(tour.ArchiveTime),
		MyOwn:         tour.MyOwn,
		KeyPoints:     KeyPointsToRpc(tour.KeyPoints),
		RequiredTimes: RequiredTimesToRpc(tour.RequiredTimes),
		Tags:          tour.Tags,
	}
}

func ToursToRpc(points []model.Tour) *tours.ToursResponse {
	result := make([]*tours.TourResponse, len(points))
	for i, e := range points {
		result[i] = TourToRpc(&e)
	}
	return &tours.ToursResponse{Tours: result}
}

func KeyPointToRpc(point *model.KeyPoint) *tours.KeyPoint {
	return &tours.KeyPoint{
		Id:          point.Id,
		TourId:      point.TourId,
		Latitude:    float32(point.Latitude),
		Longitude:   float32(point.Longitude),
		Name:        point.Name,
		Description: point.Description,
		Picture:     point.Picture,
		Public:      point.Public,
	}
}

func KeyPointsToRpc(points []model.KeyPoint) []*tours.KeyPoint {
	result := make([]*tours.KeyPoint, len(points))
	for i, e := range points {
		result[i] = KeyPointToRpc(&e)
	}
	return result
}

func RequiredTimeToRpc(reqTime *model.RequiredTime) *tours.RequiredTime {
	return &tours.RequiredTime{
		Id:            reqTime.Id,
		TourId:        reqTime.TourId,
		TransportType: tours.Transport(reqTime.Transport),
		Minutes:       int64(reqTime.Minutes),
	}
}

func RequiredTimesToRpc(times []model.RequiredTime) []*tours.RequiredTime {
	result := make([]*tours.RequiredTime, len(times))
	for i, e := range times {
		result[i] = RequiredTimeToRpc(&e)
	}
	return result
}
