package requests

import (
"encoding/json"
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type SendGame struct {
ChatId int32
DisableNotification *bool
GameShortName string
MessageThreadId *int32
ProtectContent *bool
ReplyMarkup *telegram.InlineKeyboardMarkup
ReplyParameters *telegram.ReplyParameters
}

func (r *SendGame) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendGame", r, response)
	return
}



func (r *SendGame) IsMultipart() (multipart bool) {
	return false
	}

func (r *SendGame) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	
			values["chat_id"] = strconv.FormatInt(int64(r.ChatId), 10)
			
			if r.DisableNotification != nil {
			if *r.DisableNotification {
					values["disable_notification"] = "1"
				} else {
					values["disable_notification"] = "0"
				}
			}
			
			values["game_short_name"] = r.GameShortName
			
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
			

	return
}
