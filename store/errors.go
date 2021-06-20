package store

import "errors"

var (
	// ErrNotFound indicates that a model was not found
	ErrNotFound = errors.New("store: model not found")
)
