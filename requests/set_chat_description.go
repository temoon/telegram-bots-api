package requests

import (
	"strconv"
)

type SetChatDescription struct {
	ChatId      interface{}
	Description string
}

func (r *SetChatDescription) IsMultipart() bool {
	return false
}

func (r *SetChatDescription) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.Description != "" {
		values["description"] = r.Description
	}

	return
}
