package grpcserver

import (
	"errors"
	"fmt"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestToGRPCError(t *testing.T) {
	tests := []struct {
		name     string
		inputErr error
		wantCode codes.Code
	}{
		{
			name:     "ErrMovieNotFound",
			inputErr: fmt.Errorf("Error movie not found"),
			wantCode: codes.NotFound,
		},
		{
			name:     "ErrInvalidMovieData",
			inputErr: fmt.Errorf("Invalid movie data"),
			wantCode: codes.InvalidArgument,
		},
		{
			name:     "ErrInternalServer",
			inputErr: fmt.Errorf("Erro internal server"),
			wantCode: codes.Internal,
		},
		{
			name:     "UnknownError",
			inputErr: errors.New("unknown error"),
			wantCode: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := toGRPCError(tt.inputErr)
			st, ok := status.FromError(err)
			if !ok {
				t.Errorf("Expected gRPC error, got %v", err)
				return
			}
			if st.Code() != tt.wantCode {
				t.Errorf("Expected code %v, got %v", tt.wantCode, st.Code())
			}
		})
	}
}
