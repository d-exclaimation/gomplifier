package commands

import (
	. "github.com/d-exclaimation/lineapi/bot-impl"
	"github.com/line/line-bot-sdk-go/linebot"
	"math/rand"
)

func Fortune(bot *linebot.Client, event *linebot.Event, _ string) {
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

	Send(bot, event, linebot.NewTextMessage(reply))
}
