package requests

import (
	"strconv"
)

type EditChatInviteLink struct {
	ChatId      interface{}
	ExpireDate  uint64
	InviteLink  string
	MemberLimit uint64
}

func (r *EditChatInviteLink) IsMultipart() bool {
	return false
}

func (r *EditChatInviteLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.ExpireDate != 0 {
		values["expire_date"] = strconv.FormatUint(r.ExpireDate, 10)
	}

	values["invite_link"] = r.InviteLink

	if r.MemberLimit != 0 {
		values["member_limit"] = strconv.FormatUint(r.MemberLimit, 10)
	}

	return
}
