package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/temoon/telegram-bots-api"
	"github.com/temoon/telegram-bots-api/requests"
)

func main() {
	botOpts := telegram.BotOpts{
		Token:   os.Getenv("TELEGRAM_TOKEN"),
		Timeout: time.Minute,
	}

	bot := telegram.NewBot(&botOpts)

	reqGetMe := requests.GetMe{}

	var err error
	var res interface{}
	if res, err = reqGetMe.Call(context.Background(), bot); err != nil {
		log.Fatalln(err)
	}

	if user, ok := res.(*telegram.User); ok {
		if user.Username != nil {
			log.Println(*user.Username)
		}
	}
}
