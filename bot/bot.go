package bot

import (
	"discord-go-bot/bot/src/commands" // importing commands pachage
	"fmt"                             //	to print errors
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo" // discordgo package from the repo of bwmarrin .
	"github.com/joho/godotenv"
)

var BotId string
var goBot *discordgo.Session

func Start() {

	//********* This Part replit cannot have *********///
	//*************************************************//

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//*************************************************//

	Token := os.Getenv("TOKEN")

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

	// Letting the bot have all intents because why not.
	goBot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged)

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

	BotPrefix := "."

	// Split the user message around each instance of one or more consecutive white space characters
	messageSentFull := strings.Fields(m.Content)

	if len(messageSentFull) != 0 {
		// Saving the command field of the message
		messageFirstField := strings.ToLower(messageSentFull[0])

		//Bot musn't reply to it's own messages , to confirm it we perform this check.
		if m.Author.ID == BotId {
			return
		}
		//If we message ping to our bot in our discord it will return us pong .
		if messageFirstField == BotPrefix+"ping" {
			commands.Ping(s, m)
		}
		if messageFirstField == BotPrefix+"pong" {
			commands.Pong(s, m)
		}

		if messageFirstField == BotPrefix+"giveaway" {
			commands.Giveaway(s, m)
		}

		if messageFirstField == BotPrefix+"gopher" {
			commands.Gophers(s, m)
		}

		if messageFirstField == BotPrefix+"help" {
			helpEmbed := commands.Help()
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, helpEmbed)
		}

		if messageFirstField == "hi" {
			s.ChannelMessageSend(m.ChannelID, "Hello!")
		}

		if messageFirstField == "hello" {
			s.ChannelMessageSend(m.ChannelID, "Hi!")
		}

		if messageFirstField == "bye" {
			s.ChannelMessageSend(m.ChannelID, "Sayonara 👋")
		}

	}

}
