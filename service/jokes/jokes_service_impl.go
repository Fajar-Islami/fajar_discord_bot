package jokes

import (
	"log"
	"net/http"

	"github.com/Fajar-Islami/fajar_discord_bot/helper"
	"github.com/Fajar-Islami/fajar_discord_bot/service"
	"github.com/bwmarrin/discordgo"
)

type JokesServiceImpl struct {
	s *discordgo.Session
	m *discordgo.MessageCreate
}

func NewJokesService(s *discordgo.Session, m *discordgo.MessageCreate) JokesService {
	return &JokesServiceImpl{
		s: s,
		m: m,
	}
}

func (js *JokesServiceImpl) GetRandomJokes(uri string) (*helper.ResponseImageStruct, *http.Response, string) {

	get := service.Get(uri)
	resp, err := get.Do()

	if err != nil {
		log.Println(err)

		return nil, nil, "Error Get Jokes"
	}

	respStruct := &helper.ResponseImageStruct{
		StatusCode:         resp.StatusCode,
		ExpectedStatusCode: http.StatusOK,
		Name:               "jokes-bapak2",
		Filename:           "jokes-bapak2.png",
		RespBody:           resp.Body,
	}
	return respStruct, resp, ""

}
