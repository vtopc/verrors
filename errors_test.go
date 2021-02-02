package verrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	tests := map[string]struct {
		err  Err
		want error
	}{
		"simple": {
			err: Err{
				Msg: "some error",
			},
			want: errors.New("some error"),
		},

		"with_ID": {
			err: Err{
				Msg: "some error",
				ID:  "123",
			},
			want: errors.New("ID(123): some error"),
		},

		// TODO: Add test cases.
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			assert.EqualError(t, tt.err, tt.want.Error())
		})
	}
}

//nolint:dupl
func TestNotFound(t *testing.T) {
	tests := map[string]struct {
		err  NotFound
		want error
	}{
		"with_message": {
			err: NotFound{
				Err: Err{
					ID:  "123",
					Msg: "some entity not found",
				},
			},
			want: errors.New("ID(123): some entity not found"),
		},

		"default_message": {
			err: NotFound{
				Err: Err{
					ID: "123",
				},
			},
			want: errors.New("ID(123): not found"),
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			assert.EqualError(t, tt.err, tt.want.Error())
		})
	}
}

//nolint:dupl
func TestAlreadyExists(t *testing.T) {
	tests := map[string]struct {
		err  AlreadyExists
		want error
	}{
		"with_message": {
			err: AlreadyExists{
				Err: Err{
					ID:  "123",
					Msg: "some entity duplicated",
				},
			},
			want: errors.New("ID(123): some entity duplicated"),
		},

		"default_message": {
			err: AlreadyExists{
				Err: Err{
					ID: "123",
				},
			},
			want: errors.New("ID(123): already exists"),
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			assert.EqualError(t, tt.err, tt.want.Error())
		})
	}
}
