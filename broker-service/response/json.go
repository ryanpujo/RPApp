package response

import (
	pr "github.com/spriigan/broker/product/domain"
	"github.com/spriigan/broker/user/domain"
)

type JsonRes struct {
	User     domain.User   `json:"user,omitempty"`
	Users    []domain.User `json:"users,omitempty"`
	Product  pr.Product
	Products []pr.Product
	Error    string            `json:"error,omitempty"`
	Errors   map[string]string `json:"errors,omitempty"`
}
