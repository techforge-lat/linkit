package linkit

// dependencyContainer is the container type to store all dependencies
type dependencyContainer map[string]options

// container is used to store all dependencies globally
var container dependencyContainer

type options struct {
	// dependencyName will be used as the name for the dependency to be set
	dependencyName string

	// value represents the dependency value
	// If not used, a new value will be created from the provider type
	// by dependor when setting the dependency
	value any

	// auxiliaryDependencies maps an auxiliary dependency with a parent dependency's field
	auxiliaryDependencies map[string]string

	// withCustomSetters sets AuxiliaryDependencies with custom setters
	// that must follow the next format `Set` + `TypeName`
	// e.g.
	// type User struct {
	//      role RoleUseCase
	// }
	//
	// func (u *User) SetRoleUseCase(r RoleUseCase)  {
	//      u.role = r
	// }
	//
	// if set to `false` dependor will look for this auxiliary dependencies
	// base on the value of the AuxiliaryDependencies map
	// but the parent dependency's fields must be public
	withCustomSetters bool
}

type Option func(options *options) error

func WithName(name string) Option {
	return func(options *options) error {
		if name == "" {
			return ErrInvalidDependencyName
		}
		options.dependencyName = name

		return nil
	}
}

func WithValue(value any) Option {
	return func(options *options) error {
		options.value = value
		return nil
	}
}

func WithAuxiliaryDependencies(dependencies map[string]string) Option {
	return func(options *options) error {
		options.auxiliaryDependencies = dependencies
		return nil
	}
}

func WithCustomSetters(state bool) Option {
	return func(options *options) error {
		options.withCustomSetters = state
		return nil
	}
}
