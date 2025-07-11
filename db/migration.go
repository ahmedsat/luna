package db

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/gorm"
)

var MigrationTable = []any{}

func RestDb(db *gorm.DB) error {
	if os.Getenv("ENV") == "development" {
		if err := db.Migrator().DropTable(MigrationTable...); err != nil {
			return errors.Join(err, errors.New("failed to drop migration table"))
		}
	} else {
		return fmt.Errorf("cannot drop table in %s environment", os.Getenv("ENV"))
	}
	return nil
}

func MigrateDb(db *gorm.DB) error {
	if err := db.AutoMigrate(MigrationTable...); err != nil {
		return errors.Join(err, errors.New("failed to migrate database"))
	}
	return nil
}
