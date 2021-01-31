package commands

import (
	. "github.com/d-exclaimation/lineapi/bot-impl"
	"github.com/line/line-bot-sdk-go/linebot"
)

func Prompt(bot *linebot.Client, event *linebot.Event, message string) {
	var (
		leftBtn = linebot.NewMessageAction("Install", "!virus")
		rightBtn = linebot.NewMessageAction("Cancel", "!virus")
		template = linebot.NewConfirmTemplate("Not Virus Installer", leftBtn, rightBtn)
		res = linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
	)

	Send(bot, event, res)
}

func Virus(bot *linebot.Client, event *linebot.Event, _ string) {
	Send(bot, event, linebot.NewTextMessage("Virus Installed for: " + event.Source.UserID))
}