package requests

import (
"encoding/json"
"errors"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type SetMessageReaction struct {
ChatId interface{}
IsBig *bool
MessageId int32
Reaction []interface{}
}

func (r *SetMessageReaction) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setMessageReaction", r, response)
	return
}



func (r *SetMessageReaction) IsMultipart() (multipart bool) {
	return false
	}

func (r *SetMessageReaction) GetValues() (values map[string]interface{}, err error) {
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
		
			if r.IsBig != nil {
			if *r.IsBig {
					values["is_big"] = "1"
				} else {
					values["is_big"] = "0"
				}
			}
			
			values["message_id"] = strconv.FormatInt(int64(r.MessageId), 10)
			
			if r.Reaction != nil {
			var dataReaction []byte
				if dataReaction, err = json.Marshal(r.Reaction); err != nil {
					return
				}

				values["reaction"] = string(dataReaction)
			}
			

	return
}
