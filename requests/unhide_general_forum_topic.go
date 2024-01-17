package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type UnhideGeneralForumTopic struct {
	ChatId telegram.ChatId
}

func (r *UnhideGeneralForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unhideGeneralForumTopic", r, response)
	return
}

func (r *UnhideGeneralForumTopic) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *UnhideGeneralForumTopic) GetFiles() (files map[string]io.Reader) {
	return
}
