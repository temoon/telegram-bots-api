package requests

import (
"encoding/json"
"errors"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type SendMessage struct {
ChatId interface{}
DisableNotification *bool
Entities []telegram.MessageEntity
LinkPreviewOptions *telegram.LinkPreviewOptions
MessageThreadId *int32
ParseMode *string
ProtectContent *bool
ReplyMarkup *telegram.InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
ReplyParameters *telegram.ReplyParameters
Text string
}

func (r *SendMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendMessage", r, response)
	return
}



func (r *SendMessage) IsMultipart() (multipart bool) {
	return false
	}

func (r *SendMessage) GetValues() (values map[string]interface{}, err error) {
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
			
			if r.Entities != nil {
			var dataEntities []byte
				if dataEntities, err = json.Marshal(r.Entities); err != nil {
					return
				}

				values["entities"] = string(dataEntities)
			}
			
			if r.LinkPreviewOptions != nil {
			var dataLinkPreviewOptions []byte
				if dataLinkPreviewOptions, err = json.Marshal(r.LinkPreviewOptions); err != nil {
					return
				}

				values["link_preview_options"] = string(dataLinkPreviewOptions)
			}
			
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
			
			values["text"] = r.Text
			

	return
}
