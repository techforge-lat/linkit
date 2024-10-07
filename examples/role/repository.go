package role

import "github.com/google/uuid"

type PsqlRepository struct{}

func NewPsqlRepository() *PsqlRepository {
	return &PsqlRepository{}
}

func (p *PsqlRepository) Create(user *Role) error {
	panic("not implemented") // TODO: Implement
}

func (p *PsqlRepository) Update(user *Role) error {
	panic("not implemented") // TODO: Implement
}

func (p *PsqlRepository) Delete(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (p *PsqlRepository) Get(id uuid.UUID) (Role, error) {
	panic("not implemented") // TODO: Implement
}

func (p *PsqlRepository) List() (Roles, error) {
	panic("not implemented") // TODO: Implement
}
