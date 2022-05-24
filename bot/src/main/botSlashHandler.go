package bot

import (
	"discord-go-bot/bot/src/commands"
	"fmt"
	"log"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

type InteractionResponseData struct {
	TTS             bool               `json:"tts"`
	Content         string             `json:"content"`
	Components      []MessageComponent `json:"components"`
	Embeds          []*discordgo.MessageEmbed
	AllowedMentions *MessageAllowedMentions `json:"allowed_mentions,omitempty"`
	Flags           uint64                  `json:"flags,omitempty"`
	Files           []*File                 `json:"-"`

	// NOTE: autocomplete interaction only.
	Choices []*ApplicationCommandOptionChoice `json:"choices,omitempty"`

	CustomID string `json:"custom_id,omitempty"`
	Title    string `json:"title,omitempty"`
}

func (scmSlash *SlashFeature) Ping(s *discordgo.Session, i *discordgo.InteractionCreate) {
	botLatency := s.HeartbeatLatency()
	pongMessage := fmt.Sprintf("%s ", botLatency) + "pong!"
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: pongMessage,
		},
	})
	if err != nil {
		log.Fatal("could not run ping command", err)
	}
}

func (scmSlash *SlashFeature) Pong(s *discordgo.Session, i *discordgo.InteractionCreate) {
	botLatency := s.HeartbeatLatency()
	pingMessage := fmt.Sprintf("%s ", botLatency) + "ping!"
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: pingMessage,
		},
	})
	if err != nil {
		log.Fatal("could not run pong command", err)
	}
}

func (scmSlash *SlashFeature) Jokes(s *discordgo.Session, i *discordgo.InteractionCreate) {

	jokes := commands.JokeData()

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: jokes[rand.Intn(len(jokes))],
		},
	})
	if err != nil {
		log.Fatal("could not fetch any jokes.", err)
	}
}

func (scmSlash *SlashFeature) Giveaway(s *discordgo.Session, i *discordgo.InteractionCreate) {

	giveawayEmbed := commands.Giveaway

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:    false,
			Embeds: giveawayEmbed,
		},
	})
	if err != nil {
		log.Fatal("could not fetch any jokes.", err)
	}
}
