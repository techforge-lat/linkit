package role

import (
	"github.com/google/uuid"
)

type UseCase interface {
	Repository
}

type Repository interface {
	Create(user *Role) error
	Update(user *Role) error
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (Role, error)
	List() (Roles, error)
}
