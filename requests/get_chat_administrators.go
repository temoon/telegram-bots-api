package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type GetChatAdministrators struct {
	ChatId telegram.ChatId
}

func (r *GetChatAdministrators) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]interface{})
	err = b.CallMethod(ctx, "getChatAdministrators", r, response)
	return
}

func (r *GetChatAdministrators) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *GetChatAdministrators) GetFiles() (files map[string]io.Reader) {
	return
}
