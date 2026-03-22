package recipes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ingredient struct {
	Name     string  `json:"name" bson:"name"`
	Quantity float64 `json:"quantity" bson:"quantity"`
	Unit     string  `json:"unit" bson:"unit"`
}

type Recipe struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID       string             `json:"user_id" bson:"user_id"`
	Name         string             `json:"name" bson:"name"`
	Description  string             `json:"description" bson:"description"`
	Ingredients  []Ingredient       `json:"ingredients" bson:"ingredients"`
	PrepTime     int                `json:"prep_time" bson:"prep_time"`
	CookTime     int                `json:"cook_time" bson:"cook_time"`
	Servings     int                `json:"servings" bson:"servings"`
	Calories     float64            `json:"calories" bson:"calories"`
	IsTemporary  bool               `json:"is_temporary" bson:"is_temporary"`
	ExpiresAt    *time.Time         `json:"expires_at,omitempty" bson:"expires_at,omitempty"`
	TimesCooked  int                `json:"times_cooked" bson:"times_cooked"`
	LastCookedAt *time.Time         `json:"last_cooked_at,omitempty" bson:"last_cooked_at,omitempty"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}
