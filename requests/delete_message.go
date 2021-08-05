package requests

import (
	"context"
	"github.com/temoon/go-telegram-bots-api"
	"strconv"
)

type DeleteMessage struct {
	ChatId    interface{}
	MessageId int32
}

func (r *DeleteMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteMessage", r, response)
	return
}

func (r *DeleteMessage) IsMultipart() (multipart bool) {
	return false
}

func (r *DeleteMessage) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["message_id"] = strconv.FormatInt(int64(r.MessageId), 10)

	return
}
