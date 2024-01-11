package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type GetChatAdministrators struct {
	ChatId interface{}
}

func (r *GetChatAdministrators) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]interface{})
	err = b.CallMethod(ctx, "getChatAdministrators", r, response)
	return
}

func (r *GetChatAdministrators) IsMultipart() bool {
	return false
}

func (r *GetChatAdministrators) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	default:
		err = errors.New("invalid chat_id field type")
		return
	}

	return
}
