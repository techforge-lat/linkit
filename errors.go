package linkit

import "errors"

var (
	ErrNotFound    = errors.New("dependency not found")
	ErrInvalidType = errors.New("invalid dependency type")
)

var (
	ErrInvalidDependencyName = errors.New("invalid dependency name")
)
