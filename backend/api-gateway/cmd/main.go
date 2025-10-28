package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uit-go/api-gateway/pkg/config"
	"github.com/uit-go/api-gateway/pkg/driver"
	"github.com/uit-go/api-gateway/pkg/trip"
)

func main() {
	fmt.Println("starting API Gateway")

	c := config.LoadConfig()

	r := gin.Default()

	driver.RegisterRoute(r, c.DriverSvcUrl)
	trip.RegisterRoute(r, c.TripSvcUrl)

	r.Run(c.Port)
}
