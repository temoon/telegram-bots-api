package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type HideGeneralForumTopic struct {
	ChatId telegram.ChatId
}

func (r *HideGeneralForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "hideGeneralForumTopic", r, response)
	return
}

func (r *HideGeneralForumTopic) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *HideGeneralForumTopic) GetFiles() (files map[string]io.Reader) {
	return
}
