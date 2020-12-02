package dao

import "errors"

var ErrNotFound = errors.New("Not Found")

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}
