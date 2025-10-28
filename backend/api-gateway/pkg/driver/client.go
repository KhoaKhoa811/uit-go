package driver

import (
	"fmt"
	"log"

	"github.com/uit-go/api-gateway/pkg/driver/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DriverClient struct {
	Client pb.DriverServiceClient
}

func InitDriverClient(driverSvcUrl string) DriverClient {
	fmt.Println("API Gateway: InitServiceClient")

	cc, err := grpc.NewClient(driverSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("API Gateway: Error while creating driver client %v", err)
	}

	client := pb.NewDriverServiceClient(cc)

	log.Println("API Gateway: Connected to Driver Service at", driverSvcUrl)

	return DriverClient{
		Client: client,
	}
}
