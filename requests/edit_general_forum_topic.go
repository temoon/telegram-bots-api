package requests

import (
	"context"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type EditGeneralForumTopic struct {
	ChatId interface{}
	Name   string
}

func (r *EditGeneralForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "editGeneralForumTopic", r, response)
	return
}

func (r *EditGeneralForumTopic) IsMultipart() (multipart bool) {
	return false
}

func (r *EditGeneralForumTopic) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["name"] = r.Name

	return
}