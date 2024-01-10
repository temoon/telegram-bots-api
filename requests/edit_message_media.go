package requests

import (
"encoding/json"
"errors"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type EditMessageMedia struct {
ChatId interface{}
InlineMessageId *string
Media interface{}
MessageId *int32
ReplyMarkup *telegram.InlineKeyboardMarkup
}

func (r *EditMessageMedia) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "editMessageMedia", r, response)
	return
}



func (r *EditMessageMedia) IsMultipart() (multipart bool) {
	return false
	}

func (r *EditMessageMedia) GetValues() (values map[string]interface{}, err error) {
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
			
			var dataMedia []byte
				if dataMedia, err = json.Marshal(r.Media); err != nil {
					return
				}

				values["media"] = string(dataMedia)
			
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
