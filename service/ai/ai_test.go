package ai

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"reflect"
// 	"testing"

// 	"github.com/Fajar-Islami/fajar_discord_bot/helper"
// 	"github.com/joho/godotenv"
// )

// func TestMain(m *testing.M) {
// 	fmt.Println("helper.ProjectRootPath ", helper.ProjectRootPath+"/.env")
// 	err := godotenv.Load(helper.ProjectRootPath + "/.env")
// 	if err != nil {
// 		log.Println("err read env file : ", err)
// 	}

// 	log.Println("env loaded")

// 	os.Exit(m.Run())
// }

// func TestNewAIBot(t *testing.T) {
// 	type args struct {
// 		APIKEY string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want AIBot
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewAIBot(tt.args.APIKEY); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewAIBot() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestAIBotImpl_SearchBot(t *testing.T) {
// 	type searchBot struct {
// 		name       string
// 		ai         *AIBotImpl
// 		text       string
// 		errMessage string
// 		isError    bool
// 	}
// 	tests := []searchBot{
// 		// TODO: Add test cases.
// 		{
// 			name: "Find AI",
// 			ai: &AIBotImpl{
// 				APIKEY: os.Getenv("WRITESONIC_APIKEY"),
// 			},
// 			text:       "Cari text",
// 			errMessage: "Error",
// 			isError:    false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := tt.ai.SearchBot(tt.text)
// 			t.Log("result : ", got)
// 		})
// 	}
// }
