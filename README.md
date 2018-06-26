# Telegram Bot API: Go implementation
Read an [official documentation](https://core.telegram.org/bots/api) for more information.

Currently work in progress.

## Example
### Simple bot
Listens to messages and replies to it immediately.
```go
package main

import (
    "os"
    "log"

    "github.com/temoon/go-telegram-bots-api"
    "github.com/temoon/go-telegram-bots-api/requests"
)

func main() {
    opts := &telegram.BotOpts{
        Token: os.Getenv("TOKEN"),
    }

    bot := telegram.NewBot(opts)

    if me, err := bot.GetMe(); err != nil {
        log.Fatalln("Invalid token:", err)
    } else {
        log.Println("Hello! My name is", me.FirstName)
    }

    log.Println("Listening to messages...")

    request := requests.GetUpdates{
        Timeout: 60,

        AllowedUpdates: []string{
            telegram.UpdatesAllowedMessage,
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
```

### Send photo
By upload file:
```go
photo, err := os.Open("t_logo.png")
if err != nil {
    log.Fatalln("Unable to open file:", err)
}

_, err = bot.SendPhoto(&requests.SendPhoto{
    ChatID: update.Message.Chat.ID,
    Photo:  photo,
})

if err != nil {
    log.Fatalln("Unable to send photo:", err)
}
```
By URL:
```go
_, err = bot.SendPhoto(&requests.SendPhoto{
    ChatID: update.Message.Chat.ID,
    Photo:  "https://telegram.org/img/t_logo.png",
})

if err != nil {
    log.Fatalln("Unable to send photo:", err)
}
```
