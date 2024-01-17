package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type GetMyCommands struct {
	LanguageCode *string
	Scope        interface{}
}

func (r *GetMyCommands) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.BotCommand)
	err = b.CallMethod(ctx, "getMyCommands", r, response)
	return
}

func (r *GetMyCommands) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.LanguageCode != nil {
		values["language_code"] = *r.LanguageCode
	}

	if r.Scope != nil {
		var dataScope []byte
		if dataScope, err = json.Marshal(r.Scope); err != nil {
			return
		}

		values["scope"] = string(dataScope)
	}

	return
}

func (r *GetMyCommands) GetFiles() (files map[string]io.Reader) {
	return
}
