package requests

import (
	"encoding/json"
	"strconv"
)

type RestrictChatMember struct {
	ChatId      interface{}
	Permissions interface{}
	UntilDate   uint64
	UserId      uint64
}

func (r *RestrictChatMember) IsMultipart() bool {
	return false
}

func (r *RestrictChatMember) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.Permissions != nil {
		var data []byte
		if data, err = json.Marshal(r.Permissions); err != nil {
			return
		}

		values["permissions"] = string(data)
	}

	if r.UntilDate != 0 {
		values["until_date"] = strconv.FormatUint(r.UntilDate, 10)
	}

	values["user_id"] = strconv.FormatUint(r.UserId, 10)

	return
}
