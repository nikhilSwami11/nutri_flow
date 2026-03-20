package sessions

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) StartSession(w http.ResponseWriter, r *http.Request) {
	var body struct {
		UserID     string `json:"user_id"`
		RecipeID   string `json:"recipe_id"`
		RecipeName string `json:"recipe_name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if body.UserID == "" || body.RecipeID == "" || body.RecipeName == "" {
		http.Error(w, "user_id, recipe_id, and recipe_name are required", http.StatusBadRequest)
		return
	}

	session, err := h.service.StartSession(body.UserID, body.RecipeID, body.RecipeName)
	if err != nil {
		http.Error(w, "error starting session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(session)
}

func (h *Handler) CompleteSession(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var body struct {
		ServingsEaten      float64 `json:"servings_eaten"`
		CaloriesPerServing float64 `json:"calories_per_serving"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CompleteSession(id, body.ServingsEaten, body.CaloriesPerServing); err != nil {
		http.Error(w, "error completing session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) AbandonSession(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	if err := h.service.AbandonSession(id); err != nil {
		http.Error(w, "error abandoning session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetHistory(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	history, err := h.service.GetHistory(userID)
	if err != nil {
		http.Error(w, "error fetching session history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}
