package linkit

import "errors"

var (
	ErrDependencyCoreNil       = errors.New("dependency core is nil")
	ErrNotFound                = errors.New("dependency not found")
	ErrInvalidType             = errors.New("invalid dependency type")
	ErrCouldNotBuildDependency = errors.New("could not build dependency")
)

var ErrInvalidDependencyName = errors.New("invalid dependency name")
