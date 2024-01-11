package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type DeleteForumTopic struct {
	ChatId          interface{}
	MessageThreadId int64
}

func (r *DeleteForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteForumTopic", r, response)
	return
}

func (r *DeleteForumTopic) IsMultipart() bool {
	return false
}

func (r *DeleteForumTopic) GetValues() (values map[string]interface{}, err error) {
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

	values["message_thread_id"] = strconv.FormatInt(r.MessageThreadId, 10)

	return
}
