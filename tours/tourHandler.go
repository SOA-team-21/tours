package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
	"tours.xws.com/model"
	"tours.xws.com/proto/tours"
	"tours.xws.com/service"
)

type TourHandler struct {
	tours.UnimplementedToursServiceServer
	TourService *service.TourService
}

func (handler *TourHandler) Get(ctx context.Context, request *tours.UserIdRequest) (*tours.TourResponse, error) {

	traceCtx, span := tp.Tracer(serviceName).Start(ctx, "Get Tour")
	defer func() { span.End() }() // na kraju funkcije zatvori trace

	userId := fmt.Sprint(request.UserId)

	var fromDb, err = handler.TourService.FindTour(userId, traceCtx)
	if err != nil {
		log.Println("Error while getting tour.")
		return &tours.TourResponse{}, err
	}
	log.Println("Getting tour done.")
	return TourToRpc(fromDb), nil
}

func (handler *TourHandler) GetAllByAuthor(ctx context.Context, request *tours.UserIdRequest) (*tours.ToursResponse, error) {

	traceCtx, span := tp.Tracer(serviceName).Start(ctx, "Get Tours by Author")
	defer func() { span.End() }() // na kraju funkcije zatvori trace

	userId := fmt.Sprint(request.UserId)

	var fromDb, err = handler.TourService.FindAllByAuthor(userId, traceCtx)
	if err != nil {
		log.Println("Error while getting tours by author.")
		return &tours.ToursResponse{}, err
	}
	log.Println("Getting tours by author done.")
	return ToursToRpc(fromDb), nil
}

func (handler *TourHandler) Create(ctx context.Context, request *tours.TourResponse) (*tours.TourResponse, error) {
	tour := RpcToTour(request)

	var fromDb, err = handler.TourService.Create(tour)
	if err != nil {
		log.Println("Error while creating tour.")
		return &tours.TourResponse{}, err
	}
	log.Println("Creating tour done.")
	return TourToRpc(fromDb), nil
}

func (handler *TourHandler) Update(ctx context.Context, request *tours.TourResponse) (*tours.TourResponse, error) {
	tour := RpcToTour(request)

	var fromDb, err = handler.TourService.Update(tour)
	if err != nil {
		log.Println("Error while updating tour.")
		return &tours.TourResponse{}, err
	}
	log.Println("Updating tour done.")
	return TourToRpc(fromDb), nil
}

func (handler *TourHandler) Publish(ctx context.Context, request *tours.TourResponse) (*tours.TourResponse, error) {
	tour := RpcToTour(request)

	var fromDb, err = handler.TourService.Publish(tour)
	if err != nil {
		log.Println("Error while publishing tour.")
		return &tours.TourResponse{}, err
	}
	log.Println("Publishing tour done.")
	return TourToRpc(fromDb), nil
}

func (handler *TourHandler) Archive(ctx context.Context, request *tours.TourResponse) (*tours.TourResponse, error) {
	tour := RpcToTour(request)

	var fromDb, err = handler.TourService.Archive(tour)
	if err != nil {
		log.Println("Error while archiving tour.")
		return &tours.TourResponse{}, err
	}
	log.Println("Archiving tour done.")
	return TourToRpc(fromDb), nil
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

func RpcToTour(rpcTour *tours.TourResponse) *model.Tour {
	return &model.Tour{
		Id:            rpcTour.Id,
		Name:          rpcTour.Name,
		Description:   rpcTour.Description,
		Difficult:     int64(rpcTour.Difficult),
		Price:         float64(rpcTour.Price),
		Status:        model.TourStatus(rpcTour.Status),
		AuthorId:      rpcTour.AuthorId,
		Length:        float64(rpcTour.Length),
		PublishTime:   rpcTour.PublishTime.AsTime(),
		ArchiveTime:   rpcTour.ArchiveTime.AsTime(),
		MyOwn:         rpcTour.MyOwn,
		KeyPoints:     RpcToKeyPoints(rpcTour.KeyPoints),
		RequiredTimes: RpcToRequiredTimes(rpcTour.RequiredTimes),
		Tags:          rpcTour.Tags,
	}
}

func RpcsToTours(rpcTours *tours.ToursResponse) []model.Tour {
	result := make([]model.Tour, len(rpcTours.Tours))
	for i, e := range rpcTours.Tours {
		result[i] = *RpcToTour(e)
	}
	return result
}

func RpcToKeyPoint(rpcPoint *tours.KeyPoint) *model.KeyPoint {
	return &model.KeyPoint{
		Id:          rpcPoint.Id,
		TourId:      rpcPoint.TourId,
		Latitude:    float64(rpcPoint.Latitude),
		Longitude:   float64(rpcPoint.Longitude),
		Name:        rpcPoint.Name,
		Description: rpcPoint.Description,
		Picture:     rpcPoint.Picture,
		Public:      rpcPoint.Public,
	}
}

func RpcToKeyPoints(rpcPoints []*tours.KeyPoint) []model.KeyPoint {
	result := make([]model.KeyPoint, len(rpcPoints))
	for i, e := range rpcPoints {
		result[i] = *RpcToKeyPoint(e)
	}
	return result
}

func RpcToRequiredTime(rpcTime *tours.RequiredTime) *model.RequiredTime {
	return &model.RequiredTime{
		Id:        rpcTime.Id,
		TourId:    rpcTime.TourId,
		Transport: model.TransportType(rpcTime.TransportType),
		Minutes:   int(rpcTime.Minutes),
	}
}

func RpcToRequiredTimes(rpcTimes []*tours.RequiredTime) []model.RequiredTime {
	result := make([]model.RequiredTime, len(rpcTimes))
	for i, e := range rpcTimes {
		result[i] = *RpcToRequiredTime(e)
	}
	return result
}
