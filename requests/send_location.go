package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type SendLocation struct {
	AllowSendingWithoutReply *bool
	ChatId                   interface{}
	DisableNotification      *bool
	Heading                  *int32
	HorizontalAccuracy       *float64
	Latitude                 float64
	LivePeriod               *int32
	Longitude                float64
	MessageThreadId          *int32
	ProtectContent           *bool
	ProximityAlertRadius     *int32
	ReplyMarkup              interface{}
	ReplyToMessageId         *int32
}

func (r *SendLocation) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendLocation", r, response)
	return
}

func (r *SendLocation) IsMultipart() (multipart bool) {
	return false
}

func (r *SendLocation) GetValues() (values map[string]interface{}, err error) {
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

	if r.Heading != nil {
		values["heading"] = strconv.FormatInt(int64(*r.Heading), 10)
	}

	if r.LivePeriod != nil {
		values["live_period"] = strconv.FormatInt(int64(*r.LivePeriod), 10)
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

	if r.ProximityAlertRadius != nil {
		values["proximity_alert_radius"] = strconv.FormatInt(int64(*r.ProximityAlertRadius), 10)
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

	return
}
