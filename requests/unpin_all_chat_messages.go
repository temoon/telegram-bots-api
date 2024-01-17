package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type UnpinAllChatMessages struct {
	ChatId telegram.ChatId
}

func (r *UnpinAllChatMessages) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unpinAllChatMessages", r, response)
	return
}

func (r *UnpinAllChatMessages) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *UnpinAllChatMessages) GetFiles() (files map[string]io.Reader) {
	return
}
