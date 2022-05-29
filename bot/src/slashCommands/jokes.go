package slashCommand

import (
	"discord-go-bot/bot/src/commands"
	util "discord-go-bot/bot/src/utils"
	"log"

	"github.com/bwmarrin/discordgo"
)

func DadJokesSlash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	jokes := commands.DadJoke()

	err := s.InteractionRespond(i.Interaction, util.MessageContentResponse(jokes))
	if err != nil {
		log.Fatal("could not fetch any jokes.", err)
	}
}

func JokesSlash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	jokes := commands.JokeData()

	err := s.InteractionRespond(i.Interaction, util.MessageContentResponse(jokes))
	if err != nil {
		log.Fatal("could not fetch any jokes.", err)
	}
}
