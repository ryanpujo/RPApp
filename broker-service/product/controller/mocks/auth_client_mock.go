package mocks

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/stretchr/testify/mock"
)

type MockAuth struct {
	mock.Mock
}

func (a MockAuth) VerifyIDToken(ctx context.Context, id string) (*auth.Token, error) {
	args := a.Called(ctx, id)
	token, _ := args.Get(0).(*auth.Token)
	return token, args.Error(1)
}
