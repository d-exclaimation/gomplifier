package commands

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"math/rand"
)

func Fortune(bot *linebot.Client, event *linebot.Event, message string) {
	var yesResponds = []string{"the oracle said yes", "God himself told me yes", "yes", "it's probably yes", "yes"}
	var noResponds = []string{"the oracle said no", "God said no", "No", "Probably no", "nope"}

	var respond = rand.Intn(5)
	var yesNo = rand.Intn(2)

	var reply string

	if yesNo == 0 {
		reply = yesResponds[respond]
	} else {
		reply = noResponds[respond]
	}

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do(); err != nil {
		log.Print(err)
	}
}
