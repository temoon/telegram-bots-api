package requests

import (
    "encoding/json"
    "strconv"
)

type AnswerInlineQuery struct {
    InlineQueryID     string
    Results           []interface{}
    CacheTime         uint32
    IsPersonal        bool
    NextOffset        string
    SwitchPMText      string
    SwitchPMParameter string
}

func (r *AnswerInlineQuery) IsMultipart() bool {
    return false
}

func (r *AnswerInlineQuery) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["inline_query_id"] = r.InlineQueryID

    var data []byte
    if data, err = json.Marshal(r.Results); err != nil {
        return
    }

    values["results"] = string(data)

    if r.CacheTime != 0 {
        values["cache_time"] = strconv.FormatUint(uint64(r.CacheTime), 10)
    }

    if r.IsPersonal {
        values["is_personal"] = "1"
    }

    if r.NextOffset != "" {
        values["next_offset"] = r.NextOffset
    }

    if r.SwitchPMText != "" {
        values["switch_pm_text"] = r.SwitchPMText
    }

    if r.SwitchPMParameter != "" {
        values["switch_pm_parameter"] = r.SwitchPMParameter
    }

    return
}
