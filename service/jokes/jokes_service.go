package jokes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Fajar-Islami/fajar_discord_bot/helper"
	"github.com/Fajar-Islami/fajar_discord_bot/service"
	"github.com/bwmarrin/discordgo"
)

type JokesService interface {
	GetRandomJokes() (*helper.ResponseImageStruct, *http.Response, string)
	GetRandomJokesToday() (*helper.ResponseImageStruct, *http.Response, string)
}

type JokesServiceImpl struct {
	s       *discordgo.Session
	m       *discordgo.MessageCreate
	BaseURI string
}

func NewJokesService(s *discordgo.Session, m *discordgo.MessageCreate, BaseURI string) JokesService {
	return &JokesServiceImpl{
		s:       s,
		m:       m,
		BaseURI: BaseURI,
	}
}

func (js *JokesServiceImpl) GetRandomJokes() (*helper.ResponseImageStruct, *http.Response, string) {

	get := service.Get(js.BaseURI + "/")

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

func (js *JokesServiceImpl) GetRandomJokesToday() (*helper.ResponseImageStruct, *http.Response, string) {

	get := service.Get(js.BaseURI + "/today")
	resp, err := get.Do()

	if err != nil {
		log.Println(err)
		return nil, nil, fmt.Sprintf("Error Get Jokes : %s", err.Error())
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
