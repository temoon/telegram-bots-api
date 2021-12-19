package requests

import (
	"context"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type ForwardMessage struct {
	ChatId              interface{}
	DisableNotification *bool
	FromChatId          interface{}
	MessageId           int32
}

func (r *ForwardMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "forwardMessage", r, response)
	return
}

func (r *ForwardMessage) IsMultipart() (multipart bool) {
	return false
}

func (r *ForwardMessage) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	switch value := r.FromChatId.(type) {
	case int64:
		values["from_chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["from_chat_id"] = value
	}

	values["message_id"] = strconv.FormatInt(int64(r.MessageId), 10)

	return
}
