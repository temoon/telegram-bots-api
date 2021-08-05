package requests

import (
	"encoding/json"
	"io"
	"strconv"
)

type SendDocument struct {
	AllowSendingWithoutReply    bool
	Caption                     string
	CaptionEntities             []interface{}
	ChatId                      interface{}
	DisableContentTypeDetection bool
	DisableNotification         bool
	Document                    interface{}
	ParseMode                   string
	ReplyMarkup                 interface{}
	ReplyToMessageId            uint64
	Thumb                       interface{}
}

func (r *SendDocument) IsMultipart() bool {
	return true
}

func (r *SendDocument) GetValues() (values map[string]interface{}, err error) {
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

	if r.DisableContentTypeDetection {
		values["disable_content_type_detection"] = "1"
	}

	if r.DisableNotification {
		values["disable_notification"] = "1"
	}

	switch value := r.Document.(type) {
	case io.Reader:
		values["document"] = value
	case string:
		values["document"] = value
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

	switch value := r.Thumb.(type) {
	case io.Reader:
		values["thumb"] = value
	case string:
		values["thumb"] = value
	}

	return
}
