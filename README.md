# discord-go-bot
Creating a discord bot using Go lang


### Setup
- Create a discord application and bot using [this link](https://discord.com/developers/applications)
- Create a `config.json` file in the root. It should include:
```
{
"Token"  : "PlaceYourDiscordBotTokenHere",
"BotPrefix" : "!"
}
```


### Starting
First install Go from [here](https://go.dev/dl/).  
For the first time:
- `go build main.go `

You can run the bot normally using either of the following:
- `go run main.go`
- `go run .`

### Commands
```
!help    -  A list of help commands.
!ping    -  To ping the bot!
!pong    -  To pong the bot!
!gopher  -  To show pages of Gopher images in an embed

```