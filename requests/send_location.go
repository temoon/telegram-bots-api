package requests

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SendLocation struct {
	ChatId                  telegram.ChatId
	Latitude                float64
	Longitude               float64
	AllowPaidBroadcast      *bool
	BusinessConnectionId    *string
	DirectMessagesTopicId   *int64
	DisableNotification     *bool
	Heading                 *int64
	HorizontalAccuracy      *float64
	LivePeriod              *int64
	MessageEffectId         *string
	MessageThreadId         *int64
	ProtectContent          *bool
	ProximityAlertRadius    *int64
	ReplyMarkup             interface{}
	ReplyParameters         *telegram.ReplyParameters
	SuggestedPostParameters *telegram.SuggestedPostParameters
}

func (r *SendLocation) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendLocation", r, response)
	return
}

func (r *SendLocation) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["latitude"] = strconv.FormatFloat(r.Latitude, 'f', -1, 64)

	values["longitude"] = strconv.FormatFloat(r.Longitude, 'f', -1, 64)

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

	if r.Heading != nil {
		values["heading"] = strconv.FormatInt(*r.Heading, 10)
	}

	if r.HorizontalAccuracy != nil {
		values["horizontal_accuracy"] = strconv.FormatFloat(*r.HorizontalAccuracy, 'f', -1, 64)
	}

	if r.LivePeriod != nil {
		values["live_period"] = strconv.FormatInt(*r.LivePeriod, 10)
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

	if r.ProximityAlertRadius != nil {
		values["proximity_alert_radius"] = strconv.FormatInt(*r.ProximityAlertRadius, 10)
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

	return
}

func (r *SendLocation) GetFiles() (files map[string]io.Reader) {
	return
}
