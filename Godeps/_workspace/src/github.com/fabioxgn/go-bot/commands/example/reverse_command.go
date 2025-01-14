package example

import (
	"github.com/BjornTwitchBot/BjornBot/Godeps/_workspace/src/github.com/fabioxgn/go-bot"
)

// From stackoverflow: http://stackoverflow.com/a/10030772
func reverse(command *bot.Cmd) (msg string, err error) {
	runes := []rune(command.FullArg)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	msg = string(runes)
	return
}

func init() {
	bot.RegisterCommand(
		"reverse", "Reverses a string",
		"string to be reversed",
		reverse)
}
