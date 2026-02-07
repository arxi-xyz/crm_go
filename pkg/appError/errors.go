package appError

import "errors"

var (
	ErrUserNotFound = errors.New("user_not_found")
	ErrBadPassword  = errors.New("bad_password")
	ErrUserBlocked  = errors.New("user_blocked")
	ErrInvalidCreds = errors.New("invalid_credentials")
)
