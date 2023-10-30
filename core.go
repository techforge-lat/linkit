package dependor

import (
	"errors"
	"log"
	"reflect"
	"sync"
)

var ErrNotFound = errors.New("dependency not found")
var ErrInvalidType = errors.New("invalid dependency type")

var once *sync.Once

func init() {
	once = &sync.Once{}
	setup(&container)
}

// setup initialize a dependency container once
func setup(c *dependencyContainer) {
	once.Do(func() {
		*c = make(map[string]Config, 0)
	})
}

// Set sets a dependency with the value type path as a name and defines its dependencies
func Set[T any](config Config) {
	set[T](container, config)
}

// set exists to be tested easily by receiving the dependency container
func set[T any](depContainer dependencyContainer, config Config) {
	if depContainer == nil {
		once = &sync.Once{}
		setup(&depContainer)
	}

	if config.Value == nil {
		config.Value = createValueFromType[T]()
	}

	if config.DependencyName == "" {
		config.DependencyName = Name(config.Value)
	}

	depContainer[config.DependencyName] = config
}

// GetWithName gets a dependency with a given name omitting any error
func GetWithName[T any](name string) T {
	value, err := GetWithErr[T](name)
	if err != nil {
		log.Fatalf("dependor: %v, dependencyName: %s", err, name)
	}

	return value
}

// Get gets a dependency using the type path as the name omitting any error
func Get[T any]() T {
	return GetWithName[T](Name(createValueFromType[T]()))
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

	value, ok = dependency.Value.(T)
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

func createValueFromType[T any]() T {
	var value T
	t := reflect.TypeOf(value)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	} else {
		// the populate func will throw an error if a pointer is needed for this value
		return value
	}

	return reflect.New(t).Interface().(T)
}
