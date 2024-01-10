package requests

import (
"errors"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type GetChatMenuButton struct {
ChatId *int32
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

	
			if r.ChatId != nil {
			values["chat_id"] = strconv.FormatInt(int64(*r.ChatId), 10)
			}
			

	return
}
