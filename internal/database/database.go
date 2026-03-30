package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	// "os"
	"strconv"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Service interface {
	Health() map[string]string
	Close() error

	AddNewUser(email string) map[string]string
	GetMeals(limit int) ([]Meal, error)
}

type service struct {
	db *sql.DB
}

type User struct {
	Id        uuid.UUID
	Email     string
	FirstName string
	LastName  string
}

type Meal struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

var	dbInstance *service

type Config struct {
	Database string
	Password string
	Username string
	Port string
	Host string
	Schema string
}

func New(cfg Config) Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Schema)

	db, err := sql.Open("pgx", connStr)

	if err != nil {
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}

	return dbInstance
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("%s", fmt.Sprintf("db down: %v", err))
		return stats
	}

	stats["status"] = "up"
	stats["message"] = "It's healthy"

	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	if dbStats.OpenConnections > 40 {
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

func (s *service) Close() error {
	return s.db.Close()
}

func (s *service) AddNewUser(email string) map[string]string {
	status := make(map[string]string)

	query := `INSERT INTO users (email) VALUES ($1)`
	_, err := s.db.Exec(query, email)

	if err != nil {
		log.Fatal(err)
	}

	status["write_successful"] = "true"

	return status
}

func (s *service) GetMeals(limit int) ([]Meal, error) {
	query := `
	SELECT 
		m.id,
		m.name,
		m.description,
		COALESCE(
			json_agg(
				json_build_object(
					'id', i.id,
					'name', i.name,
					'quantity', mi.quantity,
					'unit', mi.unit
				)
			) FILTER (WHERE i.id IS NOT NULL), 
			'[]'
		) as ingredients
	FROM meals m
	LEFT JOIN meal_ingredients mi ON m.id = mi.meal_id
	LEFT JOIN ingredients i ON mi.ingredient_id = i.id
	GROUP BY m.id, m.name, m.description
	ORDER BY m.id
	LIMIT $1
	`

	rows, err := s.db.Query(query, limit)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	defer rows.Close()

	var meals []Meal

	for rows.Next() {
		var m Meal
		var ingredientsJSON []byte

		err := rows.Scan(&m.Id, &m.Name, &m.Description, &ingredientsJSON)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		if len(ingredientsJSON) > 0 {
			err = json.Unmarshal(ingredientsJSON, &m.Ingredients)
			if err != nil {
				return nil, fmt.Errorf("unmarshal ingredients failed: %w", err)
			}
		}

		meals = append(meals, m)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return meals, nil
}
