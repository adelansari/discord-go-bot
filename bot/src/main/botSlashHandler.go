package bot

import (
	"discord-go-bot/bot/src/commands"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

// JSON pretty print by marshaling value
func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

// Responding to interaction with a message
func messageContentResponse(c string) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: c,
		},
	}
}

func (scmSlash *SlashFeature) Ping(s *discordgo.Session, i *discordgo.InteractionCreate) {

	botLatency := s.HeartbeatLatency()
	pongMessage := fmt.Sprintf("%s ", botLatency) + "pong!"

	err := s.InteractionRespond(i.Interaction, messageContentResponse(pongMessage))
	if err != nil {
		log.Fatal("could not run ping command", err)
	}
}

func (scmSlash *SlashFeature) Pong(s *discordgo.Session, i *discordgo.InteractionCreate) {

	botLatency := s.HeartbeatLatency()
	pingMessage := fmt.Sprintf("%s ", botLatency) + "ping!"

	err := s.InteractionRespond(i.Interaction, messageContentResponse(pingMessage))
	if err != nil {
		log.Fatal("could not run pong command", err)
	}
}

func (scmSlash *SlashFeature) Jokes(s *discordgo.Session, i *discordgo.InteractionCreate) {

	jokes := commands.JokeData()

	err := s.InteractionRespond(i.Interaction, messageContentResponse(jokes[rand.Intn(len(jokes))]))
	if err != nil {
		log.Fatal("could not fetch any jokes.", err)
	}
}

func (scmSlash *SlashFeature) Giveaway(s *discordgo.Session, i *discordgo.InteractionCreate) {

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
					messageContentResponse("The entered message ID is incorrect."),
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
					err := s.InteractionRespond(i.Interaction, messageContentResponse(winnerMessage))
					if err != nil {
						log.Fatal("could not fetch any jokes.", err)
					}
				} else {
					_ = s.InteractionRespond(
						i.Interaction,
						messageContentResponse("The message author reaction does not count."),
					)
				}
			}
		}
	}

}

func (scmSlash *SlashFeature) handleSay(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var response string
	if len(i.ApplicationCommandData().Options) > 0 {
		response = i.ApplicationCommandData().Options[0].StringValue()
	}
	if response == "" {
		response = "Say what?"
	}

	err := s.InteractionRespond(i.Interaction, messageContentResponse(response))
	if err != nil {
		log.Fatal("could not fetch any jokes.", err)
	}

}

func (scmSlash *SlashFeature) Invite(s *discordgo.Session, i *discordgo.InteractionCreate) {

	inviteLink := fmt.Sprintf("https://discord.com/oauth2/authorize?client_id=977285728042111027&permissions=1644971949425&scope=bot%%20applications.commands")
	inviteMessage := fmt.Sprintf("Please visit %s to add this bot to your server.", inviteLink)

	s.InteractionRespond(i.Interaction, messageContentResponse(inviteMessage))
}

func (scmSlash *SlashFeature) MagicBall(s *discordgo.Session, i *discordgo.InteractionCreate) {
	query := i.ApplicationCommandData().Options[0].StringValue()
	magicAnswers := commands.MagicAnswers()
	resp := fmt.Sprintf("> %s\n%s", query, magicAnswers[rand.Intn(len(magicAnswers))])

	s.InteractionRespond(i.Interaction, messageContentResponse(resp))
}
