package main

import (
	"fmt"
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
		panic(err)
	}
	dir := filepath.Dir(pathDir)
	v.AddConfigPath(dir)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	BOT_TOKEN = v.GetString("BOT_TOKEN")
	JOKESBAPAKBAPAKURI = v.GetString("JOKESBAPAKBAPAKURI")
	ENVIRONMENT = v.GetString("ENVIRONMENT")
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
		case command == "jokesbapak":
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
			str := fmt.Sprintf("Hi!, i'm running on **%s environment**g", ENVIRONMENT)
			s.ChannelMessageSend(m.ChannelID, str)
		// If the message is "ping" reply with "Poing!"
		case command == "ping":
			s.ChannelMessageSend(m.ChannelID, "Pong!")
			// If the message is "ping" reply with "Pong!"
		case command == "pong":
			s.ChannelMessageSend(m.ChannelID, "Ping!")
		default:
			s.ChannelMessageSend(m.ChannelID, "**404** Command Not Found!")

		}
	}
}
