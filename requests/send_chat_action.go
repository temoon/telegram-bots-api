package requests

import (
	"strconv"
)

type SendChatAction struct {
	Action string
	ChatId interface{}
}

func (r *SendChatAction) IsMultipart() bool {
	return false
}

func (r *SendChatAction) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["action"] = r.Action

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	return
}
