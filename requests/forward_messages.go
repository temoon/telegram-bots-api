package requests

import (
"encoding/json"
"errors"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type ForwardMessages struct {
ChatId interface{}
DisableNotification *bool
FromChatId interface{}
MessageIds []int32
MessageThreadId *int32
ProtectContent *bool
}

func (r *ForwardMessages) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.MessageId)
	err = b.CallMethod(ctx, "forwardMessages", r, response)
	return
}



func (r *ForwardMessages) IsMultipart() (multipart bool) {
	return false
	}

func (r *ForwardMessages) GetValues() (values map[string]interface{}, err error) {
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
			
			switch value := r.FromChatId.(type) {
			case int64:
					values["from_chat_id"] = strconv.FormatInt(value, 10)
				case string:
					values["from_chat_id"] = value
				default:
				err = errors.New("invalid from_chat_id field type")
				return
			}
		
			var dataMessageIds []byte
				if dataMessageIds, err = json.Marshal(r.MessageIds); err != nil {
					return
				}

				values["message_ids"] = string(dataMessageIds)
			
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
			

	return
}
