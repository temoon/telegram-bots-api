package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type EditMessageLiveLocation struct {
	Latitude             float64
	Longitude            float64
	ChatId               *telegram.ChatId
	Heading              *int64
	HorizontalAccuracy   *float64
	InlineMessageId      *string
	MessageId            *int64
	ProximityAlertRadius *int64
	ReplyMarkup          *telegram.InlineKeyboardMarkup
}

func (r *EditMessageLiveLocation) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "editMessageLiveLocation", r, response)
	return
}

func (r *EditMessageLiveLocation) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["latitude"] = strconv.FormatFloat(r.Latitude, 'f', -1, 64)

	values["longitude"] = strconv.FormatFloat(r.Longitude, 'f', -1, 64)

	if r.ChatId != nil {
		values["chat_id"] = r.ChatId.String()
	}

	if r.Heading != nil {
		values["heading"] = strconv.FormatInt(*r.Heading, 10)
	}

	if r.HorizontalAccuracy != nil {
		values["horizontal_accuracy"] = strconv.FormatFloat(*r.HorizontalAccuracy, 'f', -1, 64)
	}

	if r.InlineMessageId != nil {
		values["inline_message_id"] = *r.InlineMessageId
	}

	if r.MessageId != nil {
		values["message_id"] = strconv.FormatInt(*r.MessageId, 10)
	}

	if r.ProximityAlertRadius != nil {
		values["proximity_alert_radius"] = strconv.FormatInt(*r.ProximityAlertRadius, 10)
	}

	if r.ReplyMarkup != nil {
		var dataReplyMarkup []byte
		if dataReplyMarkup, err = json.Marshal(r.ReplyMarkup); err != nil {
			return
		}

		values["reply_markup"] = string(dataReplyMarkup)
	}

	return
}

func (r *EditMessageLiveLocation) GetFiles() (files map[string]io.Reader) {
	return
}
