package requests

import (
	"encoding/json"
	"strconv"
)

type SendContact struct {
	AllowSendingWithoutReply bool
	ChatId                   interface{}
	DisableNotification      bool
	FirstName                string
	LastName                 string
	PhoneNumber              string
	ReplyMarkup              interface{}
	ReplyToMessageId         uint64
	Vcard                    string
}

func (r *SendContact) IsMultipart() bool {
	return false
}

func (r *SendContact) GetValues() (values map[string]interface{}, err error) {
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

	values["first_name"] = r.FirstName

	if r.LastName != "" {
		values["last_name"] = r.LastName
	}

	values["phone_number"] = r.PhoneNumber

	switch value := r.ReplyMarkup.(type) {
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		}
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		}
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		}
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		}
	}

	if r.ReplyToMessageId != 0 {
		values["reply_to_message_id"] = strconv.FormatUint(r.ReplyToMessageId, 10)
	}

	if r.Vcard != "" {
		values["vcard"] = r.Vcard
	}

	return
}
