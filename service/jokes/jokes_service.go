package jokes

import (
	"net/http"

	"github.com/Fajar-Islami/fajar_discord_bot/helper"
)

type JokesService interface {
	GetRandomJokes(uri string) (*helper.ResponseImageStruct, *http.Response, string)
}
