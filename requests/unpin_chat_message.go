package requests

import (
"errors"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type UnpinChatMessage struct {
ChatId interface{}
MessageId *int32
}

func (r *UnpinChatMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unpinChatMessage", r, response)
	return
}



func (r *UnpinChatMessage) IsMultipart() (multipart bool) {
	return false
	}

func (r *UnpinChatMessage) GetValues() (values map[string]interface{}, err error) {
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
		
			if r.MessageId != nil {
			values["message_id"] = strconv.FormatInt(int64(*r.MessageId), 10)
			}
			

	return
}
