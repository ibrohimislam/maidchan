package main;

import (
    "time"
    "github.com/tucnak/telebot"
)

func main() {
    bot, err := telebot.NewBot("124223052:AAGr9lxL0Ewi50jBzEDSKy36d_Rxh-g5pUg")
    if err != nil {
        return
    }

    messages := make(chan telebot.Message)
    bot.Listen(messages, 1*time.Second)

    for message := range messages {
        if message.Text == "/hi" {
            bot.SendMessage(message.Chat,
                "Hello, "+message.Sender.FirstName+"!", nil)
        }
    }
}
