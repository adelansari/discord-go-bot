package commands

import (
	util "discord-go-bot/bot/src/utils"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, pongMessage := PingCommon(s)

	s.ChannelMessageSend(m.ChannelID, pongMessage)

}

func Pong(s *discordgo.Session, m *discordgo.MessageCreate) {
	pingMessage, _ := PingCommon(s)
	s.ChannelMessageSend(m.ChannelID, pingMessage)

}

func PingSlash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_, pongMessage := PingCommon(s)
	err := s.InteractionRespond(i.Interaction, util.MessageContentResponse(pongMessage))
	if err != nil {
		log.Fatal("could not run ping command", err)
	}
}

func PongSlash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	pingMessage, _ := PingCommon(s)
	err := s.InteractionRespond(i.Interaction, util.MessageContentResponse(pingMessage))
	if err != nil {
		log.Fatal("could not run pong command", err)
	}
}

func PingCommon(s *discordgo.Session) (string, string) {
	botLatency := s.HeartbeatLatency()
	pingMessage := fmt.Sprintf("%s ", botLatency) + "ping!"
	pongMessage := fmt.Sprintf("%s ", botLatency) + "pong!"

	return pingMessage, pongMessage
}
