package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Help() *discordgo.MessageEmbed {
	helpEmbed := &discordgo.MessageEmbed{
		Title: "Bot Commands",
		Description: fmt.Sprintf("\nHelp Command\n" +
			"`.help`    -  A list of help commands\n" +
			"\nCore Commands:\n" +
			"`.gopher`  -  To show pages of Gopher images in an embed\n" +
			"`.giveaway`- Creating a giveaway or picking a random winner!\n" +
			"`.jokes`   -  Displays a random joke.\n" +
			"`.8ball`   -  Answer to all your [yes/no] questions.\n" +

			"\nMisc. Commands:\n" +
			"`.ping`    -  To ping the bot!\n" +
			"`.pong`    -  To pong the bot!\n" +
			"`.invite`  -  To invite the bot to your server!\n"),
		Color: 3143071, // hex color to decimal
	}

	return helpEmbed

}
