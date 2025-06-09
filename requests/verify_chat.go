package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type VerifyChat struct {
	ChatId            telegram.ChatId
	CustomDescription *string
}

func (r *VerifyChat) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "verifyChat", r, response)
	return
}

func (r *VerifyChat) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	if r.CustomDescription != nil {
		values["custom_description"] = *r.CustomDescription
	}

	return
}

func (r *VerifyChat) GetFiles() (files map[string]io.Reader) {
	return
}
