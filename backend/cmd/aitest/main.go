package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/nikhilswami11/nutriflow/backend/pkg/ai"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("could not load .env:", err)
	}

	client := ai.NewClient()
	ctx := context.Background()

	// --- Test: GenerateRecipeSuggestions ---
	pantry := []ai.PantryItemInput{
		{Name: "chicken breast", Quantity: 2, Unit: "pieces"},
		{Name: "garlic", Quantity: 4, Unit: "cloves"},
		{Name: "olive oil", Quantity: 100, Unit: "ml"},
		{Name: "pasta", Quantity: 200, Unit: "g"},
	}

	profile := ai.ProfileInput{
		DietaryPreferences: []string{"high-protein"},
		Allergies:          []string{"nuts"},
		CuisinePreferences: []string{"Italian", "Mediterranean"},
		KitchenType:        "home",
	}

	fmt.Println("=== GenerateRecipeSuggestions ===")
	recipes, err := client.GenerateRecipeSuggestions(ctx, pantry, profile, nil, "")
	if err != nil {
		log.Fatal(err)
	}
	printJSON(recipes)
}

func printJSON(v any) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}
