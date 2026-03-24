package database

import (
	// "context"
	"fmt"
	"os"
	"testing"
)

var testDB *service

func TestMain(m *testing.M) {
	cfg := Config{
		Database: os.Getenv("DB_DATABASE"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
		Schema:   os.Getenv("DB_SCHEMA"),
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Schema)

	// __AUTO_GENERATED_PRINT_VAR_START__
	fmt.Println(fmt.Sprintf("TestMain dbURL: %v", dbURL)) // __AUTO_GENERATED_PRINT_VAR_END__
}

func TestGetMeals_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	// ctx := context.Background()

}
