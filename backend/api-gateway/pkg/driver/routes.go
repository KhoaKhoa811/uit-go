package driver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uit-go/api-gateway/pkg/driver/routes"
)

func RegisterRoute(r *gin.Engine, driverSvcUrl string) {
	fmt.Println("API Gateway: Register Route")

	driverClient := InitDriverClient(driverSvcUrl)

	route := r.Group("driver")
	route.POST("", driverClient.CreateDriver)
	route.PUT("/location", driverClient.UpdateDriverLocation)
}

func (driverSvc *DriverClient) CreateDriver(ctx *gin.Context) {
	fmt.Println("API Gateway: CreateDriver")

	routes.CreateDriver(ctx, driverSvc.Client)
}

func (driverSvc *DriverClient) UpdateDriverLocation(ctx *gin.Context) {
	fmt.Println("API Gateway: UpdateDriverLocation")

	routes.UpdateDriverLocation(ctx, driverSvc.Client)
}
