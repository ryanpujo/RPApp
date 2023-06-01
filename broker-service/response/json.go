package response

import (
	"firebase.google.com/go/auth"
	pr "github.com/spriigan/broker/product/domain"
	"github.com/spriigan/broker/user/domain"
)

type JsonRes struct {
	User        domain.User       `json:"user,omitempty"`
	Users       []domain.User     `json:"users,omitempty"`
	CreatedUser auth.UserRecord   `json:"createdUser,omitempty"`
	Product     pr.Product        `json:"product,omitempty"`
	Products    []pr.Product      `json:"products,omitempty"`
	Error       string            `json:"error,omitempty"`
	Errors      map[string]string `json:"errors,omitempty"`
}
