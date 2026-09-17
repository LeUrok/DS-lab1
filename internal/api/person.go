package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/LeUrok/DS-lab1/internal/model"
	"github.com/LeUrok/DS-lab1/internal/service"
)

type PersonHandler struct {
	svc *service.PersonService 
}

func NewPersonHandler(svc *service.PersonService) *PersonHandler {
	return &PersonHandler{svc: svc}
}

func (h *PersonHandler) Create (w http.ResponseWriter, r *http.Request) {
	var req model.PersonRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	id, err := h.svc.Create(req)
	if errors.Is(err, service.ErrInvalidInput) {
		writeError(w, http.StatusBadRequest, "name is required")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	w.Header().Set("Location", "/api/v1/persons/"+strconv.Itoa(int(id)))
	w.WriteHeader(http.StatusCreated)
}

func (h *PersonHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return 
	}
	 p, err := h.svc.GetByID(id)

	 if errors.Is(err, service.ErrNotFound) {
		writeError(w, http.StatusNotFound, "person not found")
		return
	 }

	 if err != nil {
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}
	writeJSON(w, http.StatusOK, p)
}

func (h *PersonHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	persons, err := h.svc.GetAll()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	writeJSON(w, http.StatusOK, persons)
}

func (h *PersonHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req model.PersonRequest
	if err := json.NewDecoder(r.Body).Decode(&req) ; err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	p, err := h.svc.Update(id, req)
	if errors.Is(err, service.ErrNotFound) {
		writeError(w, http.StatusNotFound, "person not found")
		return
	}
	if errors.Is(err, service.ErrInvalidInput) {
		writeError(w, http.StatusBadRequest, "no fields to update")
		return 
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal error")
		return 
	}
	writeJSON(w, http.StatusOK, p)
}

func (h *PersonHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.svc.Delete(id)
	if errors.Is(err, service.ErrNotFound) {
		writeError(w, http.StatusNotFound, "person not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"message": msg})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func parseId (r *http.Request) (int32, error) {
	parts := strings.Split(strings.TrimRight(r.URL.Path, "/"), "/")
	idStr := parts[len(parts) - 1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return int32(id), nil
}