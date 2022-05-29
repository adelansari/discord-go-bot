package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

const apiLink = "https://api.api-ninjas.com/v1/"
const dataLimit = "?limit="

// JSON pretty print by marshaling value
func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

// Responding to interaction with a message
func MessageContentResponse(c string) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: c,
		},
	}
}

func ApiData(category string, limit string) []byte {
	ninjaToken := os.Getenv("APININJAKEY")
	url := apiLink + category + dataLimit + limit

	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("X-Api-Key", ninjaToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	return bodyBytes
}
