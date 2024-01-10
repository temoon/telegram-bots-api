package requests

import (
"strconv"
"context"
	"github.com/temoon/telegram-bots-api"
)

type AnswerCallbackQuery struct {
CacheTime *int32
CallbackQueryId string
ShowAlert *bool
Text *string
Url *string
}

func (r *AnswerCallbackQuery) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "answerCallbackQuery", r, response)
	return
}



func (r *AnswerCallbackQuery) IsMultipart() (multipart bool) {
	return false
	}

func (r *AnswerCallbackQuery) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	
			if r.CacheTime != nil {
			values["cache_time"] = strconv.FormatInt(int64(*r.CacheTime), 10)
			}
			
			values["callback_query_id"] = r.CallbackQueryId
			
			if r.ShowAlert != nil {
			if *r.ShowAlert {
					values["show_alert"] = "1"
				} else {
					values["show_alert"] = "0"
				}
			}
			
			if r.Text != nil {
			values["text"] = *r.Text
			}
			
			if r.Url != nil {
			values["url"] = *r.Url
			}
			

	return
}
