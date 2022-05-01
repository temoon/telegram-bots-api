package requests

import (
	"context"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type GetChatMenuButton struct {
	ChatId *int64
}

func (r *GetChatMenuButton) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.MenuButton)
	err = b.CallMethod(ctx, "getChatMenuButton", r, response)
	return
}

func (r *GetChatMenuButton) IsMultipart() (multipart bool) {
	return false
}

func (r *GetChatMenuButton) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.ChatId != nil {
		values["chat_id"] = strconv.FormatInt(*r.ChatId, 10)
	}

	return
}
