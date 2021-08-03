package requests

import (
	"strconv"
)

type UnbanChatMember struct {
	ChatId       interface{}
	OnlyIfBanned bool
	UserId       uint64
}

func (r *UnbanChatMember) IsMultipart() bool {
	return false
}

func (r *UnbanChatMember) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.OnlyIfBanned {
		values["only_if_banned"] = "1"
	}

	values["user_id"] = strconv.FormatUint(r.UserId, 10)

	return
}
