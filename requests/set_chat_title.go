package requests

import (
	"strconv"
)

type SetChatTitle struct {
	ChatId interface{}
	Title  string
}

func (r *SetChatTitle) IsMultipart() bool {
	return false
}

func (r *SetChatTitle) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["title"] = r.Title

	return
}
