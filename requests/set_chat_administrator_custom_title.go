package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SetChatAdministratorCustomTitle struct {
	UserId      int64
	CustomTitle string
	ChatId      telegram.ChatId
}

func (r *SetChatAdministratorCustomTitle) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatAdministratorCustomTitle", r, response)
	return
}

func (r *SetChatAdministratorCustomTitle) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	values["custom_title"] = r.CustomTitle

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *SetChatAdministratorCustomTitle) GetFiles() (files map[string]io.Reader) {
	return
}
