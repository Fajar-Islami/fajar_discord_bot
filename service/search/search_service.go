package search

import (
	"github.com/bwmarrin/discordgo"
)

type SearchService interface {
	SearchText(query string) []*discordgo.MessageEmbed
}
