package requests

import (
"errors"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type BanChatMember struct {
ChatId interface{}
RevokeMessages *bool
UntilDate *int32
UserId int32
}

func (r *BanChatMember) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "banChatMember", r, response)
	return
}



func (r *BanChatMember) IsMultipart() (multipart bool) {
	return false
	}

func (r *BanChatMember) GetValues() (values map[string]interface{}, err error) {
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
		
			if r.RevokeMessages != nil {
			if *r.RevokeMessages {
					values["revoke_messages"] = "1"
				} else {
					values["revoke_messages"] = "0"
				}
			}
			
			if r.UntilDate != nil {
			values["until_date"] = strconv.FormatInt(int64(*r.UntilDate), 10)
			}
			
			values["user_id"] = strconv.FormatInt(int64(r.UserId), 10)
			

	return
}
