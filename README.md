# golang-linebot
 A bot for LINE Messanging API

## Overview
This is a piece of a project I am building, the project is related to bots and servers as services.

This is the code to create a bot using LINE Messaging API SDK for Golang,

## LINE API
LINE API only require a server that can take POST Request and use the sdk to do something depending on the data given. Other than that is more to do with the developer console.

I'm using the regular http library for Go that can take POST Request and parse it as a LINE Message Events

```go
// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/<LINE-API-Endpoints>", func(w http.ResponseWriter, req *http.Request) {
		var events, err = bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			// Only check for EventTypeMessage
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
                    // Do Something
				}
			}
		}
	})
```

## Bot functionalities

### Commands:
> 1. Fortune telling with a random word generator lol `!8ball`
> 2. Making simple group todo list but without any memory `!addtodo <name> && !showtodo`
> 3. Play Sokoban `!start => w, a, s, d`
> 4. Make jokes hehehe `!meme, !prompt`
> 5. Upload random images as memes `!meme`
> 6. Make a fake virus downloader to prank people `!prompt && !virus`
> 7. Send message across social media (more below) `!announce <message>`

```go
func All() map[string]Executable {
	var dict = map[string]Executable{
		"!8ball": Fortune,
		"!meme":  Memes,
		"!prompt": Prompt,
		"!virus": Virus,
		"!addtodo": AddTodo,
		"!showtodo": ShowTodo,
	}
	return dict
}
```

### Connecting to other bots

`No info here`


### LINE Flex Message UI 
LINE Provides a way to create simple application using the Flexbox system from CSS but with their APIs. There are other ways of doing this task but Flex Messages is the most customizable one.

I added some code that are UI tools (similar to something in SwiftUI or React) on top of the normal Flex Message that is provided to reduce the complexity and make them a reuseable components

```go
func MiniBubbleCard(hexColor string, head []FlexComponent, body []FlexComponent) *BubbleContainer {
	return &BubbleContainer{
		Type:      FlexContainerTypeBubble,
		Size:      FlexBubbleSizeTypeNano,
		Header:    &BoxComponent{
			Type:            FlexComponentTypeBox,
			Layout:          FlexBoxLayoutTypeVertical,
			Contents:        head,
			BackgroundColor: hexColor,
		},
		Body:      &BoxComponent{
			Type:            FlexComponentTypeBox,
			Layout:          FlexBoxLayoutTypeVertical,
			Contents:        body,
			Spacing:         FlexComponentSpacingTypeMd,
		},
	}
}

```
