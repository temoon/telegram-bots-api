package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type ReopenForumTopic struct {
	ChatId          telegram.ChatId
	MessageThreadId int64
}

func (r *ReopenForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "reopenForumTopic", r, response)
	return
}

func (r *ReopenForumTopic) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	values["message_thread_id"] = strconv.FormatInt(r.MessageThreadId, 10)

	return
}

func (r *ReopenForumTopic) GetFiles() (files map[string]io.Reader) {
	return
}
