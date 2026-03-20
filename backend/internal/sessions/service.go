package sessions

// TODO: inject recipes repository to fetch cooking tasks during session

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	sessionsRepo *Repository
	redisClient  *redis.Client
}

func NewService(sessionsRepo *Repository, redisClient *redis.Client) *Service {
	return &Service{
		sessionsRepo: sessionsRepo,
		redisClient:  redisClient,
	}
}

func (s *Service) StartSession(userID string, recipeID string, recipeName string) (*Session, error) {
	session := &Session{
		ID:            fmt.Sprintf("%s-%d", userID, time.Now().UnixNano()),
		UserID:        userID,
		RecipeID:      recipeID,
		RecipeName:    recipeName,
		Status:        "created",
		CurrentTaskID: 0,
		StartedAt:     time.Now(),
	}

	data, err := json.Marshal(session)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	key := fmt.Sprintf("session:%s", session.ID)
	if err := s.redisClient.Set(ctx, key, data, 2*time.Hour).Err(); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *Service) CompleteSession(sessionID string, servingsEaten float64, caloriesPerServing float64) error {
	// TODO: trigger pantry depletion based on recipe ingredients
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	key := fmt.Sprintf("session:%s", sessionID)
	data, err := s.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	var session Session
	if err := json.Unmarshal(data, &session); err != nil {
		return err
	}

	completed := &CompletedSession{
		UserID:           session.UserID,
		RecipeID:         session.RecipeID,
		RecipeName:       session.RecipeName,
		ServingsEaten:    servingsEaten,
		CaloriesConsumed: servingsEaten * caloriesPerServing,
		StartedAt:        session.StartedAt,
		CompletedAt:      time.Now(),
	}

	if err := s.sessionsRepo.Save(completed); err != nil {
		return err
	}

	return s.redisClient.Del(ctx, key).Err()
}

func (s *Service) AbandonSession(sessionID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	key := fmt.Sprintf("session:%s", sessionID)
	return s.redisClient.Del(ctx, key).Err()
}

func (s *Service) GetHistory(userID string) ([]CompletedSession, error) {
	return s.sessionsRepo.GetByUserID(userID)
}
