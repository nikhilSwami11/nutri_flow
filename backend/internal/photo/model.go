package photo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PhotoLog struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID            string             `json:"user_id" bson:"user_id"`
	PhotoURL          string             `json:"photo_url" bson:"photo_url"`
	DishName          string             `json:"dish_name" bson:"dish_name"`
	EstimatedCalories float64            `json:"estimated_calories" bson:"estimated_calories"`
	Confidence        float64            `json:"confidence" bson:"confidence"`
	MealType          string             `json:"meal_type" bson:"meal_type"`
	LoggedAt          time.Time          `json:"logged_at" bson:"logged_at"`
	CreatedAt         time.Time          `json:"created_at" bson:"created_at"`
}

// TODO: integrate Spoonacular API for food images
// https://api.spoonacular.com/food/images/classify
// Add image_url to pantry items when voice adds them
// Add image_url to recipes when LLM generates them
