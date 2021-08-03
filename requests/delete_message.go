package requests

import (
	"strconv"
)

type DeleteMessage struct {
	ChatId    interface{}
	MessageId uint64
}

func (r *DeleteMessage) IsMultipart() bool {
	return false
}

func (r *DeleteMessage) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["message_id"] = strconv.FormatUint(r.MessageId, 10)

	return
}
