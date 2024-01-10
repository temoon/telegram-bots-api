package requests

import (
"errors"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type UnbanChatSenderChat struct {
ChatId interface{}
SenderChatId int32
}

func (r *UnbanChatSenderChat) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unbanChatSenderChat", r, response)
	return
}



func (r *UnbanChatSenderChat) IsMultipart() (multipart bool) {
	return false
	}

func (r *UnbanChatSenderChat) GetValues() (values map[string]interface{}, err error) {
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
		
			values["sender_chat_id"] = strconv.FormatInt(int64(r.SenderChatId), 10)
			

	return
}
