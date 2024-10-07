package role

import (
	"github.com/google/uuid"
	"github.com/techforge-lat/linkit"
)

type RoleUseCase struct {
	repository Repository
}

func NewUseCase() *RoleUseCase {
	return &RoleUseCase{}
}

func (u RoleUseCase) ResolveAuxiliaryDependencies(container *linkit.DependencyContainer) error {
	repository, err := linkit.Resolve[Repository](container, linkit.DependencyName("role.repository"))
	if err != nil {
		return err
	}
	u.repository = repository

	return nil
}

func (u *RoleUseCase) Create(user *Role) error {
	panic("not implemented") // TODO: Implement
}

func (u *RoleUseCase) Update(user *Role) error {
	panic("not implemented") // TODO: Implement
}

func (u *RoleUseCase) Delete(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (u *RoleUseCase) Get(id uuid.UUID) (Role, error) {
	panic("not implemented") // TODO: Implement
}

func (u *RoleUseCase) List() (Roles, error) {
	panic("not implemented") // TODO: Implement
}
