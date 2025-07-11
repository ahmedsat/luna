package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ahmedsat/luna/db"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	err := db.Connect()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer db.DB.Close()

	err = db.DB.Ping()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = handelSubcommands()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Exiting...")
}

func handelSubcommands() error {
	if len(os.Args) < 2 {
		return errors.New("missing subcommand")
	}

	switch os.Args[1] {
	case "migrate":
		return db.Migrate()
	case "seed":
		return db.Seed()
	default:
		return fmt.Errorf("unknown subcommand: %s", os.Args[1])
	}

}
