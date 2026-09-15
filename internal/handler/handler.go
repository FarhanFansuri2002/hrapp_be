package handler

import (
	"encoding/json"
	"net/http"

	"server/internal/service"
)

type Handler struct {
	service *service.HRService
}

func New(hrService *service.HRService) *Handler {
	return &Handler{service: hrService}
}

func (handler *Handler) Health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (handler *Handler) ListEmployees(w http.ResponseWriter, request *http.Request) {
	employees, err := handler.service.ListEmployees(request.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to load employees"})
		return
	}

	writeJSON(w, http.StatusOK, employees)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
