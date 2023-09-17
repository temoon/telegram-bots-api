package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type GetChatMemberCount struct {
	ChatId interface{}
}

func (r *GetChatMemberCount) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(int32)
	err = b.CallMethod(ctx, "getChatMemberCount", r, response)
	return
}

func (r *GetChatMemberCount) IsMultipart() (multipart bool) {
	return false
}

func (r *GetChatMemberCount) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	return
}
