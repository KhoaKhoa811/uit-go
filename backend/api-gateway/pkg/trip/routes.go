package trip

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uit-go/api-gateway/pkg/trip/routes"
)

func RegisterRoute(r *gin.Engine, tripSvcUrl string) {
	fmt.Println("API Gateway: Register Trip Route")

	tripClient := InitTripClient(tripSvcUrl)

	routes := r.Group("trip")
	routes.POST("", tripClient.CreateTrip)
	routes.POST("/estimate-fare", tripClient.EstimateTripFare)
}

func (tripSvc *TripClient) CreateTrip(ctx *gin.Context) {
	fmt.Println("API Gateway: Create Trip")

	routes.CreateTrip(ctx, tripSvc.Client)
}

func (tripSvc *TripClient) EstimateTripFare(ctx *gin.Context) {
	fmt.Println("API Gateway: EstimateTripFare")

	routes.EstimateTripFare(ctx, tripSvc.Client)
}
