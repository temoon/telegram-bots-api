package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
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

func (r *UnbanChatSenderChat) IsMultipart() bool {
	return false
}

func (r *UnbanChatSenderChat) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	default:
		err = errors.New("invalid chat_id field type")
		return
	}

	values["sender_chat_id"] = strconv.FormatInt(r.SenderChatId, 10)

	return
}
