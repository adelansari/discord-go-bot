package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Say(s *discordgo.Session, m *discordgo.MessageCreate) {

	// sentMessageID := m.Reference().MessageID
	messageSentFull := strings.Fields(m.Content)
	userMessageSentArray := messageSentFull[1:] // includes the user message without the command

	fmt.Println(userMessageSentArray)

	var userMessage string

	if len(userMessageSentArray) == 0 {
		userMessage = "Say what?"
	} else {
		userMessage = strings.Join(userMessageSentArray, " ")
		// Deleting the initial user message:
		err := s.ChannelMessageDelete(m.ChannelID, m.ID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	s.ChannelMessageSend(m.ChannelID, userMessage)

}
