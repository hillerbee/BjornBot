package main

import (
	"github.com/BjornTwitchBot/BjornBot/Godeps/_workspace/src/github.com/fabioxgn/go-bot"
)

func donate(command *bot.Cmd) (msg string, err error) {
	msg = "https://www.extra-life.org/index.cfm?fuseaction=donorDrive.team&teamID=48967"
	return
}

func init() {
	bot.RegisterCommand(
		"donate",
		"Provides donation information to the channel",
		"",
		donate)
}
