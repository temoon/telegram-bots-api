package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type SendAudio struct {
	AllowSendingWithoutReply *bool
	Audio                    interface{}
	Caption                  *string
	CaptionEntities          []telegram.MessageEntity
	ChatId                   interface{}
	DisableNotification      *bool
	Duration                 *int32
	ParseMode                *string
	Performer                *string
	ReplyMarkup              interface{}
	ReplyToMessageId         *int32
	Thumb                    interface{}
	Title                    *string
}

func (r *SendAudio) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendAudio", r, response)
	return
}

func (r *SendAudio) IsMultipart() (multipart bool) {
	return true
}

func (r *SendAudio) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowSendingWithoutReply != nil {
		if *r.AllowSendingWithoutReply {
			values["allow_sending_without_reply"] = "1"
		} else {
			values["allow_sending_without_reply"] = "0"
		}
	}

	switch value := r.Audio.(type) {
	case io.Reader:
		values["audio"] = value
	case string:
		values["audio"] = value
	}

	if r.Caption != nil {
		values["caption"] = *r.Caption
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

	if r.ParseMode != nil {
		values["parse_mode"] = *r.ParseMode
	}

	if r.Performer != nil {
		values["performer"] = *r.Performer
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

	if r.ReplyToMessageId != nil {
		values["reply_to_message_id"] = strconv.FormatInt(int64(*r.ReplyToMessageId), 10)
	}

	switch value := r.Thumb.(type) {
	case io.Reader:
		values["thumb"] = value
	case string:
		values["thumb"] = value
	}

	if r.Title != nil {
		values["title"] = *r.Title
	}

	return
}
