package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type CopyMessages struct {
	ChatId              telegram.ChatId
	FromChatId          telegram.ChatId
	MessageIds          []int64
	DisableNotification *bool
	MessageThreadId     *int64
	ProtectContent      *bool
	RemoveCaption       *bool
}

func (r *CopyMessages) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.MessageId)
	err = b.CallMethod(ctx, "copyMessages", r, response)
	return
}

func (r *CopyMessages) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["from_chat_id"] = r.FromChatId.String()

	var dataMessageIds []byte
	if dataMessageIds, err = json.Marshal(r.MessageIds); err != nil {
		return
	}

	values["message_ids"] = string(dataMessageIds)

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

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

	if r.RemoveCaption != nil {
		if *r.RemoveCaption {
			values["remove_caption"] = "1"
		} else {
			values["remove_caption"] = "0"
		}
	}

	return
}

func (r *CopyMessages) GetFiles() (files map[string]io.Reader) {
	return
}
