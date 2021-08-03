package requests

import (
	"encoding/json"
	"strconv"
)

type SendMediaGroup struct {
	AllowSendingWithoutReply bool
	ChatId                   interface{}
	DisableNotification      bool
	Media                    []interface{}
	ReplyToMessageId         uint64
}

func (r *SendMediaGroup) IsMultipart() bool {
	return true
}

func (r *SendMediaGroup) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowSendingWithoutReply {
		values["allow_sending_without_reply"] = "1"
	}

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.DisableNotification {
		values["disable_notification"] = "1"
	}

	if r.Media != nil {
		var data []byte
		if data, err = json.Marshal(r.Media); err != nil {
			return
		}

		values["media"] = string(data)
	}

	if r.ReplyToMessageId != 0 {
		values["reply_to_message_id"] = strconv.FormatUint(r.ReplyToMessageId, 10)
	}

	return
}
