package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type UnpinChatMessage struct {
	ChatId    interface{}
	MessageId *int64
}

func (r *UnpinChatMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unpinChatMessage", r, response)
	return
}

func (r *UnpinChatMessage) IsMultipart() bool {
	return false
}

func (r *UnpinChatMessage) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	default:
		err = errors.New("invalid chat_id field type")
		return
	}

	if r.MessageId != nil {
		values["message_id"] = strconv.FormatInt(*r.MessageId, 10)
	}

	return
}
