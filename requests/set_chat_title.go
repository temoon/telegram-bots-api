package requests

import (
	"context"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SetChatTitle struct {
	ChatId interface{}
	Title  string
}

func (r *SetChatTitle) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatTitle", r, response)
	return
}

func (r *SetChatTitle) IsMultipart() (multipart bool) {
	return false
}

func (r *SetChatTitle) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["title"] = r.Title

	return
}
