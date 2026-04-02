package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/superdima3000/transport-auth/db"
	"strings"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user db.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password, is_admin) VALUES (?, ?, ?) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password, user.IsAdmin)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepository) GetAll() ([]db.User, error) {
	var users []db.User
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id", usersTable)
	err := r.db.Select(&users, query)
	return users, err
}

func (r *UserRepository) GetById(id int) (db.User, error) {
	var user db.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", usersTable)
	err := r.db.Get(&user, query, id)
	return user, err
}

func (r *UserRepository) Update(id int, user db.UpdateUser) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	if user.Login != nil {
		setValues = append(setValues, "login = ?")
		args = append(args, *user.Login)
	}
	if user.IsAdmin != nil {
		setValues = append(setValues, "is_admin = ?")
		args = append(args, *user.IsAdmin)
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", usersTable, setQuery)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *UserRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", usersTable)
	_, err := r.db.Exec(query, id)
	return err
}
