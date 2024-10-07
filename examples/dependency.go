package main

import (
	"github.com/techforge-lat/linkit"
	"github.com/techforge-lat/linkit/examples/role"
	"github.com/techforge-lat/linkit/examples/user"
)

func BuildDependencies() (*linkit.DependencyContainer, error) {
	container := linkit.New()

	// user defined dependencies
	userUseCase := user.NewUseCase(user.NewPsqlRepository())
	userHandler := user.NewHandler(userUseCase)

	container.Register(linkit.DependencyName("user.usecase"), userUseCase)
	container.Register(linkit.DependencyName("user.handler"), userHandler)

	// role defined dependencies
	// see that we can defined it after the user who needs it
	roleUseCase := role.NewUseCase()

	// look, we are defining a dependency after the role.usecase who needs it
	// and compared to the user, is not necesary to pass it to the NewUseCase function
	roleRepository := role.NewPsqlRepository()

	container.Register(linkit.DependencyName("role.usecase"), roleUseCase)
	container.Register(linkit.DependencyName("role.repository"), roleRepository)

	// must be after every other root dependency is added
	// this will execute every BuildDependencies function of every root dependency
	if err := container.Build(); err != nil {
		return nil, err
	}

	return container, nil
}
