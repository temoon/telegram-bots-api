package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SendLocation struct {
	ProtectContent       *bool
	ReplyMarkup          interface{}
	ChatId               telegram.ChatId
	Latitude             float64
	Heading              *int64
	ProximityAlertRadius *int64
	DisableNotification  *bool
	MessageThreadId      *int64
	Longitude            float64
	HorizontalAccuracy   *float64
	LivePeriod           *int64
	ReplyParameters      *telegram.ReplyParameters
}

func (r *SendLocation) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendLocation", r, response)
	return
}

func (r *SendLocation) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

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

	values["chat_id"] = r.ChatId.String()

	values["latitude"] = strconv.FormatFloat(r.Latitude, 'f', -1, 64)

	if r.Heading != nil {
		values["heading"] = strconv.FormatInt(*r.Heading, 10)
	}

	if r.ProximityAlertRadius != nil {
		values["proximity_alert_radius"] = strconv.FormatInt(*r.ProximityAlertRadius, 10)
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	values["longitude"] = strconv.FormatFloat(r.Longitude, 'f', -1, 64)

	if r.HorizontalAccuracy != nil {
		values["horizontal_accuracy"] = strconv.FormatFloat(*r.HorizontalAccuracy, 'f', -1, 64)
	}

	if r.LivePeriod != nil {
		values["live_period"] = strconv.FormatInt(*r.LivePeriod, 10)
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

func (r *SendLocation) GetFiles() (files map[string]io.Reader) {
	return
}
