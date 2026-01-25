package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type GetChat struct {
	ChatId telegram.ChatId
}

func (r *GetChat) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.ChatFullInfo)
	err = b.CallMethod(ctx, "getChat", r, response)
	return
}

func (r *GetChat) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *GetChat) GetFiles() (files map[string]io.Reader) {
	return
}
