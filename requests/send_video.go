package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SendVideo struct {
	AllowSendingWithoutReply *bool
	Caption                  *string
	CaptionEntities          []telegram.MessageEntity
	ChatId                   interface{}
	DisableNotification      *bool
	Duration                 *int32
	HasSpoiler               *bool
	Height                   *int32
	MessageThreadId          *int32
	ParseMode                *string
	ProtectContent           *bool
	ReplyMarkup              interface{}
	ReplyToMessageId         *int32
	SupportsStreaming        *bool
	Thumbnail                interface{}
	Video                    interface{}
	Width                    *int32
}

func (r *SendVideo) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendVideo", r, response)
	return
}

func (r *SendVideo) IsMultipart() (multipart bool) {
	return true
}

func (r *SendVideo) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowSendingWithoutReply != nil {
		if *r.AllowSendingWithoutReply {
			values["allow_sending_without_reply"] = "1"
		} else {
			values["allow_sending_without_reply"] = "0"
		}
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

	if r.HasSpoiler != nil {
		if *r.HasSpoiler {
			values["has_spoiler"] = "1"
		} else {
			values["has_spoiler"] = "0"
		}
	}

	if r.Height != nil {
		values["height"] = strconv.FormatInt(int64(*r.Height), 10)
	}

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(int64(*r.MessageThreadId), 10)
	}

	if r.ParseMode != nil {
		values["parse_mode"] = *r.ParseMode
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

	if r.SupportsStreaming != nil {
		if *r.SupportsStreaming {
			values["supports_streaming"] = "1"
		} else {
			values["supports_streaming"] = "0"
		}
	}

	switch value := r.Thumbnail.(type) {
	case io.Reader:
		values["thumbnail"] = value
	case string:
		values["thumbnail"] = value
	}

	switch value := r.Video.(type) {
	case io.Reader:
		values["video"] = value
	case string:
		values["video"] = value
	}

	if r.Width != nil {
		values["width"] = strconv.FormatInt(int64(*r.Width), 10)
	}

	return
}
