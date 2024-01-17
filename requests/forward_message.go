package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type ForwardMessage struct {
	MessageId           int64
	ChatId              telegram.ChatId
	MessageThreadId     *int64
	FromChatId          telegram.ChatId
	DisableNotification *bool
	ProtectContent      *bool
}

func (r *ForwardMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "forwardMessage", r, response)
	return
}

func (r *ForwardMessage) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["message_id"] = strconv.FormatInt(r.MessageId, 10)

	values["chat_id"] = r.ChatId.String()

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	values["from_chat_id"] = r.FromChatId.String()

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
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

func (r *ForwardMessage) GetFiles() (files map[string]io.Reader) {
	return
}
