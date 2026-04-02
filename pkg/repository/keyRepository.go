package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/superdima3000/transport-auth/db"
	"strings"
)

type KeyRepository struct {
	db *sqlx.DB
}

func NewKeyRepository(db *sqlx.DB) *KeyRepository {
	return &KeyRepository{db: db}
}

func (r *KeyRepository) Create(key db.Key) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (value, name) VALUES (?, ?) RETURNING id", keysTable)
	row := r.db.QueryRow(query, key.Value, key.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *KeyRepository) GetAll() ([]db.Key, error) {
	var keys []db.Key
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id", keysTable)
	err := r.db.Select(&keys, query)
	return keys, err
}

func (r *KeyRepository) GetById(id int) (db.Key, error) {
	var key db.Key
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", keysTable)
	err := r.db.Get(&key, query, id)
	return key, err
}

func (r *KeyRepository) Update(id int, key db.UpdateKey) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	if key.Value != nil {
		setValues = append(setValues, "value = ?")
		args = append(args, *key.Value)
	}
	if key.Name != nil {
		setValues = append(setValues, "name = ?")
		args = append(args, *key.Name)
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", keysTable, setQuery)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *KeyRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", keysTable)
	_, err := r.db.Exec(query, id)
	return err
}
