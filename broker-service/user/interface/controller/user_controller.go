package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/response"
	"github.com/spriigan/broker/user/domain"
	"github.com/spriigan/broker/user/user-proto/grpc/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserController interface {
	Create(ctx *gin.Context)
	FindUsers(ctx *gin.Context)
	FindByUsername(ctx *gin.Context)
	DeleteByUsername(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type userController struct {
	client models.UserServiceClient
}
type Uri struct {
	Username string `uri:"username" binding:"required,min=6"`
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

func (uc *userController) FindUsers(c *gin.Context) {
	var res response.JsonResponse
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	users, err := uc.client.FindUsers(ctx, &emptypb.Empty{})
	if err != nil {
		res.Error = true
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res.Error = false
	res.Data = users.User
	c.JSON(http.StatusOK, res)
}

func (uc *userController) FindByUsername(c *gin.Context) {
	var res response.JsonResponse
	var uri Uri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		res.Error = true
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	user, err := uc.client.FindByUsername(ctx, &models.Username{Username: uri.Username})
	if err != nil {
		res.Error = true
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res.Error = false
	res.Data = user
	c.JSON(http.StatusOK, res)
}

func (uc *userController) DeleteByUsername(c *gin.Context) {
	var res response.JsonResponse
	var uri Uri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		res.Error = true
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_, err = uc.client.DeleteByUsername(ctx, &models.Username{Username: uri.Username})
	if err != nil {
		res.Error = true
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res.Error = false
	res.Message = "user has been deleted"
	c.JSON(http.StatusOK, res)
}

func (uc *userController) Update(c *gin.Context) {
	var res response.JsonResponse
	var payload domain.UserPayload

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
			Id:       int64(payload.Id),
			Fname:    payload.Fname,
			Lname:    payload.Lname,
			Username: payload.Username,
			Email:    payload.Email,
		},
		Password: payload.Password,
	}

	_, err = uc.client.Update(ctx, &payloadPB)
	if err != nil {
		res.Error = true
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res.Error = false
	res.Message = "succesfully updated"
	c.JSON(http.StatusOK, res)
}
