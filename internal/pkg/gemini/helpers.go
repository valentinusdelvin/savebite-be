package gemini

import (
	"fmt"
	"os"
	"strings"

	"github.com/valentinusdelvin/savebite-be/internal/domain/dto"
)

func FormatSaveBitePrompt(req dto.AIRequest) (string, error) {

	// 1. Load prompt file
	rawPrompt, err := os.ReadFile("internal/pkg/gemini/prompt.txt")
	if err != nil {
		return "", fmt.Errorf("failed to load prompt template: %w", err)
	}

	prompt := string(rawPrompt)

	// 2. Prepare ingredients
	ingredients := ""
	if len(req.IngredientsOwned) > 0 {
		quoted := make([]string, len(req.IngredientsOwned))
		for i, v := range req.IngredientsOwned {
			quoted[i] = fmt.Sprintf(`"%s"`, v)
		}
		ingredients = strings.Join(quoted, ", ")
	}

	// 3. Safe defaults
	cookingPref := req.CookingPreference
	if cookingPref == "" {
		cookingPref = "Not specified"
	}

	notes := req.AdditionalNotes
	if notes == "" {
		notes = "None"
	}

	// 4. Inject variables
	prompt = strings.ReplaceAll(prompt, "{{ingredients}}", ingredients)
	prompt = strings.ReplaceAll(prompt, "{{cookingPreference}}", cookingPref)
	prompt = strings.ReplaceAll(prompt, "{{additionalNotes}}", notes)

	fmt.Println(prompt)

	return prompt, nil
}
