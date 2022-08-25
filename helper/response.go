package helper

import (
	"errors"
	"io"

	"github.com/bwmarrin/discordgo"
)

type ResponseStruct struct {
	StatusCode, ExpectedStatusCode int
	Name, Filename                 string
	RespBody                       io.Reader
}

func SuccessStatus(s *discordgo.Session, m *discordgo.MessageCreate, resp *ResponseStruct) error {
	if resp.StatusCode == resp.ExpectedStatusCode {
		_, err := s.ChannelFileSend(m.ChannelID, resp.Filename, resp.RespBody)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("Error: Can't get " + resp.Name + " :-(")

}
