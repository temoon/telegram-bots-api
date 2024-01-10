package requests

import (
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type GetGameHighScores struct {
ChatId *int32
InlineMessageId *string
MessageId *int32
UserId int32
}

func (r *GetGameHighScores) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.GameHighScore)
	err = b.CallMethod(ctx, "getGameHighScores", r, response)
	return
}



func (r *GetGameHighScores) IsMultipart() (multipart bool) {
	return false
	}

func (r *GetGameHighScores) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	
			if r.ChatId != nil {
			values["chat_id"] = strconv.FormatInt(int64(*r.ChatId), 10)
			}
			
			if r.InlineMessageId != nil {
			values["inline_message_id"] = *r.InlineMessageId
			}
			
			if r.MessageId != nil {
			values["message_id"] = strconv.FormatInt(int64(*r.MessageId), 10)
			}
			
			values["user_id"] = strconv.FormatInt(int64(r.UserId), 10)
			

	return
}
