package requests

import (
	"encoding/json"
	"io"
	"strconv"
)

type SendPhoto struct {
	AllowSendingWithoutReply bool
	Caption                  string
	CaptionEntities          []interface{}
	ChatId                   interface{}
	DisableNotification      bool
	ParseMode                string
	Photo                    interface{}
	ReplyMarkup              interface{}
	ReplyToMessageId         uint64
}

func (r *SendPhoto) IsMultipart() bool {
	return true
}

func (r *SendPhoto) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowSendingWithoutReply {
		values["allow_sending_without_reply"] = "1"
	}

	if r.Caption != "" {
		values["caption"] = r.Caption
	}

	if r.CaptionEntities != nil {
		var data []byte
		if data, err = json.Marshal(r.CaptionEntities); err != nil {
			return
		}

		values["caption_entities"] = string(data)
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

	if r.ParseMode != "" {
		values["parse_mode"] = r.ParseMode
	}

	switch value := r.Photo.(type) {
	case io.Reader:
		values["photo"] = value
	case string:
		values["photo"] = value
	}

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

	return
}
