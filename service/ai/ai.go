package ai

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/Fajar-Islami/fajar_discord_bot/model"
	"github.com/Fajar-Islami/fajar_discord_bot/service"
	"github.com/bwmarrin/discordgo"
)

type AIBot interface {
	SearchBot(text string) string
}

type AIBotImpl struct {
	s      *discordgo.Session
	m      *discordgo.MessageCreate
	APIKEY string
}

func NewAIBot(s *discordgo.Session, m *discordgo.MessageCreate, APIKEY string) AIBot {
	return &AIBotImpl{
		s:      s,
		m:      m,
		APIKEY: APIKEY,
	}
}

func (ai *AIBotImpl) SearchBot(text string) string {

	reqBody := model.AIReq{
		EnableGoogleResults: true,
		EnableMemory:        true,
		InputText:           text,
	}
	payload, err := json.Marshal(reqBody)
	if err != nil {
		log.Println(err)
		return "Error Marshal SearchBot AI"
	}

	post := service.Post("https://api.writesonic.com/v2/business/content/chatsonic?engine=premium", bytes.NewBuffer(payload))
	post.Header("content-type", "application/json")
	post.Header("X-RapidAPI-Key", ai.APIKEY)
	resp, err := post.Do()

	if err != nil {
		log.Println(err)
		return "Error Get Search Bot"
	}

	var r *model.AIRes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "Error Read Search Bot Body"
	}

	if err := json.Unmarshal(body, &r); err != nil || r != nil {
		log.Println(err)
		return "Error UnMarshal SearchBot AI Res Body"
	}

	return r.Message

}
