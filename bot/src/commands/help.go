package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	helpEmbed := &discordgo.MessageEmbed{
		Title: "Bot Commands",
		Description: fmt.Sprintf("\nHelp Command\n"+
			"`.help`    -  A list of help commands\n"+

			"\nCore Commands:\n"+
			"`.gopher`  -  To show pages of Gopher images in an embed\n"+
			"`.giveaway`- Creating a giveaway or picking a random winner!\n"+
			"`.joke`   -  Displays a random joke.\n"+
			"`.8ball`   -  Answer to all your [yes/no] questions.\n"+
			"`/trivia`   -  Shows general trivia miltiple choice questions.\n"+
			"`.meme`   -  Embeds a random meme image.\n"+
			"`.facts`   -  Iterate through a list of cool facts.\n"+
			"`.factstimer`   -  Sends a fact every 12 hours.\n"+

			"\nMusic Commands:\n"+
			"`.play`   -  Searches the music, joins the voice channel and plays the music.\n"+
			"`.stop`   -   Stops the music in the voice channel\n"+
			"`.leave`   -  Leaves the voice channel.\n"+
			"`.skip`   -  Skips the current music in the queue.\n"+
			"`.queue`  -  Displays the music queue.\n"+

			"\nMisc. Commands:\n"+
			"`.ping`    -  To ping the bot!\n"+
			"`.pong`    -  To pong the bot!\n"+
			"`.say`    	-  To repeat what the user says (allows gif emote usage)!\n"+
			"`.invite`  -  To invite the bot to your server!\n") +

			"\nPS: All these commands work with slash commands `/` as well.\n",

		Color: 3143071, // hex color to decimal
	}
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, helpEmbed)

}
