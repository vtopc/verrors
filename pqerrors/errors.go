package pqerrors

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/vtopc/verrors"
)

const notFoundErrCode = "23505"

// handleError maps sql/postgresql errors into the verrors
// and logs service(non-user) errors
func HandleError(_ context.Context, err error, id ...string) error {
	cErr := verrors.Err{}
	if len(id) > 0 {
		cErr.ID = id[0]
	}

	switch {
	case err == nil:
		return nil

	case errors.Is(err, sql.ErrNoRows):
		return verrors.NotFound{Err: cErr}
	}

	// handle postgresql errors:
	if pqErr := new(pq.Error); errors.As(err, &pqErr) {
		// duplicate key value violates unique constraint:
		if pqErr.Code == notFoundErrCode {
			return verrors.AlreadyExists{Err: cErr}
		}
	}

	return err
}
