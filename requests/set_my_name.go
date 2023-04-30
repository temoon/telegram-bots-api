package requests

import (
	"context"

	"github.com/temoon/telegram-bots-api"
)

type SetMyName struct {
	LanguageCode *string
	Name         *string
}

func (r *SetMyName) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setMyName", r, response)
	return
}

func (r *SetMyName) IsMultipart() (multipart bool) {
	return false
}

func (r *SetMyName) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.LanguageCode != nil {
		values["language_code"] = *r.LanguageCode
	}

	if r.Name != nil {
		values["name"] = *r.Name
	}

	return
}
