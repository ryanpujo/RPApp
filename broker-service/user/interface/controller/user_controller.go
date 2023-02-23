package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/user/user-proto/grpc/models"
)

type UserController interface {
	Create(ctx *gin.Context)
}

type userController struct {
	client models.UserServiceClient
}

func NewUserController(client models.UserServiceClient) *userController {
	return &userController{client: client}
}

func (uc *userController) Create(c *gin.Context) {
	var payload models.UserPayload
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.String(http.StatusBadRequest, "got an error")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	id, err := uc.client.RegisterUser(ctx, &payload)
	if err != nil {
		c.String(http.StatusBadRequest, "got an error")
	}

	c.String(http.StatusOK, fmt.Sprint(id.GetId()))
}
