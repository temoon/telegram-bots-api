package requests

import (
	"context"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type UnbanChatSenderChat struct {
	ChatId       interface{}
	SenderChatId int64
}

func (r *UnbanChatSenderChat) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unbanChatSenderChat", r, response)
	return
}

func (r *UnbanChatSenderChat) IsMultipart() (multipart bool) {
	return false
}

func (r *UnbanChatSenderChat) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["sender_chat_id"] = strconv.FormatInt(r.SenderChatId, 10)

	return
}
