package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/user/domain"
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
	var payload domain.UserPayload
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.String(http.StatusBadRequest, "got an error")
	}

	bio := models.UserBio{
		Fname:    payload.Fname,
		Lname:    payload.Lname,
		Username: payload.Username,
		Email:    payload.Email,
	}
	userPayload := models.UserPayload{
		Bio:      &bio,
		Password: payload.Password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	id, err := uc.client.RegisterUser(ctx, &userPayload)
	if err != nil {
		c.String(http.StatusBadRequest, "got an error")
	}

	c.String(http.StatusOK, fmt.Sprint(id.GetId()))
}
