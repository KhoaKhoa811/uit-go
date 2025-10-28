package websocket

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/qkhanh39/uit-go/driver-service/pkg/pb"
	"github.com/qkhanh39/uit-go/driver-service/pkg/services"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleDriverConnection(w http.ResponseWriter, r *http.Request, clients map[string]*websocket.Conn, server *services.Server) {
	driverId := r.URL.Query().Get("driver_id")
	if driverId == "" {
		http.Error(w, "missing driver_id", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "failed to upgrade connection", http.StatusInternalServerError)
		return
	}

	key := fmt.Sprintf("driver:%s", driverId)
	clients[key] = conn
	fmt.Printf("Driver %s connected via WebSocket\n", driverId)

	go func() {
		for {
			var msg map[string]interface{}
			err := conn.ReadJSON(&msg)
			if err != nil {
				fmt.Printf("Driver %s disconnected: %v\n", driverId, err)
				delete(clients, key)
				break
			}

			fmt.Printf("Message from driver %s: %+v\n", driverId, msg)

			if msg["type"] == "trip_response" {
				tripId := int64(msg["trip_id"].(float64))
				driverId := int64(msg["driver_id"].(float64))
				accept := msg["accept"].(bool)

				_, err := server.ResponseToTripService(
					context.Background(),
					&pb.ResponseToTripServiceRequest{
						TripId:   tripId,
						DriverId: driverId,
						Accept:   accept,
					},
				)

				if err != nil {
					fmt.Printf("failed to handle driver %d response: %v\n", driverId, err)
				} else {
					fmt.Printf("Driver %d response handled (trip %d)\n", driverId, tripId)
				}
			}
		}
	}()
}
