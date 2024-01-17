package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type ReopenGeneralForumTopic struct {
	ChatId telegram.ChatId
}

func (r *ReopenGeneralForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "reopenGeneralForumTopic", r, response)
	return
}

func (r *ReopenGeneralForumTopic) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *ReopenGeneralForumTopic) GetFiles() (files map[string]io.Reader) {
	return
}
