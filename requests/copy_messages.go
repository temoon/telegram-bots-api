package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type CopyMessages struct {
	ChatId              interface{}
	DisableNotification *bool
	FromChatId          interface{}
	MessageIds          []int32
	MessageThreadId     *int32
	ProtectContent      *bool
	RemoveCaption       *bool
}

func (r *CopyMessages) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.MessageId)
	err = b.CallMethod(ctx, "copyMessages", r, response)
	return
}

func (r *CopyMessages) IsMultipart() (multipart bool) {
	return false
}

func (r *CopyMessages) GetValues() (values map[string]interface{}, err error) {
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

	var dataMessageIds []byte
	if dataMessageIds, err = json.Marshal(r.MessageIds); err != nil {
		return
	}

	values["message_ids"] = string(dataMessageIds)

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(int64(*r.MessageThreadId), 10)
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
