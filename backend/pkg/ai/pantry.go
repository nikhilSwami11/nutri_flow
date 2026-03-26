package ai

import "context"

type ParsedPantryItem struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type parsedPantryResponse struct {
	Items []ParsedPantryItem `json:"items"`
}

func (c *Client) ParsePantryInput(ctx context.Context, text string) ([]ParsedPantryItem, error) {
	system := `You are a pantry assistant. Extract food items from natural language input.

Rules:
- Extract every food item mentioned with its quantity and unit.
- If no quantity is mentioned, default quantity to 1 and unit to "piece".
- Normalise unit names (e.g. "kgs" → "kg", "tbsp" → "tbsp").
- Return valid JSON: {"items": [{"name": "...", "quantity": 0.0, "unit": "..."}, ...]}`

	var out parsedPantryResponse
	if err := c.chat(ctx, system, text, &out); err != nil {
		return nil, err
	}

	return out.Items, nil
}
