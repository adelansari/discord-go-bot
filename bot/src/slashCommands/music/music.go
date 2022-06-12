package music

import (
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"

	"github.com/bwmarrin/discordgo"
)

//	ytVideoUrl is a constant containing endpoint for a youtube video
const ytVideoUrl string = "https://www.youtube.com/watch?v="

func InitializeRoutine() {
	SongSignal = make(chan PkgSong)
	go globalPlay(SongSignal)
}

func joinVoice(v *VoiceInstance, s *discordgo.Session, m *discordgo.MessageCreate) *VoiceInstance {
	voiceChannelID := SearchVoiceChannel(m.Author.ID)
	if voiceChannelID == "" {
		log.Println("Voice channel id not found")
		SendChannelMessage(m.ChannelID, "**[Music]** You need to join a voice channel first!")
		return nil
	}

	if v != nil {
		log.Println("INFO: Voice instance already created")
	} else {
		guildID := SearchGuild(m.ChannelID)
		mutex.Lock()
		v = new(VoiceInstance)
		VoiceInstances[guildID] = v
		v.GuildID = guildID
		v.session = s
		mutex.Unlock()
	}

	var err error
	v.voice, err = s.ChannelVoiceJoin(v.GuildID, voiceChannelID, false, true)
	if err != nil {
		sentry.CaptureException(err)
		log.Println("ERR: Error when joining voice channel")
		v.Stop()
		return nil
	}

	err = v.voice.Speaking(false)
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}

	log.Println("INFO: New voice instance created")
	return v
}

func LeaveVoice(v *VoiceInstance, m *discordgo.MessageCreate) {
	if v == nil {
		log.Println("INFO: The bot is not in a voice channel!")
		return
	}

	v.Stop()
	time.Sleep(150 * time.Millisecond)

	err := v.voice.Disconnect()
	if err != nil {
		sentry.CaptureException(err)
		return
	}

	log.Println("INFO: Voice channel left")
	mutex.Lock()
	delete(VoiceInstances, v.GuildID)
	mutex.Unlock()
	SendChannelMessage(m.ChannelID, "**[Music]** Voice channel left!")
}

func PlayMusic(n []string, v *VoiceInstance, s *discordgo.Session, m *discordgo.MessageCreate) {
	if v == nil {
		log.Println("INFO: The bot is not in a voice channel, joining now")
		v = joinVoice(v, s, m)

		if v == nil {
			log.Println("ERR: Failed to join voice channel")
			return
		}
	}

	voiceChannelID := SearchVoiceChannel(m.Author.ID)
	if v.voice.ChannelID != voiceChannelID {
		SendChannelMessage(m.ChannelID, "**[Music]** <@"+m.Author.ID+"> you need to join my voice channel first!")
		return
	}

	var videoId string
	var videoTitle string
	var err error

	// If a youtube url is sent as argument
	if strings.Contains(n[0], ytVideoUrl) {
		urlParser, err := url.Parse(n[0])
		if err != nil {
			sentry.CaptureException(err)
			log.Println("INFO: Unable to parse youtube url")
			SendChannelMessage(m.ChannelID, "**[Music]** Oops, something wrong happened when parsing youtube url")
			return
		}

		query := urlParser.Query()
		videoId = query.Get("v")

		videoTitle, err = ytFind(videoId)
		if err != nil {
			sentry.CaptureException(err)
			log.Println("INFO: unable to find title of song on youtube")
			SendChannelMessage(m.ChannelID, "**[Music]** Oops, something went wrong when fetching title on YouTube")
			return
		}

		// If argument(s) is not a youtube url
	} else {
		name := strings.Join(n, "_")
		videoId, videoTitle, err = ytSearch(name)
		if err != nil {
			sentry.CaptureException(err)
			log.Println("INFO: Unable to find song by searching youtube")
			SendChannelMessage(m.ChannelID, "**[Music]** Can't find a song with that name")
			return
		}
	}

	song, err := execYtdl(videoId, videoTitle, v, m)
	if err != nil || song.data.ID == "" {
		if err != nil {
			sentry.CaptureException(err)
		}

		log.Println("INFO: Youtube search: ", err)
		SendChannelMessage(m.ChannelID, "**[Music]** Unable to find song")
		return
	}

	SendChannelMessage(m.ChannelID, "**[Music]** "+song.data.User+" has added **"+song.data.Title+"** to the queue")

	go func() {
		SongSignal <- song
	}()
}

func SkipMusic(v *VoiceInstance, m *discordgo.MessageCreate) {
	log.Println("INFO:", m.Author.Username, "send 'skip'")
	if v == nil {
		log.Println("INFO: The bot is not in a voice channel")
		SendChannelMessage(m.ChannelID, "**[Music]** I need to join a voice channel first!")
		return
	}

	if len(v.queue) == 0 {
		log.Println("INFO: The queue is empty.")
		SendChannelMessage(m.ChannelID, "**[Music]** There is no song playing")
		return
	}

	if v.Skip() {
		SendChannelMessage(m.ChannelID, "**[Music]** I'm paused, resume first")
	}
}

func StopMusic(v *VoiceInstance, m *discordgo.MessageCreate) {
	log.Println("INFO:", m.Author.Username, "requested stopping of music")

	if v == nil {
		log.Println("INFO: The bot is not in a voice channel")
		SendChannelMessage(m.ChannelID, "**[Music]** I need to join a voice channel first!")
		return
	}
	voiceChannelID := SearchVoiceChannel(m.Author.ID)
	if v.voice.ChannelID != voiceChannelID {
		SendChannelMessage(m.ChannelID, "**[Music]** <@"+m.Author.ID+"> You need to join my voice channel to stop music!")
		return
	}

	v.Stop()
	log.Println("INFO: The bot stopped playing music")
	SendChannelMessage(m.ChannelID, "**[Music]** I have now stopped playing music!")
}
