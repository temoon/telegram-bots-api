package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type GetMyDescription struct {
	LanguageCode *string
}

func (r *GetMyDescription) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.BotDescription)
	err = b.CallMethod(ctx, "getMyDescription", r, response)
	return
}

func (r *GetMyDescription) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	if r.LanguageCode != nil {
		values["language_code"] = *r.LanguageCode
	}

	return
}

func (r *GetMyDescription) GetFiles() (files map[string]io.Reader) {
	return
}
