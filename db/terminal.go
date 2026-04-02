package db

import "errors"

type Terminal struct {
	Id           int64  `json:"id" db:"id"`
	SerialNumber string `json:"serial_number" db:"serial_number"`
	Address      string `json:"address" db:"address"`
	Name         string `json:"name" db:"name"`
}

type UpdateTerminal struct {
	SerialNumber *string `json:"serial_number" db:"serial_number"`
	Address      *string `json:"address" db:"address"`
	Name         *string `json:"name" db:"name"`
}

func (t UpdateTerminal) Validate() error {
	if t.SerialNumber == nil && t.Address == nil && t.Name == nil {
		return errors.New("all fields are nil")
	}
	return nil
}
