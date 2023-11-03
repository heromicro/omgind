package errors

import (
	oerrors "errors"
)

func Is(err, target error) bool {
	return oerrors.Is(err, target)
}

func Join(errs ...error) error {
	return oerrors.Join(errs...)
}
