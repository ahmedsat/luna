package main

import (
	"errors"
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
	farm := &models.Farm{
		ArabicName:        "مزرعة الحماة",
		EnglishName:       "Hamma Farm",
		EngineerID:        1,
		ManagerID:         1,
		Region:            "القاهرة",
		TotalArea:         4200.00,
		CultivatedArea:    2000.00,
		YearOfReclamation: 2017,
		OwnershipDocument: "123456789"}
	result := db.Create(farm)
	if result.Error != nil {
		return errors.Join(result.Error, errors.New("failed to create farm"))
	}

	fmt.Println(farm.ID)
	fmt.Println(farm.CreatedAt)
	fmt.Println(farm.UpdatedAt)

	return nil
}
