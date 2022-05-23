package bot

import (
	// importing commands pachage

	"flag"
	"fmt" //	to print errors
	"log"
	"os"

	scm "github.com/ethanent/discordgo-scm"

	"github.com/bwmarrin/discordgo" // discordgo package from the repo of bwmarrin .
	"github.com/joho/godotenv"
)

var BotId string
var goBot *discordgo.Session
var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "", "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

type SlashFeature struct {
	Session *discordgo.Session
	Manager *scm.SCM
	Token   string
	Guild   string
}

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
	// if Token == "" {
	// 	log.Fatal("Bot token was not found.")
	// }

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
	// goBot.AddHandler(slashHandler)   //  slash hander if needed

	// Letting the bot have all intents because why not.
	goBot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged)

	err = goBot.Open()
	//Error handling
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Creating slash commands with SCM

	scmSlash := &SlashFeature{}
	scmSlash.Token = Token
	scmSlash.Guild = *GuildID
	scmSlash.Session = goBot
	// creating scm manager
	scmSlash.Manager = scm.NewSCM()

	// creating and populating features
	// slashCommands := []*discordgo.ApplicationCommand{
	// 	{
	// 		Name: "ping",

	// 		Description: "To ping the bot!",
	// 	},
	// 	{
	// 		Name: "pong",

	// 		Description: "To pong the bot!",
	// 	},
	// 	{
	// 		Name: "help",

	// 		Description: "To show help embed!",
	// 	},
	// }

	PingSlash := []*scm.Feature{
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: scmSlash.Ping,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "ping",
				Description: "To ping the bot!",
			},
		},
	}
	PongSlash := []*scm.Feature{
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: scmSlash.Pong,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "pong",
				Description: "To ping the bot!",
			},
		},
	}
	JokeSlash := []*scm.Feature{
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: scmSlash.Jokes,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "joke",
				Description: "The best dad jokes on Discord!",
			},
		},
	}

	// Register ApplicationCommands
	scmSlash.Manager.AddFeatures(PingSlash)
	scmSlash.Manager.AddFeatures(PongSlash)
	scmSlash.Manager.AddFeatures(JokeSlash)

	scmSlash.Session.AddHandler(scmSlash.Manager.HandleInteraction)
	err = scmSlash.Manager.CreateCommands(scmSlash.Session, scmSlash.Guild)
	if err != nil {
		log.Fatal("Failed to create slash commands.", err)
	}

	// Command handler by sid:

	// command := &discordgo.ApplicationCommand{
	// 	Name:        "ping",
	// 	Type:        discordgo.ChatApplicationCommand,
	// 	Description: "Ping me!",
	// }

	// //registeredCommand, err := discordgo.ApplicationCommandCreate(goBot.State.User.ID, *GuildID, command)
	// registeredCommand, err := goBot.ApplicationCommandCreate(goBot.State.User.ID, *GuildID, command)
	// if err != nil {
	// 	fmt.Println(registeredCommand)
	// 	log.Panicf("Cannot create '%v' command: %v", command.Name, err)
	// }

	//If every thing works fine we will be printing this.
	fmt.Println("Bot is running !")

	// cleanly close down the Discord session

	// err = scmSlash.Manager.DeleteCommands(scmSlash.Session, scmSlash.Guild)
	// if err != nil {
	// 	log.Print("could not delete commands", err)
	// }

	// scmSlash.Session.Close()

}

// Slash handler by sid:

// func slashHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

// 	if i.Type != discordgo.InteractionApplicationCommand {
// 		return
// 	}

// 	data := i.ApplicationCommandData()
// 	switch data.Name {
// 	case "ping":
// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
// 			Data: &discordgo.InteractionResponseData{
// 				Content: fmt.Sprintf("Pong!"),
// 			},
// 		})
// 	}
// }
