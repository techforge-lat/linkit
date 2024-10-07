package user

import "github.com/google/uuid"

type PsqlRepository struct{}

func NewPsqlRepository() *PsqlRepository {
	return &PsqlRepository{}
}

func (p *PsqlRepository) Create(user *User) error {
	panic("not implemented") // TODO: Implement
}

func (p *PsqlRepository) Update(user *User) error {
	panic("not implemented") // TODO: Implement
}

func (p *PsqlRepository) Delete(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (p *PsqlRepository) Get(id uuid.UUID) (User, error) {
	panic("not implemented") // TODO: Implement
}

func (p *PsqlRepository) List() (Users, error) {
	panic("not implemented") // TODO: Implement
}
