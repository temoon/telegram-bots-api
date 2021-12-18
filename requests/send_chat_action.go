package requests

import (
	"context"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type SendChatAction struct {
	Action string
	ChatId interface{}
}

func (r *SendChatAction) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "sendChatAction", r, response)
	return
}

func (r *SendChatAction) IsMultipart() (multipart bool) {
	return false
}

func (r *SendChatAction) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["action"] = r.Action

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	return
}
