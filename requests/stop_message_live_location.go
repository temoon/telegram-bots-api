package requests

import (
"encoding/json"
"errors"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type StopMessageLiveLocation struct {
ChatId interface{}
InlineMessageId *string
MessageId *int32
ReplyMarkup *telegram.InlineKeyboardMarkup
}

func (r *StopMessageLiveLocation) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "stopMessageLiveLocation", r, response)
	return
}



func (r *StopMessageLiveLocation) IsMultipart() (multipart bool) {
	return false
	}

func (r *StopMessageLiveLocation) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	
			switch value := r.ChatId.(type) {
			case *int64:
					values["chat_id"] = strconv.FormatInt(*value, 10)
				case *string:
					values["chat_id"] = *value
				default:
				err = errors.New("invalid chat_id field type")
				return
			}
		
			if r.InlineMessageId != nil {
			values["inline_message_id"] = *r.InlineMessageId
			}
			
			if r.MessageId != nil {
			values["message_id"] = strconv.FormatInt(int64(*r.MessageId), 10)
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
