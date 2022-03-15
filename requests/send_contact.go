package requests

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type SendContact struct {
	AllowSendingWithoutReply *bool
	ChatId                   interface{}
	DisableNotification      *bool
	FirstName                string
	LastName                 *string
	PhoneNumber              string
	ProtectContent           *bool
	ReplyMarkup              interface{}
	ReplyToMessageId         *int32
	Vcard                    *string
}

func (r *SendContact) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendContact", r, response)
	return
}

func (r *SendContact) IsMultipart() (multipart bool) {
	return false
}

func (r *SendContact) GetValues() (values map[string]interface{}, err error) {
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

	values["first_name"] = r.FirstName

	if r.LastName != nil {
		values["last_name"] = *r.LastName
	}

	values["phone_number"] = r.PhoneNumber

	if r.ProtectContent != nil {
		if *r.ProtectContent {
			values["protect_content"] = "1"
		} else {
			values["protect_content"] = "0"
		}
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

	if r.Vcard != nil {
		values["vcard"] = *r.Vcard
	}

	return
}
