package telegram

import (
    "log"
    "os"

    "github.com/temoon/go-telegram-bots-api/requests"
)

func Example() {
    opts := new(BotOpts)
    opts.Token = os.Getenv("TOKEN")

    bot := NewBot(opts)

    if me, err := bot.GetMe(); err != nil {
        log.Fatalln("Invalid token:", err)
    } else {
        log.Println("Hello! My name is", me.FirstName)
    }

    log.Println("Listening to messages...")

    request := requests.GetUpdates{
        Timeout: 60,

        AllowedUpdates: []string{
            UpdatesAllowedMessage,
        },
    }

    for {
        updates, err := bot.GetUpdates(&request)

        if err != nil {
            log.Fatalln("Unable to get updates:", err)
        }

        for _, update := range updates {
            log.Println(update.Message.From.FirstName, "::", update.Message.Text)

            _, err := bot.SendMessage(&requests.SendMessage{
                ChatID:           update.Message.Chat.ID,
                ReplyToMessageID: update.Message.MessageID,
                Text:             "Got it!",
            })

            if err != nil {
                log.Fatalln("Unable to send message:", err)
            }

            request.Offset = update.UpdateID + 1
        }
    }
}
