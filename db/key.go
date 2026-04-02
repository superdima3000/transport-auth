package db

import "errors"

type Key struct {
	ID    int64   `json:"id" db:"id"`
	Value string  `json:"value" db:"value"`
	Name  *string `json:"name" db:"name"`
}

type UpdateKey struct {
	Value *string `json:"value" db:"value"`
	Name  *string `json:"name" db:"name"`
}

func (k UpdateKey) Validate() error {
	if k.Value == nil && k.Name == nil {
		return errors.New("all fields are nil")
	}
	return nil
}
