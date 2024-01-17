package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SendSticker struct {
	DisableNotification *bool
	ProtectContent      *bool
	ReplyParameters     *telegram.ReplyParameters
	ReplyMarkup         interface{}
	ChatId              telegram.ChatId
	MessageThreadId     *int64
	Sticker             telegram.InputFile
	Emoji               *string
}

func (r *SendSticker) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendSticker", r, response)
	return
}

func (r *SendSticker) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	if r.ProtectContent != nil {
		if *r.ProtectContent {
			values["protect_content"] = "1"
		} else {
			values["protect_content"] = "0"
		}
	}

	if r.ReplyParameters != nil {
		var dataReplyParameters []byte
		if dataReplyParameters, err = json.Marshal(r.ReplyParameters); err != nil {
			return
		}

		values["reply_parameters"] = string(dataReplyParameters)
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

	values["chat_id"] = r.ChatId.String()

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	values["sticker"] = r.Sticker.GetValue()

	if r.Emoji != nil {
		values["emoji"] = *r.Emoji
	}

	return
}

func (r *SendSticker) GetFiles() (files map[string]io.Reader) {
	return
}
