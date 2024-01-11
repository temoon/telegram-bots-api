package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type PinChatMessage struct {
	ChatId              interface{}
	DisableNotification *bool
	MessageId           int64
}

func (r *PinChatMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "pinChatMessage", r, response)
	return
}

func (r *PinChatMessage) IsMultipart() bool {
	return false
}

func (r *PinChatMessage) GetValues() (values map[string]interface{}, err error) {
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

	values["message_id"] = strconv.FormatInt(r.MessageId, 10)

	return
}
