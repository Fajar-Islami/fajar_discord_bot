package translate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Fajar-Islami/fajar_discord_bot/data"
	"github.com/Fajar-Islami/fajar_discord_bot/service"
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

func (ts *TranslateServiceImpl) DetectLanguage(senteces string) string {
	// payload := strings.NewReader(fmt.Sprintf("q=%s",sentece))
	// q := &DetectLang{
	// 	Q: senteces,
	// }

	q := strings.NewReader(fmt.Sprintf("q=%s", senteces))
	// buffer := make([]byte, 10)

	// _, err := q.Read(buffer)
	// if err != nil {
	// 	log.Println(err)
	// 	return "Failed to Detect Language"
	// }
	// fmt.Println("buffer", string(buffer))

	// payload, err := json.Marshal(buffer)
	// if err != nil {
	// 	log.Println(err)
	// 	return "**400** Bad Request"
	// }

	uri := ts.TRANSLATE_RapidAPI_URI + "/detect"
	// post := service.Post(uri, bytes.NewBuffer(payload))
	post := service.Post(uri, q)
	post.Header("content-type", "application/x-www-form-urlencoded")
	post.Header("Accept-Encoding", "application/gzip")
	post.Header("X-RapidAPI-Key", ts.TRANSLATE_RapidAPI_KEY)
	post.Header("X-RapidAPI-Host", ts.TRANSLATE_RapidAPI_HOST)

	resp, err := post.Do()
	if err != nil {
		log.Println(err)
		return "Failed to Detect Language"
	}

	var r *RespDetectLang
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "Failed to Detect Language"
	}

	fmt.Println("body", string(body))

	if err := json.Unmarshal(body, &r); err != nil || r != nil {
		log.Println(err)
		return "Failed to Detect Language"
	}

	fmt.Println("r", r)

	return "Language Detected"

}
