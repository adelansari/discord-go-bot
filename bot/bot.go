package bot

import (
	commands "discord-go-bot/bot/src/commands"         // importing the gophers pachage
	paginator "discord-go-bot/bot/src/utils/paginator" // importing the paginator pachage
	"discord-go-bot/config"                            //	importing our config package which we have created above
	"fmt"                                              //	to print errors
	"time"

	"github.com/bwmarrin/discordgo" // discordgo package from the repo of bwmarrin .
)

var BotId string
var goBot *discordgo.Session

func Start() {

	//creating new bot session
	goBot, err := discordgo.New("Bot " + config.Token)

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
	//Bot musn't reply to it's own messages , to confirm it we perform this check.
	if m.Author.ID == BotId {
		return
	}
	//If we message ping to our bot in our discord it will return us pong .
	if m.Content == config.BotPrefix+"ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong!")
	}
	if m.Content == config.BotPrefix+"pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "ping!")
	}

	if m.Content == config.BotPrefix+"gopher" {
		p := paginator.NewPaginator(s, m.ChannelID)

		// Add embed pages to paginator

		// returning values of gophersImages and gophersUrl from images/gophers.go
		gopherImages, gophersUrl := commands.Gophers()

		for i := 0; i < len(gopherImages); i++ {
			gopherName := fmt.Sprintf("Image %d: %s", i+1, gopherImages[i])
			p.Add(&discordgo.MessageEmbed{
				Title:       "Collectionof Gopher Images",
				Description: gopherName,
				Image:       &discordgo.MessageEmbedImage{URL: gophersUrl[i]},
			})
		}

		// p.Add(&discordgo.MessageEmbed{Description: "Page one"},
		// 	&discordgo.MessageEmbed{Description: "Page two"},
		// 	&discordgo.MessageEmbed{Description: "Page three"})

		// Sets the footers of all added pages to their page numbers.
		p.SetPageFooters()

		// When the paginator is done listening set the colour to yellow
		p.ColourWhenDone = 0xffff

		// Stop listening for reaction events after five minutes
		p.Widget.Timeout = time.Minute * 5

		// Add a custom handler for the gun reaction.
		p.Widget.Handle("🔫", func(w *paginator.Widget, r *discordgo.MessageReaction) {
			s.ChannelMessageSend(m.ChannelID, "Bang!")
		})

		p.Spawn()
	}

	if m.Content == config.BotPrefix+"help" {

		helpEmbed := commands.Help()
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, helpEmbed)
	}

}
