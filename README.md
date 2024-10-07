# Languages

- [Español](./README.es.md)
- [English](./README.md)

# Linkit
**Linkit** is a minimalistic dependency injection container for Go. It simplifies the management of types and dependencies in Go projects by collecting and initializing types.

## Motivation
In previous projects, managing dependencies lacked a comprehensive and standardized approach, making it confusing, especially for newcomers. Each project followed different patterns, which led to inconsistency and complexity in managing dependencies across teams.

**Linkit** solves this by providing a clear and defined way to handle dependencies, ensuring a consistent and structured approach that is easy to understand and maintain. It reduces confusion, speeds up onboarding, and streamlines the dependency injection process for all developers.

## Concepts

- **Container**: A collection of initialized types.
- **Root Dependency**: The main dependency in a context (handler, usecase, repository, etc).
- **Auxiliary Dependency**: Root dependencies that are required by another root dependencies.

### How to use it 
Linkit provides a 3 steps workflow

1. **Initialize the container**: Create a container instance to manage your application's dependencies.
```go
    container := linkit.New()
```

2. **Provide dependencies**: Register dependencies using the Provide method. This registers a function that will provide the dependency when needed.
```go
	userUseCase := user.NewUseCase(user.NewPsqlRepository())
	userHandler := user.NewHandler(userUseCase)

    // to provide we need to define a naming convention for every dependency name
    // here we are using `module-name.layer`= 'user.usecase'
    // this names can be defined in a single file
	container.Provide(linkit.DependencyName("user.usecase"), userUseCase)
	container.Provide(linkit.DependencyName("user.handler"), userHandler)


	// must be after every other root dependency is provided
	// this will execute every ResolveAuxiliaryDependencies method of every dependency
	if err := container.ResolveAuxiliaryDependencies(); err != nil {
		return nil, err
	}
```

3. **Resolve dependencies**: Use the `Resolve` function to retrieve a provided dependency.
```go
    // UserUseCase is a root dependency
    type UserUseCase struct {
        // repository is an auxiliary dependency which is set by the constructor
        repository Repository

        // role is an auxiliary dependency which is set by the ResolveAuxiliaryDependencies method required by linkit
        role       RoleUseCase
    }

    // NewUseCase initialize the root dependency
    func NewUseCase(repository Repository) *UserUseCase {
        return &UserUseCase{
            repository: repository,
        }
    }

    // ResolveAuxiliaryDependencies sets the auxiliary dependencies, 
    // this method is called by the ResolveAuxiliaryDependencies method of the linkit container created in the first step
    // after every root dependency has been provided so we don't run into an error when we resolve them here
    func (u UserUseCase) ResolveAuxiliaryDependencies(container *linkit.DependencyContainer) error {
        roleUseCase, err := linkit.Resolve[RoleUseCase](container, linkit.DependencyName("role.usecase"))
        if err != nil {
            return err
        }
        u.role = roleUseCase

        return nil
    }
```

*see the [examples/](./examples/) folder for more details.*

## Limitations or Known Issues
- **Pointers Requirement**: Dependencies must be passed as pointers, which may require extra care when defining and using them.
- **Manual Naming**: Each dependency must be provided with a name. While this allows flexibility, it can lead to misnaming errors, as the names aren't "typed" or enforced by the compiler. If you attempt to resolve a dependency with the wrong name, it will throw an error.
