package response

import "google.golang.org/grpc/codes"

type JsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message,omitempty"`
	Code    codes.Code  `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}
