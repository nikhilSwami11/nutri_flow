package app

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/nikhilswami11/nutriflow/backend/internal/pantry"
	"github.com/nikhilswami11/nutriflow/backend/internal/photo"
	"github.com/nikhilswami11/nutriflow/backend/internal/profile"
	"github.com/nikhilswami11/nutriflow/backend/internal/recipes"
	"github.com/nikhilswami11/nutriflow/backend/internal/sessions"
	"github.com/nikhilswami11/nutriflow/backend/pkg/db"
)

type App struct {
	database *mongo.Database
	redis    *redis.Client

	pantryRepo   *pantry.Repository
	profileRepo  *profile.Repository
	recipesRepo  *recipes.Repository
	sessionsRepo *sessions.Repository

	recipesService  *recipes.Service
	sessionsService *sessions.Service

	pantryHandler   *pantry.Handler
	profileHandler  *profile.Handler
	recipesHandler  *recipes.Handler
	sessionsHandler *sessions.Handler

	photoStorage *photo.Storage
	photoRepo    *photo.Repository
	photoService *photo.Service
	photoHandler *photo.Handler
}

func Run() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	a := &App{}
	a.database = db.Connect()
	a.redis = db.ConnectRedis()

	a.pantryRepo = pantry.NewRepository(a.database)
	a.pantryHandler = pantry.NewHandler(a.pantryRepo)

	a.profileRepo = profile.NewRepository(a.database)
	a.profileHandler = profile.NewHandler(a.profileRepo)

	a.recipesRepo = recipes.NewRepository(a.database)
	a.recipesService = recipes.NewService(a.recipesRepo, a.pantryRepo, a.profileRepo)
	a.recipesHandler = recipes.NewHandler(a.recipesService)

	a.sessionsRepo = sessions.NewRepository(a.database)
	a.sessionsService = sessions.NewService(a.sessionsRepo, a.redis)
	a.sessionsHandler = sessions.NewHandler(a.sessionsService)

	photoStorage := photo.NewStorage()
	photoRepo := photo.NewRepository(a.database)
	photoService := photo.NewService(photoRepo, photoStorage)
	photoHandler := photo.NewHandler(photoService)
	a.photoStorage = photoStorage
	a.photoRepo = photoRepo
	a.photoService = photoService
	a.photoHandler = photoHandler

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	a.RegisterRoutes(r)

	port := os.Getenv("PORT")
	log.Println("server starting on port", port)
	return http.ListenAndServe(":"+port, r)
}
