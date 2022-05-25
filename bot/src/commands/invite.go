package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Invite(s *discordgo.Session, m *discordgo.MessageCreate) {

	inviteLink := fmt.Sprintf("https://discord.com/oauth2/authorize?client_id=977285728042111027&permissions=1644971949425&scope=bot%%20applications.commands")
	inviteMessage := fmt.Sprintf("Please visit %s to add this bot to your server.", inviteLink)

	s.ChannelMessageSend(m.ChannelID, inviteMessage)

}
