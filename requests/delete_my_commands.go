package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type DeleteMyCommands struct {
	Scope        interface{}
	LanguageCode *string
}

func (r *DeleteMyCommands) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteMyCommands", r, response)
	return
}

func (r *DeleteMyCommands) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.Scope != nil {
		var dataScope []byte
		if dataScope, err = json.Marshal(r.Scope); err != nil {
			return
		}

		values["scope"] = string(dataScope)
	}

	if r.LanguageCode != nil {
		values["language_code"] = *r.LanguageCode
	}

	return
}

func (r *DeleteMyCommands) GetFiles() (files map[string]io.Reader) {
	return
}
