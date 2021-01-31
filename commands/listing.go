package commands

import (
	. "github.com/d-exclaimation/lineapi/bot-impl"
	"github.com/d-exclaimation/lineapi/components"
	"github.com/line/line-bot-sdk-go/linebot"
	"strings"
	"time"
)

var (
	todos = make([]*linebot.BubbleContainer, 0)
)

func AddTodo(bot *linebot.Client, event *linebot.Event, message string) {
	var params = strings.Split(message, " ")

	// If there are no other arguments, exit early and send an error message
	if len(params) < 2 {
		var errorMsg = linebot.NewTextMessage("Sorry, but there isn't enough arguments")
		Send(bot, event, errorMsg)
		return
	}

	// Otherwise for each extra arguments add it to the todolist
	for i := 1; i < len(params); i++ {
		todos = append(todos, components.TaskItem(params[i], time.Now().Format("2016-04-02")))
	}
	ShowTodo(bot, event, message)
}

func ShowTodo(bot *linebot.Client, event *linebot.Event, message string) {
	// Make the container and the message
	var (
		container = &linebot.CarouselContainer{
			Type:     linebot.FlexContainerTypeCarousel,
			Contents: todos,
		}
		res = linebot.NewFlexMessage("Update your app lol", container)
	)

	Send(bot, event, res)
}
