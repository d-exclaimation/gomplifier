package commands

import (
	"math/rand"
	"strings"

	. "github.com/d-exclaimation/lineapi/bot-impl"
	"github.com/line/line-bot-sdk-go/linebot"
)

func Memes(bot *linebot.Client, event *linebot.Event, message string) {
	var params = strings.Split(message, " ") // Get all the commands as slices

	// Image dictionaries
	var images = map[string]string{}

	// Get the link to an image
	var res, thumb string

	// If image is given by parameter, just use that. Otherwise, random from map
	if len(params) > 1 {
		res = params[1]
		thumb = res
	} else {
		var full = mapToSlice(images)
		var i = rand.Intn(len(full))
		res = full[i]
		thumb = images[res]
	}

	Send(bot, event, linebot.NewImageMessage(res, thumb))
}

func mapToSlice(dict map[string]string) []string {
	var res = make([]string, 0)
	for key := range dict {
		res = append(res, key)
	}
	return res
}
