package requests

import (
	"encoding/json"
	"strconv"
)

type SendMessage struct {
	AllowSendingWithoutReply bool
	ChatId                   interface{}
	DisableNotification      bool
	DisableWebPagePreview    bool
	Entities                 []interface{}
	ParseMode                string
	ReplyMarkup              interface{}
	ReplyToMessageId         uint64
	Text                     string
}

func (r *SendMessage) IsMultipart() bool {
	return false
}

func (r *SendMessage) GetValues() (values map[string]interface{}, err error) {
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

	if r.ParseMode != "" {
		values["parse_mode"] = r.ParseMode
	}

	switch value := r.ReplyMarkup.(type) {
	default:

		var data []byte
		if data, err = json.Marshal(value); err != nil {
			return
		}

		values["reply_markup"] = string(data)

	}

	if r.ReplyToMessageId != 0 {
		values["reply_to_message_id"] = strconv.FormatUint(r.ReplyToMessageId, 10)
	}

	values["text"] = r.Text

	return
}
