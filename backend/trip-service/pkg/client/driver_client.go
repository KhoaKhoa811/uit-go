package client

import (
	"context"
	"log"

	"gihub.com/uit-go/trip-service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DriverClient struct {
	Client pb.DriverServiceClient
}

func InitDriverClient(driverSvcUrl string) DriverClient {
	cc, err := grpc.NewClient(driverSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error while creating driver client %v", err)
	}

	client := pb.NewDriverServiceClient(cc)

	log.Println("Connected to Driver Service at", driverSvcUrl)

	return DriverClient{
		Client: client,
	}
}

func (c *DriverClient) GetAvailableDriverNearUser(ctx context.Context, longitude float64, latitude float64, radius int32, maxDriver int32) (*pb.GetAvailableDriversNearUserResponse, error) {
	request := &pb.GetAvailableDriversNearUserRequest{
		Longitude: longitude,
		Latitude:  latitude,
		Radius:    radius,
		MaxDriver: maxDriver,
	}
	return c.Client.GetAvailableDriversNearUser(ctx, request)
}

func (c *DriverClient) SendNotificationToDriver(ctx context.Context, request *pb.SendNotificationToDriverRequest) (*pb.SendNotificationToDriverResponse, error) {
	return c.Client.SendNotificationToDriver(ctx, request)
}
