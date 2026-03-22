package photo

import (
	"encoding/json"
	"io"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) EstimateCalories(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "file too large or invalid form", http.StatusBadRequest)
		return
	}

	userID := r.FormValue("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	mealType := r.FormValue("meal_type")
	if mealType == "" {
		http.Error(w, "meal_type is required", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "photo is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error reading photo", http.StatusInternalServerError)
		return
	}

	log, err := h.service.EstimateCalories(userID, fileBytes, header.Filename, mealType)
	if err != nil {
		http.Error(w, "error estimating calories", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(log)
}

func (h *Handler) GetHistory(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	logs, err := h.service.GetHistory(userID)
	if err != nil {
		http.Error(w, "error fetching photo history", http.StatusInternalServerError)
		return
	}

	if logs == nil {
		logs = []PhotoLog{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}
