package db

import (
	"errors"
	"regexp"
	"strings"
)

type Card struct {
	ID         int64   `json:"id"  db:"id"`
	CardNumber string  `json:"card_number" db:"card_number"`
	Balance    float64 `json:"balance" db:"balance"`
	IsBlocked  bool    `json:"is_blocked" db:"is_blocked"`
	OwnerName  *string `json:"owner_name" db:"owner_name"`
	KeyID      *int64  `json:"key_id" db:"key_id"`
}

type UpdateCard struct {
	CardNumber *string  `json:"card_number" db:"card_number"`
	Balance    *float64 `json:"balance" db:"balance"`
	IsBlocked  *bool    `json:"is_blocked" db:"is_blocked"`
	OwnerName  *string  `json:"owner_name" db:"owner_name"`
	KeyID      *int64   `json:"key_id" db:"key_id"`
}

func (k UpdateCard) Validate() error {
	if k.CardNumber == nil && k.Balance == nil && k.OwnerName == nil && k.KeyID == nil && k.IsBlocked == nil {
		return errors.New("all fields are nil")
	}
	return nil
}

var (
	hexUIDRe = regexp.MustCompile(`^[0-9A-Fa-f]+$`)
	decUIDRe = regexp.MustCompile(`^[0-9]+$`)
)

type UIDFormat string

const (
	UIDUnknown UIDFormat = "unknown"
	UIDHex4    UIDFormat = "hex_4_byte"
	UIDHex7    UIDFormat = "hex_7_byte"
	UIDHex10   UIDFormat = "hex_10_byte"
	UIDDec     UIDFormat = "decimal"
)

func (k Card) ValidateMifareNumber() (UIDFormat, error) {
	s := k.CardNumber
	v := strings.TrimSpace(s)
	v = strings.ReplaceAll(v, " ", "")
	v = strings.ReplaceAll(v, "-", "")
	v = strings.ReplaceAll(v, ":", "")

	if v == "" {
		return UIDUnknown, errors.New("invalid card number")
	}

	if hexUIDRe.MatchString(v) {
		switch len(v) {
		case 8:
			return UIDHex4, nil
		case 14:
			return UIDHex7, nil
		case 20:
			return UIDHex10, nil
		}
	}

	if decUIDRe.MatchString(v) {
		l := len(v)
		if l >= 8 && l <= 16 {
			return UIDDec, nil
		}
	}

	return UIDUnknown, errors.New("invalid card number")
}
