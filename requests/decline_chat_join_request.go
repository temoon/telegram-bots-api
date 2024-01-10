package requests

import (
"errors"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type DeclineChatJoinRequest struct {
ChatId interface{}
UserId int32
}

func (r *DeclineChatJoinRequest) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "declineChatJoinRequest", r, response)
	return
}



func (r *DeclineChatJoinRequest) IsMultipart() (multipart bool) {
	return false
	}

func (r *DeclineChatJoinRequest) GetValues() (values map[string]interface{}, err error) {
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
		
			values["user_id"] = strconv.FormatInt(int64(r.UserId), 10)
			

	return
}
