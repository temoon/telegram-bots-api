package requests

import (
	"encoding/json"
	"strconv"
)

type AnswerInlineQuery struct {
	CacheTime         uint64
	InlineQueryId     string
	IsPersonal        bool
	NextOffset        string
	Results           []interface{}
	SwitchPmParameter string
	SwitchPmText      string
}

func (r *AnswerInlineQuery) IsMultipart() bool {
	return false
}

func (r *AnswerInlineQuery) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.CacheTime != 0 {
		values["cache_time"] = strconv.FormatUint(r.CacheTime, 10)
	}

	values["inline_query_id"] = r.InlineQueryId

	if r.IsPersonal {
		values["is_personal"] = "1"
	}

	if r.NextOffset != "" {
		values["next_offset"] = r.NextOffset
	}

	if r.Results != nil {
		var data []byte
		if data, err = json.Marshal(r.Results); err != nil {
			return
		}

		values["results"] = string(data)
	}

	if r.SwitchPmParameter != "" {
		values["switch_pm_parameter"] = r.SwitchPmParameter
	}

	if r.SwitchPmText != "" {
		values["switch_pm_text"] = r.SwitchPmText
	}

	return
}
