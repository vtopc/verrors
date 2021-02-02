package verrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_toGRPCStatus(t *testing.T) {
	tests := map[string]struct {
		err      error
		wantCode codes.Code
	}{
		"not_found": {
			err:      NotFound{},
			wantCode: codes.NotFound,
		},

		"internal": {
			err:      errors.New("some internal error"),
			wantCode: codes.Internal,
		},

		"wrapped": {
			err:      fmt.Errorf("some error: %w", NotFound{}),
			wantCode: codes.NotFound,
		},

		// TODO: Add test cases.
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			got := toGRPCStatus(tt.err)
			assert.Equal(t, tt.wantCode, got.Code())

			// TODO: assert got.Message()
		})
	}
}

func TestToGRPCError(t *testing.T) {
	tests := map[string]struct {
		err            error
		wantStatusCode codes.Code
	}{
		"OK": {
			err:            nil,
			wantStatusCode: codes.OK,
		},
		"InvalidArgument": {
			err:            InvalidArgument{},
			wantStatusCode: codes.InvalidArgument,
		},
		"NotFound": {
			err:            NotFound{},
			wantStatusCode: codes.NotFound,
		},
		"AlreadyExists": {
			err:            AlreadyExists{},
			wantStatusCode: codes.AlreadyExists,
		},
		"Internal": {
			err:            errors.New("some random error"),
			wantStatusCode: codes.Internal,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			err := ToGRPCError(tt.err)
			st, ok := status.FromError(err)
			require.True(t, ok, "error is not representing status.Status")

			assert.Equal(t, tt.wantStatusCode, st.Code())
		})
	}
}
