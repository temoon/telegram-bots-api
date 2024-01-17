package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type AnswerInlineQuery struct {
	InlineQueryId string
	Results       []interface{}
	CacheTime     *int64
	IsPersonal    *bool
	NextOffset    *string
	Button        *telegram.InlineQueryResultsButton
}

func (r *AnswerInlineQuery) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "answerInlineQuery", r, response)
	return
}

func (r *AnswerInlineQuery) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["inline_query_id"] = r.InlineQueryId

	var dataResults []byte
	if dataResults, err = json.Marshal(r.Results); err != nil {
		return
	}

	values["results"] = string(dataResults)

	if r.CacheTime != nil {
		values["cache_time"] = strconv.FormatInt(*r.CacheTime, 10)
	}

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

	if r.Button != nil {
		var dataButton []byte
		if dataButton, err = json.Marshal(r.Button); err != nil {
			return
		}

		values["button"] = string(dataButton)
	}

	return
}

func (r *AnswerInlineQuery) GetFiles() (files map[string]io.Reader) {
	return
}
