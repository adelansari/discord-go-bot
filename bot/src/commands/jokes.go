package commands

import (
	util "discord-go-bot/bot/src/utils"
	"encoding/json"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

type Joke []struct {
	JokeText string `json:"joke"`
}

func BadJokes(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, JokeData())
}

func JokeData() string {
	joke := make(Joke, 1)
	json.Unmarshal(util.ApiData("jokes", "1"), &joke)
	return joke[0].JokeText
}

func DadJoke() string {
	jdata := util.DadJokeData()
	return jdata[rand.Intn(len(jdata))]
}
