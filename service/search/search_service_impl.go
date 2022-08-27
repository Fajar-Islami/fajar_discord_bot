package search

import (
	"github.com/bwmarrin/discordgo"
)

type SearchServiceImpl struct {
	s *discordgo.Session
	m *discordgo.MessageCreate
}

func NewSearchService(s *discordgo.Session, m *discordgo.MessageCreate) SearchService {
	return &SearchServiceImpl{
		s: s,
		m: m,
	}
}

func (ss *SearchServiceImpl) SearchText(query string) []*discordgo.MessageEmbed {

	messsage := []*discordgo.MessageEmbed{
		{URL: "https://en.wikipedia.org/wiki/Elon_Musk",
			Title:       "Elon Musk - Wikipedia",
			Description: ""},
		{URL: "https://www.forbes.com/profile/elon-musk/",
			Title:       "Elon Musk - Forbes",
			Description: "Elon Musk cofounded six companies including electric car maker Tesla, rocket producer SpaceX and tunneling startup Boring Company."},
	}
	return messsage

}
