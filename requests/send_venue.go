package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SendVenue struct {
	Address             string
	ChatId              telegram.ChatId
	Latitude            float64
	Longitude           float64
	Title               string
	DisableNotification *bool
	FoursquareId        *string
	FoursquareType      *string
	GooglePlaceId       *string
	GooglePlaceType     *string
	MessageThreadId     *int64
	ProtectContent      *bool
	ReplyMarkup         interface{}
	ReplyParameters     *telegram.ReplyParameters
}

func (r *SendVenue) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendVenue", r, response)
	return
}

func (r *SendVenue) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["address"] = r.Address

	values["chat_id"] = r.ChatId.String()

	values["latitude"] = strconv.FormatFloat(r.Latitude, 'f', -1, 64)

	values["longitude"] = strconv.FormatFloat(r.Longitude, 'f', -1, 64)

	values["title"] = r.Title

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	if r.FoursquareId != nil {
		values["foursquare_id"] = *r.FoursquareId
	}

	if r.FoursquareType != nil {
		values["foursquare_type"] = *r.FoursquareType
	}

	if r.GooglePlaceId != nil {
		values["google_place_id"] = *r.GooglePlaceId
	}

	if r.GooglePlaceType != nil {
		values["google_place_type"] = *r.GooglePlaceType
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

	return
}

func (r *SendVenue) GetFiles() (files map[string]io.Reader) {
	return
}
