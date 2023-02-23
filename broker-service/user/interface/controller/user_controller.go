package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/response"
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
	var res response.JsonResponse
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		res.Error = true
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	payloadPB := models.UserPayload{
		Bio: &models.UserBio{
			Fname:    payload.Fname,
			Lname:    payload.Lname,
			Username: payload.Username,
			Email:    payload.Email,
		},
		Password: payload.Password,
	}

	id, err := uc.client.RegisterUser(ctx, &payloadPB)
	if err != nil {
		res.Error = true
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res.Error = false
	res.Data = id.GetId()
	c.JSON(http.StatusOK, res)
}
