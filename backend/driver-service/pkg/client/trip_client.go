package client

import (
	"context"
	"log"

	"github.com/qkhanh39/uit-go/driver-service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TripClient struct {
	Client pb.TripServiceClient
}

func InitTripClient(tripSvcUrl string) TripClient {
	cc, err := grpc.NewClient(tripSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error while creating trip client %v", err)
	}

	client := pb.NewTripServiceClient(cc)

	log.Println("Connected to Trip Service at", tripSvcUrl)

	return TripClient{
		Client: client,
	}
}

func (c *TripClient) UpdateTripStatus(ctx context.Context, request *pb.UpdateTripStatusRequest) (*pb.UpdateTripStatusResponse, error) {
	return c.Client.UpdateTripStatus(ctx, request)
}
