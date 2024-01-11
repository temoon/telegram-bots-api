package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"strconv"
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

func (r *SetChatMenuButton) IsMultipart() bool {
	return false
}

func (r *SetChatMenuButton) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

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
