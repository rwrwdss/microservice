package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/service"
)

type SearchHandler struct {
	service service.SearchService
}

func NewSearchHandler(service service.SearchService) *SearchHandler {
	return &SearchHandler{service: service}
}

func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	results, err := h.service.Search(r.Context(), query)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}
