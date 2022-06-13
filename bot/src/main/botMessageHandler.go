package bot

import (
	"discord-go-bot/bot/src/commands"
	music "discord-go-bot/bot/src/slashCommands/music"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//Definition of messageHandler function it takes two arguments first one is discordgo.Session which is s , second one is discordgo.MessageCreate which is m.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	BotPrefix := "."
	v := music.VoiceInstances[m.GuildID]

	// Split the user message around each instance of one or more consecutive white space characters
	messageSentFull := strings.Fields(m.Content)

	if len(messageSentFull) != 0 {
		// Saving the command field of the message
		messageFirstField := strings.ToLower(messageSentFull[0])

		//Bot musn't reply to it's own messages or any other bot
		if m.Author.ID == s.State.User.ID || m.Author.Bot == true {
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
		case BotPrefix + "joke":
			commands.BadJokes(s, m)
		case BotPrefix + "meme":
			commands.Meme(s, m)
		case BotPrefix + "invite":
			commands.Invite(s, m)
		case BotPrefix + "help", s.State.User.Mention():
			commands.Help(s, m)
		case BotPrefix + "8ball":
			commands.MagicBall8(s, m)
		case BotPrefix + "trivia":
			commands.Trivia(s, m)
		case BotPrefix + "facts":
			commands.Facts(s, m)
		case BotPrefix + "factstimer":
			commands.FactsTimer(s, m)
		case BotPrefix + "play":
			if len(messageSentFull) > 1 {
				music.PlayMusic(messageSentFull[1:], v, s, m)
			} else {
				s.ChannelMessageSend(m.ChannelID, "Type the song name after the command.")
			}
		case BotPrefix + "leave":
			music.LeaveVoice(v, m)
		case BotPrefix + "skip":
			music.SkipMusic(v, m)
		case BotPrefix + "stop":
			music.StopMusic(v, m)
		case "hi":
			s.ChannelMessageSend(m.ChannelID, "Hello!")
		case "hello":
			s.ChannelMessageSend(m.ChannelID, "Hi!")
		case "bye":
			s.ChannelMessageSend(m.ChannelID, "Sayonara 👋")
		}

	}

}
