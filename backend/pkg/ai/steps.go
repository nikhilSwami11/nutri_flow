package ai

import (
	"context"
	"fmt"
	"strings"
)

type GeneratedStep struct {
	Order       int    `json:"order"`
	Instruction string `json:"instruction"`
}

type stepsResponse struct {
	Steps []GeneratedStep `json:"steps"`
}

type RecipeInput struct {
	Name        string
	Description string
	Ingredients []GeneratedIngredient
	PrepTime    int
	CookTime    int
}

func (c *Client) GenerateSteps(ctx context.Context, recipe RecipeInput) ([]GeneratedStep, error) {
	system := `You are a cooking instructor. Given a recipe, generate clear step-by-step cooking instructions.

Rules:
- Each step must be a single, actionable instruction.
- Steps should be ordered sequentially starting from 1.
- Be specific with quantities, temperatures, and timing.
- Return valid JSON: {"steps": [{"order": 1, "instruction": "..."}, ...]}`

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Recipe: %s\n", recipe.Name))
	sb.WriteString(fmt.Sprintf("Description: %s\n", recipe.Description))
	sb.WriteString(fmt.Sprintf("Prep time: %d minutes, Cook time: %d minutes\n", recipe.PrepTime, recipe.CookTime))
	sb.WriteString("Ingredients:\n")
	for _, ing := range recipe.Ingredients {
		optional := ""
		if ing.Optional {
			optional = " (optional)"
		}
		sb.WriteString(fmt.Sprintf("- %.2f %s %s%s\n", ing.Quantity, ing.Unit, ing.Name, optional))
	}

	var out stepsResponse
	if err := c.chat(ctx, system, sb.String(), &out); err != nil {
		return nil, err
	}

	return out.Steps, nil
}
