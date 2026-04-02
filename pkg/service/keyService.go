package service

import (
	"github.com/superdima3000/transport-auth/db"
	"github.com/superdima3000/transport-auth/pkg/repository"
)

type KeyService struct {
	repo repository.Key
}

func NewKeyService(repo *repository.Repository) *KeyService {
	return &KeyService{repo: repo.Key}
}

func (s *KeyService) Create(key db.Key) (int, error) {
	return s.repo.Create(key)
}

func (s *KeyService) GetAll() ([]db.Key, error) {
	return s.repo.GetAll()
}

func (s *KeyService) GetById(id int) (db.Key, error) {
	return s.repo.GetById(id)
}

func (s *KeyService) Update(id int, key db.UpdateKey) error {
	if err := key.Validate(); err != nil {
		return err
	}

	if _, err := s.repo.GetById(id); err != nil {
		return ErrKeyNotFound
	}

	return s.repo.Update(id, key)
}

func (s *KeyService) Delete(id int) error {
	return s.repo.Delete(id)
}
