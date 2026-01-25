package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type CloseGeneralForumTopic struct {
	ChatId telegram.ChatId
}

func (r *CloseGeneralForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "closeGeneralForumTopic", r, response)
	return
}

func (r *CloseGeneralForumTopic) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *CloseGeneralForumTopic) GetFiles() (files map[string]io.Reader) {
	return
}
