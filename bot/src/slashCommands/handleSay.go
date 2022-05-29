package slashCommand

import (
	util "discord-go-bot/bot/src/utils"
	"log"

	"github.com/bwmarrin/discordgo"
)

func SaySlash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var response string
	if len(i.ApplicationCommandData().Options) > 0 {
		response = i.ApplicationCommandData().Options[0].StringValue()
	}
	if response == "" {
		response = "Say what?"
	}

	err := s.InteractionRespond(i.Interaction, util.MessageContentResponse(response))
	if err != nil {
		log.Fatal("could not fetch any jokes.", err)
	}
}
