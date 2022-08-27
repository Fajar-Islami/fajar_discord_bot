package translate

import (
	"strings"

	"github.com/Fajar-Islami/fajar_discord_bot/data"
	"github.com/bwmarrin/discordgo"
)

type TranslateServiceImpl struct {
	s                       *discordgo.Session
	m                       *discordgo.MessageCreate
	TRANSLATE_RapidAPI_KEY  string
	TRANSLATE_RapidAPI_HOST string
	TRANSLATE_RapidAPI_URI  string
}

func NewTranslateService(s *discordgo.Session, m *discordgo.MessageCreate, TRANSLATE_RapidAPI_KEY, TRANSLATE_RapidAPI_HOST, TRANSLATE_RapidAPI_URI string) TranslateService {
	return &TranslateServiceImpl{
		s:                       s,
		m:                       m,
		TRANSLATE_RapidAPI_KEY:  TRANSLATE_RapidAPI_KEY,
		TRANSLATE_RapidAPI_HOST: TRANSLATE_RapidAPI_HOST,
		TRANSLATE_RapidAPI_URI:  TRANSLATE_RapidAPI_URI,
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

func (ts *TranslateServiceImpl) DetectLanguage(sentece string) string {
	// payload := strings.NewReader(fmt.Sprintf("q=%s",sentece))

	// service.GetPost(sentece,)

	return "Language not found"
	// return str.String()

}
