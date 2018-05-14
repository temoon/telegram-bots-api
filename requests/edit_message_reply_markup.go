package requests

import (
    "encoding/json"
    "errors"
    "strconv"
)

type EditMessageReplyMarkup struct {
    ChatID          interface{}
    MessageID       int
    InlineMessageID string
    ReplyMarkup     interface{}
}

func (r *EditMessageReplyMarkup) IsMultipart() bool {
    return false
}

func (r *EditMessageReplyMarkup) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID.(int64), 10)}
    case string:
        values["chat_id"] = []interface{}{r.ChatID.(string)}
    default:
        return nil, errors.New("invalid chat_id")
    }

    if r.MessageID != 0 {
        values["message_id"] = []interface{}{strconv.Itoa(r.MessageID)}
    }

    if r.InlineMessageID != "" {
        values["inline_message_id"] = []interface{}{r.InlineMessageID}
    }

    if r.ReplyMarkup != nil {
        var data []byte
        if data, err = json.Marshal(r.ReplyMarkup); err != nil {
            return
        }

        values["reply_markup"] = []interface{}{string(data)}
    }

    return
}
