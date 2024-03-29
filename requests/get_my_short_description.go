package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type GetMyShortDescription struct {
	LanguageCode *string
}

func (r *GetMyShortDescription) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.BotShortDescription)
	err = b.CallMethod(ctx, "getMyShortDescription", r, response)
	return
}

func (r *GetMyShortDescription) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.LanguageCode != nil {
		values["language_code"] = *r.LanguageCode
	}

	return
}

func (r *GetMyShortDescription) GetFiles() (files map[string]io.Reader) {
	return
}
