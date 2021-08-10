package verrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToHTTPStatusCode(t *testing.T) {
	tests := map[string]struct {
		err      error
		wantCode int
	}{
		"not_found": {
			err:      NotFound{},
			wantCode: 404,
		},

		"wrapped_not_found": {
			err:      fmt.Errorf("some error: %w", NewNotFound(errors.New("not found in table users"), "123")),
			wantCode: 404,
		},

		"bad_request": {
			err:      NewInvalidArgument(errors.New("some field is invalid")),
			wantCode: 400,
		},

		"already_exists": {
			err:      NewAlreadyExists(nil),
			wantCode: 409,
		},

		"internal": {
			err:      errors.New("some internal error"),
			wantCode: 500,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := ToHTTPStatusCode(tt.err)
			assert.Equal(t, tt.wantCode, got)
		})
	}
}
