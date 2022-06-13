package context

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var GoBot *discordgo.Session

func Initialize(discordToken string) {
	var err error
	//creating new bot session
	GoBot, err = discordgo.New("Bot " + discordToken)
	//Handling error
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Making our bot a user using User function .
	u, err := GoBot.User("@me")
	//Handlinf error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Storing our id from u to BotId .
	BotId := u.ID
	_ = BotId

}

func OpenConnection() {
	err := GoBot.Open()
	//Error handling
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
