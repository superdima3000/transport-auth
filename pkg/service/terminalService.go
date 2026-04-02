package service

import (
	"github.com/superdima3000/transport-auth/db"
	"github.com/superdima3000/transport-auth/pkg/repository"
)

type TerminalService struct {
	repo repository.Terminal
}

func NewTerminalService(repo *repository.Repository) *TerminalService {
	return &TerminalService{repo: repo.Terminal}
}

func (t *TerminalService) Create(terminal db.Terminal) (int, error) {
	return t.repo.Create(terminal)
}

func (t *TerminalService) GetAll() ([]db.Terminal, error) {
	return t.repo.GetAll()
}

func (t *TerminalService) GetById(id int) (db.Terminal, error) {
	return t.repo.GetById(id)
}

func (t *TerminalService) Update(id int, terminal db.UpdateTerminal) error {
	if err := terminal.Validate(); err != nil {
		return err
	}

	if _, err := t.repo.GetById(id); err != nil {
		return ErrTerminalNotFound
	}

	return t.repo.Update(id, terminal)
}

func (t *TerminalService) Delete(id int) error {
	return t.repo.Delete(id)
}
