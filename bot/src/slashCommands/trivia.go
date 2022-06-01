package slashCommand

import (
	util "discord-go-bot/bot/src/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
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

func TriviaSlash(s *discordgo.Session, i *discordgo.InteractionCreate) {

	question, _, allAnswers := TriviaAPI()

	btnEmoji := []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣"}
	components := []discordgo.MessageComponent{}
	for index, element := range allAnswers {
		btn := discordgo.Button{
			Emoji: discordgo.ComponentEmoji{
				Name: btnEmoji[index],
			},
			Label:    element,
			Style:    discordgo.SecondaryButton,
			CustomID: "triviaIndex_" + fmt.Sprintf("%d", index),
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

	err := s.InteractionRespond(i.Interaction, triviaMessage)
	if err != nil {
		fmt.Println("Could not send the trivia question")
	}
}

func TriviaAPI() (string, int, []string) {
	resp, err := http.Get("https://opentdb.com/api.php?amount=1&category=9&type=multiple")
	if err != nil {
		fmt.Println("Could not fetch trivia api", err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	var data triviaApi

	// unmarshall
	json.Unmarshal(body, &data)

	// General Knowledge category
	// Loop through the Results node for the Question
	var question string
	var correctAnswer string
	var allAnswers []string
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

	// Finding the correctAnswer index in allAnswers string array
	correctAnswerIndex := util.Find(allAnswers, correctAnswer)

	return question, correctAnswerIndex, allAnswers
}

func TriviaAnswer(selectedAnsIndex int) string {
	_, correctAnswerIndex, allAnswers := TriviaAPI()
	selectedAnswer := allAnswers[selectedAnsIndex]
	var btnResp string
	if selectedAnsIndex == correctAnswerIndex {
		btnResp = fmt.Sprintf("🎊 The correct answer was indeed %s.", selectedAnswer)
	} else {
		btnResp = fmt.Sprintf("%s in incorrect unfortunetlly. 😞", selectedAnswer)
	}

	return btnResp
}
