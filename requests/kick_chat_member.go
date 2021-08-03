package requests

import (
	"strconv"
)

type KickChatMember struct {
	ChatId         interface{}
	RevokeMessages bool
	UntilDate      uint64
	UserId         uint64
}

func (r *KickChatMember) IsMultipart() bool {
	return false
}

func (r *KickChatMember) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.RevokeMessages {
		values["revoke_messages"] = "1"
	}

	if r.UntilDate != 0 {
		values["until_date"] = strconv.FormatUint(r.UntilDate, 10)
	}

	values["user_id"] = strconv.FormatUint(r.UserId, 10)

	return
}
