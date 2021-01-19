package commands

import (
	"log"
	"math/rand"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func Memes(bot *linebot.Client, event *linebot.Event, message string) {
	var params = strings.Split(message, " ") // Get all the commands as slices

	// Image dictionaries
	var images = map[string]string{
		"<image-link-1>": "<thumbnail-link-1>",
		"<image-link-2>": "<thumbnail-link-2>",
		"<image-link-3>": "<thumbnail-link-3>",
	}

	// Get the link to an image
	var res, thumb string

	// If image is given by parameter, just use that. Otherwise, random from map
	if len(params) > 1 {
		res = params[1]
		thumb = res
	} else {
		var full = make([]string, 0)
		for key := range images {
			full = append(full, key)
		}
		var i = rand.Intn(len(full))
		res = full[i]
		thumb = images[res]
	}

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(res, thumb)).Do(); err != nil {
		log.Print(err)
	}
}
