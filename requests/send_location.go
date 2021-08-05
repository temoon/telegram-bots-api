package requests

import (
	"encoding/json"
	"strconv"
)

type SendLocation struct {
	AllowSendingWithoutReply bool
	ChatId                   interface{}
	DisableNotification      bool
	Heading                  uint64
	HorizontalAccuracy       float64
	Latitude                 float64
	LivePeriod               uint64
	Longitude                float64
	ProximityAlertRadius     uint64
	ReplyMarkup              interface{}
	ReplyToMessageId         uint64
}

func (r *SendLocation) IsMultipart() bool {
	return false
}

func (r *SendLocation) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

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

	if r.Heading != 0 {
		values["heading"] = strconv.FormatUint(r.Heading, 10)
	}

	if r.LivePeriod != 0 {
		values["live_period"] = strconv.FormatUint(r.LivePeriod, 10)
	}

	if r.ProximityAlertRadius != 0 {
		values["proximity_alert_radius"] = strconv.FormatUint(r.ProximityAlertRadius, 10)
	}

	switch value := r.ReplyMarkup.(type) {
	default:

		var data []byte
		if data, err = json.Marshal(value); err != nil {
			return
		}

		values["reply_markup"] = string(data)

	}

	if r.ReplyToMessageId != 0 {
		values["reply_to_message_id"] = strconv.FormatUint(r.ReplyToMessageId, 10)
	}

	return
}
