package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type UnhideGeneralForumTopic struct {
	ChatId telegram.ChatId
}

func (r *UnhideGeneralForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unhideGeneralForumTopic", r, response)
	return
}

func (r *UnhideGeneralForumTopic) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *UnhideGeneralForumTopic) GetFiles() (files map[string]io.Reader) {
	return
}
