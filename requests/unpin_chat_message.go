package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type UnpinChatMessage struct {
	MessageId *int64
	ChatId    telegram.ChatId
}

func (r *UnpinChatMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unpinChatMessage", r, response)
	return
}

func (r *UnpinChatMessage) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.MessageId != nil {
		values["message_id"] = strconv.FormatInt(*r.MessageId, 10)
	}

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *UnpinChatMessage) GetFiles() (files map[string]io.Reader) {
	return
}
