package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SendVideoNote struct {
	AllowSendingWithoutReply *bool
	ChatId                   interface{}
	DisableNotification      *bool
	Duration                 *int32
	Length                   *int32
	MessageThreadId          *int32
	ProtectContent           *bool
	ReplyMarkup              interface{}
	ReplyToMessageId         *int32
	Thumbnail                interface{}
	VideoNote                interface{}
}

func (r *SendVideoNote) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendVideoNote", r, response)
	return
}

func (r *SendVideoNote) IsMultipart() (multipart bool) {
	return true
}

func (r *SendVideoNote) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowSendingWithoutReply != nil {
		if *r.AllowSendingWithoutReply {
			values["allow_sending_without_reply"] = "1"
		} else {
			values["allow_sending_without_reply"] = "0"
		}
	}

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	if r.Duration != nil {
		values["duration"] = strconv.FormatInt(int64(*r.Duration), 10)
	}

	if r.Length != nil {
		values["length"] = strconv.FormatInt(int64(*r.Length), 10)
	}

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(int64(*r.MessageThreadId), 10)
	}

	if r.ProtectContent != nil {
		if *r.ProtectContent {
			values["protect_content"] = "1"
		} else {
			values["protect_content"] = "0"
		}
	}

	switch value := r.ReplyMarkup.(type) {
	case *telegram.InlineKeyboardMarkup:
		if value != nil {
			var dataInlineKeyboardMarkup []byte
			if dataInlineKeyboardMarkup, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(dataInlineKeyboardMarkup)
		}
	case *telegram.ReplyKeyboardMarkup:
		if value != nil {
			var dataReplyKeyboardMarkup []byte
			if dataReplyKeyboardMarkup, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(dataReplyKeyboardMarkup)
		}
	case *telegram.ReplyKeyboardRemove:
		if value != nil {
			var dataReplyKeyboardRemove []byte
			if dataReplyKeyboardRemove, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(dataReplyKeyboardRemove)
		}
	case *telegram.ForceReply:
		if value != nil {
			var dataForceReply []byte
			if dataForceReply, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(dataForceReply)
		}
	}

	if r.ReplyToMessageId != nil {
		values["reply_to_message_id"] = strconv.FormatInt(int64(*r.ReplyToMessageId), 10)
	}

	switch value := r.Thumbnail.(type) {
	case io.Reader:
		values["thumbnail"] = value
	case string:
		values["thumbnail"] = value
	}

	switch value := r.VideoNote.(type) {
	case io.Reader:
		values["video_note"] = value
	case string:
		values["video_note"] = value
	}

	return
}
