package requests

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type SendVenue struct {
	Address                  string
	AllowSendingWithoutReply *bool
	ChatId                   interface{}
	DisableNotification      *bool
	FoursquareId             *string
	FoursquareType           *string
	GooglePlaceId            *string
	GooglePlaceType          *string
	Latitude                 float64
	Longitude                float64
	ReplyMarkup              interface{}
	ReplyToMessageId         *int32
	Title                    string
}

func (r *SendVenue) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendVenue", r, response)
	return
}

func (r *SendVenue) IsMultipart() (multipart bool) {
	return false
}

func (r *SendVenue) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["address"] = r.Address

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

	values["title"] = r.Title

	return
}
