package controller

import (
	"context"
	"errors"

	"github.com/spriigan/RPMedia/domain"
	"github.com/spriigan/RPMedia/interface/repository"
	"github.com/spriigan/RPMedia/usecases/interactor"
	"github.com/spriigan/RPMedia/user-proto/grpc/models"
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

func (us *userServer) RegisterUser(ctx context.Context, req *models.UserPayload) (*models.UserId, error) {
	bio := req.GetBio()
	newUser := domain.UserPayload{
		Fname:    bio.Fname,
		Lname:    bio.Lname,
		Username: bio.Username,
		Password: req.GetPassword(),
		Email:    bio.Email,
	}

	id, err := us.interactor.Create(&newUser)
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
		Id:       int64(foundUser.Id),
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
	userBios := make([]*models.UserBio, 0, len(users))

	for _, v := range users {
		bio := models.UserBio{
			Fname:    v.Fname,
			Lname:    v.Lname,
			Username: v.Username,
			Email:    v.Email,
			Id:       int64(v.Id),
		}
		userBios = append(userBios, &bio)
	}

	return &models.Users{User: userBios}, nil
}

func (us *userServer) DeleteByUsername(ctx context.Context, username *models.Username) (*emptypb.Empty, error) {
	err := us.interactor.DeleteByUsername(username.Username)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (us *userServer) Update(ctx context.Context, payload *models.UserPayload) (*emptypb.Empty, error) {
	input := payload.GetBio()
	user := domain.UserPayload{
		Id:       int(input.GetId()),
		Fname:    input.GetFname(),
		Lname:    input.GetLname(),
		Username: input.GetUsername(),
		Email:    input.GetEmail(),
		Password: payload.GetPassword(),
	}
	err := us.interactor.Update(user)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}
	return &emptypb.Empty{}, nil
}
