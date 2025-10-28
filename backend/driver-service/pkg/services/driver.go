package services

import (
	"context"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/qkhanh39/uit-go/driver-service/pkg/client"
	"github.com/qkhanh39/uit-go/driver-service/pkg/db"
	"github.com/qkhanh39/uit-go/driver-service/pkg/models"
	"github.com/qkhanh39/uit-go/driver-service/pkg/pb"
	"github.com/qkhanh39/uit-go/driver-service/utils"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	H                db.Handler
	Redis            db.RedisHandler
	WebSocketClients map[string]*websocket.Conn
	TripClient       client.TripClient
	pb.UnimplementedDriverServiceServer
}

func (s *Server) CreateDriver(ctx context.Context, request *pb.CreateDriverRequest) (*pb.CreateDriverResponse, error) {
	var driver models.Driver
	driver.Name = request.Name
	driver.Email = request.Email
	driver.Vehicle = request.Vehicle
	driver.LicenseNumber = request.LicenseNumber
	driver.Phone = request.Phone
	driver.Password = utils.HashPassword(request.Password)

	if result := s.H.DB.Create(&driver); result.Error != nil {
		return &pb.CreateDriverResponse{
			Status: 409,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateDriverResponse{
		Status:  201,
		Id:      driver.Id,
		Name:    driver.Name,
		Vehicle: driver.Vehicle,
		Email:   driver.Email,
	}, nil
}

func (s *Server) UpdateDriverLocation(ctx context.Context, request *pb.UpdateDriverLocationRequest) (*pb.UpdateDriverLocationResponse, error) {
	driverId := fmt.Sprintf("driver:%d", request.Id)

	_, err := s.Redis.Client.GeoAdd(ctx, "drivers", &redis.GeoLocation{
		Name:      driverId,
		Longitude: request.Longitude,
		Latitude:  request.Latitude,
	}).Result()

	if err != nil {
		return &pb.UpdateDriverLocationResponse{
			Status:  "error",
			Message: fmt.Sprintf("failed to update location: %v", err),
		}, nil
	}

	_, err = s.Redis.Client.HSet(ctx, driverId, map[string]interface{}{
		"latitude":  request.Latitude,
		"longitude": request.Longitude,
		"timestamp": time.Now().Unix(),
	}).Result()

	if err != nil {
		return &pb.UpdateDriverLocationResponse{
			Status:  "err",
			Message: fmt.Sprintf("failed to store metadata: %v", err),
		}, nil
	}

	return &pb.UpdateDriverLocationResponse{
		Status:  "success",
		Message: "driver location updated successfully",
	}, nil
}

func (s *Server) GetAvailableDriversNearUser(ctx context.Context, request *pb.GetAvailableDriversNearUserRequest) (*pb.GetAvailableDriversNearUserResponse, error) {
	maxDriver := request.MaxDriver
	if maxDriver == 0 {
		maxDriver = 5
	}

	radius := request.Radius
	if radius == 0 {
		radius = 3
	}

	result, err := s.Redis.Client.GeoSearch(ctx, "drivers", &redis.GeoSearchQuery{
		Longitude:  request.Longitude,
		Latitude:   request.Latitude,
		Radius:     float64(radius),
		Count:      int(maxDriver),
		RadiusUnit: "km",
		Sort:       "ASC",
	}).Result()

	if err != nil {
		return &pb.GetAvailableDriversNearUserResponse{
			Status:  "error",
			Drivers: nil,
		}, err
	}

	var drivers []int64

	for _, driverKey := range result {
		status, err := s.Redis.Client.HGet(ctx, driverKey, "status").Result()
		if err != nil || status != "available" {
			continue
		}

		var driverId int64
		_, err = fmt.Sscanf(driverKey, "driver:%d", &driverId)
		if err != nil {
			continue
		}

		drivers = append(drivers, driverId)
	}

	return &pb.GetAvailableDriversNearUserResponse{
		Status:  "success",
		Drivers: drivers,
	}, nil
}

func (s *Server) SendNotificationToDriver(ctx context.Context, request *pb.SendNotificationToDriverRequest) (*pb.SendNotificationToDriverResponse, error) {
	driverKey := fmt.Sprintf("driver:%d", request.DriverId)

	conn, ok := s.WebSocketClients[driverKey]
	if !ok {
		return &pb.SendNotificationToDriverResponse{
			Status:  "error",
			Message: "driver not connected",
		}, nil
	}

	notification := map[string]interface{}{
		"type":              "trip_notification",
		"pickup_latitude":   request.PickupLatitude,
		"pickup_longitude":  request.PickupLongitude,
		"dropoff_latitude":  request.DropoffLatitude,
		"dropoff_longitude": request.DropoffLongitude,
		"distance":          request.Distance,
		"fare":              request.TripFare,
	}

	err := conn.WriteJSON(notification)
	if err != nil {
		return &pb.SendNotificationToDriverResponse{
			Status:  "error",
			Message: fmt.Sprintf("failed to send notification: %v", err),
		}, nil
	}

	return &pb.SendNotificationToDriverResponse{
		Status:  "success",
		Message: "notification sent successfully",
	}, nil
}

func (s *Server) ResponseToTripService(ctx context.Context, request *pb.ResponseToTripServiceRequest) (*pb.ResponseToTripServiceResponse, error) {
	if request.Accept {
		fmt.Println("response called")
		_, err := s.TripClient.UpdateTripStatus(ctx, &pb.UpdateTripStatusRequest{
			TripId:   request.TripId,
			DriverId: request.DriverId,
			Response: "accepted",
		})

		if err != nil {
			return &pb.ResponseToTripServiceResponse{
				Status:  "error",
				Message: "Failed to update trip status",
			}, nil
		}

		return &pb.ResponseToTripServiceResponse{
			Status:  "success",
			Message: "Trip accepted",
		}, nil
	} else {
		_, _ = s.TripClient.UpdateTripStatus(ctx, &pb.UpdateTripStatusRequest{
			TripId:   request.TripId,
			DriverId: request.DriverId,
			Response: "rejected",
		})

		return &pb.ResponseToTripServiceResponse{
			Status:  "rejected",
			Message: "Trip rejected",
		}, nil
	}
}
