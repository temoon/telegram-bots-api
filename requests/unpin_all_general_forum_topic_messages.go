package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type UnpinAllGeneralForumTopicMessages struct {
	ChatId telegram.ChatId
}

func (r *UnpinAllGeneralForumTopicMessages) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unpinAllGeneralForumTopicMessages", r, response)
	return
}

func (r *UnpinAllGeneralForumTopicMessages) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *UnpinAllGeneralForumTopicMessages) GetFiles() (files map[string]io.Reader) {
	return
}
