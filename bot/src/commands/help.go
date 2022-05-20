package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Help() *discordgo.MessageEmbed {
	helpEmbed := &discordgo.MessageEmbed{
		Title: "Bot Commands",
		Description: fmt.Sprintf("`!help`    -  A list of help commands\n" +
			"`!ping`    -  To ping the bot!\n" +
			"`!pong`    -  To pong the bot!\n" +
			"`!gopher`  -  To show pages of Gopher images in an embed"),
		Color: 3699351, // hex color to decimal
	}

	return helpEmbed

}
