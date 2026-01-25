package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type UnpinAllChatMessages struct {
	ChatId telegram.ChatId
}

func (r *UnpinAllChatMessages) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unpinAllChatMessages", r, response)
	return
}

func (r *UnpinAllChatMessages) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *UnpinAllChatMessages) GetFiles() (files map[string]io.Reader) {
	return
}
