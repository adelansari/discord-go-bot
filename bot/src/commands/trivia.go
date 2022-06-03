package commands

import (
	// util "discord-go-bot/bot/src/utils"
	util "discord-go-bot/bot/src/utils"
	"encoding/json"
	"fmt"
	"html"

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

func Trivia(s *discordgo.Session, m *discordgo.MessageCreate) {

	var data triviaApi

	// unmarshall
	json.Unmarshal(util.TriviaApiData(), &data)

	// General Knowledge category
	// Loop through the Results node for the Question
	var question string
	var correctAnswer string
	var allAnswers []string
	for _, rec := range data.Results {
		question = rec.Question
		correctAnswer = rec.CorrectAnswer
		allAnswers = rec.IncorrectAnswers
		allAnswers = append(allAnswers, correctAnswer)
	}

	// Unescaping entities and cleaning up the HTML special character codes:
	question = html.UnescapeString(question)
	correctAnswer = html.UnescapeString(correctAnswer)
	for index := range allAnswers {
		allAnswers[index] = html.UnescapeString(allAnswers[index])
	}

	// Suffling allAnswers elements because the previous method was not random
	util.Shuffle(allAnswers)

	// var triviaMap map[string]interface{}
	// err = json.NewDecoder(resp.Body).Decode(&triviaMap)
	// if err != nil {
	// 	fmt.Println("Could not decode the json file", err.Error())
	// }

	btnEmoji := []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣"}
	components := []discordgo.MessageComponent{}
	for index, element := range allAnswers {
		btn := discordgo.Button{
			Emoji: discordgo.ComponentEmoji{
				Name: btnEmoji[index],
			},
			Label:    element,
			Style:    discordgo.SecondaryButton,
			CustomID: "buttonIndex_" + fmt.Sprintf("%d", index),
		}
		components = append(components, btn)
	}
	triviaMessage := &discordgo.MessageSend{
		Content: question,
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: components,
			},
		},
	}

	// triviaMessage.Components[0].(*discordgo.ActionsRow).Components

	_, err := s.ChannelMessageSendComplex(m.ChannelID, triviaMessage)
	if err != nil {
		panic(err.Error())
	}
}
