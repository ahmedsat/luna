package db

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"os"
)

var DB *sql.DB

func Connect() (err error) {
	SFmt := "postgres://avnadmin:%s@%s:%s/%s?sslmode=%s"

	conn, _ := url.Parse(fmt.Sprintf(
		SFmt,
		os.Getenv("PS_PASS"),
		os.Getenv("PS_HOST"),
		os.Getenv("PS_PORT"),
		os.Getenv("PS_DB"),
		os.Getenv("PS_SSL"),
	))

	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem"

	DB, err = sql.Open("postgres", conn.String())

	if err != nil {
		return err
	}
	return
}

func Migrate() error {
	bytes, err := os.ReadFile("sql/schema.sql")
	if err != nil {
		return err
	}
	_, err = DB.Exec(string(bytes))
	if err != nil {
		return err
	}
	return nil
}

func Seed() error {
	return errors.New("not implemented")
}
