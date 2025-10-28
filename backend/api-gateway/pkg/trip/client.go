package trip

import (
	"fmt"
	"log"

	"github.com/uit-go/api-gateway/pkg/trip/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TripClient struct {
	Client pb.TripServiceClient
}

func InitTripClient(tripSvcUrl string) TripClient {
	fmt.Println("API Gateway: InitTripClient")

	cc, err := grpc.NewClient(tripSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("API Gateway: Error while creating trip client %v", err)
	}

	client := pb.NewTripServiceClient(cc)

	log.Println("API Gateway: Connected to Trip Service at", tripSvcUrl)

	return TripClient{
		Client: client,
	}
}
