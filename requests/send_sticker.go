package requests

import (
"encoding/json"
"errors"
"io"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type SendSticker struct {
ChatId interface{}
DisableNotification *bool
Emoji *string
MessageThreadId *int32
ProtectContent *bool
ReplyMarkup *telegram.InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
ReplyParameters *telegram.ReplyParameters
Sticker interface{}
}

func (r *SendSticker) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendSticker", r, response)
	return
}



func (r *SendSticker) IsMultipart() (multipart bool) {
	return false
	}

func (r *SendSticker) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	
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
			
			if r.Emoji != nil {
			values["emoji"] = *r.Emoji
			}
			
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
			
			switch value := r.Sticker.(type) {
			case int64:
					values["sticker"] = strconv.FormatInt(value, 10)
				case string:
					values["sticker"] = value
				default:
				err = errors.New("invalid sticker field type")
				return
			}
		

	return
}
