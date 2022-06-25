package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type TwitterTweet struct {
	Entities struct {
		Media []struct {
			ExpandedURL string `json:"expanded_url"`
			URL         string `json:"url"` // tweet link
		} `json:"media"`
	} `json:"entities"`
	Text string `json:"text"` // tweet content
	User struct {
		Name                 string `json:"name"`                    // user profile name
		ProfileImageURLHTTPS string `json:"profile_image_url_https"` // user profile icon
		ScreenName           string `json:"screen_name"`             // username
	} `json:"user"`
	Photos []struct {
		URL string `json:"url"`
	} `json:"photos"`
	Video struct {
		Poster   string `json:"poster"` // video thumbnail
		Variants []struct {
			Type string `json:"type"` // something like "video/mp4"
			Src  string `json:"src"`  // video link (there will be multiple resolutions)
		} `json:"variants"`
	} `json:"video"`
}

const (
	tweetJason string = "https://cdn.syndication.twimg.com/tweet?id=" // twitter json base link
)

func LinkEmbed(s *discordgo.Session, m *discordgo.MessageCreate) {

	messageSentFull := strings.Fields(m.Content)
	twitterLink := messageSentFull[0]

	cmd := exec.Command("youtube-dl", "--get-url", twitterLink)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	twitterVideoLink := string(stdout)

	s.ChannelMessageSend(m.ChannelID, twitterVideoLink)

}

func TwitterEmbed(s *discordgo.Session, m *discordgo.MessageCreate) {

	messageSentFull := strings.Fields(m.Content) // message sent by the user (stored in string array)
	twitterLink := messageSentFull[0]            // original tweet url
	twitterLinkSplit := strings.Split(twitterLink, "/")
	tweetID := twitterLinkSplit[5]

	tweetJasonLink := tweetJason + tweetID

	req, err := http.Get(tweetJasonLink)
	if err != nil {
		fmt.Print("Could not fetch tweet data", err.Error())
	}
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var twittertweet TwitterTweet
	json.Unmarshal(bodyBytes, &twittertweet)
	// fmt.Printf("%+v", twittertweet)

	// // Tweet Image
	// for _, p := range twittertweet.Photos {
	// 	if p.URL != "" {
	// 		s.ChannelMessageSend(m.ChannelID, p.URL)
	// 	}
	// }

	// Tweet Video
	if twittertweet.Video.Poster != "" {
		var TweetVideoUrl string
		for _, rec := range twittertweet.Video.Variants {
			TweetVideoUrl = rec.Src
		}
		// it will get the last video url from above iteration

		// s.ChannelMessageSend(m.ChannelID, twittertweet.Video.Poster)
		s.ChannelMessageSend(m.ChannelID, TweetVideoUrl)
	}

}
