package bot

import (
	"discord-go-bot/bot/src/commands"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

// JSON pretty print by marshaling value
func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

// Responding to interaction with a message
func messageContentResponse(c string) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: c,
		},
	}
}

func (scmSlash *SlashFeature) Ping(s *discordgo.Session, i *discordgo.InteractionCreate) {
	botLatency := s.HeartbeatLatency()
	pongMessage := fmt.Sprintf("%s ", botLatency) + "pong!"
	err := s.InteractionRespond(i.Interaction, messageContentResponse(pongMessage))
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

	fmt.Println(PrettyStruct(i.ApplicationCommandData().Options[0]))

	// var messageResponse []string = []string{
	// 	"Please enter the giveaway message!",
	// 	"Please enter the message ID cooresponding to the giveaway message!",
	// }

	for _, opt := range i.ApplicationCommandData().Options {
		switch opt.Name {
		case "create":
			userMessage := opt.StringValue()
			fmt.Println(userMessage)

		case "pick":
			userMessage := opt.StringValue()
			fmt.Println(userMessage)
		}
	}

	giveawayEmbed := []*discordgo.MessageEmbed{
		{
			Title: "Bot Commands",
		},
	}

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

func (scmSlash *SlashFeature) handleSay(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var response string
	if len(i.ApplicationCommandData().Options) > 0 {
		response = i.ApplicationCommandData().Options[0].StringValue()
	}
	if response == "" {
		response = "Say what?"
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:     false,
			Content: response,
		},
	})
	if err != nil {
		log.Fatal("could not fetch any jokes.", err)
	}

}
