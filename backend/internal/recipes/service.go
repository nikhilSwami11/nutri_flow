package recipes

import (
	"github.com/nikhilswami11/nutriflow/backend/internal/pantry"
	"github.com/nikhilswami11/nutriflow/backend/internal/profile"
)

type Service struct {
	recipesRepo *Repository
	pantryRepo  *pantry.Repository
	profileRepo *profile.Repository
}

func NewService(recipesRepo *Repository, pantryRepo *pantry.Repository, profileRepo *profile.Repository) *Service {
	return &Service{
		recipesRepo: recipesRepo,
		pantryRepo:  pantryRepo,
		profileRepo: profileRepo,
	}
}

func (s *Service) GetSuggestions(userID string) ([]Recipe, error) {
	return []Recipe{}, nil
}

func (s *Service) SaveRecipe(recipe *Recipe) error {
	recipe.IsTemporary = false
	return s.recipesRepo.Create(recipe)
}

func (s *Service) GetUserRecipes(userID string) ([]Recipe, error) {
	return s.recipesRepo.GetByUserID(userID)
}

func (s *Service) DeleteRecipe(id string, userID string) error {
	return s.recipesRepo.Delete(id, userID)
}
