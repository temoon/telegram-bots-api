package requests

import (
	"strconv"
)

type RevokeChatInviteLink struct {
	ChatId     interface{}
	InviteLink string
}

func (r *RevokeChatInviteLink) IsMultipart() bool {
	return false
}

func (r *RevokeChatInviteLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["invite_link"] = r.InviteLink

	return
}
