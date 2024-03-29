package slashCommand

import (
	util "discord-go-bot/bot/src/utils"
	"fmt"
	"log"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

func GiveawaySlash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	giveawayAuthor := i.Member.User.ID

	for _, opt := range i.ApplicationCommandData().Options {

		switch opt.Name {
		case "create":

			userMessage := opt.StringValue()
			giveawayEmbed := []*discordgo.MessageEmbed{
				{
					Title: "🎉 Giveaway",
					Color: 453611,
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "**Giveaway** by: ",
							Value: fmt.Sprintf("<@%s>", giveawayAuthor),
						},
						{
							Name:  "**Content:**",
							Value: userMessage,
						},
					},
				},
			}

			// send the giveaway embed
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					TTS:    false,
					Embeds: giveawayEmbed,
				},
			})
			if err != nil {
				log.Fatal("could not fetch any jokes.", err)
			}

			// getting the response to the giveaway interaction
			msgInteraction, _ := s.InteractionResponse(i.Interaction)

			// adding reaction
			err = s.MessageReactionAdd(msgInteraction.ChannelID, msgInteraction.ID, "❤️")
			if err != nil {
				fmt.Println("could not add an emoji to the embed.", err.Error())
				return
			}

		case "pick":
			enteredMessageID := opt.StringValue()
			fetchedMessage, err := s.ChannelMessage(i.ChannelID, enteredMessageID)
			if err != nil {
				_ = s.InteractionRespond(
					i.Interaction,
					util.MessageContentResponse("The entered message ID is incorrect."),
				)
				fmt.Println("The entered message ID is incorrect.")
				return
			}
			// message author's name
			botUsername := fetchedMessage.Author.Username

			// Preparing emotes
			allReactions := fetchedMessage.Reactions // all reaction emotes on a certain message

			reactionEmotes := []string{}
			emotesFormated := []string{}

			for iteration1 := range fetchedMessage.Reactions {
				// formatting emotes:
				if len(allReactions[iteration1].Emoji.ID) == 0 {
					reactionEmotes = append(reactionEmotes, allReactions[iteration1].Emoji.Name)
					emotesFormated = append(emotesFormated, reactionEmotes[iteration1])
				} else {
					reactionEmotes = append(reactionEmotes, allReactions[iteration1].Emoji.Name+":"+allReactions[iteration1].Emoji.ID)
					emotesFormated = append(emotesFormated, "<:"+reactionEmotes[iteration1]+">")
				}

				// formatting users:
				reactionUsers, _ := s.MessageReactions(fetchedMessage.ChannelID, fetchedMessage.ID, reactionEmotes[iteration1], 50, "", "")

				// filtering out the bot or the message author's name
				reactionUsersFiltered := []*discordgo.User{}
				for iteration2 := range reactionUsers {
					if reactionUsers[iteration2].Username != botUsername {
						reactionUsersFiltered = append(reactionUsersFiltered, reactionUsers[iteration2])
					}
				}

				if len(reactionUsersFiltered) > 0 {
					randomIndex := rand.Intn(len(reactionUsersFiltered))
					userPick := fmt.Sprintf("%s", reactionUsersFiltered[randomIndex])

					// winner message corresponding to the emote:
					winnerMessage := "The winner for " + emotesFormated[iteration1] + " reaction is " + userPick
					s.InteractionRespond(i.Interaction, util.MessageContentResponse(winnerMessage))

				} else {
					_ = s.InteractionRespond(
						i.Interaction,
						util.MessageContentResponse("The message author reaction does not count."),
					)
				}
			}

		}
	}
}
