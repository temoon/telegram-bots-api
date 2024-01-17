package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type GetChatMenuButton struct {
	ChatId *int64
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
		err = errors.New("unsupported response type")
	}

	return
}

func (r *GetChatMenuButton) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.ChatId != nil {
		values["chat_id"] = strconv.FormatInt(*r.ChatId, 10)
	}

	return
}

func (r *GetChatMenuButton) GetFiles() (files map[string]io.Reader) {
	return
}
