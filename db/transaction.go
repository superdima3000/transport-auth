package db

import "errors"

type Transaction struct {
	ID         int64   `json:"id" db:"id"`
	Amount     float64 `json:"amount" db:"amount" binding:"required"`
	CardID     *int64  `json:"card_id" db:"card_id"`
	TerminalID *int64  `json:"terminal_id" db:"terminal_id"`
	CreatedAt  string  `json:"created_at" db:"created_at"`
}

type UpdateTransaction struct {
	Amount     *float64 `json:"amount" db:"amount"`
	CardID     *int64   `json:"card_id" db:"card_id"`
	TerminalID *int64   `json:"terminal_id" db:"terminal_id"`
}

func (t UpdateTransaction) Validate() error {
	if t.Amount == nil && t.CardID == nil && t.TerminalID == nil {
		return errors.New("all fields are nil")
	}
	return nil
}
