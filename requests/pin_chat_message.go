package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type PinChatMessage struct {
	ChatId               telegram.ChatId
	MessageId            int64
	BusinessConnectionId *string
	DisableNotification  *bool
}

func (r *PinChatMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "pinChatMessage", r, response)
	return
}

func (r *PinChatMessage) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	values["message_id"] = strconv.FormatInt(r.MessageId, 10)

	if r.BusinessConnectionId != nil {
		values["business_connection_id"] = *r.BusinessConnectionId
	}

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
