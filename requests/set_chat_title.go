package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type SetChatTitle struct {
	ChatId telegram.ChatId
	Title  string
}

func (r *SetChatTitle) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatTitle", r, response)
	return
}

func (r *SetChatTitle) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["title"] = r.Title

	return
}

func (r *SetChatTitle) GetFiles() (files map[string]io.Reader) {
	return
}
