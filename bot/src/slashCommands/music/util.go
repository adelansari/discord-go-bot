package music

import (
	// "discordbot/src/context"

	"log"
	"os"

	// "discordbot/src/context"

	"github.com/bwmarrin/discordgo"
	"github.com/getsentry/sentry-go"
)

var goBot *discordgo.Session

// SearchGuild search the guild ID
func SearchGuild(textChannelID string) (guildID string) {
	channel, _ := goBot.Channel(textChannelID)
	guildID = channel.GuildID
	return guildID
}

// SearchVoiceChannel search the voice channel id into from guild.
func SearchVoiceChannel(user string) (voiceChannelID string) {
	for _, g := range goBot.State.Guilds {
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
	_, err := goBot.ChannelMessageSend(channelID, message)
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

	_, err = goBot.ChannelFileSend(channelID, name, reader)
	if err != nil {
		sentry.CaptureException(err)
	}
}
