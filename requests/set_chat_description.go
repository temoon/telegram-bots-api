package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
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

func (r *SetChatDescription) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	if r.Description != nil {
		values["description"] = *r.Description
	}

	return
}

func (r *SetChatDescription) GetFiles() (files map[string]io.Reader) {
	return
}
