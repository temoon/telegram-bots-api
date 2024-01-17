package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type UnbanChatSenderChat struct {
	ChatId       telegram.ChatId
	SenderChatId int64
}

func (r *UnbanChatSenderChat) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unbanChatSenderChat", r, response)
	return
}

func (r *UnbanChatSenderChat) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["sender_chat_id"] = strconv.FormatInt(r.SenderChatId, 10)

	return
}

func (r *UnbanChatSenderChat) GetFiles() (files map[string]io.Reader) {
	return
}
