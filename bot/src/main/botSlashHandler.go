package bot

import (
	slashCommand "discord-go-bot/bot/src/slashCommands"

	"github.com/bwmarrin/discordgo"
)

var (
	componentHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"triviaIndex_0": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			slashCommand.TriviaSlash(s, i)
		},
		"triviaIndex_1": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			slashCommand.TriviaSlash(s, i)
		},
		"triviaIndex_2": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			slashCommand.TriviaSlash(s, i)
		},
		"triviaIndex_3": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			slashCommand.TriviaSlash(s, i)
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			slashCommand.PingSlash(s, i)
		},
		"pong": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			slashCommand.PongSlash(s, i)
		},
		"joke": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			slashCommand.JokesSlash(s, i)
		},
		"dad-joke": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			slashCommand.DadJokesSlash(s, i)
		},
		"meme": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			slashCommand.MemeSlash(s, i)
		},
		"invite": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			slashCommand.InviteSlash(s, i)
		},
		"giveaway": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			slashCommand.GiveawaySlash(s, i)
		},
		"say": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			slashCommand.SaySlash(s, i)
		},
		"8-ball": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			slashCommand.MagicBallSlash(s, i)
		},
		"trivia": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			slashCommand.TriviaSlash(s, i)
		},
	}
)
