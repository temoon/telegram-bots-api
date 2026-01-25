package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type RemoveChatVerification struct {
	ChatId telegram.ChatId
}

func (r *RemoveChatVerification) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "removeChatVerification", r, response)
	return
}

func (r *RemoveChatVerification) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *RemoveChatVerification) GetFiles() (files map[string]io.Reader) {
	return
}
