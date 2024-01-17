package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type UnpinAllForumTopicMessages struct {
	ChatId          telegram.ChatId
	MessageThreadId int64
}

func (r *UnpinAllForumTopicMessages) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unpinAllForumTopicMessages", r, response)
	return
}

func (r *UnpinAllForumTopicMessages) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["message_thread_id"] = strconv.FormatInt(r.MessageThreadId, 10)

	return
}

func (r *UnpinAllForumTopicMessages) GetFiles() (files map[string]io.Reader) {
	return
}
