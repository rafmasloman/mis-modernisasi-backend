package errors

import "errors"

var (
	ErrFilterNotFound       = errors.New("filter not found")
	ErrInvalidFilterType    = errors.New("invalid filter type")
	ErrDuplicateFilter      = errors.New("filter already exists")
	ErrInvalidFilterPayload = errors.New("invalid request payload")
)
