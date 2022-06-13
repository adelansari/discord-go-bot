# discord-go-bot
Creating a discord bot using Go lang

![Alt](https://repobeats.axiom.co/api/embed/76f75f56b773f742fd9f9df54443be7fea06186b.svg)


## Setup
- Create a discord application and bot using [this link](https://discord.com/developers/applications)
- Create a `.env` file in the root. Follow the same format as `.env.Example`


## Starting
First install Go from [here](https://go.dev/dl/).

Install the following dependencies:
- [DiscordGo](https://github.com/bwmarrin/discordgo) : &nbsp; `go get github.com/bwmarrin/discordgo`
- [GoDotEnv](https://github.com/joho/godotenv) : &nbsp; `go get github.com/joho/godotenv`
- `go mod download all`

To find and install all required packages recursively for a project:
- `go get -u -v ./...`

To ensure that the go.mod file matches the source code in the module:
- `go mod tidy`  

You can run the bot normally using either of the following:
- `go run main.go`
- `go run .`

## Music setup
### Requirements:
- [youtube-dl.exe](http://ytdl-org.github.io/youtube-dl/download.html)
- [ffmpeg.exe & ffprobe.exe](https://www.gyan.dev/ffmpeg/builds/)
### Installation:
- Put the `.exe` files in the root directory.

## Commands:
Message Commands:
```
Help Command: 
.help     -  A list of help commands.

Core Commands:
.gopher     -   To show pages of Gopher images in an embed.
.giveaway   -   Creating a giveaway or picking a random winner.
.joke       -   Displays a random joke.
.8ball      -   Answer to all your [yes/no] questions.
/trivia     -   Shows general trivia miltiple choice questions.
.meme       -   Embeds a random meme image.
.facts      -   Iterate through a list of cool facts.
.factstimer -   Sends a fact every 12 hours.

Music Commands:
.play       -   Searches the music, joins the voice channel and plays the music.
.stop       -   Stops the music in the voice channel
.leave      -   Leaves the voice channel
.skip       -   Skips the current music in the queue
.queue      -   Displays the music queue

Misc. Commands:
.ping       -  To ping the bot!
.pong       -  To pong the bot!
.invite     -  To invite the bot to your server!
```
Interaction/slash commands:  
Slash commands are accecible by using forward slash  `/`