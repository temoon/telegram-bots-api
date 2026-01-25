package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type BanChatSenderChat struct {
	ChatId       telegram.ChatId
	SenderChatId int64
}

func (r *BanChatSenderChat) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "banChatSenderChat", r, response)
	return
}

func (r *BanChatSenderChat) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	values["sender_chat_id"] = strconv.FormatInt(r.SenderChatId, 10)

	return
}

func (r *BanChatSenderChat) GetFiles() (files map[string]io.Reader) {
	return
}
