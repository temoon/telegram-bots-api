package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type PinChatMessage struct {
	ChatId              telegram.ChatId
	MessageId           int64
	DisableNotification *bool
}

func (r *PinChatMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "pinChatMessage", r, response)
	return
}

func (r *PinChatMessage) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["message_id"] = strconv.FormatInt(r.MessageId, 10)

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	return
}

func (r *PinChatMessage) GetFiles() (files map[string]io.Reader) {
	return
}
