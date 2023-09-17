package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type UnpinAllForumTopicMessages struct {
	ChatId          interface{}
	MessageThreadId int32
}

func (r *UnpinAllForumTopicMessages) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unpinAllForumTopicMessages", r, response)
	return
}

func (r *UnpinAllForumTopicMessages) IsMultipart() (multipart bool) {
	return false
}

func (r *UnpinAllForumTopicMessages) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["message_thread_id"] = strconv.FormatInt(int64(r.MessageThreadId), 10)

	return
}
