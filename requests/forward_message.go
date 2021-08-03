package requests

import (
	"strconv"
)

type ForwardMessage struct {
	ChatId              interface{}
	DisableNotification bool
	FromChatId          interface{}
	MessageId           uint64
}

func (r *ForwardMessage) IsMultipart() bool {
	return false
}

func (r *ForwardMessage) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.DisableNotification {
		values["disable_notification"] = "1"
	}

	switch value := r.FromChatId.(type) {
	case uint64:
		values["from_chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["from_chat_id"] = value
	}

	values["message_id"] = strconv.FormatUint(r.MessageId, 10)

	return
}
