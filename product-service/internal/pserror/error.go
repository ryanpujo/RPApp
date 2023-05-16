package pserror

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PsError struct {
	err      error
	Code     string
	Message  string
	HttpCode int
}

func (ps PsError) Error() string {
	return fmt.Sprintf("%s:%s", ps.Message, ps.err.Error())
}

func ParseErrors(err error) error {
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("product not found")
		}
		pgErr, ok := err.(*pgconn.PgError)
		if ok {
			switch pgErr.Code {
			case "23503":
				return PsError{
					err:     err,
					Code:    pgErr.Code,
					Message: "make sure you have the store and specify the category for the product",
				}
			case "23505":
				return PsError{
					err:     err,
					Code:    pgErr.Code,
					Message: "data already exist",
				}
			case "23502":
				return PsError{
					err:     err,
					Code:    pgErr.Code,
					Message: "make sure all required field is filled",
				}
			default:
				return PsError{
					err:     err,
					Code:    pgErr.Code,
					Message: "unknown error",
				}
			}
		}
		return err
	}
	return nil
}
