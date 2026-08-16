package grpcserver

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func toGRPCError(err error) error {
	if errors.Is(err, fmt.Errorf("movie not found")) {
		return status.Error(codes.NotFound, err.Error())
	}

	if errors.Is(err, fmt.Errorf("invalid movie ID")) {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	if errors.Is(err, fmt.Errorf("internal server error")) {
		return status.Error(codes.Internal, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
