package services

import (
	"context"
	"fmt"
	"log"

	"gihub.com/uit-go/trip-service/pkg/client"
	"gihub.com/uit-go/trip-service/pkg/db"
	"gihub.com/uit-go/trip-service/pkg/models"
	"gihub.com/uit-go/trip-service/pkg/pb"
	"gihub.com/uit-go/trip-service/utils"
)

type Server struct {
	H            db.Handler
	DriverClient client.DriverClient
	pb.UnimplementedTripServiceServer
}

func (s *Server) CreateTrip(ctx context.Context, request *pb.CreateTripRequest) (*pb.CreateTripResponse, error) {
	radius := 3
	maxDriver := 5
	driverResp, err := s.DriverClient.GetAvailableDriverNearUser(ctx, request.PickupLongitude, request.PickupLatitude, int32(radius), int32(maxDriver))

	if err != nil || len(driverResp.Drivers) == 0 {
		return &pb.CreateTripResponse{
			Status:  "error",
			Message: "No available drivers nearby",
		}, nil
	}

	driverId := driverResp.Drivers[0]

	fareResp, err := s.EstimateTripFare(ctx, &pb.EstimateTripFareRequest{
		PickupLatitude:   request.PickupLatitude,
		PickupLongitude:  request.PickupLongitude,
		DropoffLatitude:  request.DropoffLatitude,
		DropoffLongitude: request.DropoffLongitude,
	})

	if err != nil {
		return &pb.CreateTripResponse{
			Status:  "err",
			Message: fmt.Sprintf("failed to estimate trip fare: %v", err),
		}, nil
	}

	sendNotificationRequest := pb.SendNotificationToDriverRequest{
		DriverId:         driverId,
		PickupLatitude:   request.PickupLatitude,
		PickupLongitude:  request.PickupLongitude,
		DropoffLatitude:  request.DropoffLatitude,
		DropoffLongitude: request.DropoffLongitude,
		Distance:         fareResp.Distance,
		TripFare:         fareResp.EstimateTripFare,
	}

	sendNotificationResponse, err := s.DriverClient.SendNotificationToDriver(ctx, &sendNotificationRequest)

	if err != nil {
		return &pb.CreateTripResponse{
			Status:  "error",
			Message: fmt.Sprintf("gRPC internal error: %v", err),
		}, nil
	}

	if sendNotificationResponse.Status != "success" {
		return &pb.CreateTripResponse{
			Status:  "error",
			Message: fmt.Sprintf("failed to send notification: %s", sendNotificationResponse.Message),
		}, nil
	}

	trip := models.Trip{
		UserId:      request.UserId,
		PickupLat:   request.PickupLatitude,
		PickupLong:  request.PickupLongitude,
		DropoffLat:  request.DropoffLatitude,
		DropoffLong: request.DropoffLongitude,
		Status:      "pending",
	}

	s.H.DB.Create(&trip)

	return &pb.CreateTripResponse{
		Status:   "success",
		Message:  "Trip created successfully",
		DriverId: trip.DriverId,
		TripId:   trip.Id,
	}, nil
}

func (s *Server) EstimateTripFare(ctx context.Context, request *pb.EstimateTripFareRequest) (*pb.EstimateTripFareResponse, error) {
	baseFare := 10000.0
	perKm := 8000.0

	distance := utils.HaversineDistance(request.PickupLatitude, request.PickupLongitude, request.DropoffLatitude, request.DropoffLongitude)

	totalFare := baseFare + perKm*distance

	return &pb.EstimateTripFareResponse{
		Distance:         distance,
		EstimateTripFare: int64(totalFare),
	}, nil
}

func (s *Server) UpdateTripStatus(ctx context.Context, request *pb.UpdateTripStatusRequest) (*pb.UpdateTripStatusResponse, error) {
	var trip models.Trip
	if result := s.H.DB.First(&trip, request.TripId); result.Error != nil {
		return &pb.UpdateTripStatusResponse{
			Status:  "error",
			Message: "Trip not found",
		}, nil
	}

	if request.Response != "accepted" {

	}

	trip.DriverId = request.DriverId
	trip.Status = request.Response

	s.H.DB.Save(&trip)

	log.Printf("Trip %d updated by driver %d: %s\n", trip.Id, request.DriverId, request.Response)

	return &pb.UpdateTripStatusResponse{
		Status:  "successful",
		Message: fmt.Sprintf("Trip status updated to %s", request.Response),
	}, nil
}
