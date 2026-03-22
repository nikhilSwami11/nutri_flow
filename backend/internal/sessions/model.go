package sessions

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Step struct {
	Order       int    `json:"order"`
	Instruction string `json:"instruction"`
}

type CookingTask struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Instruction string `json:"instruction"`
	Completed   bool   `json:"completed"`
	Todo        bool   `json:"todo"`
}

type Session struct {
	ID            string        `json:"id"`
	UserID        string        `json:"user_id"`
	RecipeID      string        `json:"recipe_id"`
	RecipeName    string        `json:"recipe_name"`
	Status        string        `json:"status"`
	CurrentTaskID int           `json:"current_task_id"`
	StartedAt     time.Time     `json:"started_at"`
	CookingTasks  []CookingTask `json:"cooking_tasks"`
	Steps         []Step        `json:"steps"`
}

type CompletedSession struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID            string             `json:"user_id" bson:"user_id"`
	RecipeID          string             `json:"recipe_id" bson:"recipe_id"`
	RecipeName        string             `json:"recipe_name" bson:"recipe_name"`
	ServingsEaten     float64            `json:"servings_eaten" bson:"servings_eaten"`
	CaloriesConsumed  float64            `json:"calories_consumed" bson:"calories_consumed"`
	StartedAt         time.Time          `json:"started_at" bson:"started_at"`
	CompletedAt       time.Time          `json:"completed_at" bson:"completed_at"`
}
