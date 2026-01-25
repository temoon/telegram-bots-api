package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type CloseForumTopic struct {
	ChatId          telegram.ChatId
	MessageThreadId int64
}

func (r *CloseForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "closeForumTopic", r, response)
	return
}

func (r *CloseForumTopic) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	values["message_thread_id"] = strconv.FormatInt(r.MessageThreadId, 10)

	return
}

func (r *CloseForumTopic) GetFiles() (files map[string]io.Reader) {
	return
}
