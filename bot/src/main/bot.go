package bot

import (
	// importing commands pachage

	"discord-go-bot/bot/src/context"
	music "discord-go-bot/bot/src/slashCommands/music"
	"flag"
	"fmt" //	to print errors
	"log"
	"os"
	"os/signal"

	scm "github.com/ethanent/discordgo-scm"

	"github.com/bwmarrin/discordgo" // discordgo package from the repo of bwmarrin .
	"github.com/joho/godotenv"
)

var (
	BotId          string
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "", "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
	scmSlash       *SlashFeature
)

type SlashFeature struct {
	Session *discordgo.Session
	Manager *scm.SCM
	Token   string
	Guild   string
}

func Start() {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Replit doesn't need to read .env files.")
	}

	Token := os.Getenv("TOKEN")
	// if Token == "" {
	// 	log.Fatal("Bot token was not found.")
	// }

	context.Initialize(Token)

	// Adding handler function to handle our messages using AddHandler from discordgo package.

	context.GoBot.AddHandler(messageHandler)

	// Letting the bot have all intents because why not.
	context.GoBot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	// opening discord connection
	context.OpenConnection()

	// Creating slash commands with SCM

	scmSlash = &SlashFeature{}
	scmSlash.Token = Token
	scmSlash.Guild = *GuildID
	scmSlash.Session = context.GoBot
	// creating scm manager
	scmSlash.Manager = scm.NewSCM()

	// Register ApplicationCommands
	slashCommandFeatures(slashfeatures)
	scmSlash.Manager.AddFeatures(*slashfeatures)

	scmSlash.Session.AddHandler(scmSlash.Manager.HandleInteraction)
	err = scmSlash.Manager.CreateCommands(scmSlash.Session, scmSlash.Guild)
	if err != nil {
		log.Fatal("Failed to create slash commands.", err)
	}

	//If every thing works fine we will be printing this.
	fmt.Println("Bot is running !")

	music.InitializeRoutine()

	// Cleanly close down the Discord session.
	defer context.GoBot.Close()

	stop := make(chan os.Signal, 1)
	// signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	// signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

}
