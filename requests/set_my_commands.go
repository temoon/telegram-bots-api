package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
)

type SetMyCommands struct {
	Commands     []telegram.BotCommand
	LanguageCode *string
	Scope        interface{}
}

func (r *SetMyCommands) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setMyCommands", r, response)
	return
}

func (r *SetMyCommands) IsMultipart() bool {
	return false
}

func (r *SetMyCommands) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	var dataCommands []byte
	if dataCommands, err = json.Marshal(r.Commands); err != nil {
		return
	}

	values["commands"] = string(dataCommands)

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
