package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var (
	BOT_TOKEN          string
	JOKESBAPAKBAPAKURI string
	ENVIRONMENT        string
)

func init() {
	v := viper.New()

	v.SetConfigFile(".env")
	pathDir, err := os.Executable()
	if err != nil {
		log.Println("pathDir", err)
	}
	dir := filepath.Dir(pathDir)
	v.AddConfigPath(dir)

	// Tell viper to automatically override values that it has read from config file
	// with the values of the corresponding environment variables if they exist.
	viper.AutomaticEnv()
	v.BindEnv("TMP")

	if err := v.ReadInConfig(); err != nil {
		log.Println("ReadInConfig", err)
		BOT_TOKEN = os.Getenv("BOT_TOKEN")
		JOKESBAPAKBAPAKURI = os.Getenv("JOKESBAPAKBAPAKURI")
		ENVIRONMENT = os.Getenv("ENVIRONMENT")
	} else {
		BOT_TOKEN = v.GetString("BOT_TOKEN")
		JOKESBAPAKBAPAKURI = v.GetString("JOKESBAPAKBAPAKURI")
		ENVIRONMENT = v.GetString("ENVIRONMENT")

	}

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

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.Split(m.Content, ".")

	if (content[0] == "!fb" && ENVIRONMENT == "PROD") || (content[0] == "!fbdev" && ENVIRONMENT == "DEV") {
		// fmt.Printf("Command from %s ENVIRONMENT\n", ENVIRONMENT)

		command := content[1]

		switch {
		case command == "jokes":
			//Call the JOKESBAPAKBAPAKURI API and retrieve our jokes
			response, err := http.Get(JOKESBAPAKBAPAKURI)
			if err != nil {
				fmt.Println(err)
			}
			defer response.Body.Close()

			if response.StatusCode == 200 {
				_, err = s.ChannelFileSend(m.ChannelID, "jokes-bapak2.png", response.Body)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Error: Can't get jokes-bapak2 ! :-(")
			}

		case command == "env":
			str := fmt.Sprintf("Hi!, i'm running on **%s environment**", ENVIRONMENT)
			s.ChannelMessageSend(m.ChannelID, str)

		case command == "sholat":
			s.ChannelMessageSend(m.ChannelID, "COMING SOON!!")

		case command == "search":
			s.ChannelMessageSend(m.ChannelID, "COMING SOON!!")

		case command == "translate":
			s.ChannelMessageSend(m.ChannelID, "COMING SOON!!")

		// If the message is "ping" reply with "Poing!"
		case command == "ping":
			s.ChannelMessageSend(m.ChannelID, "Pong!")

			// If the message is "ping" reply with "Pong!"
		case command == "pong":
			s.ChannelMessageSend(m.ChannelID, "Ping!")

		case command == "command":
			var str strings.Builder
			str.WriteString("Fajar Bot command list :\n")
			str.WriteString(fmt.Sprint("- `", content[0], ".jokes` = Get single random joke\n"))
			str.WriteString(fmt.Sprint("- `", content[0], ".env`= check environment\n"))
			str.WriteString(fmt.Sprint("- `", content[0], ".sholat` = COMING SOON!!\n"))
			str.WriteString(fmt.Sprint("- `", content[0], ".search` = COMING SOON!!\n"))
			str.WriteString(fmt.Sprint("- `", content[0], ".translate` = COMING SOON!!\n"))
			str.WriteString(fmt.Sprint("- `", content[0], ".ping` = test ping\n"))
			str.WriteString(fmt.Sprint("- `", content[0], ".pong` = test ping\n"))

			s.ChannelMessageSend(m.ChannelID, str.String())

		default:
			s.ChannelMessageSend(m.ChannelID, "**404** Command Not Found!")

		}
	}
}
