//
//  controller.rust
//  games
//
//  Created by d-exclaimation on 1:14 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package games

import (
	. "github.com/d-exclaimation/lineapi/bot-impl"
	"github.com/d-exclaimation/lineapi/components"
	"github.com/line/line-bot-sdk-go/linebot"
	"strings"
)

func NewGame(bot *linebot.Client, event *linebot.Event, games map[string]*Sokoban, id string) {
	// Create a new instance of the game
	games[id] = NewSokoban(10, 5)
	var flex = linebot.NewFlexMessage(games[id].Show(), components.GameView(games[id].ShowParts()))
	Send(bot, event, flex)
}

func MoveGame(bot *linebot.Client, event *linebot.Event, message string, games map[string]*Sokoban, id string) {
	// Continue with the existing instance of the game
	games[id].Move(strings.ToLower(message))
	var flex = linebot.NewFlexMessage(games[id].Show(), components.GameView(games[id].ShowParts()))
	Send(bot, event, flex)
}
