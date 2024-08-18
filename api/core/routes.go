package core

import (
	"api/ctrl"
	"net/http"
)

// InitRoutes is a centralized place to map all of our HTTP endpoints to controller methods.
func InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /v1/questions", ctrl.GetQuestions())
	mux.HandleFunc("PUT /v1/questions", ctrl.Answer)
}
