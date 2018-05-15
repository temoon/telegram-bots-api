package requests

import (
    "encoding/json"
    "errors"
    "strconv"
)

type EditMessageCaption struct {
    ChatID          interface{}
    MessageID       int
    InlineMessageID string
    Caption         string
    ParseMode       string
    ReplyMarkup     interface{}
}

func (r *EditMessageCaption) IsMultipart() bool {
    return false
}

func (r *EditMessageCaption) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = strconv.FormatInt(r.ChatID.(int64), 10)
    case string:
        values["chat_id"] = r.ChatID.(string)
    default:
        return nil, errors.New("invalid chat_id")
    }

    if r.MessageID != 0 {
        values["message_id"] = strconv.Itoa(r.MessageID)
    }

    if r.InlineMessageID != "" {
        values["inline_message_id"] = r.InlineMessageID
    }

    if r.Caption != "" {
        values["caption"] = r.Caption
    }

    if r.ParseMode != "" {
        values["parse_mode"] = r.ParseMode
    }

    if r.ReplyMarkup != nil {
        var data []byte
        if data, err = json.Marshal(r.ReplyMarkup); err != nil {
            return
        }

        values["reply_markup"] = string(data)
    }

    return
}
