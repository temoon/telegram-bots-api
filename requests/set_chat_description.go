package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type SetChatDescription struct {
	ChatId      telegram.ChatId
	Description *string
}

func (r *SetChatDescription) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatDescription", r, response)
	return
}

func (r *SetChatDescription) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	if r.Description != nil {
		values["description"] = *r.Description
	}

	return
}

func (r *SetChatDescription) GetFiles() (files map[string]io.Reader) {
	return
}
