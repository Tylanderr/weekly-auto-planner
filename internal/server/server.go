package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/tylanderr/home-operations-portal/internal/database"
)

type Server struct {
	port int
	db   database.Service
}

// Returning a pointer to a struct to allow modification to the structs fields by the caller
// A Server struct has a Handler, which contains ResponseWriters and Requests
func NewServer() *http.Server {
	//Atoi is equivalent to ParseInt. PORT is defined in the .env file
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	cfg := database.Config{
		Database: os.Getenv("DB_DATABASE"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port: os.Getenv("DB_PORT"),
		Host: os.Getenv("DB_HOST"),
		Schema: os.Getenv("DB_SCHEMA"),
	}

	NewServer := &Server{
		port: port,
		db:   database.New(cfg),
	}

	// Declare Server config
	// Return a pointer to the struct
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
