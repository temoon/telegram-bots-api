package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type DeleteMessage struct {
	ChatId    interface{}
	MessageId int64
}

func (r *DeleteMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteMessage", r, response)
	return
}

func (r *DeleteMessage) IsMultipart() bool {
	return false
}

func (r *DeleteMessage) GetValues() (values map[string]interface{}, err error) {
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

	values["message_id"] = strconv.FormatInt(r.MessageId, 10)

	return
}
