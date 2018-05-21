package requests

import (
    "strconv"
)

type AnswerCallbackQuery struct {
    CallbackQueryID string
    Text            string
    ShowAlert       bool
    URL             string
    CacheTime       uint32
}

func (r *AnswerCallbackQuery) IsMultipart() bool {
    return false
}

func (r *AnswerCallbackQuery) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    values["callback_query_id"] = r.CallbackQueryID

    if r.Text != "" {
        values["text"] = r.Text
    }

    if r.ShowAlert {
        values["show_alert"] = "1"
    }

    if r.URL != "" {
        values["url"] = r.URL
    }

    if r.CacheTime != 0 {
        values["cache_time"] = strconv.FormatUint(uint64(r.CacheTime), 10)
    }

    return
}
