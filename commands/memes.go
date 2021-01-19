package commands

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"math/rand"
	"strings"
)

func Memes(bot *linebot.Client, event *linebot.Event, message string) {
	var params = strings.Split(message, " ") // Get all the commands as slices

	// Image dictionaries
	var images = map[string]string{
		"https://i.imgur.com/lfUAvZw.png": "https://i.imgur.com/KgfNw0E.png",
		"https://i.imgur.com/euvFB6k.jpg": "https://i.imgur.com/euvFB6k.jpg",
		"https://i.imgur.com/f6pPlXJ.jpg": "https://i.imgur.com/f6pPlXJ.jpg",
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
