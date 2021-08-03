package requests

import (
	"encoding/json"
	"strconv"
)

type EditMessageLiveLocation struct {
	ChatId               interface{}
	Heading              uint64
	HorizontalAccuracy   float64
	InlineMessageId      string
	Latitude             float64
	Longitude            float64
	MessageId            uint64
	ProximityAlertRadius uint64
	ReplyMarkup          interface{}
}

func (r *EditMessageLiveLocation) IsMultipart() bool {
	return false
}

func (r *EditMessageLiveLocation) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.Heading != 0 {
		values["heading"] = strconv.FormatUint(r.Heading, 10)
	}

	if r.InlineMessageId != "" {
		values["inline_message_id"] = r.InlineMessageId
	}

	if r.MessageId != 0 {
		values["message_id"] = strconv.FormatUint(r.MessageId, 10)
	}

	if r.ProximityAlertRadius != 0 {
		values["proximity_alert_radius"] = strconv.FormatUint(r.ProximityAlertRadius, 10)
	}

	if r.ReplyMarkup != nil {
		var data []byte
		if data, err = json.Marshal(r.ReplyMarkup); err != nil {
			return
		}

		values["reply_markup"] = string(data)
	}

	return
}
