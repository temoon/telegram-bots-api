package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type ReopenGeneralForumTopic struct {
	ChatId interface{}
}

func (r *ReopenGeneralForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "reopenGeneralForumTopic", r, response)
	return
}

func (r *ReopenGeneralForumTopic) IsMultipart() (multipart bool) {
	return false
}

func (r *ReopenGeneralForumTopic) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	return
}
