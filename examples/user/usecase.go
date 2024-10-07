package user

import (
	"github.com/google/uuid"
	"github.com/techforge-lat/linkit"
	"github.com/techforge-lat/linkit/examples/role"
)

type UserUseCase struct {
	repository Repository
	role       RoleUseCase
}

func NewUseCase(repository Repository) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (u UserUseCase) BuildDependencies(container *linkit.DependencyContainer) error {
	roleUseCase, err := linkit.Get[RoleUseCase](container, linkit.DependencyName("role.usecase"))
	if err != nil {
		return err
	}
	u.role = roleUseCase

	return nil
}

func (u *UserUseCase) Create(user *User) error {
	if err := u.repository.Create(user); err != nil {
		return err
	}

	if err := u.role.Create(&role.Role{Name: "admin"}); err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) Update(user *User) error {
	panic("not implemented") // TODO: Implement
}

func (u *UserUseCase) Delete(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (u *UserUseCase) Get(id uuid.UUID) (User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *UserUseCase) List() (Users, error) {
	panic("not implemented") // TODO: Implement
}
