package slashCommand

import (
	util "discord-go-bot/bot/src/utils"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func PingSlash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	botLatency := s.HeartbeatLatency()
	pongMessage := fmt.Sprintf("%s ", botLatency) + "pong!"

	err := s.InteractionRespond(i.Interaction, util.MessageContentResponse(pongMessage))
	if err != nil {
		log.Fatal("could not run ping command", err)
	}
}

func PongSlash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	botLatency := s.HeartbeatLatency()
	pingMessage := fmt.Sprintf("%s ", botLatency) + "ping!"

	err := s.InteractionRespond(i.Interaction, util.MessageContentResponse(pingMessage))
	if err != nil {
		log.Fatal("could not run pong command", err)
	}
}
