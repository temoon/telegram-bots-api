package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type UnpinChatMessage struct {
	ChatId               telegram.ChatId
	BusinessConnectionId *string
	MessageId            *int64
}

func (r *UnpinChatMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unpinChatMessage", r, response)
	return
}

func (r *UnpinChatMessage) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	if r.BusinessConnectionId != nil {
		values["business_connection_id"] = *r.BusinessConnectionId
	}

	if r.MessageId != nil {
		values["message_id"] = strconv.FormatInt(*r.MessageId, 10)
	}

	return
}

func (r *UnpinChatMessage) GetFiles() (files map[string]io.Reader) {
	return
}
