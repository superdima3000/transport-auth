package service

import (
	"github.com/superdima3000/transport-auth/db"
	"github.com/superdima3000/transport-auth/pkg/repository"
)

type CardService struct {
	repo    repository.Card
	keyRepo repository.Key
}

func NewCardService(repo *repository.Repository) *CardService {
	return &CardService{
		repo:    repo.Card,
		keyRepo: repo.Key,
	}
}

func (s *CardService) Create(card db.Card) (int, error) {
	if card.KeyID != nil {
		_, err := s.keyRepo.GetById(int(*card.KeyID))
		if err != nil {
			return 0, ErrKeyNotFound
		}
	}

	return s.repo.Create(card)
}

func (s *CardService) GetAll() ([]db.Card, error) {
	return s.repo.GetAll()
}

func (s *CardService) GetById(id int) (db.Card, error) {
	return s.repo.GetById(id)
}

func (s *CardService) Update(id int, card db.UpdateCard) error {
	if err := card.Validate(); err != nil {
		return err
	}

	if _, err := s.repo.GetById(id); err != nil {
		return ErrCardNotFound
	}

	if card.KeyID != nil {
		_, err := s.keyRepo.GetById(int(*card.KeyID))
		if err != nil {
			return ErrKeyNotFound
		}
	}

	return s.repo.Update(id, card)
}

func (s *CardService) Delete(id int) error {
	return s.repo.Delete(id)
}
