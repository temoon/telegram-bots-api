package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SetChatAdministratorCustomTitle struct {
	ChatId      telegram.ChatId
	CustomTitle string
	UserId      int64
}

func (r *SetChatAdministratorCustomTitle) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatAdministratorCustomTitle", r, response)
	return
}

func (r *SetChatAdministratorCustomTitle) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["custom_title"] = r.CustomTitle

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *SetChatAdministratorCustomTitle) GetFiles() (files map[string]io.Reader) {
	return
}
