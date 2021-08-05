package requests

import (
	"encoding/json"
	"io"
	"strconv"
)

type SendVideoNote struct {
	AllowSendingWithoutReply bool
	ChatId                   interface{}
	DisableNotification      bool
	Duration                 uint64
	Length                   uint64
	ReplyMarkup              interface{}
	ReplyToMessageId         uint64
	Thumb                    interface{}
	VideoNote                interface{}
}

func (r *SendVideoNote) IsMultipart() bool {
	return true
}

func (r *SendVideoNote) GetValues() (values map[string]interface{}, err error) {
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

	if r.Duration != 0 {
		values["duration"] = strconv.FormatUint(r.Duration, 10)
	}

	if r.Length != 0 {
		values["length"] = strconv.FormatUint(r.Length, 10)
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

	switch value := r.VideoNote.(type) {
	case io.Reader:
		values["video_note"] = value
	case string:
		values["video_note"] = value
	}

	return
}
