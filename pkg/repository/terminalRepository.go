package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/superdima3000/transport-auth/db"
)

type TerminalRepository struct {
	db *sqlx.DB
}

func NewTerminalRepository(db *sqlx.DB) *TerminalRepository {
	return &TerminalRepository{db: db}
}

func (t *TerminalRepository) Create(terminal db.Terminal) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (serial_number, address, name) VALUES (?, ?, ?) RETURNING id", terminalsTable)
	row := t.db.QueryRow(query, terminal.SerialNumber, terminal.Address, terminal.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TerminalRepository) GetAll() ([]db.Terminal, error) {
	var terminals []db.Terminal

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY serial_number", terminalsTable)

	err := t.db.Select(&terminals, query)
	return terminals, err
}

func (t *TerminalRepository) GetById(id int) (db.Terminal, error) {
	var terminal db.Terminal

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", terminalsTable)
	err := t.db.Get(&terminal, query, id)

	return terminal, err
}

func (t *TerminalRepository) Update(id int, terminal db.UpdateTerminal) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if terminal.SerialNumber != nil {
		setValues = append(setValues, "serial_number = ?")
		args = append(args, *terminal.SerialNumber)
	}
	if terminal.Address != nil {
		setValues = append(setValues, "address = ?")
		args = append(args, *terminal.Address)
	}
	if terminal.Name != nil {
		setValues = append(setValues, "name = ?")
		args = append(args, *terminal.Name)
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", terminalsTable, setQuery)

	args = append(args, id)
	_, err := t.db.Exec(query, args...)
	return err
}

func (t *TerminalRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", terminalsTable)
	_, err := t.db.Exec(query, id)

	return err
}
