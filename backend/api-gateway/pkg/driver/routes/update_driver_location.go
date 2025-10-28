package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uit-go/api-gateway/pkg/driver/pb"
)

type UpdateDriverLocationRequestBody struct {
	DriverId  int64   `json:"driver_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func UpdateDriverLocation(ctx *gin.Context, c pb.DriverServiceClient) {
	fmt.Println("API Gateway: Update Driver Location")

	body := UpdateDriverLocationRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdateDriverLocation(context.Background(), &pb.UpdateDriverLocationRequest{
		Id:        body.DriverId,
		Latitude:  body.Latitude,
		Longitude: body.Longitude,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
