package requests

import (
	"context"
	"github.com/temoon/go-telegram-bots-api"
	"strconv"
)

type DeleteChatPhoto struct {
	ChatId interface{}
}

func (r *DeleteChatPhoto) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteChatPhoto", r, response)
	return
}

func (r *DeleteChatPhoto) IsMultipart() (multipart bool) {
	return false
}

func (r *DeleteChatPhoto) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	return
}
