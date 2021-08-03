package requests

import (
	"encoding/json"
	"strconv"
)

type SetChatPermissions struct {
	ChatId      interface{}
	Permissions interface{}
}

func (r *SetChatPermissions) IsMultipart() bool {
	return false
}

func (r *SetChatPermissions) GetValues() (values map[string]interface{}, err error) {
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

	return
}
