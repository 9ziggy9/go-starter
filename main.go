package main

// WARNING: gorilla is in archived now and no longer maintained!
// TODO: csrf protection via gorilla seems slightly suss, research pls.
// I should probably switch to gin immediately.
import (
	"fmt"
	"os"
	"errors"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
	"github.com/9ziggy9/go-starter/config"
	"github.com/9ziggy9/go-starter/schema"
	"github.com/9ziggy9/go-starter/seeders"
)

func main() {
	// HANDLE CLI ARGS
	arg, err := processArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}

	// LOAD ENV AND SET PORT
	if err := config.LoadEnv(".env"); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file: %s\n", err)
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Fprintf(os.Stderr, "Error loading port env variable.\n")
		os.Exit(1)
	}

	// OPEN DB CONNECTION
	dsn := config.BuildDSN(
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_PORT"),
	)

	fmt.Printf("\n\nCONNECTING TO DB VIA DSN\n%s\n\n\n", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening DB connection: %s\n", err)
		os.Exit(1)
	}

	// TODO: replace with long opts
	switch arg {
	case "seed":
		db.AutoMigrate(&schema.User{})
		db.Create(seeders.Users)
	case "purge":
		db.Migrator().DropTable(&schema.User{})
		fallthrough
	default:
		db.AutoMigrate(&schema.User{})
	}
}

func processArgs() (string, error) {
	args := os.Args

	if len(args) == 1 {
		return "", nil
	}

	if len(args) > 2 {
		fmt.Fprintf(os.Stderr, "Warning: user supplied too many CLI args.\n")
		fmt.Println("API will only use the first provided.")
	}

	arg := args[1]

	validArgs := [2]string{"seed", "purge"}
	for _, validArg := range validArgs {
		if arg == validArg {
			return arg, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Invalid arg provided -- %s\n", arg))
}
