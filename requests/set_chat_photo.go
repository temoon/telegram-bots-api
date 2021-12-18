package requests

import (
	"context"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type SetChatPhoto struct {
	ChatId interface{}
	Photo  interface{}
}

func (r *SetChatPhoto) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatPhoto", r, response)
	return
}

func (r *SetChatPhoto) IsMultipart() (multipart bool) {
	return true
}

func (r *SetChatPhoto) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["photo"] = r.Photo

	return
}
