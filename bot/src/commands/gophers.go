package commands

func Gophers() ([]string, []string) {
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
