package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidationErrorUnwrap(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string, len(verr))
	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err
	}
	return errs
}
