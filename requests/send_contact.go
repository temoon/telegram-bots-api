package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SendContact struct {
	DisableNotification *bool
	ReplyMarkup         interface{}
	ChatId              telegram.ChatId
	MessageThreadId     *int64
	FirstName           string
	LastName            *string
	PhoneNumber         string
	Vcard               *string
	ProtectContent      *bool
	ReplyParameters     *telegram.ReplyParameters
}

func (r *SendContact) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendContact", r, response)
	return
}

func (r *SendContact) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
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

	values["chat_id"] = r.ChatId.String()

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	values["first_name"] = r.FirstName

	if r.LastName != nil {
		values["last_name"] = *r.LastName
	}

	values["phone_number"] = r.PhoneNumber

	if r.Vcard != nil {
		values["vcard"] = *r.Vcard
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

	return
}

func (r *SendContact) GetFiles() (files map[string]io.Reader) {
	return
}
