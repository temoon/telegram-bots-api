package helpers

import (
	"context"
	"errors"

	"github.com/temoon/go-telegram-bots-api"
	"github.com/temoon/go-telegram-bots-api/requests"
)

func NewBotWithCheck(ctx context.Context, opts *telegram.BotOpts) (bot *telegram.Bot, me *telegram.User, err error) {
	bot = telegram.NewBot(opts)

	var ok bool
	var req = requests.GetMe{}
	var res interface{}
	if res, err = req.Call(ctx, bot); err != nil {
		return nil, nil, err
	} else if me, ok = res.(*telegram.User); ok {
		return
	} else {
		return nil, nil, errors.New("not found")
	}
}
