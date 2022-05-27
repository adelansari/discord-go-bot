package commands

import (
	paginator "discord-go-bot/bot/src/utils/paginator"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Gophers(s *discordgo.Session, m *discordgo.MessageCreate) {
	p := paginator.NewPaginator(s, m.ChannelID)

	// Add embed pages to paginator

	// returning values of gophersImages and gophersUrl from images/gophers.go
	gopherImages, gophersUrl := GopherList()

	for i := 0; i < len(gopherImages); i++ {
		gopherName := fmt.Sprintf("Image %d: %s", i+1, gopherImages[i])
		p.Add(&discordgo.MessageEmbed{
			Title:       "Collection of Gopher Images",
			Description: gopherName,
			Image:       &discordgo.MessageEmbedImage{URL: gophersUrl[i]},
		})
	}

	// Sets the footers of all added pages to their page numbers.
	p.SetPageFooters()

	// When the paginator is done listening set the colour to yellow
	p.ColourWhenDone = 0xffff

	// Stop listening for reaction events after five minutes
	p.Widget.Timeout = time.Minute * 5

	// Add a custom handler for the gun reaction.
	p.Widget.Handle("🔫", func(w *paginator.Widget, r *discordgo.MessageReaction) {
		s.ChannelMessageSend(m.ChannelID, "Bang!")
	})

	p.Spawn()
}

func GopherList() ([]string, []string) {
	gophersImages := []string{
		"5th-element.png",
		"arrow-gopher.png",
		"back-to-the-future-v2.png",
		"baywatch.png",
		"big-bang-theory.png",
		"bike-gopher.png",
		"blues-gophers.png",
		"buffy-the-gopher-slayer.png",
		"chandleur-gopher.png",
		"cherry-gopher.png",
		"devnation-france-gopher.png",
		"dr-who.png",
		"fire-gopher.png",
		"firefly-gopher.png",
		"fort-boyard.png",
		"friends.png",
		"gandalf-colored.png",
		"gandalf.png",
		"gladiator-gopher.png",
		"gopher-dead.png",
		"gopher-johnny.jpg",
		"gopher-open.png",
		"gopher-speaker.png",
		"gopher.png",
		"graffiti-devfest-nantes-2021.png",
		"halloween-spider.png",
		"happy-gopher.png",
		"harry-gopher.png",
		"idea-gopher.png",
		"indiana-jones.png",
		"jedi-gopher.png",
		"jurassic-park.png",
		"love-gopher.png",
		"luigi-gopher.png",
		"mac-gopher.png",
		"marshal.png",
		"men-in-black-v2.png",
		"mojito-gopher.png",
		"paris-gopher.png",
		"sandcastle-gopher.png",
		"saved-by-the-bell.png",
		"star-wars.png",
		"stargate.png",
		"tadx-gopher.png",
		"urgences.png",
		"vampire-xmas.png",
		"wired-gopher.png",
		"x-files.png",
		"yoda-gopher.png",
	}
	imgUrl := "https://raw.githubusercontent.com/scraly/gophers/main/"

	gophersUrl := []string{}

	// for i := 0; i < len(gophersImages); i++ {
	// 	gophersUrl = append(gophersUrl, imgUrl+gophersImages[i])
	// }

	for index := range gophersImages {
		gophersUrl = append(gophersUrl, imgUrl+gophersImages[index])
	}

	// fmt.Println(gophersUrl)

	return gophersImages, gophersUrl

}
