package requests

import (
	"context"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type BanChatSenderChat struct {
	ChatId       interface{}
	SenderChatId int64
}

func (r *BanChatSenderChat) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "banChatSenderChat", r, response)
	return
}

func (r *BanChatSenderChat) IsMultipart() (multipart bool) {
	return false
}

func (r *BanChatSenderChat) GetValues() (values map[string]interface{}, err error) {
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
