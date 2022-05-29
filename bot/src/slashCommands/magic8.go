package slashCommand

import (
	"discord-go-bot/bot/src/commands"
	util "discord-go-bot/bot/src/utils"
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

func MagicBallSlash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	query := i.ApplicationCommandData().Options[0].StringValue()
	magicAnswers := commands.MagicAnswers()
	resp := fmt.Sprintf("> %s\n%s", query, magicAnswers[rand.Intn(len(magicAnswers))])

	s.InteractionRespond(i.Interaction, util.MessageContentResponse(resp))
}
