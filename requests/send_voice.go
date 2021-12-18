package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type SendVoice struct {
	AllowSendingWithoutReply bool
	Caption                  string
	CaptionEntities          []telegram.MessageEntity
	ChatId                   interface{}
	DisableNotification      bool
	Duration                 int32
	ParseMode                string
	ReplyMarkup              interface{}
	ReplyToMessageId         int32
	Voice                    interface{}
}

func (r *SendVoice) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendVoice", r, response)
	return
}

func (r *SendVoice) IsMultipart() (multipart bool) {
	return true
}

func (r *SendVoice) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowSendingWithoutReply {
		values["allow_sending_without_reply"] = "1"
	}

	if r.Caption != "" {
		values["caption"] = r.Caption
	}

	if r.CaptionEntities != nil {
		var dataCaptionEntities []byte
		if dataCaptionEntities, err = json.Marshal(r.CaptionEntities); err != nil {
			return
		}

		values["caption_entities"] = string(dataCaptionEntities)
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

	if r.ParseMode != "" {
		values["parse_mode"] = r.ParseMode
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

	switch value := r.Voice.(type) {
	case io.Reader:
		values["voice"] = value
	case string:
		values["voice"] = value
	}

	return
}
