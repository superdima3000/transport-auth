package db

import "errors"

type User struct {
	Id       int64  `json:"id"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsAdmin  int64  `json:"is_admin" db:"is_admin"`
}

type UpdateUser struct {
	Login   *string `json:"login" binding:"required"`
	IsAdmin *int64  `json:"is_admin" db:"is_admin"`
}

func (k UpdateUser) Validate() error {
	if k.Login == nil && k.IsAdmin == nil {
		return errors.New("all fields are nil")
	}
	return nil
}
