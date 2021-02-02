package verrors

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO: move to a separate package?

// ToGRPCError maps err to gRPC status codes error.
func ToGRPCError(err error) error {
	if err == nil {
		return nil
	}

	return toGRPCStatus(err).Err()
}

// toGRPCStatus maps err to gRPC status code.
// Read more about gRPC status codes - https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
func toGRPCStatus(err error) *status.Status {
	if e := new(InvalidArgument); errors.As(err, e) {
		return status.New(codes.InvalidArgument, e.Error())
	}

	if e := new(NotFound); errors.As(err, e) {
		return status.New(codes.NotFound, e.Error())
	}

	if e := new(AlreadyExists); errors.As(err, e) {
		return status.New(codes.AlreadyExists, e.Error())
	}

	// default:
	return status.New(codes.Internal, err.Error())
}
