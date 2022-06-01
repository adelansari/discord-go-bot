package bot

import (
	"github.com/bwmarrin/discordgo"
)

var (
	slashCmd = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "To ping the bot!",
		},
		{
			Name:        "pong",
			Description: "To ping the bot!",
		},
		{
			Name:        "joke",
			Description: "The best jokes on Discord!",
		},
		{
			Name:        "dad-joke",
			Description: "The best dad jokes on Discord!",
		},
		{
			Name:        "meme",
			Description: "A random meme!",
		},
		{
			Name:        "invite",
			Description: "Display an invite link for this bot.",
		},
		{
			Name:        "giveaway",
			Description: "giveaway embed",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "create",
					Description: "Please enter the giveaway message here:",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "pick",
					Description: "Please enter the message ID corresponding to the giveaway message.",
					Required:    false,
				},
			},
		},
		{
			Name:        "say",
			Description: "You will never never force me to talk...",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "stuff",
					Description: "Stuff to say",
					Required:    false,
				},
			},
		},
		{
			Name:        "8-ball",
			Description: "I can answer all your [yes/no] questions!...",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "question",
					Description: "The life-changing question to ask",
					Required:    true,
				},
			},
		},
		{
			Name:        "trivia",
			Description: "Displays a multiple choice trivia questions from General Knowledge category.",
		},
	}
)
