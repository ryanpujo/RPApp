package response

import "github.com/spriigan/broker/user/domain"

type JsonRes struct {
	User   domain.User       `json:"user,omitempty"`
	Users  []domain.User     `json:"users,omitempty"`
	Error  string            `json:"error,omitempty"`
	Errors map[string]string `json:"errors,omitempty"`
}
