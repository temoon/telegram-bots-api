package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type RemoveChatVerification struct {
	ChatId telegram.ChatId
}

func (r *RemoveChatVerification) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "removeChatVerification", r, response)
	return
}

func (r *RemoveChatVerification) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *RemoveChatVerification) GetFiles() (files map[string]io.Reader) {
	return
}
