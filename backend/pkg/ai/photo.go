package ai

import (
	"context"
	"fmt"
)

type CalorieEstimate struct {
	DishName          string  `json:"dish_name"`
	EstimatedCalories float64 `json:"estimated_calories"`
	Confidence        float64 `json:"confidence"`
}

func (c *Client) EstimateCalories(ctx context.Context, photoURL string, mealType string) (*CalorieEstimate, error) {
	system := `You are a nutrition expert. Given a photo URL of a meal and the meal type, estimate the calorie content.

Rules:
- Identify the dish name from the photo URL context and meal type.
- Provide a realistic calorie estimate for a typical serving.
- Confidence should be between 0.0 and 1.0 (1.0 = very confident).
- Return valid JSON: {"dish_name": "...", "estimated_calories": 0.0, "confidence": 0.0}`

	user := fmt.Sprintf("Photo URL: %s\nMeal type: %s", photoURL, mealType)

	var out CalorieEstimate
	if err := c.chat(ctx, system, user, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
