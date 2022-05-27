# discord-go-bot
Creating a discord bot using Go lang

![Alt](https://repobeats.axiom.co/api/embed/76f75f56b773f742fd9f9df54443be7fea06186b.svg)


## Setup
- Create a discord application and bot using [this link](https://discord.com/developers/applications)
- Create a `.env` file in the root. It should include:
```
TOKEN=AddYourDiscordBotTokenHere
APININJAKEY=AddYourAPITokenHere
```


## Starting
First install Go from [here](https://go.dev/dl/).

Install the following dependencies:
- [DiscordGo](https://github.com/bwmarrin/discordgo) : &nbsp; `go get github.com/bwmarrin/discordgo`
- [GoDotEnv](https://github.com/joho/godotenv) : &nbsp; `go get github.com/joho/godotenv`

To find and install all required packages recursively for a project:
- `go get -u -v ./...`

To ensure that the go.mod file matches the source code in the module:
- `go mod tidy`  



You can run the bot normally using either of the following:
- `go run main.go`
- `go run .`

## Commands:
Message Commands:
```
Help Command: 
.help     -  A list of help commands.

Core Commands:
.gopher   -  To show pages of Gopher images in an embed.
.giveaway -  Creating a giveaway or picking a random winner.
.joke	  -  Displays a random joke.
.8ball    -  Answer to all your [yes/no] questions.

Misc. Commands:
.ping     -  To ping the bot!
.pong     -  To pong the bot!
.invite   -  To invite the bot to your server!
```
Interaction/slash commands:  
Slash commands are accecible by using forward slash  `/`