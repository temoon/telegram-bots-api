package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/go-telegram-bots-api"
	"io"
	"strconv"
)

type SendVideoNote struct {
	AllowSendingWithoutReply bool
	ChatId                   interface{}
	DisableNotification      bool
	Duration                 int32
	Length                   int32
	ReplyMarkup              interface{}
	ReplyToMessageId         int32
	Thumb                    interface{}
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

	if r.AllowSendingWithoutReply {
		values["allow_sending_without_reply"] = "1"
	}

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.DisableNotification {
		values["disable_notification"] = "1"
	}

	if r.Duration != 0 {
		values["duration"] = strconv.FormatInt(int64(r.Duration), 10)
	}

	if r.Length != 0 {
		values["length"] = strconv.FormatInt(int64(r.Length), 10)
	}

	switch value := r.ReplyMarkup.(type) {
	case *telegram.InlineKeyboardMarkup:
		var dataInlineKeyboardMarkup []byte
		if dataInlineKeyboardMarkup, err = json.Marshal(value); err != nil {
			return
		}

		values["reply_markup"] = string(dataInlineKeyboardMarkup)
	case *telegram.ReplyKeyboardMarkup:
		var dataReplyKeyboardMarkup []byte
		if dataReplyKeyboardMarkup, err = json.Marshal(value); err != nil {
			return
		}

		values["reply_markup"] = string(dataReplyKeyboardMarkup)
	case *telegram.ReplyKeyboardRemove:
		var dataReplyKeyboardRemove []byte
		if dataReplyKeyboardRemove, err = json.Marshal(value); err != nil {
			return
		}

		values["reply_markup"] = string(dataReplyKeyboardRemove)
	case *telegram.ForceReply:
		var dataForceReply []byte
		if dataForceReply, err = json.Marshal(value); err != nil {
			return
		}

		values["reply_markup"] = string(dataForceReply)
	}

	if r.ReplyToMessageId != 0 {
		values["reply_to_message_id"] = strconv.FormatInt(int64(r.ReplyToMessageId), 10)
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
