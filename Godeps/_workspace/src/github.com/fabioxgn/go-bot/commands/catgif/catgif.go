package catgif

import (
	"github.com/BjornTwitchBot/BjornBot/Godeps/_workspace/src/github.com/fabioxgn/go-bot"
	"net/http"
)

func gif(command *bot.Cmd) (msg string, err error) {
	res, err := http.Get("http://thecatapi.com/api/images/get?format=src&type=gif")
	if err != nil {
		return "", err
	}
	return res.Request.URL.String(), nil
}

func init() {
	bot.RegisterCommand(
		"catgif",
		"Returns a random cat gif.",
		"",
		gif)
}
