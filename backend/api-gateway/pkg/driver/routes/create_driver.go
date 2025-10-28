package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uit-go/api-gateway/pkg/driver/pb"
)

type CreateDriverRequestBody struct {
	Name          string `json:"name"`
	Vehicle       string `json:"vehicle"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Phone         string `json:"phone"`
	LicenseNumber string `json:"license_number"`
}

func CreateDriver(ctx *gin.Context, c pb.DriverServiceClient) {
	fmt.Println("API Gateway: Create Driver")

	body := CreateDriverRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateDriver(context.Background(), &pb.CreateDriverRequest{
		Name:          body.Name,
		Vehicle:       body.Vehicle,
		Email:         body.Email,
		Password:      body.Password,
		Phone:         body.Phone,
		LicenseNumber: body.LicenseNumber,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
