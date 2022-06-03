package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

const apiLink = "https://api.api-ninjas.com/v1/"
const dataLimit = "?limit="
const memeApiLink = "https://meme-api.herokuapp.com/gimme"
const triviaApiLink = "https://opentdb.com/api.php?amount=1&category=9&type=multiple"

// JSON pretty print by marshaling value
func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

// Find element in slice/array with linear search
func Find(a []string, x string) int {
	// Return the smallest index i at which x == a[i]
	for i, n := range a {
		if x == n {
			return i
		}
	}
	// Return len(a) if there is no such index
	return len(a)
}

// Shuffle a string array elements
func Shuffle(vals []string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
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

func JokeApiData(category string, limit string) []byte {
	ninjaToken := os.Getenv("APININJAKEY")
	url := apiLink + category + dataLimit + limit

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("X-Api-Key", ninjaToken)

	return getFromUrl(req)
}

func MemeApiData() []byte {
	req, err := http.NewRequest("GET", memeApiLink, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	return getFromUrl(req)
}

func TriviaApiData() []byte {
	req, err := http.NewRequest("GET", triviaApiLink, nil)
	if err != nil {
		fmt.Print("Could not fetch trivia api", err.Error())
	}
	return getFromUrl(req)
}

func getFromUrl(req *http.Request) []byte {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
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
