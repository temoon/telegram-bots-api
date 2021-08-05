package requests

import (
	"context"
	"github.com/temoon/go-telegram-bots-api"
	"strconv"
)

type PinChatMessage struct {
	ChatId              interface{}
	DisableNotification bool
	MessageId           int32
}

func (r *PinChatMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "pinChatMessage", r, response)
	return
}

func (r *PinChatMessage) IsMultipart() (multipart bool) {
	return false
}

func (r *PinChatMessage) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.DisableNotification {
		values["disable_notification"] = "1"
	}

	values["message_id"] = strconv.FormatInt(int64(r.MessageId), 10)

	return
}
