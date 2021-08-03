package requests

import (
	"strconv"
)

type GetChatMember struct {
	ChatId interface{}
	UserId uint64
}

func (r *GetChatMember) IsMultipart() bool {
	return false
}

func (r *GetChatMember) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["user_id"] = strconv.FormatUint(r.UserId, 10)

	return
}
