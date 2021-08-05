package requests

import (
	"encoding/json"
	"strconv"
)

type CopyMessage struct {
	AllowSendingWithoutReply bool
	Caption                  string
	CaptionEntities          []interface{}
	ChatId                   interface{}
	DisableNotification      bool
	FromChatId               interface{}
	MessageId                uint64
	ParseMode                string
	ReplyMarkup              interface{}
	ReplyToMessageId         uint64
}

func (r *CopyMessage) IsMultipart() bool {
	return false
}

func (r *CopyMessage) GetValues() (values map[string]interface{}, err error) {
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

	switch value := r.FromChatId.(type) {
	case uint64:
		values["from_chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["from_chat_id"] = value
	}

	values["message_id"] = strconv.FormatUint(r.MessageId, 10)

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

	return
}
