package commands

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// func GiveawayExport() {
// 	giveawayHelpEmbed := Giveaway(s *discordgo.Session, m *discordgo.MessageCreate)
// }

func Giveaway(s *discordgo.Session, m *discordgo.MessageCreate) {

	messageSentFull := strings.Fields(m.Content)
	// the message ID inputted by the user
	if len(messageSentFull) == 1 {
		s := s
		m := m
		giveawayHelpEmbed := &discordgo.MessageEmbed{
			Title:       "Giveaway Help",
			Description: fmt.Sprintf("To create a giveaway you must use one of the following commands:"),
			Fields: []*discordgo.MessageEmbedField{
				{
					Name: "`.giveaway create YourMessageHere`",
					Value: "To create a message reaction embed with the giveaway content.\n" +
						"Example:\n*.giveaway create This is a giveaway for Scarlet Nexus!*",
				},
				{
					Name: "\n`.giveaway pick MessageID`",
					Value: "To pick a winner from the giveaway with the Message ID.\n" +
						"Example:\n*.giveaway pick 978202141602742302*",
				},
			},
			Color: 9589448, // hex color to decimal
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, giveawayHelpEmbed)

	} else {
		giveawayCommand := strings.ToLower(messageSentFull[1])

		switch giveawayCommand {
		case "create":
			GiveawayCreate(s, m)
		case "pick":
			PickWinner(s, m)
		}
	}

}

func GiveawayCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	messageSentFull := strings.Fields(m.Content)
	// A message reated to the giveaway
	givewayMessage := []string{}
	for i := 2; i < len(messageSentFull); i++ {
		givewayMessage = append(givewayMessage, messageSentFull[i])
	}
	givewayMessageJoined := strings.Join(givewayMessage, " ")

	giveawayAuthor := m.Author.ID

	giveawayEmbed := &discordgo.MessageEmbed{
		Title: "🎉 Giveaway",
		Color: 453611,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "**Giveaway** by: ",
				Value: fmt.Sprintf("<@%s>", giveawayAuthor),
			},
			{
				Name:  "**Content:**",
				Value: givewayMessageJoined,
			},
		},
	}

	embedObject, err := s.ChannelMessageSendEmbed(m.ChannelID, giveawayEmbed)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = s.MessageReactionAdd(embedObject.ChannelID, embedObject.ID, "❤️")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Deleting the initial user message:
	err = s.ChannelMessageDelete(m.ChannelID, m.ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func PickWinner(s *discordgo.Session, m *discordgo.MessageCreate) {

	messageSentFull := strings.Fields(m.Content)
	// the message ID inputted by the user
	messageIDField := strings.ToLower(messageSentFull[2])

	// Getting a the message by ID
	fetchedMessage, err := s.ChannelMessage(m.ChannelID, messageIDField)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Bot or the message author's name
	botUsername := fetchedMessage.Author.Username

	// Preparing emotes
	allReactions := fetchedMessage.Reactions // all reaction emotes on a certain message
	reactionEmotes := []string{}
	emotesFormated := []string{}
	// fetchedUsers := []string{}

	for i := range fetchedMessage.Reactions {
		// formatting emotes:
		if len(allReactions[i].Emoji.ID) == 0 {
			reactionEmotes = append(reactionEmotes, allReactions[i].Emoji.Name)
			emotesFormated = append(emotesFormated, reactionEmotes[i])
		} else {
			reactionEmotes = append(reactionEmotes, allReactions[i].Emoji.Name+":"+allReactions[i].Emoji.ID)
			emotesFormated = append(emotesFormated, "<:"+reactionEmotes[i]+">")
		}

		// formatting users:
		reactionUsers, _ := s.MessageReactions(m.ChannelID, fetchedMessage.ID, reactionEmotes[i], 50, "", "")

		// filtering out the bot or the message author's name
		reactionUsersFiltered := []*discordgo.User{}
		for j := range reactionUsers {
			if reactionUsers[j].Username != botUsername {
				reactionUsersFiltered = append(reactionUsersFiltered, reactionUsers[j])
			}
		}

		if len(reactionUsersFiltered) > 0 {
			randomIndex := rand.Intn(len(reactionUsersFiltered))
			userPick := fmt.Sprintf("%s", reactionUsersFiltered[randomIndex])

			// winner message corresponding to the emote:
			winnerMessage := "The winner for " + emotesFormated[i] + " reaction is " + userPick
			s.ChannelMessageSend(m.ChannelID, winnerMessage)
		} else {
			s.ChannelMessageSend(
				m.ChannelID,
				"The message author reaction does not count.",
			)
		}
	}

}
