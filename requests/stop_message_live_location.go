package requests

import (
	"encoding/json"
	"strconv"
)

type StopMessageLiveLocation struct {
	ChatId          interface{}
	InlineMessageId string
	MessageId       uint64
	ReplyMarkup     interface{}
}

func (r *StopMessageLiveLocation) IsMultipart() bool {
	return false
}

func (r *StopMessageLiveLocation) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.InlineMessageId != "" {
		values["inline_message_id"] = r.InlineMessageId
	}

	if r.MessageId != 0 {
		values["message_id"] = strconv.FormatUint(r.MessageId, 10)
	}

	if r.ReplyMarkup != nil {
		var data []byte
		if data, err = json.Marshal(r.ReplyMarkup); err != nil {
			return
		}

		values["reply_markup"] = string(data)
	}

	return
}
