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
		// TODO: call Pipeline 2 (steps generation) to populate Steps
		Steps: []Step{},
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

func (s *Service) AbandonSession(sessionID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	key := fmt.Sprintf("session:%s", sessionID)
	return s.redisClient.Del(ctx, key).Err()
}
