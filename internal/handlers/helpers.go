package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func readJSON(body io.ReadCloser, out any) error {
	defer body.Close()
	dec := json.NewDecoder(body)
	dec.DisallowUnknownFields()
	return dec.Decode(out)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func checkID(r *http.Request) (int, error) {
	idStr := r.PathValue("id")
	if idStr == "" {
		return 0, errors.New("invalid id")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("invalid id")
	}
	return id, nil
}

func writeBadRequest(w http.ResponseWriter, err error) {
	http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusBadRequest)
}

func writeNotFound(w http.ResponseWriter) {
	http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
}

func writeInternalServerError(w http.ResponseWriter) {
	http.Error(w, `{"error":"internal"}`, http.StatusInternalServerError)
}
