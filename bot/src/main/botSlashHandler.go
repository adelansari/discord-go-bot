package bot

import (
	"discord-go-bot/bot/src/commands"
	slash "discord-go-bot/bot/src/slashCommands"

	"github.com/bwmarrin/discordgo"
)

func (scmSlash *SlashFeature) Ping(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slash.PingSlash(s, i)
}

func (scmSlash *SlashFeature) Pong(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slash.PongSlash(s, i)
}

func (scmSlash *SlashFeature) DadJoke(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slash.DadJokesSlash(s, i)
}

func (scmSlash *SlashFeature) Jokes(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slash.JokesSlash(s, i)
}

func (scmSlash *SlashFeature) Meme(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commands.MemeSlash(s, i)
}

func (scmSlash *SlashFeature) Giveaway(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slash.GiveawaySlash(s, i)
}

func (scmSlash *SlashFeature) handleSay(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slash.SaySlash(s, i)
}

func (scmSlash *SlashFeature) Invite(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slash.InviteSlash(s, i)
}

func (scmSlash *SlashFeature) MagicBall(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slash.MagicBallSlash(s, i)
}

func (scmSlash *SlashFeature) Trivia(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slash.TriviaSlash(s, i)
}

// func (scmSlash *SlashFeature) Music(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	music.MusicSlash(s, i)
// }
