package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/superdima3000/transport-auth/db"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user db.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password, is_admin) values (?, ?, ?) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Login, user.Password, user.IsAdmin)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthRepository) GetUserByUsernameAndPassword(username string, password string) (db.User, error) {
	var user db.User

	query := fmt.Sprintf("SELECT id, login, password, is_admin FROM %s WHERE login = ? AND password = ?", usersTable)
	err := r.db.Get(&user, query, username, password)

	logrus.Infof("user: %v", user)

	return user, err
}
