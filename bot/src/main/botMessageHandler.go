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

		switch messageFirstField {
		case BotPrefix + "ping":
			commands.Ping(s, m)
		case BotPrefix + "pong":
			commands.Pong(s, m)
		case BotPrefix + "giveaway":
			commands.Giveaway(s, m)
		case BotPrefix + "gopher":
			commands.Gophers(s, m)
		case BotPrefix + "jokes":
			commands.BadJokes(s, m)
		case BotPrefix + "invite":
			commands.Invite(s, m)
		case BotPrefix + "help":
			helpEmbed := commands.Help()
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, helpEmbed)
		case BotPrefix + "8ball":
			commands.MagicBall8(s, m)
		case "hi":
			s.ChannelMessageSend(m.ChannelID, "Hello!")
		case "hello":
			s.ChannelMessageSend(m.ChannelID, "Hi!")
		case "bye":
			s.ChannelMessageSend(m.ChannelID, "Sayonara 👋")
		}

	}

}
