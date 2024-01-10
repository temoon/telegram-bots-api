package requests

import (
"encoding/json"
"errors"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type SendVenue struct {
Address string
ChatId interface{}
DisableNotification *bool
FoursquareId *string
FoursquareType *string
GooglePlaceId *string
GooglePlaceType *string
Latitude float64
Longitude float64
MessageThreadId *int32
ProtectContent *bool
ReplyMarkup *telegram.InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
ReplyParameters *telegram.ReplyParameters
Title string
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
			
			switch value := r.ChatId.(type) {
			case int64:
					values["chat_id"] = strconv.FormatInt(value, 10)
				case string:
					values["chat_id"] = value
				default:
				err = errors.New("invalid chat_id field type")
				return
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
			
			values["latitude"] = strconv.FormatFloat(r.Latitude, 'f', -1, 64)
			
			values["longitude"] = strconv.FormatFloat(r.Longitude, 'f', -1, 64)
			
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
			
			if r.ReplyMarkup != nil {
			var dataReplyMarkup []byte
				if dataReplyMarkup, err = json.Marshal(r.ReplyMarkup); err != nil {
					return
				}

				values["reply_markup"] = string(dataReplyMarkup)
			}
			
			if r.ReplyParameters != nil {
			var dataReplyParameters []byte
				if dataReplyParameters, err = json.Marshal(r.ReplyParameters); err != nil {
					return
				}

				values["reply_parameters"] = string(dataReplyParameters)
			}
			
			values["title"] = r.Title
			

	return
}
