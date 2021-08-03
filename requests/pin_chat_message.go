package requests

import (
	"strconv"
)

type PinChatMessage struct {
	ChatId              interface{}
	DisableNotification bool
	MessageId           uint64
}

func (r *PinChatMessage) IsMultipart() bool {
	return false
}

func (r *PinChatMessage) GetValues() (values map[string]interface{}, err error) {
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

	values["message_id"] = strconv.FormatUint(r.MessageId, 10)

	return
}
