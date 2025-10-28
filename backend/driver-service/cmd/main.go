package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/qkhanh39/uit-go/driver-service/pkg/client"
	"github.com/qkhanh39/uit-go/driver-service/pkg/config"
	"github.com/qkhanh39/uit-go/driver-service/pkg/db"
	"github.com/qkhanh39/uit-go/driver-service/pkg/pb"
	"github.com/qkhanh39/uit-go/driver-service/pkg/services"
	ws "github.com/qkhanh39/uit-go/driver-service/pkg/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	c := config.LoadConfig()

	h := db.Init(c.DBUrl)

	redis := db.InitRedis(c.Redis)

	wsClients := make(map[string]*websocket.Conn)

	tripClient := client.InitTripClient(c.TripSvcUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listening: ", err)
	}

	fmt.Println("Driver service on ", c.Port)

	s := services.Server{
		H:                h,
		Redis:            redis,
		WebSocketClients: wsClients,
		TripClient:       tripClient,
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.HandleDriverConnection(w, r, wsClients, &s)
	})

	go func() {
		fmt.Println("WebSocket server running on :8080/ws")
		log.Fatal(http.ListenAndServe(c.SocketPort, nil))
	}()

	grpcServer := grpc.NewServer()

	pb.RegisterDriverServiceServer(grpcServer, &s)
	reflection.Register(grpcServer)

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("error while serving %v", err)
	}
}
