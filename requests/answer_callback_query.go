package requests

import (
	"strconv"
)

type AnswerCallbackQuery struct {
	CacheTime       uint64
	CallbackQueryId string
	ShowAlert       bool
	Text            string
	Url             string
}

func (r *AnswerCallbackQuery) IsMultipart() bool {
	return false
}

func (r *AnswerCallbackQuery) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.CacheTime != 0 {
		values["cache_time"] = strconv.FormatUint(r.CacheTime, 10)
	}

	values["callback_query_id"] = r.CallbackQueryId

	if r.ShowAlert {
		values["show_alert"] = "1"
	}

	if r.Text != "" {
		values["text"] = r.Text
	}

	if r.Url != "" {
		values["url"] = r.Url
	}

	return
}
