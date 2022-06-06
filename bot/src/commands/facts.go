package commands

import (
	util "discord-go-bot/bot/src/utils"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

type FactsStruct struct {
	Factcount string   `json:"factcount"`
	Facts     []string `json:"facts"`
}

func FactsTimer(s *discordgo.Session, m *discordgo.MessageCreate) {

	// send fact message repeatedly at regular interval. Here I chose every 12 hours.

	tick := time.Tick(12 * time.Hour)
	for range tick {
		Facts(s, m)
	}

	// // if I want to stop the ticket:
	// go func() {
	// 	for range ticker.C {
	// 		fmt.Println("Tick")
	// 	}
	// }()
	// time.Sleep(1600 * time.Hour)
	// ticker.Stop()

}

func Facts(s *discordgo.Session, m *discordgo.MessageCreate) {

	// read `GuildID_facts_iteration.txt` file
	itrFilePath := "./bot/src/utils/localDB/facts/" + m.GuildID + "_facts_iteration.txt"
	itrFileContentStr, err := ioutil.ReadFile(itrFilePath)
	if err != nil {
		// if the file doesn't exist, create a new one with value 0
		_ = ioutil.WriteFile(itrFilePath, []byte("0"), 0644)
		itrFileContentStr, _ = ioutil.ReadFile(itrFilePath)
	}
	iterationInt, _ := strconv.Atoi(string(itrFileContentStr))

	// initializing our factsData array
	var factsData FactsStruct

	// unmarshall
	json.Unmarshal(util.FactsData(), &factsData)

	var factsStrings []string

	for _, rec := range factsData.Facts {
		factsStrings = append(factsStrings, rec)
	}

	// printing the fact corresponding to the iteration value
	factMsg := fmt.Sprintf("**Fact #%d**: %s", iterationInt+1, factsStrings[iterationInt])
	s.ChannelMessageSend(m.ChannelID, factMsg)

	// preparing the next iteration:
	var newIteration string
	// reseting the iteration to zero if it reaches the end of the array. Else we add +1 to the iteration
	if iterationInt == (len(factsStrings) - 1) {
		newIteration = "0"
	} else {
		newIteration = strconv.Itoa(iterationInt + 1)
	}

	// update `facts_teration.txt` file with the new iteration

	err = ioutil.WriteFile(itrFilePath, []byte(newIteration), 0)
	if err != nil {
		fmt.Println("Could not update facts iteration", err)
	}

}
