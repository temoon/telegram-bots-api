package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type GetChatMenuButton struct {
	ChatId interface{}
}

func (r *GetChatMenuButton) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(interface{})
	err = b.CallMethod(ctx, "getChatMenuButton", r, response)
	return
}

func (r *GetChatMenuButton) CallWithResponse(ctx context.Context, b *telegram.Bot, response interface{}) (err error) {
	switch response.(type) {
	case *telegram.MenuButtonCommands, *telegram.MenuButtonWebApp, *telegram.MenuButtonDefault:
		err = b.CallMethod(ctx, "getChatMenuButton", r, response)
	default:
		err = errors.New("unexpected response type")
	}

	return
}

func (r *GetChatMenuButton) IsMultipart() (multipart bool) {
	return false
}

func (r *GetChatMenuButton) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	return
}
