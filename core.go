package dependor

import (
	"errors"
)

var ErrNotFound = errors.New("dependor: dependency not found")
var ErrInvalidType = errors.New("dependor: invalid dependency type")

func init() {
	container = make(map[string]metadata)
}

// metadata stores the dependency value and the dependencies it has
type metadata struct {
	value     any
	dependsOn map[string]string
}

// dependencyContainer is the container type to store all dependencies
type dependencyContainer map[string]metadata

// container is used to store all dependencies globally
var container dependencyContainer

// Set sets a dependency with a name and defines its dependencies
func Set[T any](name string, value T, dependsOn map[string]string) {
	set[T](container, name, value, dependsOn)
}

// set exists to be tested easily by receiving the dependency container
func set[T any](depContainer dependencyContainer, name string, value T, dependsOn map[string]string) {
	if depContainer == nil {
		depContainer = make(dependencyContainer)
	}

	depContainer[name] = metadata{
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
	return getWithErr[T](container, name)
}

// getWithErr exists to be tested easily by receiving the dependency container
func getWithErr[T any](container dependencyContainer, name string) (T, error) {
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
