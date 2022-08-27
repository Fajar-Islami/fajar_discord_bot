package helper

import (
	"errors"
	"io"

	"github.com/bwmarrin/discordgo"
)

type (
	ResponseImageStruct struct {
		StatusCode, ExpectedStatusCode int
		Name, Filename                 string
		RespBody                       io.Reader
	}
)

func ResponseImage(s *discordgo.Session, m *discordgo.MessageCreate, resp *ResponseImageStruct) error {
	if resp.StatusCode == resp.ExpectedStatusCode {
		_, err := s.ChannelFileSend(m.ChannelID, resp.Filename, resp.RespBody)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("Error: Can't get " + resp.Name + " :-(")
}
