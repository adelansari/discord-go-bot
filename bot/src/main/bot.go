package bot

import (
	// importing commands pachage

	"flag"
	"fmt" //	to print errors
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo" // discordgo package from the repo of bwmarrin .
	"github.com/joho/godotenv"
)

var (
	BotId    string
	goBot    *discordgo.Session
	AppID    = flag.String("app", "", "Application ID")
	GuildID  = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken = flag.String("token", "", "Bot access token")
	Cleanup  = flag.Bool("cleanup", true, "Cleanup of commands after shutting down")
)

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

	// creating new bot session
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

	// Adding slash command handler:
	goBot.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:

			if h, ok := componentHandlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})

	// Letting the bot have all intents because why not.
	goBot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	err = goBot.Open()
	//Error handling
	if err != nil {
		fmt.Println("error opening connection", err.Error())
		return
	}

	cmdIDs := make(map[string]string, len(slashCmd))

	for _, v := range slashCmd {
		for _, guild := range goBot.State.Guilds {
			rcmd, err := goBot.ApplicationCommandCreate(goBot.State.User.ID, guild.ID, v)
			if err != nil {
				fmt.Println(fmt.Sprintf("Cannot create '%v' command: %v", v.Name, err))
			}
			cmdIDs[rcmd.ID] = rcmd.Name
		}
	}

	//If every thing works fine we will be printing this.
	fmt.Println("Bot is running !")

	// Cleanly close down the Discord session.
	defer goBot.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
	log.Println("Graceful shutdown")

	if !*Cleanup {
		return
	}

	for _, guild := range goBot.State.Guilds {
		for id, name := range cmdIDs {
			err := goBot.ApplicationCommandDelete(goBot.State.User.ID, guild.ID, id)
			if err != nil {
				fmt.Println(fmt.Sprintf("Cannot delete slash command %q: %v", name, err))
			}
		}
	}

	log.Println("Removed all slash commands.")

}
