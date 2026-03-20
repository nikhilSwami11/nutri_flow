package profile

import "time"

type Profile struct {
	UserID             string    `json:"user_id" bson:"_id"`
	Name               string    `json:"name" bson:"name"`
	Age                int       `json:"age" bson:"age"`
	Height             float64   `json:"height" bson:"height"`
	Weight             float64   `json:"weight" bson:"weight"`
	DailyCalorieGoal   int       `json:"daily_calorie_goal" bson:"daily_calorie_goal"`
	DietaryPreferences []string  `json:"dietary_preferences" bson:"dietary_preferences"`
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" bson:"updated_at"`
}
