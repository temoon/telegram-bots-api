package requests

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type AnswerInlineQuery struct {
	CacheTime         *int32
	InlineQueryId     string
	IsPersonal        *bool
	NextOffset        *string
	Results           []interface{}
	SwitchPmParameter *string
	SwitchPmText      *string
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

	if r.SwitchPmParameter != nil {
		values["switch_pm_parameter"] = *r.SwitchPmParameter
	}

	if r.SwitchPmText != nil {
		values["switch_pm_text"] = *r.SwitchPmText
	}

	return
}
