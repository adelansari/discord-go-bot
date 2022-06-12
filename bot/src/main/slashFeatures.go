package bot

import (
	scm "github.com/ethanent/discordgo-scm"

	"github.com/bwmarrin/discordgo" // discordgo package from the repo of bwmarrin .
)

var slashfeatures *[]*scm.Feature

func slashCommandFeatures(featurePointer *[]*scm.Feature) {

	scmSlash := &SlashFeature{}

	// creating and populating features

	appInteraction := discordgo.InteractionApplicationCommand

	slashfeatures = &[]*scm.Feature{
		{
			Type:    appInteraction,
			Handler: scmSlash.Ping,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "ping",
				Description: "To ping the bot!",
			},
		},
		{
			Type:    appInteraction,
			Handler: scmSlash.Pong,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "pong",
				Description: "To ping the bot!",
			},
		},
		{
			Type:    appInteraction,
			Handler: scmSlash.Jokes,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "joke",
				Description: "The best jokes on Discord!",
			},
		},
		{
			Type:    appInteraction,
			Handler: scmSlash.DadJoke,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "dad-joke",
				Description: "The best dad jokes on Discord!",
			},
		},
		{
			Type:    appInteraction,
			Handler: scmSlash.Meme,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "meme",
				Description: "A random meme!",
			},
		},
		{
			Type:    appInteraction,
			Handler: scmSlash.Invite,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "invite",
				Description: "Display an invite link for this bot.",
			},
		},
		{
			Type:    appInteraction,
			Handler: scmSlash.Giveaway,
			ApplicationCommand: &discordgo.ApplicationCommand{
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
		},
		{
			Type:    appInteraction,
			Handler: scmSlash.handleSay,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "say",
				Description: "You will never never force me to talk...",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "stuff",
						Description: "Stuff to say",
						Required:    true,
					},
				},
			},
		},
		{
			Type:    appInteraction,
			Handler: scmSlash.MagicBall,
			ApplicationCommand: &discordgo.ApplicationCommand{
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
		},
		{
			Type:    appInteraction,
			Handler: scmSlash.Trivia,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "trivia",
				Description: "Displays a multiple choice trivia questions from General Knowledge category.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        "number",
						Description: "Number of trivia questions?",
						Required:    true,
					},
				},
			},
		},
		{
			Type:     discordgo.InteractionMessageComponent,
			Handler:  scmSlash.Trivia,
			CustomID: "triviaIndex_0",
		},
		{
			Type:     discordgo.InteractionMessageComponent,
			Handler:  scmSlash.Trivia,
			CustomID: "triviaIndex_1",
		},
		{
			Type:     discordgo.InteractionMessageComponent,
			Handler:  scmSlash.Trivia,
			CustomID: "triviaIndex_2",
		},
		{
			Type:     discordgo.InteractionMessageComponent,
			Handler:  scmSlash.Trivia,
			CustomID: "triviaIndex_3",
		},
	}

	// featurePointer := &slashfeatures
	// fmt.Println(*featurePointer)
}
