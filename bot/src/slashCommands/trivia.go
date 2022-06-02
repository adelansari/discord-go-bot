package slashCommand

import (
	util "discord-go-bot/bot/src/utils"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type triviaApi struct {
	ResponseCode int `json:"response_code"`
	Results      []struct {
		Category         string   `json:"category"`
		Type             string   `json:"type"`
		Difficulty       string   `json:"difficulty"`
		Question         string   `json:"question"`
		CorrectAnswer    string   `json:"correct_answer"`
		IncorrectAnswers []string `json:"incorrect_answers"`
	} `json:"results"`
}

var (
	correctAnswer string
	allAnswers    []string
	triviaBtn     []string
	btnEmoji      []string
	question      string
)

func createQuestion(s *discordgo.Session, i *discordgo.InteractionCreate) *discordgo.InteractionResponse {

	var data triviaApi

	// unmarshall
	json.Unmarshal(util.TriviaApiData(), &data)

	// General Knowledge category
	// Loop through the Results node for the Question

	for _, rec := range data.Results {
		question = rec.Question
		question = strings.ReplaceAll(question, "&quot;", "`")
		question = strings.ReplaceAll(question, "&#039;", "'")
		correctAnswer = rec.CorrectAnswer
		allAnswers = rec.IncorrectAnswers
		allAnswers = append(allAnswers, correctAnswer)
		rand.Shuffle(len(allAnswers), func(i, j int) { allAnswers[i], allAnswers[j] = allAnswers[j], allAnswers[i] })
	}

	fmt.Println(correctAnswer)

	for index := range allAnswers {
		triviaCustomID := "triviaIndex_" + fmt.Sprintf("%d", index)
		triviaBtn = append(triviaBtn, triviaCustomID)
	}
	btnEmoji = []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣"}
	components := []discordgo.MessageComponent{}
	for index, element := range allAnswers {
		btn := discordgo.Button{
			Emoji: discordgo.ComponentEmoji{
				Name: btnEmoji[index],
			},
			Label:    element,
			Style:    discordgo.SecondaryButton,
			CustomID: triviaBtn[index],
		}
		components = append(components, btn)
	}
	triviaMessage := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: question,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: components,
				},
			},
		},
	}
	return triviaMessage
}

func TriviaSlash(s *discordgo.Session, i *discordgo.InteractionCreate) {

	switch i.Type {
	case discordgo.InteractionApplicationCommand:

		err := s.InteractionRespond(i.Interaction, createQuestion(s, i))
		if err != nil {
			fmt.Println("Could not send the trivia question", err.Error())
		}
	case discordgo.InteractionMessageComponent:

		//s.ChannelMessageDelete(i.ChannelID, i.Message.ID)

		// storing the custom ID after the user clicks on any button.
		btnCustomID := i.MessageComponentData().CustomID

		// Finding element index
		correctAnswerIndex := util.Find(allAnswers, correctAnswer)
		btnCustomIDIndex := util.Find(triviaBtn, btnCustomID)

		components := []discordgo.MessageComponent{}
		var btn discordgo.Button
		for index, element := range allAnswers {
			if index == correctAnswerIndex {
				btn = discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: btnEmoji[index],
					},
					Label:    element,
					Disabled: true,
					Style:    discordgo.SuccessButton,
					CustomID: triviaBtn[index],
				}
			} else {
				btn = discordgo.Button{
					Emoji: discordgo.ComponentEmoji{
						Name: btnEmoji[index],
					},
					Label:    element,
					Style:    discordgo.DangerButton,
					Disabled: true,
					CustomID: triviaBtn[index],
				}
			}

			components = append(components, btn)
		}

		var btnResp string
		if btnCustomIDIndex == correctAnswerIndex {
			btnResp = fmt.Sprintf(question+"\n🎊 The correct answer was indeed %s.", correctAnswer)
			s.ChannelMessageEditComplex(&discordgo.MessageEdit{
				Content: &btnResp,
				ID:      i.Message.ID,
				Channel: i.ChannelID,
				Flags:   discordgo.MessageFlagsLoading,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: components,
					},
				},
			})
			err := s.InteractionRespond(i.Interaction, createQuestion(s, i))
			if err != nil {
				fmt.Println("Could not send the trivia question", err.Error())
			}
		} else {
			btnResp = fmt.Sprintf(question+"\n%s is incorrect unfortunately. 😞", allAnswers[btnCustomIDIndex])
			s.ChannelMessageEditComplex(&discordgo.MessageEdit{
				Content: &btnResp,
				ID:      i.Message.ID,
				Channel: i.ChannelID,
				Flags:   discordgo.MessageFlagsLoading,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: components,
					},
				},
			})
			err := s.InteractionRespond(i.Interaction, createQuestion(s, i))
			if err != nil {
				fmt.Println("Could not send the trivia question", err.Error())
			}
		}
	}

}
