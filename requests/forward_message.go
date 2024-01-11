package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type ForwardMessage struct {
	ChatId              interface{}
	DisableNotification *bool
	FromChatId          interface{}
	MessageId           int64
	MessageThreadId     *int64
	ProtectContent      *bool
}

func (r *ForwardMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "forwardMessage", r, response)
	return
}

func (r *ForwardMessage) IsMultipart() bool {
	return false
}

func (r *ForwardMessage) GetValues() (values map[string]interface{}, err error) {
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
	default:
		err = errors.New("invalid from_chat_id field type")
		return
	}

	values["message_id"] = strconv.FormatInt(r.MessageId, 10)

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	if r.ProtectContent != nil {
		if *r.ProtectContent {
			values["protect_content"] = "1"
		} else {
			values["protect_content"] = "0"
		}
	}

	return
}
