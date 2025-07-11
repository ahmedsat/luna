package main

import (
	"fmt"
	"os"

	"github.com/ahmedsat/luna/db"
	"github.com/ahmedsat/luna/models"
	"gorm.io/gorm"
)

func main() {

	db, err := db.Connect()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = handelSubCommands(db)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func handelSubCommands(gdb *gorm.DB) error {
	if len(os.Args) < 2 {
		return fmt.Errorf("please provide a sub-command")
	}

	switch os.Args[1] {
	default:
		return fmt.Errorf("invalid sub-command")
	case "migrate":
		return db.MigrateDb(gdb)
	case "reset":
		return db.RestDb(gdb)
	case "test":
		return test(gdb)
	}
}

func test(db *gorm.DB) error {
	farm := &models.Farm{}
	err := db.Last(farm).Error
	if err != nil {
		return err
	}

	fmt.Println(farm)

	return nil
}
