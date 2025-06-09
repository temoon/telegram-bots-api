package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SendVideoNote struct {
	ChatId               telegram.ChatId
	VideoNote            telegram.InputFile
	AllowPaidBroadcast   *bool
	BusinessConnectionId *string
	DisableNotification  *bool
	Duration             *int64
	Length               *int64
	MessageEffectId      *string
	MessageThreadId      *int64
	ProtectContent       *bool
	ReplyMarkup          interface{}
	ReplyParameters      *telegram.ReplyParameters
	Thumbnail            *telegram.InputFile
}

func (r *SendVideoNote) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendVideoNote", r, response)
	return
}

func (r *SendVideoNote) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["video_note"] = r.VideoNote.GetValue()

	if r.AllowPaidBroadcast != nil {
		if *r.AllowPaidBroadcast {
			values["allow_paid_broadcast"] = "1"
		} else {
			values["allow_paid_broadcast"] = "0"
		}
	}

	if r.BusinessConnectionId != nil {
		values["business_connection_id"] = *r.BusinessConnectionId
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	if r.Duration != nil {
		values["duration"] = strconv.FormatInt(*r.Duration, 10)
	}

	if r.Length != nil {
		values["length"] = strconv.FormatInt(*r.Length, 10)
	}

	if r.MessageEffectId != nil {
		values["message_effect_id"] = *r.MessageEffectId
	}

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	if r.ProtectContent != nil {
		if *r.ProtectContent {
			values["protect_content"] = "1"
		} else {
			values["protect_content"] = "0"
		}
	}

	if r.ReplyMarkup != nil {
		switch value := r.ReplyMarkup.(type) {
		case telegram.InlineKeyboardMarkup, telegram.ReplyKeyboardMarkup, telegram.ReplyKeyboardRemove, telegram.ForceReply:
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		default:
			err = errors.New("unsupported reply_markup field type")
			return
		}
	}

	if r.ReplyParameters != nil {
		var dataReplyParameters []byte
		if dataReplyParameters, err = json.Marshal(r.ReplyParameters); err != nil {
			return
		}

		values["reply_parameters"] = string(dataReplyParameters)
	}

	if r.Thumbnail != nil {
		values["thumbnail"] = r.Thumbnail.GetValue()
	}

	return
}

func (r *SendVideoNote) GetFiles() (files map[string]io.Reader) {
	return
}
