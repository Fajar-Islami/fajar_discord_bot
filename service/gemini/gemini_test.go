package gemini

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var geminiSvc AIGemini

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("err read env file : ", err)
	}

	apiKey := os.Getenv("GOOGLE_GEMINI")
	geminiSvc = NewAIGemini(apiKey)
	os.Exit(m.Run())
}

func TestGenerateText(t *testing.T) {
	// fmt.Println("halo")
	geminiSvc.GenerateText("Ceritkan tentang kamu")
}
