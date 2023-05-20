package pserror

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PsError struct {
	err     error
	Code    string
	Message string
}

var ErrNotFound = errors.New("product not found")

func (ps PsError) Error() string {
	return fmt.Sprintf("%s:%s", ps.Message, ps.err.Error())
}

func ParseErrors(err error) error {
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
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

func ToGrpcError(err error) error {
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}
		psErr, ok := err.(*PsError)
		if ok {
			switch psErr.Code {
			case "23503":
				return status.Error(codes.InvalidArgument, psErr.Message)
			case "23505":
				return status.Error(codes.AlreadyExists, psErr.Message)
			case "23502":
				return status.Error(codes.InvalidArgument, psErr.Message)
			default:
				return status.Error(codes.Unknown, psErr.Message)
			}
		}
		return status.Error(codes.Unknown, err.Error())
	}
	return nil
}
