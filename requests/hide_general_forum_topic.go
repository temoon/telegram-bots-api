package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type HideGeneralForumTopic struct {
	ChatId telegram.ChatId
}

func (r *HideGeneralForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "hideGeneralForumTopic", r, response)
	return
}

func (r *HideGeneralForumTopic) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *HideGeneralForumTopic) GetFiles() (files map[string]io.Reader) {
	return
}
