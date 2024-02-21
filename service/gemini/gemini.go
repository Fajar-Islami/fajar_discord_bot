package gemini

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AIGemini interface {
	GenerateText(search string) string
}

type AIGeminiImpl struct {
	client *genai.Client
}

func NewAIGemini(apikey string) AIGemini {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(apikey))
	if err != nil {
		log.Fatal(err)
		// return
	}
	// defer client.Close()

	return &AIGeminiImpl{client: client}
}

func (a *AIGeminiImpl) GenerateText(search string) (result string) {
	ctx := context.Background()
	fmt.Println("search", search)

	// For text-only input, use the gemini-pro model
	model := a.client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text("Cerita tentang marcos."))
	if err != nil {
		log.Fatal(err)
		result = "FAILED TO GET DATA"
		fmt.Println("result", result)
		return
	}

	for _, v := range resp.Candidates {
		for _, parts := range v.Content.Parts {
			result += fmt.Sprintf("%v", parts)
		}
	}

	fmt.Println("result", result)
	return result
}
