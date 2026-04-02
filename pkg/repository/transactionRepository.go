package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/superdima3000/transport-auth/db"
	"strings"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(transaction db.Transaction) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (amount, card_id, terminal_id, created_at) VALUES (?, ?, ?, ?) RETURNING id", transactionsTable)
	row := r.db.QueryRow(query, transaction.Amount, transaction.CardID, transaction.TerminalID, transaction.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TransactionRepository) GetAll() ([]db.Transaction, error) {
	var transactions []db.Transaction
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY created_at DESC", transactionsTable)
	err := r.db.Select(&transactions, query)
	return transactions, err
}

func (r *TransactionRepository) GetById(id int) (db.Transaction, error) {
	var transaction db.Transaction
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", transactionsTable)
	err := r.db.Get(&transaction, query, id)
	return transaction, err
}

func (r *TransactionRepository) Update(id int, transaction db.UpdateTransaction) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	if transaction.Amount != nil {
		setValues = append(setValues, "amount = ?")
		args = append(args, *transaction.Amount)
	}
	if transaction.CardID != nil {
		setValues = append(setValues, "card_id = ?")
		args = append(args, *transaction.CardID)
	}
	if transaction.TerminalID != nil {
		setValues = append(setValues, "terminal_id = ?")
		args = append(args, *transaction.TerminalID)
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", transactionsTable, setQuery)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TransactionRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", transactionsTable)
	_, err := r.db.Exec(query, id)
	return err
}
