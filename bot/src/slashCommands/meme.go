package slashCommand

import (
	"discord-go-bot/bot/src/commands"
	"log"

	"github.com/bwmarrin/discordgo"
)

func MemeSlash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	meme := commands.MemeEmbed()
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{meme},
		},
	})
	if err != nil {
		log.Fatal("could not fetch any memes.", err)
	}
}
