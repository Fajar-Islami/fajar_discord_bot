package model

import "github.com/bwmarrin/discordgo"

type (
	ServiceCommon struct {
		S *discordgo.Session
		M *discordgo.MessageCreate
	}

	ListRunTimes struct {
		Runtime []struct {
			Language string   `json:"language,omitempty"`
			Version  string   `json:"version,omitempty"`
			Aliases  []string `json:"aliases,omitempty"`
			Compiled bool     `json:"compiled,omitempty"`
		} `json:"runtime,omitempty"`
	}
)
