package commands

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

type Executable func(bot *linebot.Client, event *linebot.Event, message string)

func All() map[string]Executable {
	var dict = map[string]Executable{
		"!8ball":    Fortune,
		"!meme":     Memes,
		"!prompt":   Prompt,
		"!virus":    Virus,
		"!addtodo":  AddTodo,
		"!showtodo": ShowTodo,
	}
	return dict
}
