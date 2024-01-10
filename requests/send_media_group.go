package requests

import (
"encoding/json"
"errors"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type SendMediaGroup struct {
ChatId interface{}
DisableNotification *bool
Media []telegram.InputMediaAudio , InputMediaDocument , InputMediaPhoto and InputMediaVideo
MessageThreadId *int32
ProtectContent *bool
ReplyParameters *telegram.ReplyParameters
}

func (r *SendMediaGroup) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.Message)
	err = b.CallMethod(ctx, "sendMediaGroup", r, response)
	return
}



func (r *SendMediaGroup) IsMultipart() (multipart bool) {
	return false
	}

func (r *SendMediaGroup) GetValues() (values map[string]interface{}, err error) {
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
			
			var dataMedia []byte
				if dataMedia, err = json.Marshal(r.Media); err != nil {
					return
				}

				values["media"] = string(dataMedia)
			
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
			
			if r.ReplyParameters != nil {
			var dataReplyParameters []byte
				if dataReplyParameters, err = json.Marshal(r.ReplyParameters); err != nil {
					return
				}

				values["reply_parameters"] = string(dataReplyParameters)
			}
			

	return
}
