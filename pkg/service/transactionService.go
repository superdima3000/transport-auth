package service

import (
	"time"

	"github.com/superdima3000/transport-auth/db"
	"github.com/superdima3000/transport-auth/pkg/repository"
)

type TransactionService struct {
	repo         repository.Transaction
	cardRepo     repository.Card
	terminalRepo repository.Terminal
}

func NewTransactionService(repo *repository.Repository) *TransactionService {
	return &TransactionService{
		repo:         repo.Transaction,
		cardRepo:     repo.Card,
		terminalRepo: repo.Terminal,
	}
}

func (s *TransactionService) Create(transaction db.Transaction) (int, error) {
	if transaction.CreatedAt == "" {
		transaction.CreatedAt = time.Now().Format(time.RFC3339)
	}

	if transaction.CardID != nil {
		_, err := s.cardRepo.GetById(int(*transaction.CardID))
		if err != nil {
			return 0, ErrCardNotFound
		}
	}

	if transaction.TerminalID != nil {
		_, err := s.terminalRepo.GetById(int(*transaction.TerminalID))
		if err != nil {
			return 0, ErrTerminalNotFound
		}
	}

	return s.repo.Create(transaction)
}

func (s *TransactionService) Authorize(id int) (string, error) {
	transaction, err := s.repo.GetById(id)
	if err != nil {
		return "", ErrTransactionNotFound
	}

	if transaction.CardID == nil {
		return "", ErrCardNotFound
	}

	card, err := s.cardRepo.GetById(int(*transaction.CardID))
	if err != nil {
		return "", ErrCardNotFound
	}

	if _, err := card.ValidateMifareNumber(); err != nil {
		return "", ErrCardNumberInvalid
	}

	if transaction.Amount <= 0 {
		return "", ErrInvalidAmount
	}

	if card.Balance < transaction.Amount {
		return "declined", ErrNotEnoughFund
	}

	if transaction.TerminalID != nil {
		if _, err := s.terminalRepo.GetById(int(*transaction.TerminalID)); err != nil {
			return "", ErrTerminalNotFound
		}
	}

	return "approved", nil
}

func (s *TransactionService) GetAll() ([]db.Transaction, error) {
	return s.repo.GetAll()
}

func (s *TransactionService) GetById(id int) (db.Transaction, error) {
	return s.repo.GetById(id)
}

func (s *TransactionService) Update(id int, transaction db.UpdateTransaction) error {
	if err := transaction.Validate(); err != nil {
		return err
	}

	if _, err := s.repo.GetById(id); err != nil {
		return ErrTransactionNotFound
	}

	return s.repo.Update(id, transaction)
}

func (s *TransactionService) Delete(id int) error {
	return s.repo.Delete(id)
}
