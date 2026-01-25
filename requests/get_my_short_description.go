package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type GetMyShortDescription struct {
	LanguageCode *string
}

func (r *GetMyShortDescription) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.BotShortDescription)
	err = b.CallMethod(ctx, "getMyShortDescription", r, response)
	return
}

func (r *GetMyShortDescription) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	if r.LanguageCode != nil {
		values["language_code"] = *r.LanguageCode
	}

	return
}

func (r *GetMyShortDescription) GetFiles() (files map[string]io.Reader) {
	return
}
