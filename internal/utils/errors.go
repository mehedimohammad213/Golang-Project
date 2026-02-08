package utils

import "errors"

var (
	ErrNotFound      = errors.New("resource not found")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrBadRequest    = errors.New("bad request")
	ErrInternal      = errors.New("internal server error")
	ErrAlreadyExists = errors.New("resource already exists")
)
