package grpcserver

import (
	"errors"
	errD "movies-service/internal/core/domain/err-d"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func toGRPCError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, errD.ErrMovieNotFound):
		return status.Error(codes.NotFound, err.Error())

	case errors.Is(err, errD.ErrIDNotValid),
		errors.Is(err, errD.ErrInvalidMovieData):
		return status.Error(codes.InvalidArgument, err.Error())

	default:
		return status.Error(codes.Internal, err.Error())
	}
}
