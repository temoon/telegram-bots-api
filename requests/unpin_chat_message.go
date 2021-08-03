package requests

import (
	"strconv"
)

type UnpinChatMessage struct {
	ChatId    interface{}
	MessageId uint64
}

func (r *UnpinChatMessage) IsMultipart() bool {
	return false
}

func (r *UnpinChatMessage) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.MessageId != 0 {
		values["message_id"] = strconv.FormatUint(r.MessageId, 10)
	}

	return
}
