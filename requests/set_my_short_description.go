package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type SetMyShortDescription struct {
	LanguageCode     *string
	ShortDescription *string
}

func (r *SetMyShortDescription) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setMyShortDescription", r, response)
	return
}

func (r *SetMyShortDescription) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.LanguageCode != nil {
		values["language_code"] = *r.LanguageCode
	}

	if r.ShortDescription != nil {
		values["short_description"] = *r.ShortDescription
	}

	return
}

func (r *SetMyShortDescription) GetFiles() (files map[string]io.Reader) {
	return
}
