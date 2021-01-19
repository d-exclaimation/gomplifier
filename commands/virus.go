package commands

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

func Prompt(bot *linebot.Client, event *linebot.Event, message string) {
	leftBtn := linebot.NewMessageAction("Install", "!virus")
	rightBtn := linebot.NewMessageAction("Cancel", "!virus")

	template := linebot.NewConfirmTemplate("Not Virus Installer", leftBtn, rightBtn)

	res := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)

	if _, err := bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
		log.Print(err)
	}
}

func Virus(bot *linebot.Client, event *linebot.Event, message string) {
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Virus Installed for: " + event.Source.UserID)).Do(); err != nil {
		log.Print(err)
	}
}