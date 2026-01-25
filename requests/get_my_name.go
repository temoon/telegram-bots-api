package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type GetMyName struct {
	LanguageCode *string
}

func (r *GetMyName) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.BotName)
	err = b.CallMethod(ctx, "getMyName", r, response)
	return
}

func (r *GetMyName) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.LanguageCode != nil {
		values["language_code"] = *r.LanguageCode
	}

	return
}

func (r *GetMyName) GetFiles() (files map[string]io.Reader) {
	return
}
