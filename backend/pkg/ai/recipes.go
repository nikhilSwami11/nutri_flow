package ai

import (
	"context"
	"fmt"
	"strings"
)

// Input types — keeps pkg/ai decoupled from internal domain packages

type PantryItemInput struct {
	Name     string
	Quantity float64
	Unit     string
}

type ProfileInput struct {
	DietaryPreferences []string
	Allergies          []string
	CuisinePreferences []string
	KitchenType        string
}

// Output types

type GeneratedIngredient struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
	Optional bool    `json:"optional"`
}

type GeneratedRecipe struct {
	IsSaved     bool                  `json:"is_saved"`
	SavedName   string                `json:"saved_name,omitempty"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Ingredients []GeneratedIngredient `json:"ingredients"`
	PrepTime    int                   `json:"prep_time"`
	CookTime    int                   `json:"cook_time"`
	Servings    int                   `json:"servings"`
	Calories    float64               `json:"calories"`
}

type recipeSuggestionsResponse struct {
	Recipes []GeneratedRecipe `json:"recipes"`
}

func (c *Client) GenerateRecipeSuggestions(
	ctx context.Context,
	pantry []PantryItemInput,
	profile ProfileInput,
	savedRecipeNames []string,
	query string,
) ([]GeneratedRecipe, error) {
	system := `You are a personal chef assistant. Given a user's pantry, dietary profile, and optionally a specific dish request, suggest exactly 3 recipes.

Rules:
- If a saved recipe fits, set is_saved=true and saved_name to the exact saved recipe name.
- Mark ingredients the user is missing or that are optional as optional=true.
- Respect allergies strictly — never include allergens.
- Return valid JSON: {"recipes": [...]}`

	var sb strings.Builder
	sb.WriteString("Pantry items:\n")
	for _, p := range pantry {
		sb.WriteString(fmt.Sprintf("- %s: %.2f %s\n", p.Name, p.Quantity, p.Unit))
	}

	sb.WriteString(fmt.Sprintf("\nDietary preferences: %s", strings.Join(profile.DietaryPreferences, ", ")))
	sb.WriteString(fmt.Sprintf("\nAllergies: %s", strings.Join(profile.Allergies, ", ")))
	sb.WriteString(fmt.Sprintf("\nCuisine preferences: %s", strings.Join(profile.CuisinePreferences, ", ")))
	sb.WriteString(fmt.Sprintf("\nKitchen type: %s", profile.KitchenType))

	if len(savedRecipeNames) > 0 {
		sb.WriteString(fmt.Sprintf("\nUser's saved recipes: %s", strings.Join(savedRecipeNames, ", ")))
	}

	if query != "" {
		sb.WriteString(fmt.Sprintf("\nUser specifically requested: %s", query))
	}

	var out recipeSuggestionsResponse
	if err := c.chat(ctx, system, sb.String(), &out); err != nil {
		return nil, err
	}

	return out.Recipes, nil
}
