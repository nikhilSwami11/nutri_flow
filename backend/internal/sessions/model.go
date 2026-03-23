package sessions

import "time"

type Step struct {
	Order       int    `json:"order"`
	Instruction string `json:"instruction"`
	Completed   bool   `json:"completed"`
}

type Session struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	RecipeID      string    `json:"recipe_id"`
	RecipeName    string    `json:"recipe_name"`
	Status        string    `json:"status"`
	CurrentTaskID int       `json:"current_task_id"`
	StartedAt     time.Time `json:"started_at"`
	Steps         []Step    `json:"steps"`
}
