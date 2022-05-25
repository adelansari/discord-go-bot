package commands

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MagicBall8(s *discordgo.Session, m *discordgo.MessageCreate) {

	messageSentFull := strings.Fields(m.Content)
	// A message reated to the giveaway
	magicBallQuestion := []string{}
	for i := 1; i < len(messageSentFull); i++ {
		magicBallQuestion = append(magicBallQuestion, messageSentFull[i])
	}
	magicBallJoined := strings.Join(magicBallQuestion, " ")

	if len(messageSentFull) == 1 {
		s := s
		m := m
		magicBallEmbed := &discordgo.MessageEmbed{
			Title: "8ball Help",
			Description: fmt.Sprintf("\u200B\nI can answer all your [yes/no] questions!\n" +
				"Use the following command to be enlightened on your life-changing question.\n\u200B"),
			Fields: []*discordgo.MessageEmbedField{
				{
					Name: "`.8ball YourQuestion`",
					Value: "Example:\n" +
						"*.8ball Should I try bungee jumping?*",
				},
			},
			Color: 9589448, // hex color to decimal
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, magicBallEmbed)

	} else {
		magicAnswers := MagicAnswers()
		magicBallResponse := fmt.Sprintf("> %s\n%s", magicBallJoined, magicAnswers[rand.Intn(len(magicAnswers))])
		s.ChannelMessageSend(m.ChannelID, magicBallResponse)
	}

}

func MagicAnswers() []string {
	var magicAnswers = []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}

	return magicAnswers
}
