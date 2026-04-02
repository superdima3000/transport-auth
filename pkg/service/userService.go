package service

import (
	"github.com/superdima3000/transport-auth/db"
	"github.com/superdima3000/transport-auth/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo.User}
}

func (s *UserService) Create(user db.User) (int, error) {
	return s.repo.Create(user)
}

func (s *UserService) GetAll() ([]db.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetById(id int) (db.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Update(id int, user db.UpdateUser) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if _, err := s.repo.GetById(id); err != nil {
		return ErrUserNotFound
	}

	return s.repo.Update(id, user)
}

func (s *UserService) Delete(id int) error {
	return s.repo.Delete(id)
}
