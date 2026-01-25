package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SetChatMenuButton struct {
	ChatId     *int64
	MenuButton interface{}
}

func (r *SetChatMenuButton) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatMenuButton", r, response)
	return
}

func (r *SetChatMenuButton) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	if r.ChatId != nil {
		values["chat_id"] = strconv.FormatInt(*r.ChatId, 10)
	}

	if r.MenuButton != nil {
		var dataMenuButton []byte
		if dataMenuButton, err = json.Marshal(r.MenuButton); err != nil {
			return
		}

		values["menu_button"] = string(dataMenuButton)
	}

	return
}

func (r *SetChatMenuButton) GetFiles() (files map[string]io.Reader) {
	return
}
