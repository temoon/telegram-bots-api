package requests

import (
	"io"
	"strconv"
)

type SetChatPhoto struct {
	ChatId interface{}
	Photo  interface{}
}

func (r *SetChatPhoto) IsMultipart() bool {
	return true
}

func (r *SetChatPhoto) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["photo"] = r.Photo

	return
}
