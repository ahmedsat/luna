package db

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/joho/godotenv/autoload"
)

func Connect() (db *gorm.DB, err error) {
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

	var log logger.Interface
	if os.Getenv("ENV") == "development" {
		log = logger.Default.LogMode(logger.Warn)
	} else {
		log = logger.Default.LogMode(logger.Silent)
	}

	db, err = gorm.Open(postgres.Open(conn.String()), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		return db, errors.Join(err, errors.New("failed to connect to database"))
	}

	return
}
