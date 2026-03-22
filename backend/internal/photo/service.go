package photo

import "time"

type Service struct {
	repo    *Repository
	storage *Storage
}

func NewService(repo *Repository, storage *Storage) *Service {
	return &Service{repo: repo, storage: storage}
}

func (s *Service) EstimateCalories(userID string, fileBytes []byte, filename string, mealType string) (*PhotoLog, error) {
	photoURL, err := s.storage.UploadPhoto(userID, fileBytes, filename)
	if err != nil {
		return nil, err
	}

	// TODO: call Pipeline 3 (calorie estimation) passing:
	// photoURL, userID, mealType, user kitchen context
	// pipeline returns DishName, EstimatedCalories, Confidence

	log := &PhotoLog{
		UserID:            userID,
		PhotoURL:          photoURL,
		DishName:          "Unknown",
		EstimatedCalories: 0,
		Confidence:        0,
		MealType:          mealType,
		LoggedAt:          time.Now(),
	}

	if err := s.repo.Save(log); err != nil {
		return nil, err
	}

	return log, nil
}

func (s *Service) GetHistory(userID string) ([]PhotoLog, error) {
	return s.repo.GetByUserID(userID)
}
