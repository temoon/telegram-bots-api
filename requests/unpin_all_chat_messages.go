package requests

import (
	"context"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type UnpinAllChatMessages struct {
	ChatId interface{}
}

func (r *UnpinAllChatMessages) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unpinAllChatMessages", r, response)
	return
}

func (r *UnpinAllChatMessages) IsMultipart() (multipart bool) {
	return false
}

func (r *UnpinAllChatMessages) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	return
}
