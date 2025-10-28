package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uit-go/api-gateway/pkg/trip/pb"
)

type EstimateTripFareRequestBody struct {
	PickupLatitude   float64 `json:"pickup_latitude"`
	PickupLongitude  float64 `json:"pickup_longitude"`
	DropoffLatitude  float64 `json:"dropoff_latitude"`
	DropoffLongitude float64 `json:"dropoff_longitude"`
}

func EstimateTripFare(ctx *gin.Context, c pb.TripServiceClient) {
	fmt.Println("API Gateway: Estimate Trip Fare")

	body := EstimateTripFareRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.EstimateTripFare(context.Background(), &pb.EstimateTripFareRequest{
		PickupLatitude:   body.PickupLatitude,
		PickupLongitude:  body.PickupLongitude,
		DropoffLatitude:  body.DropoffLatitude,
		DropoffLongitude: body.DropoffLongitude,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
