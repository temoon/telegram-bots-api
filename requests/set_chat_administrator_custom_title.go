package requests

import (
	"context"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SetChatAdministratorCustomTitle struct {
	ChatId      interface{}
	CustomTitle string
	UserId      int64
}

func (r *SetChatAdministratorCustomTitle) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatAdministratorCustomTitle", r, response)
	return
}

func (r *SetChatAdministratorCustomTitle) IsMultipart() (multipart bool) {
	return false
}

func (r *SetChatAdministratorCustomTitle) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["custom_title"] = r.CustomTitle

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
