package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type EditGeneralForumTopic struct {
	ChatId telegram.ChatId
	Name   string
}

func (r *EditGeneralForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "editGeneralForumTopic", r, response)
	return
}

func (r *EditGeneralForumTopic) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["name"] = r.Name

	return
}

func (r *EditGeneralForumTopic) GetFiles() (files map[string]io.Reader) {
	return
}
