package router

import (
	"net/http"

	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/handler"
)

func New(searchHandler *handler.SearchHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.HealthHandler)
	mux.HandleFunc("GET /search", searchHandler.Search)

	return mux
}
