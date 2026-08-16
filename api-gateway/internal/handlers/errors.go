package handlers

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func grpcErrorToHTTP(err error) int {
	code := status.Code(err)

	if code == codes.NotFound {
		return 404
	}

	if code == codes.InvalidArgument {
		return 400
	}
	return 500
}
