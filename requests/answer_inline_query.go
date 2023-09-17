package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type AnswerInlineQuery struct {
	Button        *telegram.InlineQueryResultsButton
	CacheTime     *int32
	InlineQueryId string
	IsPersonal    *bool
	NextOffset    *string
	Results       []interface{}
}

func (r *AnswerInlineQuery) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "answerInlineQuery", r, response)
	return
}

func (r *AnswerInlineQuery) IsMultipart() (multipart bool) {
	return false
}

func (r *AnswerInlineQuery) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.Button != nil {
		var dataButton []byte
		if dataButton, err = json.Marshal(r.Button); err != nil {
			return
		}

		values["button"] = string(dataButton)
	}

	if r.CacheTime != nil {
		values["cache_time"] = strconv.FormatInt(int64(*r.CacheTime), 10)
	}

	values["inline_query_id"] = r.InlineQueryId

	if r.IsPersonal != nil {
		if *r.IsPersonal {
			values["is_personal"] = "1"
		} else {
			values["is_personal"] = "0"
		}
	}

	if r.NextOffset != nil {
		values["next_offset"] = *r.NextOffset
	}

	var dataResults []byte
	if dataResults, err = json.Marshal(r.Results); err != nil {
		return
	}

	values["results"] = string(dataResults)

	return
}
