package main

import (
	"fmt"
	"log"
	"net"

	"gihub.com/uit-go/trip-service/pkg/client"
	"gihub.com/uit-go/trip-service/pkg/config"
	"gihub.com/uit-go/trip-service/pkg/db"
	"gihub.com/uit-go/trip-service/pkg/pb"
	"gihub.com/uit-go/trip-service/pkg/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	c := config.LoadConfig()

	h := db.Init(c.DBUrl)
	driverClient := client.InitDriverClient(c.DriverSvcUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listening: ", err)
	}

	fmt.Println("Trip service on ", c.Port)

	s := services.Server{
		H:            h,
		DriverClient: driverClient,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTripServiceServer(grpcServer, &s)
	reflection.Register(grpcServer)

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("err while serving %v", err)
	}
}
