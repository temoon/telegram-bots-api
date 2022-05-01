package requests

import (
	"context"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type GetChat struct {
	ChatId interface{}
}

func (r *GetChat) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Chat)
	err = b.CallMethod(ctx, "getChat", r, response)
	return
}

func (r *GetChat) IsMultipart() (multipart bool) {
	return false
}

func (r *GetChat) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	return
}
