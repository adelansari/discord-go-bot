package bot

import (
	"discord-go-bot/bot/src/commands"      // importing commands pachage
	"fmt"                                    //	to print errors
  "os"
  
	"github.com/bwmarrin/discordgo" // discordgo package from the repo of bwmarrin .
)

var BotId string
var goBot *discordgo.Session

func Start() {
  
  Token := os.Getenv("Token")

	//creating new bot session
	goBot, err := discordgo.New("Bot " + Token)

	//Handling error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Making our bot a user using User function .
	u, err := goBot.User("@me")
	//Handlinf error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Storing our id from u to BotId .
	BotId = u.ID

	// Adding handler function to handle our messages using AddHandler from discordgo package.
	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	//Error handling
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//If every thing works fine we will be printing this.
	fmt.Println("Bot is running !")
}

//Definition of messageHandler function it takes two arguments first one is discordgo.Session which is s , second one is discordgo.MessageCreate which is m.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

  BotPrefix := os.Getenv("BotPrefix")
  
	//Bot musn't reply to it's own messages , to confirm it we perform this check.
	if m.Author.ID == BotId {
		return
	}
	//If we message ping to our bot in our discord it will return us pong .
	if m.Content == BotPrefix+"ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong!")
	}
	if m.Content == BotPrefix+"pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "ping!")
	}

	if m.Content == BotPrefix+"gopher" {
    commands.Gophers(s,m)
	}

	if m.Content == BotPrefix+"help" {
    helpEmbed:= commands.Help()
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, helpEmbed)
	}

}
