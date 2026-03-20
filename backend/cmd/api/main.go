package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/nikhilswami11/nutriflow/backend/internal/pantry"
	"github.com/nikhilswami11/nutriflow/backend/internal/profile"
	"github.com/nikhilswami11/nutriflow/backend/pkg/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	database := db.Connect()

	pantryRepo := pantry.NewRepository(database)
	pantryHandler := pantry.NewHandler(pantryRepo)

	profileRepo := profile.NewRepository(database)
	profileHandler := profile.NewHandler(profileRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	r.Route("/pantry", func(r chi.Router) {
		r.Get("/", pantryHandler.GetAll)
		r.Post("/", pantryHandler.Create)
		r.Put("/{id}", pantryHandler.Update)
		r.Delete("/{id}", pantryHandler.Delete)
	})

	r.Route("/profile", func(r chi.Router) {
		r.Get("/", profileHandler.Get)
		r.Post("/", profileHandler.Create)
		r.Put("/", profileHandler.Update)
	})

	port := os.Getenv("PORT")
	log.Println("server starting on port", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("server error: ", err)
	}
}
