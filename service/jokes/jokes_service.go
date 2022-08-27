package jokes

import (
	"net/http"

	"github.com/Fajar-Islami/go_discord_bot/helper"
)

type JokesService interface {
	GetRandomJokes(uri string) (*helper.ResponseImageStruct, *http.Response)
}
