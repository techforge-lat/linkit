package linkit

type Dependency interface {
	SetDependencies(*DependencyContainer) error
}

type DependencyName string
