//
//  main.go
//  main
//
//  Created by d-exclaimation on 5:08 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	. "github.com/d-exclaimation/lineapi/bot-impl"
	"github.com/d-exclaimation/lineapi/commands"
	arcade "github.com/d-exclaimation/lineapi/games"

	"github.com/line/line-bot-sdk-go/linebot"
)

var prefix = "!"

type HubMessage struct {
	Author string
	Msg    string
	Key    string
}

func main() {

	// Get bot with all the env variables
	var bot, err = linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	var games = make(map[string]*arcade.Sokoban)

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
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

					// Get id of the room, group, or user
					var id string
					switch event.Source.Type {
					case "user":
						id = event.Source.UserID
					case "group":
						id = event.Source.GroupID
					case "room":
						id = event.Source.RoomID
					}

					// Let the action function handle the rest
					actions(bot, event, message.Text, games, id)
				}
			}
		}
	})

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

// Command function
func actions(bot *linebot.Client, event *linebot.Event, message string, games map[string]*arcade.Sokoban, id string) {
	fmt.Println(message) // Logging for debugging

	_, ok := games[id] // Check whether the given id has a game in the map

	// If the message sent is a movement {w, a, s, d} then try to play the game, if game exist
	if isWASD(message) && ok {
		arcade.MoveGame(bot, event, message, games, id)
		return
	}

	// A non command is ignored (start with a prefix)
	if !strings.HasPrefix(message, prefix) {
		return
	}

	// Special cases
	if strings.HasPrefix(message, "!start") {
		arcade.NewGame(bot, event, games, id)
		return
	}

	// Get the command itself without any parameter
	var name = strings.Split(message, " ")[0]

	// Get the appropriate functions hashmaps from command directory
	var commandMap = commands.All()
	_, ok = commandMap[name] // Check if found
	if ok {
		// Execute the function given the required parameters
		commandMap[name](bot, event, message)
	}
}

func isWASD(msg string) bool {
	var lower = strings.ToLower(msg)
	return lower == "w" || lower == "a" || lower == "s" || lower == "d"
}
