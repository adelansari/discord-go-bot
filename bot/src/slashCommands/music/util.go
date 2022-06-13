package music

import (
	"discord-go-bot/bot/src/context"
	"log"
	"os"

	"github.com/getsentry/sentry-go"
)

// SearchGuild search the guild ID
func SearchGuild(textChannelID string) (guildID string) {

	channel, _ := context.GoBot.Channel(textChannelID)
	guildID = channel.GuildID
	return guildID
}

// SearchVoiceChannel search the voice channel id into from guild.
func SearchVoiceChannel(user string) (voiceChannelID string) {
	for _, g := range context.GoBot.State.Guilds {
		for _, v := range g.VoiceStates {
			if v.UserID == user {
				return v.ChannelID
			}
		}
	}
	return ""
}

// SendChannelMessage sends a channel message to channel with channel id equal to m.ChannelID
func SendChannelMessage(channelID string, message string) {
	_, err := context.GoBot.ChannelMessageSend(channelID, message)
	if err != nil {
		sentry.CaptureException(err)
	}
}

func SendChannelFile(channelID string, filepath string, name string) {
	reader, err := os.Open(filepath)
	if err != nil {
		sentry.CaptureException(err)
		log.Println(err)
		return
	}

	_, err = context.GoBot.ChannelFileSend(channelID, name, reader)
	if err != nil {
		sentry.CaptureException(err)
	}
}
