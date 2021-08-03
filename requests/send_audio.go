package requests

import (
	"encoding/json"
	"io"
	"strconv"
)

type SendAudio struct {
	AllowSendingWithoutReply bool
	Audio                    interface{}
	Caption                  string
	CaptionEntities          []interface{}
	ChatId                   interface{}
	DisableNotification      bool
	Duration                 uint64
	ParseMode                string
	Performer                string
	ReplyMarkup              interface{}
	ReplyToMessageId         uint64
	Thumb                    interface{}
	Title                    string
}

func (r *SendAudio) IsMultipart() bool {
	return true
}

func (r *SendAudio) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowSendingWithoutReply {
		values["allow_sending_without_reply"] = "1"
	}

	switch value := r.Audio.(type) {
	case io.Reader:
		values["audio"] = value
	case string:
		values["audio"] = value
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

	if r.Duration != 0 {
		values["duration"] = strconv.FormatUint(r.Duration, 10)
	}

	if r.ParseMode != "" {
		values["parse_mode"] = r.ParseMode
	}

	if r.Performer != "" {
		values["performer"] = r.Performer
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

	switch value := r.Thumb.(type) {
	case io.Reader:
		values["thumb"] = value
	case string:
		values["thumb"] = value
	}

	if r.Title != "" {
		values["title"] = r.Title
	}

	return
}
