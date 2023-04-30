package requests

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SetChatMenuButton struct {
	ChatId     interface{}
	MenuButton interface{}
}

func (r *SetChatMenuButton) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatMenuButton", r, response)
	return
}

func (r *SetChatMenuButton) IsMultipart() (multipart bool) {
	return false
}

func (r *SetChatMenuButton) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	switch value := r.MenuButton.(type) {
	case *telegram.MenuButtonCommands:
		if value != nil {
			var dataMenuButtonCommands []byte
			if dataMenuButtonCommands, err = json.Marshal(value); err != nil {
				return
			}

			values["menu_button"] = string(dataMenuButtonCommands)
		}
	case *telegram.MenuButtonWebApp:
		if value != nil {
			var dataMenuButtonWebApp []byte
			if dataMenuButtonWebApp, err = json.Marshal(value); err != nil {
				return
			}

			values["menu_button"] = string(dataMenuButtonWebApp)
		}
	case *telegram.MenuButtonDefault:
		if value != nil {
			var dataMenuButtonDefault []byte
			if dataMenuButtonDefault, err = json.Marshal(value); err != nil {
				return
			}

			values["menu_button"] = string(dataMenuButtonDefault)
		}
	}

	return
}
