package requests

import (
	"encoding/json"
	"strconv"
)

type SendVenue struct {
	Address                  string
	AllowSendingWithoutReply bool
	ChatId                   interface{}
	DisableNotification      bool
	FoursquareId             string
	FoursquareType           string
	GooglePlaceId            string
	GooglePlaceType          string
	Latitude                 float64
	Longitude                float64
	ReplyMarkup              interface{}
	ReplyToMessageId         uint64
	Title                    string
}

func (r *SendVenue) IsMultipart() bool {
	return false
}

func (r *SendVenue) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["address"] = r.Address

	if r.AllowSendingWithoutReply {
		values["allow_sending_without_reply"] = "1"
	}

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.DisableNotification {
		values["disable_notification"] = "1"
	}

	if r.FoursquareId != "" {
		values["foursquare_id"] = r.FoursquareId
	}

	if r.FoursquareType != "" {
		values["foursquare_type"] = r.FoursquareType
	}

	if r.GooglePlaceId != "" {
		values["google_place_id"] = r.GooglePlaceId
	}

	if r.GooglePlaceType != "" {
		values["google_place_type"] = r.GooglePlaceType
	}

	switch value := r.ReplyMarkup.(type) {
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		}
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		}
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		}
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		}
	}

	if r.ReplyToMessageId != 0 {
		values["reply_to_message_id"] = strconv.FormatUint(r.ReplyToMessageId, 10)
	}

	values["title"] = r.Title

	return
}
