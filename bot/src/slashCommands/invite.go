package slashCommand

import (
	util "discord-go-bot/bot/src/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func InviteSlash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	inviteLink := fmt.Sprintf("https://discord.com/oauth2/authorize?client_id=977285728042111027&permissions=1644971949425&scope=bot%%20applications.commands")
	inviteMessage := fmt.Sprintf("Please visit %s to add this bot to your server.", inviteLink)

	s.InteractionRespond(i.Interaction, util.MessageContentResponse(inviteMessage))
}
