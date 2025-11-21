package server

import (
	"net/http"

	"gorm.io/gorm"

	"test_qa-api/internal/handlers"
	"test_qa-api/internal/repository"
	"test_qa-api/internal/services"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer(db *gorm.DB) *Server {
	mux := http.NewServeMux()

	repo := repository.NewRepo(db)
	svc := services.NewQAService(repo)
	h := handlers.NewHandlers(svc)

	mux.HandleFunc("GET /questions", h.ListQuestions)
	mux.HandleFunc("POST /questions", h.CreateQuestion)
	mux.HandleFunc("GET /questions/{id}", h.GetQuestion)
	mux.HandleFunc("DELETE /questions/{id}", h.DeleteQuestion)

	mux.HandleFunc("POST /questions/{id}/answers", h.CreateAnswer)
	mux.HandleFunc("GET /answers/{id}", h.GetAnswer)
	mux.HandleFunc("DELETE /answers/{id}", h.DeleteAnswer)

	return &Server{mux: mux}
}

func (s *Server) GetServerMux() *http.ServeMux {
	return s.mux
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}
