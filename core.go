package linkit

import (
	"log"
	"reflect"
	"sync"
)

var once *sync.Once

func init() {
	once = &sync.Once{}
	setup(&container)
}

// setup initialize a dependency container once
func setup(c *dependencyContainer) {
	once.Do(func() {
		*c = make(map[string]options, 0)
	})
}

// Set sets a dependency with the value type path as a name and defines its dependencies
func Set[T any](opts ...Option) {
	set[T](container, opts...)
}

// set exists to be tested easily by receiving the dependency container
func set[T any](depContainer dependencyContainer, opts ...Option) {
	var config options
	for _, opt := range opts {
		if err := opt(&config); err != nil {
			log.Fatalf("linkit.set(): %v\n", err)
		}
	}

	if depContainer == nil {
		once = &sync.Once{}
		setup(&depContainer)
	}

	if config.value == nil {
		config.value = createValueFromType[T]()
	}

	if config.dependencyName == "" {
		config.dependencyName = Name(config.value)
	}

	depContainer[config.dependencyName] = config
}

// GetWithName gets a dependency with a given name omitting any error
func GetWithName[T any](name string) T {
	value, err := GetWithErr[T](name)
	if err != nil {
		log.Fatalf("linkit: %v, dependencyName: %s", err, name)
	}

	return value
}

// Get gets a dependency using the type path as the name omitting any error
func Get[T any]() T {
	return GetWithName[T](Name(createValueFromType[T]()))
}

// GetWithErr gets a dependency and a possible error
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
