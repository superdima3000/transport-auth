package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/superdima3000/transport-auth/db"
	"strings"
)

type CardRepository struct {
	db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) *CardRepository {
	return &CardRepository{db: db}
}

func (r *CardRepository) Create(card db.Card) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (card_number, balance, is_blocked, owner_name, key_id) VALUES (?, ?, ?, ?, ?) RETURNING id", cardsTable)
	row := r.db.QueryRow(query, card.CardNumber, card.Balance, card.IsBlocked, card.OwnerName, card.KeyID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CardRepository) GetAll() ([]db.Card, error) {
	var cards []db.Card
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id", cardsTable)
	err := r.db.Select(&cards, query)
	return cards, err
}

func (r *CardRepository) GetById(id int) (db.Card, error) {
	var card db.Card
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", cardsTable)
	err := r.db.Get(&card, query, id)
	return card, err
}

func (r *CardRepository) Update(id int, card db.UpdateCard) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	if card.CardNumber != nil {
		setValues = append(setValues, "card_number = ?")
		args = append(args, *card.CardNumber)
	}
	if card.Balance != nil {
		setValues = append(setValues, "balance = ?")
		args = append(args, *card.Balance)
	}
	if card.IsBlocked != nil {
		setValues = append(setValues, "is_blocked = ?")
		args = append(args, *card.IsBlocked)
	}
	if card.OwnerName != nil {
		setValues = append(setValues, "owner_name = ?")
		args = append(args, *card.OwnerName)
	}
	if card.KeyID != nil {
		setValues = append(setValues, "key_id = ?")
		args = append(args, *card.KeyID)
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", cardsTable, setQuery)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *CardRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", cardsTable)
	_, err := r.db.Exec(query, id)
	return err
}
