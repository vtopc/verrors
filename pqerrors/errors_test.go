package pqerrors

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vtopc/verrors"
)

func Test_handleError(t *testing.T) {
	tests := map[string]struct {
		err error
		id  []string

		wantErr error
	}{
		"no_error": {
			err:     nil,
			wantErr: nil,
		},
		"not_found": {
			err:     sql.ErrNoRows,
			wantErr: verrors.NotFound{Err: verrors.Err{}},
		},
		"not_found_with_id": {
			err:     sql.ErrNoRows,
			id:      []string{"foo"},
			wantErr: verrors.NotFound{Err: verrors.Err{ID: "foo"}},
		},
		"already_exists": {
			err:     &pq.Error{Code: notFoundErrCode},
			wantErr: verrors.AlreadyExists{Err: verrors.Err{}},
		},
		"already_exists_with_id": {
			err:     &pq.Error{Code: notFoundErrCode},
			id:      []string{"foo"},
			wantErr: verrors.AlreadyExists{Err: verrors.Err{ID: "foo"}},
		},
		"unknown_error": {
			err:     errors.New("some error"),
			wantErr: errors.New("some error"),
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			err := HandleError(context.Background(), tt.err, tt.id...)
			if tt.wantErr == nil {
				require.NoError(t, err)
				return
			}

			assert.EqualError(t, err, tt.wantErr.Error())
		})
	}
}
