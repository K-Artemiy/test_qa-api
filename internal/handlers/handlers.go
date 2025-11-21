package handlers

import (
	"net/http"
	"strings"

	"test_qa-api/internal/models"
	"test_qa-api/internal/services"
)

type Handlers struct {
	svc *services.QAService
}

func NewHandlers(svc *services.QAService) *Handlers {
	return &Handlers{svc: svc}
}

func (h *Handlers) ListQuestions(w http.ResponseWriter, r *http.Request) {
	qs, err := h.svc.ListQuestions()
	if err != nil {
		writeInternalServerError(w)
		return
	}
	writeJSON(w, http.StatusOK, qs)
}

func (h *Handlers) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var req models.QuestionDTO
	if err := readJSON(r.Body, &req); err != nil {
		writeBadRequest(w, err)
		return
	}
	created, err := h.svc.CreateQuestion(&req)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

func (h *Handlers) GetQuestion(w http.ResponseWriter, r *http.Request) {
	id, err := checkID(r)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	q, err := h.svc.GetQuestion(id)
	if err != nil {
		writeNotFound(w)
		return
	}
	writeJSON(w, http.StatusOK, q)
}

func (h *Handlers) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	id, err := checkID(r)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	if err := h.svc.DeleteQuestion(id); err != nil {
		writeNotFound(w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	id, err := checkID(r)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	var req models.AnswerDTO
	if err := readJSON(r.Body, &req); err != nil {
		writeBadRequest(w, err)
		return
	}
	created, err := h.svc.CreateAnswer(id, &req)
	if err != nil {
		if strings.Contains(err.Error(), "question not found") {
			writeNotFound(w)
			return
		}
		writeBadRequest(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

func (h *Handlers) GetAnswer(w http.ResponseWriter, r *http.Request) {
	id, err := checkID(r)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	a, err := h.svc.GetAnswer(id)
	if err != nil {
		writeNotFound(w)
		return
	}
	writeJSON(w, http.StatusOK, a)
}

func (h *Handlers) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	id, err := checkID(r)
	if err != nil {
		writeBadRequest(w, err)
		return
	}
	if err := h.svc.DeleteAnswer(id); err != nil {
		writeNotFound(w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
