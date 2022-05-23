package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {

	botLatency := s.HeartbeatLatency()
	pongMessage := fmt.Sprintf("%s ", botLatency) + "pong!"

	s.ChannelMessageSend(m.ChannelID, pongMessage)

}

func Pong(s *discordgo.Session, m *discordgo.MessageCreate) {

	botLatency := s.HeartbeatLatency()
	pingMessage := fmt.Sprintf("%s ", botLatency) + "ping!"

	s.ChannelMessageSend(m.ChannelID, pingMessage)

}
