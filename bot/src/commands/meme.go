package commands

import (
	util "discord-go-bot/bot/src/utils"
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

type MemeStruct struct {
	PostLink  string   `json:"postLink"`
	Subreddit string   `json:"subreddit"`
	Title     string   `json:"title"`
	URL       string   `json:"url"`
	Nsfw      bool     `json:"nsfw"`
	Spoiler   bool     `json:"spoiler"`
	Author    string   `json:"author"`
	Ups       int      `json:"ups"`
	Preview   []string `json:"preview"`
}

func Meme(s *discordgo.Session, m *discordgo.MessageCreate) {
	ms := &discordgo.MessageSend{
		Embed: MemeEmbed(),
	}
	s.ChannelMessageSendComplex(m.ChannelID, ms)
}

func MemeEmbed() *discordgo.MessageEmbed {
	meme := MemeData()
	embed := &discordgo.MessageEmbed{
		Image: &discordgo.MessageEmbedImage{
			URL: meme.URL,
		},
		Title: meme.Title,
	}

	return embed
}

func MemeData() MemeStruct {
	meme := MemeStruct{}
	json.Unmarshal(util.MemeApiData(), &meme)
	return meme
}
