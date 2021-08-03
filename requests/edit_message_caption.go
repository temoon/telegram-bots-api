package requests

import (
	"encoding/json"
	"strconv"
)

type EditMessageCaption struct {
	Caption         string
	CaptionEntities []interface{}
	ChatId          interface{}
	InlineMessageId string
	MessageId       uint64
	ParseMode       string
	ReplyMarkup     interface{}
}

func (r *EditMessageCaption) IsMultipart() bool {
	return false
}

func (r *EditMessageCaption) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

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

	if r.InlineMessageId != "" {
		values["inline_message_id"] = r.InlineMessageId
	}

	if r.MessageId != 0 {
		values["message_id"] = strconv.FormatUint(r.MessageId, 10)
	}

	if r.ParseMode != "" {
		values["parse_mode"] = r.ParseMode
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
