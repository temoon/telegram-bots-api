package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SendChatAction struct {
	Action               string
	ChatId               telegram.ChatId
	BusinessConnectionId *string
	MessageThreadId      *int64
}

func (r *SendChatAction) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "sendChatAction", r, response)
	return
}

func (r *SendChatAction) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["action"] = r.Action

	values["chat_id"] = r.ChatId.String()

	if r.BusinessConnectionId != nil {
		values["business_connection_id"] = *r.BusinessConnectionId
	}

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	return
}

func (r *SendChatAction) GetFiles() (files map[string]io.Reader) {
	return
}
