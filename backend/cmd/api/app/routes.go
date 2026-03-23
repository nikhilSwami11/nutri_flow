package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *App) RegisterRoutes(r *chi.Mux) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	r.Route("/pantry", func(r chi.Router) {
		r.Get("/", a.pantryHandler.GetAll)
		r.Post("/", a.pantryHandler.Create)
		r.Put("/{id}", a.pantryHandler.Update)
		r.Delete("/{id}", a.pantryHandler.Delete)
	})

	r.Route("/sessions", func(r chi.Router) {
		r.Post("/start", a.sessionsHandler.StartSession)
		r.Post("/{id}/abandon", a.sessionsHandler.AbandonSession)
	})

	r.Route("/recipes", func(r chi.Router) {
		r.Get("/suggestions", a.recipesHandler.GetSuggestions)
		r.Get("/", a.recipesHandler.GetUserRecipes)
		r.Post("/", a.recipesHandler.SaveRecipe)
		r.Delete("/{id}", a.recipesHandler.DeleteRecipe)
	})

	r.Route("/profile", func(r chi.Router) {
		r.Get("/", a.profileHandler.Get)
		r.Post("/", a.profileHandler.Create)
		r.Put("/", a.profileHandler.Update)
	})

	r.Route("/photo", func(r chi.Router) {
		r.Post("/estimate", a.photoHandler.EstimateCalories)
		r.Get("/history", a.photoHandler.GetHistory)
	})
}
