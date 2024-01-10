package requests

import (
"encoding/json"
"errors"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type CopyMessage struct {
Caption *string
CaptionEntities []telegram.MessageEntity
ChatId interface{}
DisableNotification *bool
FromChatId interface{}
MessageId int32
MessageThreadId *int32
ParseMode *string
ProtectContent *bool
ReplyMarkup *telegram.InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
ReplyParameters *telegram.ReplyParameters
}

func (r *CopyMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.MessageId)
	err = b.CallMethod(ctx, "copyMessage", r, response)
	return
}



func (r *CopyMessage) IsMultipart() (multipart bool) {
	return false
	}

func (r *CopyMessage) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	
			if r.Caption != nil {
			values["caption"] = *r.Caption
			}
			
			if r.CaptionEntities != nil {
			var dataCaptionEntities []byte
				if dataCaptionEntities, err = json.Marshal(r.CaptionEntities); err != nil {
					return
				}

				values["caption_entities"] = string(dataCaptionEntities)
			}
			
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
			
			switch value := r.FromChatId.(type) {
			case int64:
					values["from_chat_id"] = strconv.FormatInt(value, 10)
				case string:
					values["from_chat_id"] = value
				default:
				err = errors.New("invalid from_chat_id field type")
				return
			}
		
			values["message_id"] = strconv.FormatInt(int64(r.MessageId), 10)
			
			if r.MessageThreadId != nil {
			values["message_thread_id"] = strconv.FormatInt(int64(*r.MessageThreadId), 10)
			}
			
			if r.ParseMode != nil {
			values["parse_mode"] = *r.ParseMode
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
			

	return
}
