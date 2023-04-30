package requests

import (
	"context"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type CloseForumTopic struct {
	ChatId          interface{}
	MessageThreadId int32
}

func (r *CloseForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "closeForumTopic", r, response)
	return
}

func (r *CloseForumTopic) IsMultipart() (multipart bool) {
	return false
}

func (r *CloseForumTopic) GetValues() (values map[string]interface{}, err error) {
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
