package dependor

import (
	"errors"
)

var ErrNotFound = errors.New("dependor: dependency not found")
var ErrInvalidType = errors.New("dependor: invalid dependency type")

func init() {
	container = make(map[string]dependencyMetadata)
}

type dependencyMetadata struct {
	value     any
	dependsOn map[string]string
}

var container map[string]dependencyMetadata

// Set sets a dependency with a name and defines its dependencies
func Set[T any](name string, value T, dependsOn map[string]string) {
	if container == nil {
		container = make(map[string]dependencyMetadata)
	}

	container[name] = dependencyMetadata{
		value:     value,
		dependsOn: dependsOn,
	}
}

// Get gets the dependency omitting any error
func Get[T any](name string) T {
	value, _ := GetWithErr[T](name)
	return value
}

// GetWithErr gets a dependency and a posible error
func GetWithErr[T any](name string) (T, error) {
	var value T
	dependency, ok := container[name]
	if !ok {
		return value, ErrNotFound
	}

	value, ok = dependency.value.(T)
	if !ok {
		return value, ErrInvalidType
	}

	return value, nil
}
