package main

import (
	bot "discord-go-bot/bot/src/main"
)

func main() {
	// err := config.ReadConfig()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	bot.Start()

	// <-make(chan struct{})
	// return
}
