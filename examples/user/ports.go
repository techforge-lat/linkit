package user

import (
	"github.com/google/uuid"
	"github.com/techforge-lat/linkit/examples/role"
)

type UseCase interface {
	Repository
}

type Repository interface {
	Create(user *User) error
	Update(user *User) error
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (User, error)
	List() (Users, error)
}

type RoleUseCase interface {
	Create(role *role.Role) error
}
