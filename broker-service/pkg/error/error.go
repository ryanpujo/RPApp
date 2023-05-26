package error

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(c *gin.Context, err error) {

	var verr validator.ValidationErrors
	if errors.As(err, &verr) {
		errs := make(map[string]string)

		for _, f := range verr {
			err := f.ActualTag()
			if f.Param() != "" {
				err = fmt.Sprintf("%s=%s", err, f.Param())
			}
			errs[f.Field()] = err
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}
	st, ok := status.FromError(err)
	if ok {
		gerr := gin.H{"error": st.Message()}
		switch st.Code() {
		case codes.AlreadyExists:
			c.AbortWithStatusJSON(http.StatusConflict, gerr)
			return
		case codes.NotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gerr)
			return
		case codes.InvalidArgument:
			c.AbortWithStatusJSON(http.StatusBadRequest, gerr)
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gerr)
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
