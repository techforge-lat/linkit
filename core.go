package linkit

import (
	"fmt"
)

type DependencyContainer struct {
	dependencies map[DependencyName]any
}

func New() *DependencyContainer {
	return &DependencyContainer{
		dependencies: make(map[DependencyName]any),
	}
}

func (d *DependencyContainer) Register(name DependencyName, dependency any) {
	d.dependencies[name] = dependency
}

// Build builds the core with its dependencies
// NOTE: must be called after all root dependencies are registered
func (d *DependencyContainer) Build() error {
	for name, dependency := range d.dependencies {
		rootDependency, ok := dependency.(Dependency)
		if !ok {
			continue
		}

		if err := rootDependency.SetDependencies(d); err != nil {
			return fmt.Errorf("%w %s", ErrCouldNotBuildDependency, name)
		}
	}

	return nil
}

func Get[T any](core *DependencyContainer, name DependencyName) (T, error) {
	var dependency T

	if core == nil {
		return dependency, ErrDependencyCoreNil
	}

	dependencyAbstract, ok := core.dependencies[name]
	if !ok {
		return dependency, fmt.Errorf("%w: %s", ErrNotFound, name)
	}

	rootDependency, ok := dependencyAbstract.(T)
	if !ok {
		return dependency, fmt.Errorf("%w: %s", ErrInvalidType, name)
	}

	dependency = rootDependency

	return dependency, nil
}
