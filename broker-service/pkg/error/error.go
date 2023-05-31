package error

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spriigan/broker/response"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(c *gin.Context, err error) {

	var verr validator.ValidationErrors
	var res response.JsonRes
	if errors.As(err, &verr) {
		errs := make(map[string]string)

		for _, f := range verr {
			err := f.ActualTag()
			if f.Param() != "" {
				err = fmt.Sprintf("%s=%s", err, f.Param())
			}
			errs[f.Field()] = err
		}
		res.Errors = errs
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	st, ok := status.FromError(err)
	if ok {
		res.Error = st.Message()
		switch st.Code() {
		case codes.AlreadyExists:
			c.AbortWithStatusJSON(http.StatusConflict, res)
			return
		case codes.NotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, res)
			return
		case codes.InvalidArgument:
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
			return
		}
	}
	res.Error = err.Error()
	c.AbortWithStatusJSON(http.StatusInternalServerError, res)
}
