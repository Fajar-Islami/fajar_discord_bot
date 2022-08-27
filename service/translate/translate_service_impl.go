package translate

import (
	"strings"

	"github.com/Fajar-Islami/go_discord_bot/data"
	"github.com/bwmarrin/discordgo"
)

type TranslateServiceImpl struct {
	s *discordgo.Session
	m *discordgo.MessageCreate
}

func NewTranslateService(s *discordgo.Session, m *discordgo.MessageCreate) TranslateService {
	return &TranslateServiceImpl{
		s: s,
		m: m,
	}
}

func (ts *TranslateServiceImpl) LanguageList() string {
	var str strings.Builder

	for _, v := range data.TranslateListLang {
		str.WriteString(v)
	}

	return str.String()
}
func (ts *TranslateServiceImpl) LanguageCode(lang string) string {
	var str strings.Builder

	for _, v := range data.TranslateListLang {
		if strings.Contains(strings.ToLower(v), strings.ToLower(lang)) {
			str.WriteString(v)
		}
	}
	if len(str.String()) < 1 {
		return "Language not found"
	}

	return str.String()

}
