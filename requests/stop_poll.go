package requests

import (
	"encoding/json"
	"strconv"
)

type StopPoll struct {
	ChatId      interface{}
	MessageId   uint64
	ReplyMarkup interface{}
}

func (r *StopPoll) IsMultipart() bool {
	return false
}

func (r *StopPoll) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["message_id"] = strconv.FormatUint(r.MessageId, 10)

	if r.ReplyMarkup != nil {
		var data []byte
		if data, err = json.Marshal(r.ReplyMarkup); err != nil {
			return
		}

		values["reply_markup"] = string(data)
	}

	return
}
