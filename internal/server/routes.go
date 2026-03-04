package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
}

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	return mux
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
