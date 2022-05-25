package bot

import (
	// importing commands pachage

	//	to print errors

	scm "github.com/ethanent/discordgo-scm"

	"github.com/bwmarrin/discordgo" // discordgo package from the repo of bwmarrin .
)

var slashfeatures *[]*scm.Feature

func slashCommandFeatures(featurePointer *[]*scm.Feature) {

	scmSlash := &SlashFeature{}

	// creating and populating features

	slashfeatures = &[]*scm.Feature{
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: scmSlash.Ping,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "ping",
				Description: "To ping the bot!",
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: scmSlash.Pong,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "pong",
				Description: "To ping the bot!",
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
			Handler: scmSlash.Jokes,
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "joke",
				Description: "The best dad jokes on Discord!",
			},
		},
		{
			Type:    discordgo.InteractionApplicationCommand,
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
			Type:    discordgo.InteractionApplicationCommand,
			Handler: scmSlash.handleSay,
			ApplicationCommand: &discordgo.ApplicationCommand{
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
		},
	}

	// featurePointer := &slashfeatures
	// fmt.Println(*featurePointer)
}
