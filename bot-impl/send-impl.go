//
//  send-impl.rust
//  bot_impl
//
//  Created by d-exclaimation on 12:42 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package bot_impl

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

func Send(bot *linebot.Client, event *linebot.Event, messages ...linebot.SendingMessage) {
	if _, err := bot.ReplyMessage(event.ReplyToken, messages...).Do(); err != nil {
		log.Print(err)
	}
}


func PushSend(bot *linebot.Client, to string, messages ...linebot.SendingMessage) {
	if _, err := bot.PushMessage(to, messages...).Do(); err != nil {
		log.Print(err)
	}
}
