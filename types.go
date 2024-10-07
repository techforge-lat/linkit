package linkit

type Dependency interface {
	ResolveAuxiliaryDependencies(*DependencyContainer) error
}

type DependencyName string
