package dependor

import (
	"errors"
	"reflect"
	"sync"
)

var ErrNotFound = errors.New("dependor: dependency not found")
var ErrInvalidType = errors.New("dependor: invalid dependency type")

var once *sync.Once

func init() {
	once = &sync.Once{}
	setup(container)
}

// setup initialize a dependency container once
func setup(c dependencyContainer) {
	once.Do(func() {
		c = make(map[string]metadata)
	})
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

// SetWithName sets a dependency with a name and defines its dependencies
func SetWithName[T any](name string, value T, dependsOn map[string]string) {
	set[T](container, name, value, dependsOn)
}

// Set sets a dependency with the value type path as a name and defines its dependencies
func Set[T any](value T, dependsOn map[string]string) {
	set[T](container, Name(value), value, dependsOn)
}

// set exists to be tested easily by receiving the dependency container
func set[T any](depContainer dependencyContainer, name string, value T, dependsOn map[string]string) {
	if depContainer == nil {
		once = &sync.Once{}
		setup(depContainer)
	}

	depContainer[name] = metadata{
		value:     value,
		dependsOn: dependsOn,
	}
}

// GetWithName gets a dependency with a given name omitting any error
func GetWithName[T any](name string) T {
	value, _ := GetWithErr[T](name)
	return value
}

// Get gets a dependency using the type path as the name omitting any error
func Get[T any]() T {
	var value T
	return GetWithName[T](Name(value))
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

// Name gets a dependency name by using its type path
func Name(v any) string {
	value := reflect.ValueOf(v)
	if isPointer(value) {
		value = value.Elem()
	}

	return value.Type().String()
}
