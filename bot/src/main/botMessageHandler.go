package bot

import (
	"discord-go-bot/bot/src/commands"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//Definition of messageHandler function it takes two arguments first one is discordgo.Session which is s , second one is discordgo.MessageCreate which is m.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	BotPrefix := "+"

	// Split the user message around each instance of one or more consecutive white space characters
	messageSentFull := strings.Fields(m.Content)

	if len(messageSentFull) != 0 {
		// Saving the command field of the message
		messageFirstField := strings.ToLower(messageSentFull[0])

		//Bot musn't reply to it's own messages , to confirm it we perform this check.
		if m.Author.ID == BotId {
			return
		}
		//If we message ping to our bot in our discord it will return us pong .
		if messageFirstField == BotPrefix+"ping" {
			commands.Ping(s, m)
		}
		if messageFirstField == BotPrefix+"pong" {
			commands.Pong(s, m)
		}

		if messageFirstField == BotPrefix+"giveaway" {
			commands.Giveaway(s, m)
		}

		if messageFirstField == BotPrefix+"gopher" {
			commands.Gophers(s, m)
		}
		if messageFirstField == BotPrefix+"jokes" {
			commands.BadJokes(s, m)
		}

		if messageFirstField == BotPrefix+"help" {
			helpEmbed := commands.Help()
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, helpEmbed)
		}

		if messageFirstField == "hi" {
			s.ChannelMessageSend(m.ChannelID, "Hello!")
		}

		if messageFirstField == "hello" {
			s.ChannelMessageSend(m.ChannelID, "Hi!")
		}

		if messageFirstField == "bye" {
			s.ChannelMessageSend(m.ChannelID, "Sayonara 👋")
		}

	}

}
