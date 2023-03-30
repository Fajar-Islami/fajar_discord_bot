package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Fajar-Islami/fajar_discord_bot/helper"
	"github.com/Fajar-Islami/fajar_discord_bot/service"
	"github.com/Fajar-Islami/fajar_discord_bot/service/ai"
	"github.com/Fajar-Islami/fajar_discord_bot/service/jokes"
	"github.com/Fajar-Islami/fajar_discord_bot/service/search"
	"github.com/Fajar-Islami/fajar_discord_bot/service/translate"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	BOT_TOKEN               string
	JOKESBAPAKBAPAKURI      string
	ENVIRONMENT             string
	PESTO_TOKEN             string
	PESTO_URI               string
	TRANSLATE_RapidAPI_KEY  string
	TRANSLATE_RapidAPI_HOST string
	TRANSLATE_RapidAPI_URI  string
	WRITESONIC_APIKEY       string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("err read env file : ", err)
	}

	log.Println("ReadInConfig", err)
	BOT_TOKEN = os.Getenv("BOT_TOKEN")
	JOKESBAPAKBAPAKURI = os.Getenv("JOKESBAPAKBAPAKURI")
	ENVIRONMENT = os.Getenv("ENVIRONMENT")
	PESTO_TOKEN = os.Getenv("PESTO_TOKEN")
	PESTO_URI = os.Getenv("PESTO_URI")
	TRANSLATE_RapidAPI_KEY = os.Getenv("TRANSLATE_RapidAPI_KEY")
	TRANSLATE_RapidAPI_HOST = os.Getenv("TRANSLATE_RapidAPI_HOST")
	TRANSLATE_RapidAPI_URI = os.Getenv("TRANSLATE_RapidAPI_URI")
	WRITESONIC_APIKEY = os.Getenv("WRITESONIC_APIKEY")

}

func main() {

	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + BOT_TOKEN)
	if err != nil {
		fmt.Println("Error creating Discord session", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate Events
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	//  Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection", err)
		return
	}

	// Wait here untul CTRC-C or other term sign is received.
	fmt.Printf("Bot is now running on %s environtment. Press CTRL+C to exit\n", ENVIRONMENT)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly  close down the Discord session
	dg.Close()

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	defer func() {
		if r := recover(); r != nil {
			s.ChannelMessageSend(m.ChannelID, "Oh no, something went wrong with me! Can you guys help me to ping my masters? @FajarIslami#8186")
		}
	}()

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	translateService := translate.NewTranslateService(s, m, TRANSLATE_RapidAPI_KEY, TRANSLATE_RapidAPI_HOST, TRANSLATE_RapidAPI_URI)
	jokesService := jokes.NewJokesService(s, m, JOKESBAPAKBAPAKURI)
	aiService := ai.NewAIBot(s, m, WRITESONIC_APIKEY)
	searchService := search.NewSearchService(s, m)

	botname := strings.Split(m.Content, " ")

	if (botname[0] == "!fb" && ENVIRONMENT == "PROD") || (botname[0] == "!fbdev" && ENVIRONMENT == "DEV") {
		// fmt.Printf("Command from %s ENVIRONMENT\n", ENVIRONMENT)

		command := botname[1]

		switch {

		// example command = !fb jokes
		case command == "jokes":
			//Call the JOKESBAPAKBAPAKURI API and retrieve our jokes
			resBody, resp, errGet := jokesService.GetRandomJokes()

			if errGet != "" {
				s.ChannelMessageSend(m.ChannelID, errGet)
				return
			}

			defer resp.Body.Close()
			err := helper.ResponseImage(s, m, resBody)
			if err != nil {
				log.Println(err)
				s.ChannelMessageSend(m.ChannelID, err.Error())
			}

		// example command = !fb joktod
		case command == "joktod":
			//Call the JOKESBAPAKBAPAKURI API and retrieve our jokes
			resBody, resp, errGet := jokesService.GetRandomJokesToday()

			if errGet != "" {
				s.ChannelMessageSend(m.ChannelID, errGet)
				return
			}

			defer resp.Body.Close()

			err := helper.ResponseImage(s, m, resBody)
			if err != nil {
				log.Println(err)
				s.ChannelMessageSend(m.ChannelID, err.Error())
			}

		// example command = !fb search=<text you want to search>
		case strings.Contains(command, "search="):
			content := strings.Join(botname[1:], " ")
			fmt.Println("content", content)
			searchArr := strings.Split(content, "=")[1:]
			searchStr := strings.Join(searchArr, "")

			if searchStr == "" {
				s.ChannelMessageSend(m.ChannelID, "Invliad search text")
				return
			}

			resp := aiService.SearchBot(searchStr)
			s.ChannelMessageSend(m.ChannelID, resp)

		// example command = !fb rcelist
		case command == "rcelist":
			s.ChannelMessageSend(m.ChannelID, "[Link text](http://example.com)")

		// example command = !fb env
		case command == "env":
			str := fmt.Sprintf("Hi!, i'm running on **%s environment**", ENVIRONMENT)
			s.ChannelMessageSend(m.ChannelID, str)

		// example command = !fb sholat
		case command == "sholat":
			s.ChannelMessageSend(m.ChannelID, "COMING SOON!!")

		// example command = !fb search
		case command == "search":
			res := searchService.SearchText("anything")
			s.ChannelMessageSendEmbeds(m.ChannelID, res)

		// example command = !fb translate-langlist
		case command == "translate-langlist":
			str := translateService.LanguageList()
			s.ChannelMessageSend(m.ChannelID, str)

		// example command = !fb translate-codelang eng
		case command == "translate-codelang":
			if len(botname) < 3 {
				s.ChannelMessageSend(m.ChannelID, "400 Bad request for language")
				return
			}

			var lang = botname[2]

			str := translateService.LanguageCode(lang)

			s.ChannelMessageSend(m.ChannelID, str)

		// example command = !fb translate-detectlang English is hard but
		case command == "translate-detectlang":
			if len(botname) < 3 {
				s.ChannelMessageSend(m.ChannelID, "400 Bad request for language")
				return
			}

			var sentences = botname[2]

			str := translateService.DetectLanguage(sentences)

			s.ChannelMessageSend(m.ChannelID, str)

		// example command = !fb translate-codelang eng
		case command == "translate":
			s.ChannelMessageSend(m.ChannelID, "COMING SOON!!")

		// If the message is "ping" reply with "Poing!"
		// example command = !fb ping
		case command == "ping":
			s.ChannelMessageSend(m.ChannelID, "Pong!")

		// If the message is "ping" reply with "Pong!"
		// example command = !fb pong
		case command == "pong":
			s.ChannelMessageSend(m.ChannelID, "Ping!")

		// example command = !fb intro
		case command == "intro":
			s.ChannelMessageSend(m.ChannelID, "Fajar BOT v1.0.0\n Update Breaking Changes, check `!fb command` (no longer use dots for commands) ")

		// example command = !fb contribute
		case command == "contribute":
			s.ChannelMessageSend(m.ChannelID, "Feel free to contribute here https://github.com/Fajar-Islami/fajar_discord_bot")

		// example command = !fb commands
		case command == "commands":
			res := service.ListCommand(botname[0])
			s.ChannelMessageSend(m.ChannelID, res)

		default:
			s.ChannelMessageSend(m.ChannelID, "**404** Command Not Found!")

		}
	}
}
