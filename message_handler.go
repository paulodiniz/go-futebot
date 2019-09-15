package main

import (
	"bytes"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func messageCreateEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Message.Content == "!tabela" {
		tabela := FetchTable()
		buf := bytes.Buffer{}
		for _, score := range tabela {
			buf.WriteString(fmt.Sprintf("%s: %d", score.team, score.points) + "\n")
		}

		s.ChannelMessageSend(m.ChannelID, buf.String())
	}
}
