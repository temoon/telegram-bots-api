package requests

import (
	"strconv"
)

type UnpinAllChatMessages struct {
	ChatId interface{}
}

func (r *UnpinAllChatMessages) IsMultipart() bool {
	return false
}

func (r *UnpinAllChatMessages) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	return
}
