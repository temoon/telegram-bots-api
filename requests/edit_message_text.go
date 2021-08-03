package requests

import (
	"encoding/json"
	"strconv"
)

type EditMessageText struct {
	ChatId                interface{}
	DisableWebPagePreview bool
	Entities              []interface{}
	InlineMessageId       string
	MessageId             uint64
	ParseMode             string
	ReplyMarkup           interface{}
	Text                  string
}

func (r *EditMessageText) IsMultipart() bool {
	return false
}

func (r *EditMessageText) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.DisableWebPagePreview {
		values["disable_web_page_preview"] = "1"
	}

	if r.Entities != nil {
		var data []byte
		if data, err = json.Marshal(r.Entities); err != nil {
			return
		}

		values["entities"] = string(data)
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

	values["text"] = r.Text

	return
}
