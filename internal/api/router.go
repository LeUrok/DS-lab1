package api

import (
	"net/http"
)

func NewRouter (h *PersonHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/persons", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetAll(w, r)
		case http.MethodPost:
			h.Create(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)

		}
	})
	
	mux.HandleFunc("/api/v1/persons/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetByID(w, r)
		case http.MethodPatch:
			h.Update(w, r)
		case http.MethodDelete:
			h.Delete(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)

		}	
	})
	return mux
}