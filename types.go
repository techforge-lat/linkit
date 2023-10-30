package dependor

// dependencyContainer is the container type to store all dependencies
type dependencyContainer map[string]Config

// container is used to store all dependencies globally
var container dependencyContainer

type Config struct {
	// DependencyName will be use as the name for the dependency to be set
	DependencyName string

	// Value represents the dependency value
	// If not used, a new value will be created from the provider type
	// by dependor when setting the dependency
	Value any

	// AuxiliaryDependencies maps an auxiliary dependency with a parent dependency's field
	AuxiliaryDependencies map[string]string

	// WithCustomSetters sets AuxiliaryDependencies with custom setters
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
	WithCustomSetters bool
}
