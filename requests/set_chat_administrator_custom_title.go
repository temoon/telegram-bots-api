package requests

import (
	"strconv"
)

type SetChatAdministratorCustomTitle struct {
	ChatId      interface{}
	CustomTitle string
	UserId      uint64
}

func (r *SetChatAdministratorCustomTitle) IsMultipart() bool {
	return false
}

func (r *SetChatAdministratorCustomTitle) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["custom_title"] = r.CustomTitle

	values["user_id"] = strconv.FormatUint(r.UserId, 10)

	return
}
