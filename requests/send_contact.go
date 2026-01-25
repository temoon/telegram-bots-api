package requests

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SendContact struct {
	ChatId                  telegram.ChatId
	FirstName               string
	PhoneNumber             string
	AllowPaidBroadcast      *bool
	BusinessConnectionId    *string
	DirectMessagesTopicId   *int64
	DisableNotification     *bool
	LastName                *string
	MessageEffectId         *string
	MessageThreadId         *int64
	ProtectContent          *bool
	ReplyMarkup             interface{}
	ReplyParameters         *telegram.ReplyParameters
	SuggestedPostParameters *telegram.SuggestedPostParameters
	Vcard                   *string
}

func (r *SendContact) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendContact", r, response)
	return
}

func (r *SendContact) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	values["first_name"] = r.FirstName

	values["phone_number"] = r.PhoneNumber

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

	if r.DirectMessagesTopicId != nil {
		values["direct_messages_topic_id"] = strconv.FormatInt(*r.DirectMessagesTopicId, 10)
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	if r.LastName != nil {
		values["last_name"] = *r.LastName
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

	if r.SuggestedPostParameters != nil {
		var dataSuggestedPostParameters []byte
		if dataSuggestedPostParameters, err = json.Marshal(r.SuggestedPostParameters); err != nil {
			return
		}

		values["suggested_post_parameters"] = string(dataSuggestedPostParameters)
	}

	if r.Vcard != nil {
		values["vcard"] = *r.Vcard
	}

	return
}

func (r *SendContact) GetFiles() (files map[string]io.Reader) {
	return
}
