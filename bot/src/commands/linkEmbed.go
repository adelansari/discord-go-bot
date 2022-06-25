package commands

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
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
