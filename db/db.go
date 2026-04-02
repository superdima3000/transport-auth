package db

import (
	"embed"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var migrations embed.FS

func New(dsn string) *sqlx.DB {
	db, err := sqlx.Open("sqlite", dsn)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("sqlite"); err != nil {
		logrus.Fatal(err.Error())
	}

	if err := goose.Up(db.DB, "migrations"); err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
