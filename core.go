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

func (d *DependencyContainer) Provide(name DependencyName, dependency any) {
	d.dependencies[name] = dependency
}

// ResolveAuxiliaryDependencies builds the core with its dependencies
// NOTE: must be called after all root dependencies are registered
func (d *DependencyContainer) ResolveAuxiliaryDependencies() error {
	for name, dependency := range d.dependencies {
		rootDependency, ok := dependency.(Dependency)
		if !ok {
			continue
		}

		if err := rootDependency.ResolveAuxiliaryDependencies(d); err != nil {
			return fmt.Errorf("%w %s. %s", ErrCouldNotBuildDependency, name, err.Error())
		}
	}

	return nil
}

func Resolve[T any](core *DependencyContainer, name DependencyName) (T, error) {
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
