package bot

import (
	"discord-go-bot/bot/src/commands"
	slash "discord-go-bot/bot/src/slashCommands"
	music "discord-go-bot/bot/src/slashCommands/music"

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

func (scmSlash *SlashFeature) Music(s *discordgo.Session, i *discordgo.InteractionCreate) {
	music.MusicSlash(s, i)
}

// func (scmSlash *SlashFeature) TriviaComponent(s *discordgo.Session, i *discordgo.InteractionCreate) {

// 	// storing the custom ID after the user clicks on any button.
// 	btnCustomID := i.MessageComponentData().CustomID

// 	// importing data
// 	allAnswers, correctAnswer, triviaBtn := slash.TriviaSlash(s, i)

// 	// Finding element index
// 	correctAnswerIndex := util.Find(allAnswers, correctAnswer)
// 	btnCustomIDIndex := util.Find(triviaBtn, btnCustomID)

// 	var btnResp string
// 	if btnCustomIDIndex == correctAnswerIndex {
// 		btnResp = fmt.Sprintf("🎊 The correct answer was indeed %s.", correctAnswer)
// 		s.InteractionRespond(i.Interaction, util.MessageContentResponse(btnResp))
// 	} else {
// 		btnResp = fmt.Sprintf("%s in incorrect unfortunetlly. 😞", allAnswers[btnCustomIDIndex])
// 		s.InteractionRespond(i.Interaction, util.MessageContentResponse(btnResp))
// 	}

// }
