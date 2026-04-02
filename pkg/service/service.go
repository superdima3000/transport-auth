package service

import (
	"github.com/superdima3000/transport-auth/db"
	"github.com/superdima3000/transport-auth/pkg/repository"
)

type Authorization interface {
	CreateUser(user db.User) (int, error)
	GenerateToken(login string, password string) (string, error)
	ParseToken(token string) (*TokenClaims, error)
}

type Key interface {
	Create(key db.Key) (int, error)
	GetAll() ([]db.Key, error)
	GetById(id int) (db.Key, error)
	Update(id int, key db.UpdateKey) error
	Delete(id int) error
}

type Card interface {
	Create(card db.Card) (int, error)
	GetAll() ([]db.Card, error)
	GetById(id int) (db.Card, error)
	Update(id int, card db.UpdateCard) error
	Delete(id int) error
}

type Terminal interface {
	Create(terminal db.Terminal) (int, error)
	GetAll() ([]db.Terminal, error)
	GetById(id int) (db.Terminal, error)
	Update(id int, terminal db.UpdateTerminal) error
	Delete(id int) error
}

type Transaction interface {
	Create(transaction db.Transaction) (int, error)
	GetAll() ([]db.Transaction, error)
	GetById(id int) (db.Transaction, error)
	Update(id int, transaction db.UpdateTransaction) error
	Delete(id int) error
	Authorize(id int) (string, error)
}

type User interface {
	Create(user db.User) (int, error)
	GetAll() ([]db.User, error)
	GetById(id int) (db.User, error)
	Update(id int, user db.UpdateUser) error
	Delete(id int) error
}

type Service struct {
	Authorization
	Key
	Card
	Terminal
	Transaction
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Terminal:      NewTerminalService(repos),
		User:          NewUserService(repos),
		Key:           NewKeyService(repos),
		Card:          NewCardService(repos),
		Transaction:   NewTransactionService(repos),
	}
}
