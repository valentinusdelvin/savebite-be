package dto

type AIRequest struct {
	IngredientsOwned  []string `json:"ingredients_owned"`
	CookingPreference string   `json:"cooking_preference"`
	AdditionalNotes   string   `json:"additional_notes"`
}

type Recipe struct {
	DishName           string `json:"dish_name"`
	CookingTimeMinutes int    `json:"cooking_time_minutes"`
	Servings           int    `json:"servings"`
	DishType           string `json:"dish_type"`
	Ingredients        []struct {
		Name     string  `json:"name"`
		Quantity string  `json:"quantity"`
		Notes    *string `json:"notes"`
	} `json:"ingredients"`
	CookingSteps []string `json:"cooking_steps"`
	RecipeNotes  *string  `json:"recipe_notes"`
}
