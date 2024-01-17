package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SendChatAction struct {
	Action          string
	ChatId          telegram.ChatId
	MessageThreadId *int64
}

func (r *SendChatAction) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "sendChatAction", r, response)
	return
}

func (r *SendChatAction) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["action"] = r.Action

	values["chat_id"] = r.ChatId.String()

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	return
}

func (r *SendChatAction) GetFiles() (files map[string]io.Reader) {
	return
}
