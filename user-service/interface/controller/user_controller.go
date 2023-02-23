package controller

import (
	"context"
	"errors"

	"github.com/spriigan/RPApp/interface/repository"
	"github.com/spriigan/RPApp/usecases/interactor"
	"github.com/spriigan/RPApp/user-proto/grpc/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userServer struct {
	models.UnimplementedUserServiceServer
	interactor interactor.UserInteractor
}

func NewUserServer(i interactor.UserInteractor) *userServer {
	return &userServer{interactor: i}
}

func (us *userServer) RegisterUser(ctx context.Context, payload *models.UserPayload) (*models.UserId, error) {
	id, err := us.interactor.Create(payload)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}
	return &models.UserId{Id: int64(id)}, nil
}

func (us *userServer) FindByUsername(ctx context.Context, username *models.Username) (*models.UserBio, error) {
	input := username.GetUsername()
	foundUser, err := us.interactor.FindByUsername(input)
	if err != nil {
		if errors.Is(err, repository.ErrNoUserFound) {
			return nil, status.Error(codes.NotFound, repository.ErrNoUserFound.Error())
		}
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}
	bio := models.UserBio{
		Id:       foundUser.Id,
		Fname:    foundUser.Fname,
		Lname:    foundUser.Lname,
		Username: foundUser.Username,
		Email:    foundUser.Email,
	}
	return &bio, nil
}

func (us *userServer) FindUsers(context.Context, *emptypb.Empty) (*models.Users, error) {
	users, err := us.interactor.FindUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *userServer) DeleteByUsername(ctx context.Context, username *models.Username) (*emptypb.Empty, error) {
	err := us.interactor.DeleteByUsername(username.Username)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (us *userServer) Update(ctx context.Context, payload *models.UserPayload) (*emptypb.Empty, error) {
	err := us.interactor.Update(payload)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}
	return &emptypb.Empty{}, nil
}
