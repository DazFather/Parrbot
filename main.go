package main

import (
	"parrbot/message"
	"parrbot/robot"

	"github.com/NicoNex/echotron/v3"
)

func main() {
	var commandList = []robot.Command{
		{Name: "Start", Trigger: "/start", ReplyAt: message.MESSAGE, Scope: helloHandler},
		{Name: "Credits", Trigger: "/info", ReplyAt: message.CALLBACK_QUERY, Scope: infoHandler},
	}

	robot.Start(commandList)
}

var helloHandler robot.CommandFunc = func(bot *robot.Bot, update *message.Update) message.Any {
	var kbd = [][]echotron.InlineKeyboardButton{{
		{Text: "ℹ️ more info", CallbackData: "/info"},
	}}

	var msg = message.Text{"🦜 Hello World!", nil}
	return *msg.ClipInlineKeyboard(kbd)
}

var infoHandler robot.CommandFunc = func(bot *robot.Bot, update *message.Update) message.Any {
	update.CallbackQuery.Message.EditText("Made with ❤️ by @DazFather", nil)
	return nil
}
