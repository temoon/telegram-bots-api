package requests

import (
	"encoding/json"
	"strconv"
)

type EditMessageMedia struct {
	ChatId          interface{}
	InlineMessageId string
	Media           interface{}
	MessageId       uint64
	ReplyMarkup     interface{}
}

func (r *EditMessageMedia) IsMultipart() bool {
	return true
}

func (r *EditMessageMedia) GetValues() (values map[string]interface{}, err error) {
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

	if r.Media != nil {
		var data []byte
		if data, err = json.Marshal(r.Media); err != nil {
			return
		}

		values["media"] = string(data)
	}

	switch value := r.Media.(type) {
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["media"] = string(data)
		}
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["media"] = string(data)
		}
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["media"] = string(data)
		}
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["media"] = string(data)
		}
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["media"] = string(data)
		}
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
