package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type UnpinAllGeneralForumTopicMessages struct {
	ChatId interface{}
}

func (r *UnpinAllGeneralForumTopicMessages) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unpinAllGeneralForumTopicMessages", r, response)
	return
}

func (r *UnpinAllGeneralForumTopicMessages) IsMultipart() (multipart bool) {
	return false
}

func (r *UnpinAllGeneralForumTopicMessages) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	return
}
