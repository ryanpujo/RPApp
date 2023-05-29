package controller

import (
	"context"

	"github.com/spriigan/RPApp/pkg/usererror"
	"github.com/spriigan/RPApp/repository"
	"github.com/spriigan/RPApp/service"
	"github.com/spriigan/RPApp/user-proto/userpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userController struct {
	userpb.UnimplementedUserServiceServer
	userService service.UserService
}

func NewUserController(serv service.UserService) *userController {
	return &userController{userService: serv}
}

func (c userController) CreateUser(ctx context.Context, payload *userpb.UserPayload) (*userpb.User, error) {
	arg := repository.CreateUserParams{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Username:  payload.Username,
	}

	createdUser, err := c.userService.CreateUser(ctx, arg)
	if err != nil {
		return nil, usererror.ToGrpcError(err)
	}
	user := &userpb.User{
		Id:        createdUser.ID,
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		Username:  createdUser.Username,
		CreatedAt: timestamppb.New(createdUser.CreatedAt.Time),
	}

	return user, nil
}

func (c userController) DeleteById(ctx context.Context, id *userpb.UserId) (*emptypb.Empty, error) {
	err := c.userService.DeleteByID(ctx, id.Id)
	return &emptypb.Empty{}, usererror.ToGrpcError(err)
}

func (c userController) GetById(ctx context.Context, id *userpb.UserId) (*userpb.User, error) {
	user, err := c.userService.GetById(ctx, id.Id)
	if err != nil {
		return nil, usererror.ToGrpcError(err)
	}
	found := &userpb.User{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		CreatedAt: timestamppb.New(user.CreatedAt.Time),
	}
	return found, nil
}

func (c userController) GetMany(ctx context.Context, limit *userpb.Limit) (*userpb.Users, error) {
	found, err := c.userService.GetMany(ctx, limit.Limit)
	if err != nil {
		return nil, usererror.ToGrpcError(err)
	}

	users := &userpb.Users{
		Users: make([]*userpb.User, 0, len(found)),
	}
	for _, v := range found {
		user := userpb.User{
			Id:        v.ID,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Username:  v.Username,
			CreatedAt: timestamppb.New(v.CreatedAt.Time),
		}
		users.Users = append(users.Users, &user)
	}

	return users, nil
}

func (c userController) UpdateById(ctx context.Context, payload *userpb.UserPayload) (*emptypb.Empty, error) {
	arg := repository.UpdateByIDParams{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Username:  payload.Username,
		ID:        payload.Id,
	}

	err := c.userService.UpdateByID(ctx, arg)

	return &emptypb.Empty{}, usererror.ToGrpcError(err)
}
