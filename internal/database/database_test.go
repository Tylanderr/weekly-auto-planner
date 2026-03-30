package database

import (
	// "context"
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

	dbService := New(cfg)
	testDB = dbService.(*service)

	code := m.Run()

	dbService.Close()

	os.Exit(code)
}

// func TestGetMeals_Integration(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("skipping integration test")
// 	}
//
// 	ctx := context.Background()
//
// 	// Setup: Insert test data
// 	mealName := "Test Pasta"
// 	mealDesc := "A delicious test pasta"
//
// 	// Insert test meal
// 	var mealID int
// 	err := testDB.db.QueryRowContext(ctx, `
// 		INSERT INTO meals (name, description) 
// 		VALUES ($1, $2) 
// 		RETURNING id
// 	`, mealName, mealDesc).Scan(&mealID)
// 	if err != nil {
// 		t.Fatalf("failed to insert test meal: %v", err)
// 	}
// }
